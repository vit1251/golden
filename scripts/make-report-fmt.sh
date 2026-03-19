#!/bin/sh

SCRIPT_DIR="$(readlink -f $(dirname "$0"))"
SRC_DIR="$(dirname "${SCRIPT_DIR}")"

gofmt -w ${SRC_DIR}
