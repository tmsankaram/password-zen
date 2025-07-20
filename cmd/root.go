/*
Copyright © 2025 Mahadeva Sankaram Telidevara <tmsankaram@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tmsankaram/password-zen/internal/version"
)

var rootCmd = &cobra.Command{
	Use:   "password-zen",
	Short: "A modern CLI tool for secure password generation and analysis",
	Long: `Password Zen is a command-line tool for generating secure passwords and analyzing password strength.

Features:
• Generate cryptographically secure passwords with customizable options
• Analyze password strength against security criteria
• Batch analysis of password files
• Colorful output with animations (can be disabled)
• Cross-platform support (Windows, Linux, macOS)

Use 'password-zen <command> --help' for detailed command information.`,
	Version: version.Short(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Version flag
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	// Set custom version template
	rootCmd.SetVersionTemplate(version.Info() + "\n")
}
