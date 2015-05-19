// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: media.vdl

package media_sharing

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// MediaSharingClientMethods is the client interface
// containing MediaSharing methods.
type MediaSharingClientMethods interface {
	// DisplayURL will cause the server to display whatever media is at
	// the given URL.  The server will rely on the ContentType response
	// header it gets when fetching the url to decide how to display
	// the media.
	DisplayUrl(ctx *context.T, url string, opts ...rpc.CallOpt) error
	// DisplayBytes will cause the server to display whatever media is
	// sent in the stream.  In the case of audio or movie media, the
	// media should be played while the data is streaming.  The mediaType
	// can be used by the server to decide how to display the media.
	DisplayBytes(ctx *context.T, mediaType string, opts ...rpc.CallOpt) (MediaSharingDisplayBytesClientCall, error)
}

// MediaSharingClientStub adds universal methods to MediaSharingClientMethods.
type MediaSharingClientStub interface {
	MediaSharingClientMethods
	rpc.UniversalServiceMethods
}

// MediaSharingClient returns a client stub for MediaSharing.
func MediaSharingClient(name string) MediaSharingClientStub {
	return implMediaSharingClientStub{name}
}

type implMediaSharingClientStub struct {
	name string
}

func (c implMediaSharingClientStub) DisplayUrl(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "DisplayUrl", []interface{}{i0}, nil, opts...)
	return
}

func (c implMediaSharingClientStub) DisplayBytes(ctx *context.T, i0 string, opts ...rpc.CallOpt) (ocall MediaSharingDisplayBytesClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "DisplayBytes", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implMediaSharingDisplayBytesClientCall{ClientCall: call}
	return
}

// MediaSharingDisplayBytesClientStream is the client stream for MediaSharing.DisplayBytes.
type MediaSharingDisplayBytesClientStream interface {
	// SendStream returns the send side of the MediaSharing.DisplayBytes client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item []byte) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// MediaSharingDisplayBytesClientCall represents the call returned from MediaSharing.DisplayBytes.
type MediaSharingDisplayBytesClientCall interface {
	MediaSharingDisplayBytesClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implMediaSharingDisplayBytesClientCall struct {
	rpc.ClientCall
}

func (c *implMediaSharingDisplayBytesClientCall) SendStream() interface {
	Send(item []byte) error
	Close() error
} {
	return implMediaSharingDisplayBytesClientCallSend{c}
}

type implMediaSharingDisplayBytesClientCallSend struct {
	c *implMediaSharingDisplayBytesClientCall
}

func (c implMediaSharingDisplayBytesClientCallSend) Send(item []byte) error {
	return c.c.Send(item)
}
func (c implMediaSharingDisplayBytesClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implMediaSharingDisplayBytesClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// MediaSharingServerMethods is the interface a server writer
// implements for MediaSharing.
type MediaSharingServerMethods interface {
	// DisplayURL will cause the server to display whatever media is at
	// the given URL.  The server will rely on the ContentType response
	// header it gets when fetching the url to decide how to display
	// the media.
	DisplayUrl(ctx *context.T, call rpc.ServerCall, url string) error
	// DisplayBytes will cause the server to display whatever media is
	// sent in the stream.  In the case of audio or movie media, the
	// media should be played while the data is streaming.  The mediaType
	// can be used by the server to decide how to display the media.
	DisplayBytes(ctx *context.T, call MediaSharingDisplayBytesServerCall, mediaType string) error
}

// MediaSharingServerStubMethods is the server interface containing
// MediaSharing methods, as expected by rpc.Server.
// The only difference between this interface and MediaSharingServerMethods
// is the streaming methods.
type MediaSharingServerStubMethods interface {
	// DisplayURL will cause the server to display whatever media is at
	// the given URL.  The server will rely on the ContentType response
	// header it gets when fetching the url to decide how to display
	// the media.
	DisplayUrl(ctx *context.T, call rpc.ServerCall, url string) error
	// DisplayBytes will cause the server to display whatever media is
	// sent in the stream.  In the case of audio or movie media, the
	// media should be played while the data is streaming.  The mediaType
	// can be used by the server to decide how to display the media.
	DisplayBytes(ctx *context.T, call *MediaSharingDisplayBytesServerCallStub, mediaType string) error
}

// MediaSharingServerStub adds universal methods to MediaSharingServerStubMethods.
type MediaSharingServerStub interface {
	MediaSharingServerStubMethods
	// Describe the MediaSharing interfaces.
	Describe__() []rpc.InterfaceDesc
}

// MediaSharingServer returns a server stub for MediaSharing.
// It converts an implementation of MediaSharingServerMethods into
// an object that may be used by rpc.Server.
func MediaSharingServer(impl MediaSharingServerMethods) MediaSharingServerStub {
	stub := implMediaSharingServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implMediaSharingServerStub struct {
	impl MediaSharingServerMethods
	gs   *rpc.GlobState
}

func (s implMediaSharingServerStub) DisplayUrl(ctx *context.T, call rpc.ServerCall, i0 string) error {
	return s.impl.DisplayUrl(ctx, call, i0)
}

func (s implMediaSharingServerStub) DisplayBytes(ctx *context.T, call *MediaSharingDisplayBytesServerCallStub, i0 string) error {
	return s.impl.DisplayBytes(ctx, call, i0)
}

func (s implMediaSharingServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implMediaSharingServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{MediaSharingDesc}
}

// MediaSharingDesc describes the MediaSharing interface.
var MediaSharingDesc rpc.InterfaceDesc = descMediaSharing

// descMediaSharing hides the desc to keep godoc clean.
var descMediaSharing = rpc.InterfaceDesc{
	Name:    "MediaSharing",
	PkgPath: "v.io/x/media_sharing",
	Methods: []rpc.MethodDesc{
		{
			Name: "DisplayUrl",
			Doc:  "// DisplayURL will cause the server to display whatever media is at\n// the given URL.  The server will rely on the ContentType response\n// header it gets when fetching the url to decide how to display\n// the media.",
			InArgs: []rpc.ArgDesc{
				{"url", ``}, // string
			},
		},
		{
			Name: "DisplayBytes",
			Doc:  "// DisplayBytes will cause the server to display whatever media is\n// sent in the stream.  In the case of audio or movie media, the \n// media should be played while the data is streaming.  The mediaType\n// can be used by the server to decide how to display the media.",
			InArgs: []rpc.ArgDesc{
				{"mediaType", ``}, // string
			},
		},
	},
}

// MediaSharingDisplayBytesServerStream is the server stream for MediaSharing.DisplayBytes.
type MediaSharingDisplayBytesServerStream interface {
	// RecvStream returns the receiver side of the MediaSharing.DisplayBytes server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// MediaSharingDisplayBytesServerCall represents the context passed to MediaSharing.DisplayBytes.
type MediaSharingDisplayBytesServerCall interface {
	rpc.ServerCall
	MediaSharingDisplayBytesServerStream
}

// MediaSharingDisplayBytesServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements MediaSharingDisplayBytesServerCall.
type MediaSharingDisplayBytesServerCallStub struct {
	rpc.StreamServerCall
	valRecv []byte
	errRecv error
}

// Init initializes MediaSharingDisplayBytesServerCallStub from rpc.StreamServerCall.
func (s *MediaSharingDisplayBytesServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the MediaSharing.DisplayBytes server stream.
func (s *MediaSharingDisplayBytesServerCallStub) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implMediaSharingDisplayBytesServerCallRecv{s}
}

type implMediaSharingDisplayBytesServerCallRecv struct {
	s *MediaSharingDisplayBytesServerCallStub
}

func (s implMediaSharingDisplayBytesServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implMediaSharingDisplayBytesServerCallRecv) Value() []byte {
	return s.s.valRecv
}
func (s implMediaSharingDisplayBytesServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}
