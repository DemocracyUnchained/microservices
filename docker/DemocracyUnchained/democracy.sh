#!/bin/bash

# Change to the app directory.
cd /go/src/democracy
# Get the module dependencies.
go get .
# Build the go code.
go build .

# Start the go code.
./democracy
