#!/bin/bash
# Copyright 2015 The Vanadium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


hostname="rpi2media"

set -x
set -e

vbecome -role=identity/role/vlab/admin device associate add vlab/devices/${hostname}/devmgr/device testuser1 dev.v.io/u/mattr@google.com

installation=$(vbecome -role=identity/role/vlab/admin device install vlab/devices/${hostname}/devmgr/apps applications/mediaserver)

vbecome --role=identity/role/vlab/admin device acl set -f ${installation} \
  dev.v.io/role/vlab/admin Admin,Debug,Read,Resolve,Write \
  dev.v.io/u/mattr@google.com Read

instance=$(device instantiate ${installation} ${hostname})
device run ${instance}
