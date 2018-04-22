#!/bin/sh
set -ex
curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
curl -sSL https://install.goreleaser.com/github.com/alecthomas/gometalinter.sh | sh
curl -sSL https://install.goreleaser.com/github.com/client9/misspell | sh
