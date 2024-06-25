#!/bin/bash

# Build for Windows (x86_64)
echo "Building for Windows (x86_64)..."
GOOS=windows GOARCH=amd64 go build -o vet-clinic-windows.exe

# Build for Linux (x86_64)
echo "Building for Linux (x86_64)..."
GOOS=linux GOARCH=amd64 go build -o vet-clinic-linux

# Build for macOS (x86_64)
echo "Building for macOS (x86_64)..."
GOOS=darwin GOARCH=amd64 go build -o vet-clinic-macos

# Build for macOS (ARM64, M1 chip)
echo "Building for macOS (ARM64)..."
GOOS=darwin GOARCH=arm64 go build -o vet-clinic-macos-arm64

# Build for Linux (ARM)
echo "Building for Linux (ARM)..."
GOOS=linux GOARCH=arm go build -o vet-clinic-linux-arm

echo "Build process completed!"
