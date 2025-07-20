# Password Zen 🔐

A modern, secure CLI tool for password generation and analysis.

[![Go Version](https://img.shields.io/badge/Go-1.24.4+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/badge/Release-v1.0.0-blue.svg)](https://github.com/tmsankaram/password-zen/releases)

## Features ✨

- 🔒 **Secure Generation**: Cryptographically secure random password generation
- 📊 **Password Analysis**: Analyze password strength with detailed feedback
- 📁 **Batch Processing**: Analyze multiple passwords from files
- 🎨 **Beautiful Output**: Colorful terminal output with animations
- ⚙️ **Customizable**: Extensive configuration options
- 🖥️ **Cross-Platform**: Works on Windows, Linux, and macOS

## Quick Start 🚀

### Generate a Password

```bash
# Generate a 12-character password (default)
password-zen generate

# Generate with symbols and specific length
password-zen generate --length 16 --include-symbols

# Exclude ambiguous characters
password-zen generate --exclude-ambiguous
```

### Analyze Password Strength

```bash
# Analyze a single password
password-zen analyze --password "mypassword123"

# Analyze passwords from a file
password-zen analyze --file passwords.txt

# Save analysis report
password-zen analyze --file passwords.txt --output report.txt
```

## Installation 📦

### Quick Install (Recommended)

**Linux/macOS - One-line install:**

```bash
curl -sSL https://raw.githubusercontent.com/tmsankaram/password-zen/main/scripts/install/install.sh | bash
```

**Windows - PowerShell (Run as Administrator):**

```powershell
# Download and run install script
iwr -useb https://raw.githubusercontent.com/tmsankaram/password-zen/main/scripts/install/install.ps1 | iex
```

### Manual Installation

**Windows:**

```powershell
# Download from GitHub releases
$version = "v1.0.0"
$arch = if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "amd64" }
Invoke-WebRequest -Uri "https://github.com/tmsankaram/password-zen/releases/download/$version/password-zen-windows-$arch.exe" -OutFile "password-zen.exe"

# Add to PATH or move to a directory in PATH
```

**Linux:**

```bash
# Download and install
VERSION="v1.0.0"
ARCH="amd64"  # or "arm64" for ARM systems
curl -L -o password-zen "https://github.com/tmsankaram/password-zen/releases/download/$VERSION/password-zen-linux-$ARCH"
chmod +x password-zen
sudo mv password-zen /usr/local/bin/

# Verify installation
password-zen --version
```

**macOS:**

```bash
# Download and install
VERSION="v1.0.0"
ARCH="amd64"  # or "arm64" for Apple Silicon
curl -L -o password-zen "https://github.com/tmsankaram/password-zen/releases/download/$VERSION/password-zen-darwin-$ARCH"
chmod +x password-zen
sudo mv password-zen /usr/local/bin/

# Verify installation
password-zen --version
```

### Build from Source

```bash
# Clone repository
git clone https://github.com/tmsankaram/password-zen.git
cd password-zen

# Build
go build -o password-zen .

# Install globally (optional)
go install .
```

## Usage Examples 📚

### Password Generation

```bash
# Basic generation
password-zen generate
# Output: aB3kL9mN2pQr

# Custom length with symbols
password-zen generate --length 20 --include-symbols
# Output: aB3!kL9@mN2#pQr$5sT^

# Exclude confusing characters
password-zen generate --exclude-ambiguous --length 16
# Output: aBcDeFgHjKmNpQrS

# Custom character set
password-zen generate --charset "ABC123!@#" --length 10
# Output: A3!B@1C#2A
```

### Password Analysis

```bash
# Analyze with default criteria
password-zen analyze --password "MyPassword123"
```

**Output:**

```
🔍 Analyzing password 1...
Password 1: STRONG ✓
  ✓ Length: 13 characters
  ✓ Contains digits
  ✓ Contains uppercase letters
  ✓ Contains lowercase letters

🎉 Excellent! All 1 passwords are strong!
```

```bash
# Custom analysis criteria
password-zen analyze --password "weak" --min-length 12 --require-symbols
```

**Output:**

```
Password 1: WEAK ✗
  ✗ Too short (4 < 12 characters)
  ✗ Missing digits
  ✗ Missing uppercase letters
  ✗ Missing special characters
  ✓ Contains lowercase letters

⚠️ Warning! Only 0/1 passwords meet criteria
```

### Batch Analysis

Create a file `passwords.txt`:

```
password123
MyStrongPassword456
weak
AnotherGoodPassword789
```

Analyze the file:

```bash
password-zen analyze --file passwords.txt --output report.txt
```

## Command Reference 📖

### Global Flags

- `--help, -h`: Show help information
- `--version, -v`: Show version information

### Generate Command

```bash
password-zen generate [flags]
```

**Flags:**

- `--length, -l`: Password length (default: 12)
- `--include-symbols, -s`: Include special characters
- `--include-digits, -d`: Include digits (default: true)
- `--exclude-ambiguous, -e`: Exclude ambiguous characters (il1Lo0O)
- `--charset, -c`: Custom character set

### Analyze Command

```bash
password-zen analyze [flags]
```

**Required (one of):**

- `--password, -p`: Single password to analyze
- `--file, -f`: File containing passwords (one per line)

**Optional:**

- `--output, -o`: Save report to file
- `--min-length, -m`: Minimum required length (default: 8)
- `--require-symbols, -s`: Require special characters
- `--require-digits, -d`: Require digits (default: true)
- `--require-uppercase, -u`: Require uppercase letters (default: true)
- `--require-lowercase, -l`: Require lowercase letters (default: true)
- `--no-color`: Disable colored output
- `--no-animation`: Disable animations

## Configuration 🔧

Password Zen supports configuration files and environment variables for setting defaults.

**Config file location:** `~/.password-zen/config.yaml`

**Example config:**

```yaml
defaults:
  length: 16
  include-symbols: true
  exclude-ambiguous: true
  min-length: 10

output:
  no-color: false
  no-animation: false
```

## Security Features 🛡️

- **Cryptographic Randomness**: Uses `crypto/rand` for secure password generation
- **Memory Safety**: Passwords are not logged or stored unnecessarily
- **No Network**: Completely offline operation
- **Uniform Distribution**: Ensures all characters have equal probability

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support 💬

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/tmsankaram/password-zen/issues)
- 💡 **Feature Requests**: [GitHub Issues](https://github.com/tmsankaram/password-zen/issues)
- 📧 **Email**: [tmsankaram@gmail.com](mailto:tmsankaram@gmail.com)

## Changelog 📝

### v1.0.0 (2025-01-21)

- Initial release
- Password generation with customizable options
- Password strength analysis
- Batch file processing
- Colorful terminal output with animations
- Cross-platform support

---

**Made with ❤️ by [Mahadeva Sankaram](https://github.com/tmsankaram)**
