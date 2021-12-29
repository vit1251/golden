#!/bin/bash

export GOOS="darwin"
export GOARCH="amd64"
export CGO_ENABLED="1"
go build -o golden-darwin-amd64
