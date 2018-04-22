#!/bin/sh
set -ex
curl -sSL https://install.goreleaser.com/github.com/alecthomas/gometalinter.sh | sh
curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
