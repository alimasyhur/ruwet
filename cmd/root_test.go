package cmd

import (
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	// Test that root command executes without error
	rootCmd.SetArgs([]string{"--help"})

	// Cobra will print to stdout, we just check it doesn't error
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Root command --help failed: %v", err)
	}
}

func TestScanCommand(t *testing.T) {
	rootCmd.SetArgs([]string{"scan", "--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Scan command --help failed: %v", err)
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
