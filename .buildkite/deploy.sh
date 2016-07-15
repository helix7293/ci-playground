#!/bin/bash

set -euxo pipefail

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go

buildkite-agent artifact download ci-playground ./
sudo /opt/ci-playground/update.sh ci-playground
