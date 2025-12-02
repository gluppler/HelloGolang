package main

import (
	"bytes"
	"os"
	"testing"
)

// TestParseELF tests ELF parsing
func TestParseELF(t *testing.T) {
	// Create minimal ELF file for testing
	elfData := []byte{
		0x7f, 'E', 'L', 'F', // Magic
		2,                   // 64-bit
		1,                   // Little endian
		1,                   // Version
		0,                   // OS/ABI
		0, 0, 0, 0, 0, 0, 0, // Padding
		2, 0,                 // Type: ET_EXEC
		0x3e, 0x00,           // Machine: EM_X86_64
		1, 0, 0, 0,           // Version
		0, 0, 0, 0, 0, 0, 0, 0, // Entry
		0, 0, 0, 0, 0, 0, 0, 0, // PhOff
		0, 0, 0, 0, 0, 0, 0, 0, // ShOff
		0, 0, 0, 0,           // Flags
		64, 0,                 // EhSize
		56, 0,                 // PhentSize
		0, 0,                  // PhNum
		64, 0,                 // ShentSize
		0, 0,                  // ShNum
		0, 0,                  // ShStrndx
	}

	reader := bytes.NewReader(elfData)
	elf, err := ParseELF(reader)

	if err != nil {
		t.Fatalf("Failed to parse ELF: %v", err)
	}

	if elf.Class != "ELF64" {
		t.Errorf("Expected ELF64, got %s", elf.Class)
	}

	if elf.Machine != "EM_X86_64" {
		t.Errorf("Expected EM_X86_64, got %s", elf.Machine)
	}
}

// TestReadCString tests C string reading
func TestReadCString(t *testing.T) {
	data := []byte{'h', 'e', 'l', 'l', 'o', 0, 'w', 'o', 'r', 'l', 'd'}
	result := readCString(data)

	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
}

// TestGetELFType tests ELF type conversion
func TestGetELFType(t *testing.T) {
	if getELFType(2) != "ET_EXEC" {
		t.Errorf("Expected ET_EXEC, got %s", getELFType(2))
	}
}

// TestGetMachine tests machine type conversion
func TestGetMachine(t *testing.T) {
	if getMachine(0x3e) != "EM_X86_64" {
		t.Errorf("Expected EM_X86_64, got %s", getMachine(0x3e))
	}
}
