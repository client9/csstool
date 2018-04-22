#!/bin/sh
set -ex
PATH="./bin:$PATH"
go build ./...
go test .
./bin/gometalinter \
    --vendor \
    --deadline=60s \
    --disable-all \
    --enable=vet \
    --enable=golint \
    --enable=gofmt \
    --enable=goimports \
    --enable=ineffassign \
    ./...
go install ./css
