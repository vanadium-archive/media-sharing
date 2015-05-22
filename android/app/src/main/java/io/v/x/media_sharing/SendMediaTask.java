// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io.v.x.media_sharing;

import android.app.Activity;
import android.net.Uri;
import android.os.AsyncTask;

import org.apache.commons.io.IOUtils;

import java.io.IOException;
import java.io.InputStream;

import io.v.v23.OptionDefs;
import io.v.v23.Options;
import io.v.v23.context.VContext;
import io.v.v23.vdl.TypedClientStream;
import io.v.v23.verror.VException;

/**
 * Background task to stream media without blocking the UI thread.
 */
public class SendMediaTask extends AsyncTask<Void, Void, Void> {
    Activity activity;
    VContext vContext;
    String targetName;
    Uri uri;
    String mimeType;

    public SendMediaTask(Activity activity, VContext vContext, String targetName, Uri uri, String mimeType) {
        this.activity = activity;
        this.vContext = vContext;
        this.targetName = targetName;
        this.uri = uri;
        this.mimeType = mimeType;
    }

    @Override
    protected Void doInBackground(Void... params) {
        try {
            InputStream is = activity.getContentResolver().openInputStream(uri);

            MediaSharingClient client = MediaSharingClientFactory.getMediaSharingClient(targetName);
            
            // TODO(bprosnitz) Remove this option when possible. It is allows the app to connect
            // without having the proper blessings.
            Options opts = new Options();
            opts.set(OptionDefs.SKIP_SERVER_ENDPOINT_AUTHORIZATION, true);

            String mimeType = activity.getIntent().getType();
            TypedClientStream<byte[], Void, Void> stream = client.displayBytes(vContext, mimeType, opts);

            ClientByteOutputStream os = new ClientByteOutputStream(stream);
            IOUtils.copy(is, os);
            stream.finish();
            return null;
        } catch (IOException e) {
            throw new RuntimeException(e);
        } catch (VException e) {
            throw new RuntimeException(e);
        }
    }
}
