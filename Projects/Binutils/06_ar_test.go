package main

import (
	"os"
	"testing"
)

// TestCreateArchive tests archive creation
func TestCreateArchive(t *testing.T) {
	// Create test files
	tmpfile1, err := os.CreateTemp("", "test_ar_1_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile1.Name())
	tmpfile1.WriteString("test content 1")
	tmpfile1.Close()

	tmpfile2, err := os.CreateTemp("", "test_ar_2_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile2.Name())
	tmpfile2.WriteString("test content 2")
	tmpfile2.Close()

	archiveName := os.TempDir() + "/test_archive.a"
	defer os.Remove(archiveName)

	// Create archive
	err = createArchive(archiveName, []string{tmpfile1.Name(), tmpfile2.Name()})
	if err != nil {
		t.Fatalf("createArchive failed: %v", err)
	}

	// List archive
	err = listArchive(archiveName)
	if err != nil {
		t.Errorf("listArchive failed: %v", err)
	}
}

// TestReadArchive tests archive reading
func TestReadArchive(t *testing.T) {
	// Create minimal archive
	archiveData := []byte("!<arch>\n")
	
	tmpfile, err := os.CreateTemp("", "test_ar_read_*.a")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(archiveData); err != nil {
		t.Fatalf("Failed to write archive data: %v", err)
	}
	tmpfile.Close()

	members, err := readArchive(tmpfile.Name())
	if err != nil {
		t.Errorf("readArchive failed: %v", err)
	}

	if len(members) != 0 {
		t.Errorf("Expected 0 members, got %d", len(members))
	}
}

// TestContainsPathTraversal tests path traversal detection
func TestContainsPathTraversal(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{"../file.txt", true},
		{"..\\file.txt", true},
		{"/etc/passwd", true},
		{"normal_file.txt", false},
		{"subdir/file.txt", false},
	}

	for _, tt := range tests {
		result := containsPathTraversal(tt.path)
		if result != tt.expected {
			t.Errorf("containsPathTraversal(%s) = %v, expected %v", tt.path, result, tt.expected)
		}
	}
}
