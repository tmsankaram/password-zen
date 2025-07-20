package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password",
	Long: `Generate a secure password with customizable options.
You can specify the length of the password and whether to include digits and symbols.`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Define flags for the generate command
	generateCmd.Flags().IntP("length", "l", 12, "Length of the password (default is 12)")
	generateCmd.Flags().BoolP("digits", "d", true, "Include digits in the password (default is true)")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password (default is false)")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSymbols, _ := cmd.Flags().GetBool("symbols")

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if isDigits {
		charset += "0123456789"
	}
	if isSymbols {
		charset += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	cmd.Println("Generated Password:", string(password))
}
