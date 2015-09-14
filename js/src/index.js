// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var domready = require('domready');
var format = require('format');
var vanadium = require('vanadium');

var FileEmitter = require('./file-emitter');

domready(onDomReady);

// Entry point of the app.
function onDomReady() {
  vanadium.init(function(err, rt) {
    if (err) {
      appendStatus('ERROR: ' + err);
      return;
    }

    // Connect the file input to a FileEmitter.
    var fileInput = document.getElementById('file-input');
    var fileEmitter = new FileEmitter(fileInput);

    fileEmitter.on('error', function(text) {
      appendStatus('ERROR: ' + text);
    });

    fileEmitter.on('file', function(file) {
      appendStatus(format('Sending file %s of type %s and %d bytes.',
                          file.name, file.type, file.size));
      sendToDisplay(rt, file, function(err) {
        if (err) {
          appendStatus('ERROR: ' + err);
          return;
        }
        appendStatus('Success.');
      });
    });
  });
}

// Append text to the status div.
function appendStatus(text) {
  var status = document.getElementById('status');
  status.innerHTML += text + '<br>';
}

function sendToDisplay(rt, file, cb) {
  var ctx = rt.getContext();
  var client = rt.getClient();
  client.bindTo(ctx, 'users/mattr@google.com/media2', function(err, s) {
    if (err) {
      return cb(err);
    }

    var promise = s.displayBytes(ctx, file.type);
    promise.catch(cb);
    promise.stream.on('error', cb);
    promise.stream.on('finish', cb);

    file.stream.pipe(promise.stream);
  });
}
