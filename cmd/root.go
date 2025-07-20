/*
Copyright © 2025 Mahadeva Sankaram Telidevara <tmsankaram@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "password-zen",
	Short: "A password generation tool",
	Long: `Password Zen is a command-line tool for generating secure passwords.
	It allows users to create strong, random passwords with various options for customization.
	Use 'password-zen --help' or 'password-zen <command> --help' for more details on usage and available commands.`,
	// Run:
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add any initialization code here if needed
	// For example, setting up flags or configuration
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of Password Zen")
	rootCmd.Flags().BoolP("help", "h", true, "Display help for Password Zen")
}
