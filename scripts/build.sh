#!/bin/bash

# build.sh - Cross-platform build script for Password Zen

set -e

# Configuration
APP_NAME="password-zen"
VERSION="1.0.0"
BUILD_DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS="-w -s -X github.com/tmsankaram/password-zen/internal/version.Version=${VERSION} -X github.com/tmsankaram/password-zen/internal/version.BuildDate=${BUILD_DATE} -X github.com/tmsankaram/password-zen/internal/version.GitCommit=${GIT_COMMIT}"

# Create dist directory
mkdir -p dist

echo "Building Password Zen v${VERSION}..."
echo "Build Date: ${BUILD_DATE}"
echo "Git Commit: ${GIT_COMMIT}"
echo ""

# Build for different platforms
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o "dist/${APP_NAME}-windows-amd64.exe" .

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o "dist/${APP_NAME}-linux-amd64" .

echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o "dist/${APP_NAME}-linux-arm64" .

echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o "dist/${APP_NAME}-darwin-amd64" .

echo "Building for macOS (arm64 - Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o "dist/${APP_NAME}-darwin-arm64" .

echo ""
echo "Build completed! Binaries are in the 'dist' directory:"
ls -la dist/

# Create checksums
echo ""
echo "Generating checksums..."
cd dist
sha256sum * > checksums.txt
cd ..

echo "Done! Release artifacts:"
echo "- Binaries: dist/"
echo "- Checksums: dist/checksums.txt"
