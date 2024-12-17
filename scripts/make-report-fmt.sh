#!/bin/sh

SCRIPT_DIR="$(readlink -f $(dirname "$0"))"
SRC_DIR="$(dirname "${SCRIPT_DIR}")"

go fmt ${SRC_DIR}
