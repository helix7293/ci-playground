#!/bin/bash

set -euxo pipefail

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go

go get github.com/jteeuwen/go-bindata/go-bindata
go-bindata static

go test
