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
	For more information, visit the project's GitHub repository.
	Use 'password-zen --help' or 'password-zen <command> --help' for more details on usage and available commands.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
