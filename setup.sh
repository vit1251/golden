#!/bin/sh

set -x

# Step 0. Move executable
mv ./golden ./golden.bak

# Step 1. Download source code dependencies
go get -v

# Step 2. Compile Golden
go build -o golden .

# Step 3. Start reader
./golden service >golden_service.log 2>golden_service_err.log
