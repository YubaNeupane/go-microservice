#!/bin/bash

# Set environment variables for cross-compilation
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# Build the Go binary
go build -o build/api-gateway ./services/api-gateway

# Print a success message
echo "Build complete."
