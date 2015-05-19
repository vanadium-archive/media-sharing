// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/x/lib/cmdline"
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
	}

	server, err := v23.NewServer(ctx)
	if err != nil {
		return err
	}
	eps, err := server.Listen(v23.GetListenSpec(ctx))
	if err != nil {
		return err
	}
	if err := server.Serve(name, media_sharing.MediaSharingServer(&media{}), nil); err != nil {
		return err
	}
	fmt.Printf("Listening at: %s", eps[0].Name())

	<-signals.ShutdownOnSignals(ctx)
	return nil
}

type media struct{}

// DisplayURL will cause the server to display whatever media is at
// the given URL.  The server will rely on the ContentType response
// header it gets when fetching the url to decide how to display
// the media.
func (m *media) DisplayUrl(ctx *context.T, call rpc.ServerCall, url string) error {
	return nil
}

// DisplayBytes will cause the server to display whatever media is
// sent in the stream.  In the case of audio or movie media, the
// media should be played while the data is streaming.  The mediaType
// can be used by the server to decide how to display the media.
func (m *media) DisplayBytes(ctx *context.T, call media_sharing.MediaSharingDisplayBytesServerCall, mediaType string) error {
	return nil
}
