#!/bin/sh

go install github.com/gordonklaus/ineffassign@latest

~/go/bin/ineffassign ../...

