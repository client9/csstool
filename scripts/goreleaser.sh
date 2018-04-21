#!/bin/sh -e
# autorelease based on tag
test -n "$TRAVIS_TAG" && curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | bash
