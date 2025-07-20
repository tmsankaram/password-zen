package cmd

import (
	"strings"
	"testing"
)

func TestBuildCharset(t *testing.T) {
	tests := []struct {
		name              string
		includeDigits     bool
		includeSymbols    bool
		excludeAmbiguous  bool
		expectContains    []string
		expectNotContains []string
	}{
		{
			name:           "Basic charset with digits",
			includeDigits:  true,
			includeSymbols: false,
			expectContains: []string{"a", "Z", "0", "9"},
		},
		{
			name:           "Charset with symbols",
			includeDigits:  true,
			includeSymbols: true,
			expectContains: []string{"a", "Z", "0", "!", "@"},
		},
		{
			name:              "Exclude ambiguous characters",
			includeDigits:     true,
			includeSymbols:    false,
			excludeAmbiguous:  true,
			expectNotContains: []string{"i", "l", "1", "L", "o", "0", "O"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charset := buildCharset(tt.includeDigits, tt.includeSymbols, tt.excludeAmbiguous)

			for _, char := range tt.expectContains {
				if !strings.Contains(charset, char) {
					t.Errorf("Expected charset to contain '%s', but it didn't. Charset: %s", char, charset)
				}
			}

			for _, char := range tt.expectNotContains {
				if strings.Contains(charset, char) {
					t.Errorf("Expected charset to NOT contain '%s', but it did. Charset: %s", char, charset)
				}
			}
		})
	}
}

func TestGenerateSecurePassword(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charset string
		wantErr bool
	}{
		{
			name:    "Valid password generation",
			length:  12,
			charset: "abcdefghijklmnopqrstuvwxyz",
			wantErr: false,
		},
		{
			name:    "Zero length should error",
			length:  0,
			charset: "abc",
			wantErr: true,
		},
		{
			name:    "Empty charset should error",
			length:  12,
			charset: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := generateSecurePassword(tt.length, tt.charset)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if len(password) != tt.length {
				t.Errorf("Expected password length %d, got %d", tt.length, len(password))
			}

			// Check that all characters in password are from charset
			for _, char := range password {
				if !strings.Contains(tt.charset, string(char)) {
					t.Errorf("Password contains character '%c' not in charset '%s'", char, tt.charset)
				}
			}
		})
	}
}
