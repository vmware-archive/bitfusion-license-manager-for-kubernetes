# Copyright 2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#
from ubuntu:16.04
maintainer Tom Hite <thite@vmware.com>

RUN apt-get update \
  && apt-get install -y libjson-c2 \
    libjsoncpp1 \
    librdmacm1 \
    libssl-dev \
    libibverbs1 \
    libnuma1 \
    libcapstone3 \
    libnl-3-200 \
    libnl-route-3-200 \
    curl \
    jq \
    wget

RUN wget -O installfd getfd.bitfusion.io && bash installfd -- -m binaries

COPY start /start

# We leave no default command for the container. To get help, run /start as:
#
#   CMD ["/start", "help"]
#
#   Then check logs to see the output of the help text.
# 
# Running this image follows this protocol:
#
#   CMD ["/start", "daemonset"]
#   CMD ["/start", "run"]
#
# or someting of the form:
#   CMD ["/bin/bash"]
#
