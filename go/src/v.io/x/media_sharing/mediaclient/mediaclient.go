// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"v.io/v23/context"
	"v.io/x/lib/cmdline"
	"v.io/x/media_sharing"
	"v.io/x/ref/lib/v23cmd"
	_ "v.io/x/ref/runtime/factories/static"
)

func main() {
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

func display(ctx *context.T, env *cmdline.Env, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Both a server and url must be specified.")
	}
	name, url := args[0], args[1]
	client := media_sharing.MediaSharingClient(name)

	return client.DisplayUrl(ctx, url)
}
