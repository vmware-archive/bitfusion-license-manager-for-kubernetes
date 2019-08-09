#!/bin/bash
#
# Copyright 2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

# setup the version to build
CONTAINER=${CONTAINER:=flexdirect}
VERSION=${VERSION:=1.0.0}

echo "Building ${CONTAINER}:${VERSION} . . ."

# build the image
docker build --rm --tag ${CONTAINER}:${VERSION} .


# Push the container if there happens to be a command line argument
if [ -n "$1" ]; then
    docker push ${CONTAINER}:${VERSION}
fi
