#!/bin/sh

set +x

# Step 1. Download source code dependencies
go get -v

# Step 2. Compile Golden
go build -o golden .

# Step 3. Start reader
./golden reader
