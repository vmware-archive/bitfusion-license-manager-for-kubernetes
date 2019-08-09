# Copyright (c) 2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

default: all

all: web/static/Skeleton/index.html tcontainer fdcontainer

web/static/Skeleton/index.html:
	git submodule init
	git submodule update --recursive

# flexdirect container
fdcontainer:
	cd build/docker/flexdirect; ./build.sh
.PHONY: fdcontainer

# token container
tcontainer: cmd/token/token
	cd build/docker/token; ./build.sh
.PHONY: tcontainer

cmd/token/token: go.mod $(GOFILES)
	cd cmd/token; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a --installsuffix cgo token.go

cmd/token/token-darwin: go.mod $(GOFILES)
	cd cmd/token; GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -o token-darwin --installsuffix cgo token.go

go.mod:
	go mod init github.com/tdhite/bitfusion-k8s
	for m in $$(cat forcemodules); do go get "$$m"; done
	go get ./...

clean:
	rm -f cmd/token/token
	rm -f cmd/token/token-darwin
	rm -f build/docker/token/token
	rm -rf build/docker/token/html
.PHONY: clean

test:
	go test ./...
.PHONY: test
