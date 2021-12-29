#!/bin/bash

echo "Step 1. Generate assets..."
go generate

echo "Step 2. Compile executables..."
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="1"
go build -o golden-linux-amd64
