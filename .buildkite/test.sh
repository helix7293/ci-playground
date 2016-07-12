#!/bin/bash

set -euxo pipefail

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go

env

go test
