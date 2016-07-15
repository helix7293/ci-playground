#!/bin/bash

set -euxo pipefail

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go

mkdir -p tmp 
buildkite-agent artifact download ci-playground tmp/
sudo /opt/ci-playground/update.sh tmp/ci-playground
