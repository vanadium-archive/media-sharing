#!/bin/bash
# Copyright 2015 The Vanadium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -x
set -e

JIRI_PROFILE=arm jiri go install v.io/x/media_sharing/...
mkdir -p $JIRI_ROOT/release/go/bin/linux_arm
cp $JIRI_ROOT/release/projects/media-sharing/go/bin/linux_arm/* $JIRI_ROOT/release/go/bin/linux_arm
vbecome --role=identity/role/vprod/publisher device publish --goos=linux --goarch=arm mediaserver

#installation=$(device ls --installation-state=Active vlab/devices/rpi2media/devmgr/apps/mediaserver/* | tail -n 1)
installation=$(namespace glob vlab/devices/rpi2media/devmgr/apps/mediaserver/* | tail -n 1)
instance=$(namespace glob $installation/* | tail -n 1)

device update -parallelism=BYKIND $installation
device update -parallelism=BYKIND $instance
