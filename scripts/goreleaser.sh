#!/bin/sh -e
# autorelease based on tag
if [ -n "$TRAVIS_TAG" ]; then
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | bash
	./bin/goreleaser
fi
