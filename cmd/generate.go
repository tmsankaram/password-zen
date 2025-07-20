/*
Copyright © 2025 Mahadeva Sankaram
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password",
	Long: `Generate a secure password with customizable options.
You can specify the length, character sets, and other parameters to create a password that meets your security needs.
For more details, use 'password-zen generate --help'.`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 12, "Length of the generated password")
	generateCmd.Flags().BoolP("include-symbols", "s", false, "Include special characters like !@#$%^&*()")
	generateCmd.Flags().BoolP("include-digits", "d", true, "Include digits defaults to true")
	generateCmd.Flags().BoolP("exclude-ambiguous", "e", false, "Exclude ambiguous characters like il1Lo0O")
	generateCmd.Flags().StringP("charset", "c", "", "Custom character set to use for password generation. If not specified, defaults to alphanumeric characters with optional symbols and digits.")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	includeSymbols, _ := cmd.Flags().GetBool("include-symbols")
	includeDigits, _ := cmd.Flags().GetBool("include-digits")
	excludeAmbiguous, _ := cmd.Flags().GetBool("exclude-ambiguous")
	customCharset, _ := cmd.Flags().GetString("charset")

	// Validate input
	if length <= 0 {
		cmd.PrintErr("Error: Password length must be greater than 0\n")
		return
	} else if length > 128 {
		cmd.PrintErr("Error: Password length must not exceed 128 characters\n")
		return
	}

	var charset string
	if customCharset != "" {
		charset = customCharset
	} else {
		charset = buildCharset(includeDigits, includeSymbols, excludeAmbiguous)
	}

	if len(charset) == 0 {
		cmd.PrintErr("Error: No valid characters available for password generation\n")
		return
	}

	password, err := generateSecurePassword(length, charset)
	if err != nil {
		cmd.PrintErr(fmt.Sprintf("Error generating password: %v\n", err))
		return
	}

	fmt.Println(password)
}

func buildCharset(includeDigits, includeSymbols, excludeAmbiguous bool) string {
	// Start with base letters
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Add digits if requested
	if includeDigits {
		charset += "0123456789"
	}

	// Add symbols if requested
	if includeSymbols {
		charset += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}

	// Remove ambiguous characters if requested
	if excludeAmbiguous {
		ambiguous := "il1Lo0O"
		for _, char := range ambiguous {
			charset = strings.ReplaceAll(charset, string(char), "")
		}
	}

	return charset
}

func generateSecurePassword(length int, charset string) (string, error) {
	if length <= 0 || len(charset) == 0 {
		return "", fmt.Errorf("invalid parameters")
	}
	// use crypto/rand for secure random generation
	result := make([]byte, length)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error generating random index: %v", err)
		}
		result[i] = charset[index.Int64()]
	}
	return string(result), nil
}
