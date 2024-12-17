#!/bin/bash

SCRIPT_DIR="$(readlink -f $(dirname "$0"))"
SRC_DIR="$(dirname "${SCRIPT_DIR}")"

echo "=== Golden Point Compile Script ==="

# Step 0. Prepare environemnt
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="1"
#
echo " Src: ${SRC_DIR}"
echo "  OS: ${GOOS}"
echo "Arch: ${GOARCH}"

# Step 1. Get Go modules
#
echo "==> Step 1. Get Go modules..."
go get -v ${SRC_DIR}/...

# Step 2. Generate assets
#
echo "==> Step 2. Generate assets..."
go generate

# Step 3. Run unittest
#
echo "==> Step 3. Run unittest..."
go test ${SRC_DIR}/...

