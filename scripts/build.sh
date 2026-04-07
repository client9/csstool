#!/bin/sh
set -ex
PATH="./bin:$PATH"
go build ./...
go test .
go install ./css
