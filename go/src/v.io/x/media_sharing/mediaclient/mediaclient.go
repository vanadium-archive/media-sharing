// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"v.io/v23/context"
	"v.io/x/lib/cmdline"
	"v.io/x/media_sharing"
	"v.io/x/ref/lib/v23cmd"
	_ "v.io/x/ref/runtime/factories/static"
)

var stream *bool

func main() {
	stream = root.Flags.Bool("stream", true,
		"If true stream the data instead of sending the URL.")

	cmdline.Main(root)
}

var root = &cmdline.Command{
	Name:     os.Args[0],
	ArgsName: "<server> <url>",
	ArgsLong: ("<server> is the Vanadium name of a media server.  " +
		"<url> is an url to some content."),
	Runner: v23cmd.RunnerFunc(display),
	Short:  "Share media with a remote display.",
}

type streamWriter struct {
	buf    []byte
	stream interface {
		Send(item []byte) error
		Close() error
	}
}

func (w *streamWriter) Write(p []byte) (n int, err error) {
	return len(p), w.stream.Send(p)
}

func display(ctx *context.T, env *cmdline.Env, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Both a server and url must be specified.")
	}
	name, urlStr := args[0], args[1]
	client := media_sharing.MediaSharingClient(name)

	if *stream {
		parsed, err := url.Parse(urlStr)
		if err != nil {
			return err
		}

		var reader io.Reader
		var mimeType string

		if parsed.Scheme == "file" {
			file, err := os.Open(parsed.Path)
			if err != nil {
				return err
			}
			defer file.Close()
			reader = file
			mimeType = mime.TypeByExtension(filepath.Ext(parsed.Path))
		} else {
			resp, err := http.Get(urlStr)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			reader = resp.Body
			mimeType = resp.Header.Get("Content-Type")
		}

		call, err := client.DisplayBytes(ctx, mimeType)
		if err != nil {
			return fmt.Errorf("Failed to start RPC: %v", err)
		}
		sw := &streamWriter{stream: call.SendStream()}
		if _, err = io.Copy(sw, reader); err != nil {
			return fmt.Errorf("Failed copy stream: %v", err)
		}
		return call.Finish()
	} else {
		return client.DisplayUrl(ctx, urlStr)
	}
}
