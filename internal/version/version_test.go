package version

import (
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	info := Info()

	// Check that info contains expected components
	expectedParts := []string{
		"Password Zen v",
		"Built:",
		"Commit:",
		"Go:",
	}

	for _, part := range expectedParts {
		if !strings.Contains(info, part) {
			t.Errorf("Info() should contain '%s', but got: %s", part, info)
		}
	}

	// Check that version is included
	if !strings.Contains(info, Version) {
		t.Errorf("Info() should contain version '%s', but got: %s", Version, info)
	}
}

func TestShort(t *testing.T) {
	short := Short()

	if short != Version {
		t.Errorf("Short() = %v, want %v", short, Version)
	}
}

func TestVersionConstants(t *testing.T) {
	// Test that version constants are not empty
	if Version == "" {
		t.Error("Version should not be empty")
	}

	// BuildDate and GitCommit can be "development" in tests
	if BuildDate == "" {
		t.Error("BuildDate should not be empty")
	}

	if GitCommit == "" {
		t.Error("GitCommit should not be empty")
	}
}
