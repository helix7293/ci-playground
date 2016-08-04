#!/bin/bash

set -euxo pipefail

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

go build
