// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var EE = require('events').EventEmitter;
var format = require('format');
var inherits = require('inherits');

module.exports = FileEmitter;

function FileEmitter(input) {
  EE.call(this);

  // TODO(nlacasse): Consider making these an argument to FileEmitter;
  this._allowedTypes = ['audio', 'image', 'video'];
  this._maxSize = 10 * 1000 * 1000; // 10MB

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

  // TODO(nlacasse): Consider removing the file size limit.
  if (file.size > this._maxSize) {
    this.emit('error', format('File too large.'));
    return;
  }

  var reader = new FileReader();

  var self = this;
  reader.addEventListener('error', function(ev) {
    self.emit('error', ev);
  });

  reader.addEventListener('load', function() {
    self.emit('file', {
      name: file.name,
      type: file.type,
      size: file.size,
      bytes: new Uint8Array(reader.result)
    });
  });

  reader.readAsArrayBuffer(file);
};
