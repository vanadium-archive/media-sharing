// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/x/lib/cmdline"
	"v.io/x/lib/vlog"
	"v.io/x/media_sharing"
	"v.io/x/ref/lib/signals"
	"v.io/x/ref/lib/v23cmd"
	_ "v.io/x/ref/runtime/factories/static"
)

func main() {
	cmdline.Main(root)
}

var root = &cmdline.Command{
	Name:     os.Args[0],
	ArgsName: "<name>",
	ArgsLong: "<name> Is the name to mount the server under.",
	Runner:   v23cmd.RunnerFunc(serve),
	Short:    "Serve a display to remote clients.",
}

func serve(ctx *context.T, env *cmdline.Env, args []string) error {
	name := ""
	if len(args) > 0 {
		name = args[0]
		fmt.Printf("mounting under: %s\n", name)
	}

	server, err := v23.NewServer(ctx)
	if err != nil {
		return err
	}
	eps, err := server.Listen(v23.GetListenSpec(ctx))
	if err != nil {
		return err
	}
	if err := server.Serve(name, defaultMediaServer(), security.AllowEveryone()); err != nil {
		return err
	}
	fmt.Printf("Listening at: %s", eps[0].Name())

	<-signals.ShutdownOnSignals(ctx)
	return nil
}

// handler defines the interface for a media handler.
type handler interface {
	// Matches returns true if this handler supports mimetype.
	Matches(mimetype string) bool

	// Display starts displaying the content in r.
	// This method should not be blocking and shuould return quickly.
	// The returned function should be called to completely stop and clean up
	// the media playback before starting a new display.
	Display(ctx *context.T, mimetype string, r io.ReadCloser) (func(), error)
}

type eogHandler struct{}

func (eogHandler) Matches(mimetype string) bool {
	return strings.HasPrefix(mimetype, "image/")
}

func (eogHandler) Display(ctx *context.T, mimetype string, r io.ReadCloser) (func(), error) {
	// eog cannot read from a pipe, so we have to write the file to
	// the filesystem before displaying it.
	defer r.Close()
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(tmp, r); err != nil {
		return nil, err
	}

	cmd := exec.Command("eog", "-f", tmp.Name())
	stop := func() {
		if cmd.Process != nil {
			if err := cmd.Process.Kill(); err != nil {
				vlog.Errorf("Could not kill eog: %v", err)
			}
		}
		os.Remove(tmp.Name())
	}
	if err := cmd.Start(); err != nil {
		return stop, err
	}
	return stop, nil
}

type vlcHandler struct{}

func (vlcHandler) Matches(mimetype string) bool {
	return strings.HasPrefix(mimetype, "audio/") || strings.HasPrefix(mimetype, "video/")
}

func (vlcHandler) Display(ctx *context.T, mimetype string, r io.ReadCloser) (func(), error) {
	args := []string{
		"--no-video-title-show",
		"--fullscreen",
		"-",
		"vlc://quit",
	}

	if strings.HasPrefix(mimetype, "audio/") {
		args = append([]string{"--audio-visual=visual"}, args...)
	}

	cmd := exec.Command("vlc", args...)
	cmd.Stdin = r
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return func() {
		r.Close()
		if err := cmd.Process.Kill(); err != nil {
			vlog.Errorf("Could not kill vlc: %v", err)
		}
	}, nil
}

type media struct {
	handlers []handler

	stopMu sync.Mutex
	stop   func()
}

func defaultMediaServer() media_sharing.MediaSharingServerStub {
	m := &media{
		handlers: []handler{
			&eogHandler{},
			&vlcHandler{},
		},
	}
	return media_sharing.MediaSharingServer(m)
}

func detectMimeType(r *bufio.Reader) string {
	// If the server didn't tell us the content type, we will have to guess.
	// Note that I am purposefully ignoring the error from Peek.  The guesser looks
	// at up to 512 bytes.  We just get as many as we can and make our best
	// guess with that.
	data, _ := r.Peek(512)
	return http.DetectContentType(data)
}

type bufferedReadCloser struct {
	*bufio.Reader
	close func() error
}

func (bc *bufferedReadCloser) Close() error {
	return bc.close()
}

func (m *media) findHandler(mimetype string) (handler, error) {
	for _, h := range m.handlers {
		if h.Matches(mimetype) {
			return h, nil
		}
	}
	return nil, fmt.Errorf("Unsupported content type: %s", mimetype)
}

func (m *media) display(ctx *context.T, mimetype string, r io.ReadCloser) error {
	rc := &bufferedReadCloser{
		Reader: bufio.NewReader(r),
		close:  r.Close,
	}

	if mimetype == "" {
		mimetype = detectMimeType(rc.Reader)
	}

	// Find a suitable handler.
	h, err := m.findHandler(mimetype)
	if err != nil {
		return err
	}

	m.stopMu.Lock()
	if m.stop != nil {
		m.stop()
	}
	m.stop, err = h.Display(ctx, mimetype, r)
	m.stopMu.Unlock()
	return err
}

// DisplayURL will cause the server to display whatever media is at
// the given URL.  The server will rely on the ContentType response
// header it gets when fetching the url to decide how to display
// the media.
func (m *media) DisplayUrl(ctx *context.T, call rpc.ServerCall, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return m.display(ctx, resp.Header.Get("Content-Type"), resp.Body)
}

// The streamReader adapts a vanadium steam of byte slices to an io.ReadCloser.
type streamReader struct {
	last   []byte
	stream interface {
		Advance() bool
		Value() []byte
		Err() error
	}
}

func (s *streamReader) Read(p []byte) (int, error) {
	if len(s.last) == 0 {
		if !s.stream.Advance() {
			err := s.stream.Err()
			if err == nil {
				err = io.EOF
			}
			return 0, err
		}
		s.last = s.stream.Value()
	}
	if len(s.last) == 0 {
		return 0, nil
	}
	n := copy(p, s.last)
	s.last = s.last[n:]
	return n, nil
}

func (s *streamReader) Close() error {
	// Do nothing.
	return nil
}

// DisplayBytes will cause the server to display whatever media is
// sent in the stream.  In the case of audio or movie media, the
// media should be played while the data is streaming.  The mediaType
// can be used by the server to decide how to display the media.
func (m *media) DisplayBytes(ctx *context.T, call media_sharing.MediaSharingDisplayBytesServerCall, mediaType string) error {
	r := &streamReader{stream: call.RecvStream()}
	return m.display(ctx, mediaType, r)
}
