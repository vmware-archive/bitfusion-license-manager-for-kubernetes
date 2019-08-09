#!/bin/bash
#
# Copyright (c) 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

CONTAINER=${CONTAINER:=tokenmgr}
VERSION=${VERSION:-1.0.0}

echo "Building ${CONTAINER}:${VERSION} . . ."

# Grab the latest build output
cp -a ../../../cmd/token/token .

# Copy the relevant skeleton code
mkdir -p html/skeleton
cp -a ../../../web/static/Skeleton/css html/skeleton/
cp -a ../../../web/static/Skeleton/images html/skeleton/

# Copy the template files
cp -a ../../../web/templates/tmpl html/

# Build and push the container
docker build --rm -t ${CONTAINER}:${VERSION} .

# Push the container if there happens to be a command line argument
if [ -n "$1" ]; then
    docker push ${CONTAINER}:${VERSION}
fi
