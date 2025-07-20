#!/bin/bash

# Installation script for Password Zen
# Usage: curl -sSL https://raw.githubusercontent.com/tmsankaram/password-zen/main/scripts/install/install.sh | bash

set -e

# Configuration
REPO="tmsankaram/password-zen"
BINARY_NAME="password-zen"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    case "$OS" in
        linux)
            OS="linux"
            ;;
        darwin)
            OS="darwin"
            ;;
        *)
            log_error "Unsupported operating system: $OS"
            exit 1
            ;;
    esac

    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            log_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac

    PLATFORM="${OS}-${ARCH}"
    log_info "Detected platform: $PLATFORM"
}

# Get latest release version
get_latest_version() {
    log_info "Fetching latest release information..."

    VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$VERSION" ]; then
        log_error "Failed to get latest version"
        exit 1
    fi

    log_info "Latest version: $VERSION"
}

# Download binary
download_binary() {
    BINARY_URL="https://github.com/$REPO/releases/download/$VERSION/${BINARY_NAME}-${PLATFORM}"
    TEMP_DIR=$(mktemp -d)
    TEMP_FILE="$TEMP_DIR/$BINARY_NAME"

    >&2 log_info "Downloading $BINARY_NAME from $BINARY_URL"

    if command -v curl >/dev/null 2>&1; then
        if ! curl -L -o "$TEMP_FILE" "$BINARY_URL"; then
            >&2 log_error "Failed to download binary"
            exit 1
        fi
    elif command -v wget >/dev/null 2>&1; then
        if ! wget -O "$TEMP_FILE" "$BINARY_URL"; then
            >&2 log_error "Failed to download binary"
            exit 1
        fi
    else
        >&2 log_error "Neither curl nor wget is available. Please install one of them."
        exit 1
    fi

    if [ ! -f "$TEMP_FILE" ]; then
        >&2 log_error "Failed to download binary - file not found"
        exit 1
    fi

    chmod +x "$TEMP_FILE"

    # Only output the temp file path to stdout (no log!)
    echo "$TEMP_FILE"
}


# Install binary
install_binary() {
    local temp_file="$1"
    local target="$INSTALL_DIR/$BINARY_NAME"

    log_info "Installing $BINARY_NAME to $target"

    # Check if we need sudo
    if [ ! -w "$INSTALL_DIR" ]; then
        if command -v sudo >/dev/null 2>&1; then
            sudo mv "$temp_file" "$target"
        else
            log_error "No write permission to $INSTALL_DIR and sudo not available"
            log_info "Please run the following command manually:"
            log_info "  mv $temp_file $target"
            exit 1
        fi
    else
        mv "$temp_file" "$target"
    fi

    log_success "$BINARY_NAME installed successfully!"
}

# Verify installation
verify_installation() {
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        local version_output=$($BINARY_NAME --version 2>&1)
        log_success "Installation verified: $version_output"
    else
        log_warning "Binary installed but not found in PATH"
        log_info "You may need to add $INSTALL_DIR to your PATH or restart your shell"
    fi
}

# Main installation process
main() {
    echo "🔐 Password Zen Installer"
    echo "========================"
    echo

    detect_platform
    get_latest_version

    temp_file=$(download_binary)
    install_binary "$temp_file"
    verify_installation

    echo
    log_success "Installation complete!"
    echo
    echo "📚 Get started:"
    echo "  $BINARY_NAME generate --help      # Generate passwords"
    echo "  $BINARY_NAME analyze --help       # Analyze passwords"
    echo
    echo "🔗 Documentation: https://github.com/$REPO"
    echo "🐛 Report issues: https://github.com/$REPO/issues"
}

# Check if running as root (not recommended for security tools)
if [ "$EUID" -eq 0 ]; then
    log_warning "Running as root is not recommended for security tools"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        log_info "Installation cancelled"
        exit 1
    fi
fi

main "$@"
