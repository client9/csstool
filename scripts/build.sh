#!/bin/sh
set -ex

go build ./...
go install ./cmd/cssformat ./cmd/csscut ./cmd/csscount
gometalinter \
    --vendor \
    --deadline=60s \
    --disable-all \
    --enable=vet \
    --enable=golint \
    --enable=gofmt \
    --enable=goimports \
    --enable=gosimple \
    --enable=staticcheck \
    --enable=ineffassign \
    --exclude=/usr/local/go/src/net/lookup_unix.go \
    ./...
go test .
