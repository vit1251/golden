#!/bin/bash

echo "=== Golden Point Compile Script - Linux ==="

# Section 1. Setup depenencies...
#
echo "Step 1. Setup depenencies..."
go get -v -u

# Section 2. Generate assets...
#
echo "Step 2. Generate assets..."
go generate

# Section 3. Compile...
#
echo "Step 3. Compile amd64 executable..."
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="1"
go build -o golden-linux-amd64

# Section 4. Make ZIP portable distribution package...
#
#echo "Step 4. Make ZIP portable distribution package..."
