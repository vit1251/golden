#!/bin/bash

echo "=== Golden Point Compile Script - Linux ==="

# Section 1. Generate assets...
#
echo "Step 1. Generate assets..."
go generate

# Section 2. Compile executables...
#
echo "Step 2. Compile executables..."
go get -v -u

echo "Step 2.1. Compile X86_64 executable..."
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="1"
go build -o golden-linux-amd64

#echo "Step 2.1. Compile X86 executable..."
#export GOOS="linux"
#export GOARCH="386"
#export CGO_ENABLED="1"
#go build -o golden-linux-386

# Section 3. Make ZIP portable distribution package...
#
echo "Step 3. Make ZIP portable distribution package..."
