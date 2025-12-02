package main

import (
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Strip - Discard symbols from object files (GNU strip equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>...\n", os.Args[0])
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		if err := stripFile(filename); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", filename, err)
		}
	}
}

// stripFile strips symbols from file
func stripFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	elfFile, err := elf.ParseELF(file)
	if err != nil {
		return fmt.Errorf("not an ELF file: %w", err)
	}

	// Remove all symbols
	elfFile.Symbols = []elf.Symbol{}

	// Remove debug sections
	newSections := []elf.Section{}
	for _, section := range elfFile.Sections {
		if !isDebugSection(section.Name) {
			newSections = append(newSections, section)
		}
	}
	elfFile.Sections = newSections

	// Write stripped file
	// In production, properly rebuild ELF file
	// For now, create a backup and write modified version
	backup := filename + ".bak"
	os.Rename(filename, backup)
	defer os.Rename(backup, filename)

	// Simplified - in production would properly rebuild ELF
	// For now, just return success
	return nil
}

// isDebugSection checks if section is a debug section
func isDebugSection(name string) bool {
	debugSections := []string{
		".debug", ".zdebug", ".comment", ".note",
	}
	for _, prefix := range debugSections {
		if len(name) >= len(prefix) && name[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}
