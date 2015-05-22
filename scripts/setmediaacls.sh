#!/bin/bash
# Copyright 2015 The Vanadium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x
set -e

namespace permissions set users/mattr@google.com - << EOF
{
 "Admin":{
   "In":["dev.v.io/u/mattr@google.com"]
 },
 "Resolve":{
   "In":["..."]
 },
 "Read":{
   "In":["..."]
 }
}
EOF

namespace permissions set users/mattr@google.com/media2 - << EOF
{
 "Admin":{
   "In":["dev.v.io/u/mattr@google.com"]
 },
 "Resolve":{
   "In":["..."]
 },
 "Read":{
   "In":["..."]
 }
}
EOF
