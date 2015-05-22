// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io.v.x.media_sharing;

import android.os.AsyncTask;

import io.v.v23.OptionDefs;
import io.v.v23.Options;
import io.v.v23.context.VContext;
import io.v.v23.verror.VException;

/**
 * Background task to send a URL without blocking the UI thread.
 */
public class SendUrlTask extends AsyncTask<Void, Void, Void> {
    VContext vContext;
    String targetName;
    String url;

    public SendUrlTask(VContext vContext, String targetName, String url) {
        this.vContext = vContext;
        this.targetName = targetName;
        this.url = url;
    }

    @Override
    protected Void doInBackground(Void... params) {
        try {
            MediaSharingClient client = MediaSharingClientFactory.getMediaSharingClient(targetName);

            // TODO(bprosnitz) Remove this option when possible. It is allows the app to connect
            // without having the proper blessings.
            Options opts = new Options();
            opts.set(OptionDefs.SKIP_SERVER_ENDPOINT_AUTHORIZATION, true);

            client.displayUrl(vContext, url, opts);
            return null;
        } catch (VException e) {
            throw new RuntimeException(e);
        }
    }
}
