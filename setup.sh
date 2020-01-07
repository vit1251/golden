#!/bin/sh

set -x

# Step 0. Move executable
mv ./golden ./golden.bak

# Step 1. Download source code dependencies
go get -v

# Step 2. Compile Golden
go build -o golden .

# Step 3. Start reader
./golden mailer >golden_mailer.log 2>golden_mailer_err.log
./golden toss >golden_toss.log 2>golden_toss_err.log
./golden reader >golden_reader.log 2>golden_reader_err.log
