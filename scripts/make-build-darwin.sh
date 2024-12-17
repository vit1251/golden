#!/bin/bash

echo "=== Golden Point Compile Script ==="

# Step 0. Prepare environemnt
export GOOS="darwin"
export GOARCH="amd64"
export CGO_ENABLED="1"
#
echo "OS: ${GOOS}"
echo "Arch: ${GOARCH}"

# Step 1. Get Go modules
#
echo "==> Step 1. Get Go modules..."
go get -v ../...

# Step 2. Generate assets
#
echo "==> Step 2. Generate assets..."
go generate

# Step 3. Run unittest
#
echo "==> Step 3. Run unittest..."
go test ../...

# Step 4. Compile executable
#
echo "==> Step 4. Compile executable..."
go build -o ../golden-${GOOS}-${GOARCH} ../cmd/golden
