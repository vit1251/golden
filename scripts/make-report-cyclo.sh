#!/bin/sh

go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

~/go/bin/gocyclo ..

