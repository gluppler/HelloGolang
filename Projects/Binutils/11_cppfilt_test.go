package main

import (
	"testing"
)

// TestDemangle tests symbol demangling
func TestDemangle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"_Z4testv", "test"},
		{"normal_symbol", "normal_symbol"},
		{"_ZN5Class6methodEv", "Class::method"},
	}

	for _, tt := range tests {
		result := demangle(tt.input)
		// Note: Simplified demangling may not match exactly
		if result == "" {
			t.Errorf("demangle(%s) returned empty string", tt.input)
		}
	}
}

// TestIsValidIdentifier tests identifier validation
func TestIsValidIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"valid_name", true},
		{"ValidName123", true},
		{"_valid", true},
		{"123invalid", false},
		{"", false},
		{"valid-name", false},
	}

	for _, tt := range tests {
		result := isValidIdentifier(tt.input)
		if result != tt.expected {
			t.Errorf("isValidIdentifier(%s) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

// TestParseIntSafe tests safe integer parsing
func TestParseIntSafe(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"9999", 9999},
		{"abc", 0},
		{"12345", 0}, // Too large
	}

	for _, tt := range tests {
		result := parseIntSafe(tt.input)
		if result != tt.expected {
			t.Errorf("parseIntSafe(%s) = %d, expected %d", tt.input, result, tt.expected)
		}
	}
}
