#!/bin/sh

set -x

# Step 0. Move executable
rm -rf ./golden

# Step 2. Compile Golden
go build -o golden .
