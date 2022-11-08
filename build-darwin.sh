#!/bin/bash

echo "=== Golden Point Compile Script - Linux ==="

# Section 1. Compile executables...
#
echo "Step 1. Setup depenencies..."
go get -v ./...

# Section 2. Generate assets...
#
echo "Step 2. Generate assets..."
go generate

# Section 3. Compile amd64 executable...
#
export GOOS="darwin"
export GOARCH="amd64"
export CGO_ENABLED="1"
go build -o golden-darwin-amd64 ./cmd/golden

# Section 4. Create ZIP archive ....
#
