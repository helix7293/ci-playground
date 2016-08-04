#!/bin/bash

set -euxo pipefail

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

mkdir -p tmp 
buildkite-agent artifact download ci-playground tmp/
sudo /opt/ci-playground/update.sh tmp/ci-playground
