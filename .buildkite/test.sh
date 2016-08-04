#!/bin/bash

set -euxo pipefail

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

go get github.com/jteeuwen/go-bindata/go-bindata
go-bindata static

go test
