#!/bin/bash
set -e

ROOT="$(dirname "$(readlink -f "$0")")/.."

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"
OUTPUT="golden${GOOS:+-${GOOS}-${GOARCH}}"

echo "Build: ${GOOS}/${GOARCH} -> ${OUTPUT}"

go mod download
go test ./...
go build -o "${OUTPUT}" ./cmd/golden
