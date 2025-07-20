#!/bin/bash

# Quick release script for Password Zen
# Usage: ./scripts/release.sh v1.0.0

set -e

VERSION="$1"

if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Validate version format
if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    log_error "Invalid version format. Use semantic versioning (e.g., v1.0.0)"
    exit 1
fi

echo "🚀 Password Zen Release Script"
echo "=============================="
echo

log_info "Preparing release $VERSION"

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ] && [ "$CURRENT_BRANCH" != "master" ]; then
    log_error "Not on main/master branch. Current branch: $CURRENT_BRANCH"
    exit 1
fi

# Check for uncommitted changes
if [[ -n $(git status --porcelain) ]]; then
    log_error "Uncommitted changes found. Please commit or stash them first."
    git status --short
    exit 1
fi

# Pull latest changes
log_info "Pulling latest changes..."
git pull origin "$CURRENT_BRANCH"

# Run tests
log_info "Running tests..."
go test -v ./...
if [ $? -ne 0 ]; then
    log_error "Tests failed. Aborting release."
    exit 1
fi

# Check formatting
log_info "Checking code formatting..."
if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
    log_error "Code formatting issues found:"
    gofmt -s -l .
    exit 1
fi

# Run vet
log_info "Running go vet..."
go vet ./...
if [ $? -ne 0 ]; then
    log_error "go vet found issues. Aborting release."
    exit 1
fi

# Update version in version.go
VERSION_WITHOUT_V="${VERSION#v}"
log_info "Updating version to $VERSION_WITHOUT_V in internal/version/version.go"

# Create a backup
cp internal/version/version.go internal/version/version.go.backup

# Update version
sed -i.tmp "s/Version = \".*\"/Version = \"$VERSION_WITHOUT_V\"/" internal/version/version.go
rm internal/version/version.go.tmp

# Build and test
log_info "Testing build..."
go build -o password-zen-test .
./password-zen-test --version
rm password-zen-test

# Create commit for version bump
log_info "Creating version bump commit..."
git add internal/version/version.go
git commit -m "Bump version to $VERSION"

# Create and push tag
log_info "Creating and pushing tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION

Features:
- Secure password generation with cryptographic randomness
- Password strength analysis with detailed feedback
- Batch processing for multiple passwords
- Beautiful colored output with animations
- Cross-platform support (Windows, Linux, macOS)"

git push origin "$CURRENT_BRANCH"
git push origin "$VERSION"

log_success "Release $VERSION created successfully!"
echo
log_info "GitHub Actions will now:"
echo "  ✓ Run tests and security scans"
echo "  ✓ Build binaries for all platforms"
echo "  ✓ Create GitHub release"
echo "  ✓ Upload release artifacts"
echo
log_info "Monitor the release at: https://github.com/tmsankaram/password-zen/actions"
log_info "Release will be available at: https://github.com/tmsankaram/password-zen/releases/tag/$VERSION"
