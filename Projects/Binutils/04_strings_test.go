package main

import (
	"os"
	"testing"
)

// TestExtractStrings tests string extraction
func TestExtractStrings(t *testing.T) {
	// Create test file
	testData := []byte("Hello\x00World\x00Test123\x00\x00\x00\x00")
	
	tmpfile, err := os.CreateTemp("", "test_strings_*.bin")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	tmpfile.Close()

	// Test extraction
	err = extractStrings(tmpfile.Name(), 4)
	if err != nil {
		t.Errorf("extractStrings failed: %v", err)
	}
}

// TestParseInt tests integer parsing
func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"4", 4, false},
		{"10", 10, false},
		{"100", 100, false},
		{"abc", 0, true},
		{"101", 0, true}, // Too large
	}

	for _, tt := range tests {
		result, err := parseInt(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("parseInt(%s) expected error, got %d", tt.input, result)
			}
		} else {
			if err != nil {
				t.Errorf("parseInt(%s) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseInt(%s) = %d, expected %d", tt.input, result, tt.expected)
			}
		}
	}
}
