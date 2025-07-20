package cmd

import (
	"testing"
)

func TestContainsSymbol(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Password with symbol", "test@123", true},
		{"Password without symbol", "test123", false},
		{"Password with multiple symbols", "test!@#$", true},
		{"Empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsSymbol(tt.password); got != tt.want {
				t.Errorf("containsSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDigit(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Password with digit", "test123", true},
		{"Password without digit", "testABC", false},
		{"Password with single digit", "test1", true},
		{"Empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDigit(tt.password); got != tt.want {
				t.Errorf("containsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUppercase(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Password with uppercase", "Test123", true},
		{"Password without uppercase", "test123", false},
		{"Password with multiple uppercase", "TEST", true},
		{"Empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsUppercase(tt.password); got != tt.want {
				t.Errorf("containsUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsLowercase(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Password with lowercase", "Test123", true},
		{"Password without lowercase", "TEST123", false},
		{"Password with multiple lowercase", "test", true},
		{"Empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsLowercase(tt.password); got != tt.want {
				t.Errorf("containsLowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}
