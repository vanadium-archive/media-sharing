// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var EE = require('events').EventEmitter;
var format = require('format');
var inherits = require('inherits');
var ReadableBlobStream = require('readable-blob-stream');

module.exports = FileEmitter;

function FileEmitter(input) {
  EE.call(this);

  // TODO(nlacasse): Consider making these an argument to FileEmitter;
  this._allowedTypes = ['audio', 'image', 'video'];

  input.addEventListener('change', this._onFileChange.bind(this), false);
}

inherits(FileEmitter, EE);

FileEmitter.prototype._isAllowedType = function(type) {
  for (var i = 0; i < this._allowedTypes.length; i++) {
    if (type.indexOf(this._allowedTypes[i]) === 0) {
      return true;
    }
  }
  return false;
};

FileEmitter.prototype._onFileChange = function(ev) {
  var files = ev.target.files;
  if (files.length === 0) {
    this.emit('error', 'No files selected.');
    return;
  }

  // TODO(nlacasse): Consider handling multiple files?  For now we just take
  // the first.  Perhaps if multiple files are selected, we send each to a
  // different media server?
  var file = files[0];

  if (!this._isAllowedType(file.type)) {
    this.emit('error', format('Filetype %s is not supported.', file.type));
    return;
  }

  this.emit('file', {
    name: file.name,
    size: file.size,
    type: file.type,
    stream: new ReadableBlobStream(file)
  });
};
