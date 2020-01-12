#!/bin/sh

set -x

# Step 1. Start reader
./golden service >golden_service.log 2>golden_service_err.log
