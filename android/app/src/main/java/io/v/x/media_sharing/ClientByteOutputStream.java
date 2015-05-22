// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io.v.x.media_sharing;

import java.io.IOException;
import java.io.OutputStream;
import java.util.Arrays;

import io.v.v23.vdl.TypedClientStream;
import io.v.v23.verror.VException;

/**
 * Output stream that writes to a Vanadium byte stream.
 */
class ClientByteOutputStream extends OutputStream {
    private TypedClientStream<byte[], Void, Void> clientStream;
    public ClientByteOutputStream(TypedClientStream<byte[], Void, Void> clientStream) {
        this.clientStream = clientStream;
    }

    @Override
    public void write(int oneByte) throws IOException {
        this.write(new byte[]{(byte)oneByte});
    }

    @Override
    public void write(byte[] b) throws IOException {
        try {
            this.clientStream.send(b);
        } catch (VException e) {
            throw new IOException("Failed to write data to client stream", e);
        }
    }

    @Override
    public void write(byte[] b, int off, int len) throws IOException {
        this.write(Arrays.copyOfRange(b, off, off + len));
    }

    @Override
    public void close() {
        // stream.finish() should be called separately
    }

    @Override
    public void flush(){
        // ignore
    }
}
