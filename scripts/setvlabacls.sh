#!/bin/bash
# Copyright 2015 The Vanadium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x
set -e

vbecome --role=identity/role/vlab/admin namespace permissions set vlab - << EOF
{
 "Admin":{
   "In":["dev.v.io/role/vlab"]
 },
 "Resolve":{
   "In":["..."]
 }
}
EOF

vbecome --role=identity/role/vlab/admin namespace permissions set vlab/devices - << EOF
{
 "Admin":{
   "In":["dev.v.io/role/vlab"]
 },
 "Resolve":{
   "In":["..."]
 }
}
EOF

vbecome --role=identity/role/vlab/admin namespace permissions set vlab/tunnel - << EOF
{
 "Admin":{
   "In":["dev.v.io/role/vlab"]
 },
 "Resolve":{
   "In":["..."]
 }
}
EOF

vbecome --role=identity/role/vlab/admin namespace permissions set vlab/devices/rpi2media - << EOF
{
 "Admin":{
   "In":["dev.v.io/role/vlab"]
 },
 "Resolve":{
   "In":["..."]
 },
 "Read":{
   "In":["..."]
 }
}
EOF

vbecome --role=identity/role/vlab/admin namespace permissions set vlab/tunnel/rpi2media - << EOF
{
 "Admin":{
   "In":["dev.v.io/role/vlab"]
 },
 "Resolve":{
   "In":["..."]
 },
 "Read":{
   "In":["..."]
 }
}
EOF
