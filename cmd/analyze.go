/*
Copyright © 2025 Mahadeva Sankaram
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze the passwords",
	Long: `Analyze the passwords to ensure they meet security standards and requirements.
This command will check various aspects of the passwords such as length, character diversity.
For example, it can verify if the passwords contain required character types (uppercase, lowercase, digits, symbols) and if they meet the specified length criteria.
You can also use this command to generate reports on password strength and compliance.
For more information on the analysis results, you can use the --output flag to specify a file for the report.`,
	Run: analyzePassword,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringP("password", "p", "", "Password to analyze")
	analyzeCmd.Flags().StringP("file", "f", "", "Text file containing passwords to analyze")

	// Mark that one of password or file is required
	analyzeCmd.MarkFlagsMutuallyExclusive("password", "file")
	analyzeCmd.MarkFlagsOneRequired("password", "file")

	// Optional flags for analysis criteria
	analyzeCmd.Flags().StringP("output", "o", "", "Output file for the analysis report")
	analyzeCmd.Flags().IntP("min-length", "m", 8, "Minimum length for passwords")
	analyzeCmd.Flags().BoolP("require-symbols", "s", false, "Require passwords to contain special characters")
	analyzeCmd.Flags().BoolP("require-digits", "d", true, "Require passwords to contain digits")
	analyzeCmd.Flags().BoolP("require-uppercase", "u", true, "Require passwords to contain uppercase letters")
	analyzeCmd.Flags().BoolP("require-lowercase", "l", true, "Require passwords to contain lowercase letters")
	analyzeCmd.Flags().BoolP("no-color", "", false, "Disable colored output")
	analyzeCmd.Flags().BoolP("no-animation", "", false, "Disable animations")
}

// Color definitions for consistent output
var (
	greenCheck = color.New(color.FgGreen, color.Bold).SprintFunc()
	redCross   = color.New(color.FgRed, color.Bold).SprintFunc()
	greenText  = color.New(color.FgGreen).SprintFunc()
	redText    = color.New(color.FgRed).SprintFunc()
	yellowText = color.New(color.FgYellow).SprintFunc()
	cyanText   = color.New(color.FgCyan, color.Bold).SprintFunc()
)

// Simple animation for analysis
func animateAnalysis(passwordNum int) {
	dots := []string{".", "..", "..."}
	for i := 0; i < 3; i++ {
		fmt.Printf("\r%s Analyzing password %d%s   ", cyanText("🔍"), passwordNum, dots[i])
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Print("\r" + strings.Repeat(" ", 50) + "\r") // Clear the line
}

func analyzePassword(cmd *cobra.Command, args []string) {
	filepath, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	minLength, _ := cmd.Flags().GetInt("min-length")
	requireSymbols, _ := cmd.Flags().GetBool("require-symbols")
	requireDigits, _ := cmd.Flags().GetBool("require-digits")
	requireUppercase, _ := cmd.Flags().GetBool("require-uppercase")
	requireLowercase, _ := cmd.Flags().GetBool("require-lowercase")
	noColor, _ := cmd.Flags().GetBool("no-color")
	noAnimation, _ := cmd.Flags().GetBool("no-animation")

	// Disable color if requested
	if noColor {
		color.NoColor = true
	}

	var passwords []string
	// check if the file exists, is text file and is readable
	if filepath != "" {
		if err := checkFileExists(filepath); err != nil {
			cmd.PrintErrf("Error reading file: %v\n", err)
			return
		} else {
			cmd.Printf("Analyzing passwords from file: %s\n", filepath)
			passwords = analyzeFile(filepath)
		}
	}
	if len(passwords) == 0 {
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			cmd.PrintErr("Error: No password provided for analysis. Use --password or --file to specify a password or file.\n")
			return
		}
		passwords = []string{password}
	}

	if len(passwords) == 0 {
		cmd.PrintErr("Error: No password provided for analysis. Use --password or --file to specify a password or file.\n")
		return
	}

	var allResults []string
	passCount := 0

	for i, password := range passwords {
		// Show animated analysis
		if !noAnimation {
			animateAnalysis(i + 1)
		} else {
			fmt.Printf("Analyzing password %d...\n", i+1)
		}

		// Reset analysis for each password
		var currentAnalysis []string
		passed := true

		if len(password) < minLength {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Too short (%d < %d characters)", redCross("✗"), len(password), minLength))
			passed = false
		} else {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Length: %d characters", greenCheck("✓"), len(password)))
		}

		if requireSymbols && !containsSymbol(password) {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Missing special characters", redCross("✗")))
			passed = false
		} else if requireSymbols {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Contains special characters", greenCheck("✓")))
		}

		if requireDigits && !containsDigit(password) {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Missing digits", redCross("✗")))
			passed = false
		} else if requireDigits {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Contains digits", greenCheck("✓")))
		}

		if requireUppercase && !containsUppercase(password) {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Missing uppercase letters", redCross("✗")))
			passed = false
		} else if requireUppercase {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Contains uppercase letters", greenCheck("✓")))
		}

		if requireLowercase && !containsLowercase(password) {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Missing lowercase letters", redCross("✗")))
			passed = false
		} else if requireLowercase {
			currentAnalysis = append(currentAnalysis, fmt.Sprintf("  %s Contains lowercase letters", greenCheck("✓")))
		}

		// Format result for this password
		statusText := func() string {
			if passed {
				passCount++
				return greenText("STRONG") + " " + greenCheck("✓")
			}
			return redText("WEAK") + " " + redCross("✗")
		}()

		result := fmt.Sprintf("%s %d: %s\n", cyanText("Password"), i+1, statusText)

		for _, check := range currentAnalysis {
			result += fmt.Sprintf("%s\n", check)
		}
		result += "\n"

		// For file output (without colors)
		plainResult := fmt.Sprintf("Password %d: %s\n", i+1, func() string {
			if passed {
				return "STRONG ✓"
			}
			return "WEAK ✗"
		}())
		for _, check := range currentAnalysis {
			// Remove color codes for file output
			plainCheck := strings.ReplaceAll(check, greenCheck("✓"), "✓")
			plainCheck = strings.ReplaceAll(plainCheck, redCross("✗"), "✗")
			plainResult += fmt.Sprintf("%s\n", plainCheck)
		}
		plainResult += "\n"

		allResults = append(allResults, plainResult)
		fmt.Print(result)
	}

	// Summary
	summaryText := func() string {
		if noColor {
			return fmt.Sprintf("Summary: %d/%d passwords meet all criteria", passCount, len(passwords))
		}
		if passCount == len(passwords) {
			return greenText(fmt.Sprintf("🎉 Excellent! All %d passwords are strong!", len(passwords)))
		} else if passCount > len(passwords)/2 {
			return yellowText(fmt.Sprintf("👍 Good! %d/%d passwords meet criteria", passCount, len(passwords)))
		} else {
			return redText(fmt.Sprintf("⚠️  Warning! Only %d/%d passwords meet criteria", passCount, len(passwords)))
		}
	}()

	plainSummary := fmt.Sprintf("Summary: %d/%d passwords meet all criteria\n", passCount, len(passwords))
	allResults = append(allResults, plainSummary)

	fmt.Println(summaryText)

	// Write to file if specified
	if output != "" {
		fullOutput := strings.Join(allResults, "")
		if err := os.WriteFile(output, []byte(fullOutput), 0644); err != nil {
			cmd.PrintErrf("Error writing to output file: %v\n", err)
			return
		}
		cmd.Printf("Analysis results written to %s\n", output)
	}
}
func containsSymbol(password string) bool {
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	for _, char := range password {
		if strings.ContainsRune(symbols, char) {
			return true
		}
	}
	return false
}

func containsDigit(password string) bool {
	for _, char := range password {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func containsUppercase(password string) bool {
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func containsLowercase(password string) bool {
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

// analyzeFile reads passwords from a file and returns them as a slice
func analyzeFile(filepath string) []string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	lines := strings.Split(string(file), "\n")
	var passwords []string

	// Filter out empty lines
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			passwords = append(passwords, line)
		}
	}

	return passwords
}

func checkFileExists(filepath string) error {
	// Check if file exists and is readable
	info, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filepath)
		}
		return fmt.Errorf("cannot access file: %v", err)
	}

	if info.IsDir() {
		return fmt.Errorf("path is a directory, not a file: %s", filepath)
	}

	// Try to read the file to check if it's readable
	_, err = os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("cannot read file: %v", err)
	}

	return nil
}
