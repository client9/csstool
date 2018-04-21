#!/bin/sh
echo "Real publishing is done by travis-ci"
curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | bash

set -ex
echo "TAG:= $(git tag | tail -1)"
rm -rf ./dist
./bin/goreleaser --skip-publish --skip-validate
