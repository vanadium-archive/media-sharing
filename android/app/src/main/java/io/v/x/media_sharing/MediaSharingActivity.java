// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io.v.x.media_sharing;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.net.Uri;
import android.support.v7.app.ActionBarActivity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import io.v.v23.V;
import io.v.v23.context.VContext;


public class MediaSharingActivity extends ActionBarActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_media_sharing);

        final VContext vContext = V.init();

        final EditText namebox = (EditText)findViewById(R.id.share_name);
        final SharedPreferences pref = this.getPreferences(Context.MODE_PRIVATE);
        String lastName = pref.getString(getString(R.string.preference_last_name), "");
        namebox.setText(lastName);

        final Button button = (Button)findViewById(R.id.share_button);
        button.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String name = namebox.getText().toString();
                SharedPreferences.Editor editor = pref.edit();
                editor.putString(getString(R.string.preference_last_name), name);
                editor.commit();
                onButtonClick(vContext, name);
                MediaSharingActivity.this.finish();
            }
        });
    }

    protected void onButtonClick(final VContext vContext, final String targetName) {
        String action = getIntent().getAction();
        String type = getIntent().getType();
        if (Intent.ACTION_SEND.equals(action) && type != null) {
            if (type.startsWith("image/") || type.startsWith("video/") || type.startsWith("audio/")) {
                Uri uri = (Uri)getIntent().getExtras().get(Intent.EXTRA_STREAM);
                String mimeType = getIntent().getType();

                new SendMediaTask(MediaSharingActivity.this, vContext, targetName, uri, mimeType).execute();
            } else if ("text/plain".equals(type)) {
                String url = (String)getIntent().getExtras().get(Intent.EXTRA_TEXT);
                new SendUrlTask(this, vContext, targetName, url).execute();
            }
        }
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_media_sharing, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }
}
