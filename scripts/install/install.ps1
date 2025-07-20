@echo off
REM Installation script for Password Zen on Windows
REM Usage: Run this script as Administrator or ensure the target directory is writable

setlocal EnableDelayedExpansion

REM Configuration
set "REPO=tmsankaram/password-zen"
set "BINARY_NAME=password-zen"
set "INSTALL_DIR=%ProgramFiles%\PasswordZen"

echo 🔐 Password Zen Windows Installer
echo =================================
echo.

REM Detect architecture
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
    set "ARCH=amd64"
) else if "%PROCESSOR_ARCHITECTURE%"=="ARM64" (
    set "ARCH=arm64"
) else (
    echo [ERROR] Unsupported architecture: %PROCESSOR_ARCHITECTURE%
    pause
    exit /b 1
)

echo [INFO] Detected architecture: %ARCH%

REM Create installation directory
if not exist "%INSTALL_DIR%" (
    echo [INFO] Creating installation directory: %INSTALL_DIR%
    mkdir "%INSTALL_DIR%" 2>NUL
    if errorlevel 1 (
        echo [ERROR] Failed to create installation directory. Please run as Administrator.
        pause
        exit /b 1
    )
)

REM Get latest version (simplified - you might want to use PowerShell for better API handling)
echo [INFO] Please enter the version to install (e.g., v1.0.0):
set /p VERSION="Version: "

if "%VERSION%"=="" (
    set "VERSION=v1.0.0"
    echo [INFO] Using default version: %VERSION%
)

REM Download URL
set "BINARY_URL=https://github.com/%REPO%/releases/download/%VERSION%/%BINARY_NAME%-windows-%ARCH%.exe"
set "TARGET_FILE=%INSTALL_DIR%\%BINARY_NAME%.exe"

echo [INFO] Downloading from: %BINARY_URL%
echo [INFO] Installing to: %TARGET_FILE%

REM Download using PowerShell
powershell -Command "& {try { Invoke-WebRequest -Uri '%BINARY_URL%' -OutFile '%TARGET_FILE%' -UseBasicParsing } catch { Write-Host '[ERROR] Download failed:' $_.Exception.Message; exit 1 }}"

if errorlevel 1 (
    echo [ERROR] Download failed
    pause
    exit /b 1
)

REM Verify download
if not exist "%TARGET_FILE%" (
    echo [ERROR] Download verification failed
    pause
    exit /b 1
)

echo [SUCCESS] Download complete!

REM Add to PATH (requires PowerShell)
echo [INFO] Adding to system PATH...
powershell -Command "& {$env:Path = [Environment]::GetEnvironmentVariable('Path','Machine'); if ($env:Path -notlike '*%INSTALL_DIR%*') { [Environment]::SetEnvironmentVariable('Path', $env:Path + ';%INSTALL_DIR%', 'Machine'); Write-Host '[SUCCESS] Added to PATH' } else { Write-Host '[INFO] Already in PATH' }}"

echo.
echo [SUCCESS] Installation complete!
echo.
echo 📚 Get started:
echo   %BINARY_NAME% generate --help      # Generate passwords
echo   %BINARY_NAME% analyze --help       # Analyze passwords
echo.
echo 🔗 Documentation: https://github.com/%REPO%
echo 🐛 Report issues: https://github.com/%REPO%/issues
echo.
echo NOTE: You may need to restart your command prompt for PATH changes to take effect.
echo.
pause
