#!/bin/sh
set -ex
curl -sfSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
curl -sfSL https://install.goreleaser.com/github.com/alecthomas/gometalinter.sh | sh
curl -sfSL https://install.goreleaser.com/github.com/client9/misspell.sh | sh
