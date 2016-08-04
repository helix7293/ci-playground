#!/bin/bash

set -euxo pipefail

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go

go-bindata static

go test
