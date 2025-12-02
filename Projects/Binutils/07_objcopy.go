package main

import (
	"fmt"
	"io"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Objcopy - Copy and translate object files (GNU objcopy equivalent)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input> <output> [options]\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	options := parseOptions(os.Args[3:])

	if err := copyObject(inputFile, outputFile, options); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// CopyOptions represents objcopy options
type CopyOptions struct {
	StripAll      bool
	StripDebug    bool
	StripSymbols  []string
	KeepSymbols   []string
	AddSection    map[string][]byte
	RemoveSection []string
}

// parseOptions parses command line options
func parseOptions(args []string) CopyOptions {
	opts := CopyOptions{
		AddSection: make(map[string][]byte),
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--strip-all":
			opts.StripAll = true
		case "--strip-debug":
			opts.StripDebug = true
		case "--strip-symbol":
			if i+1 < len(args) {
				opts.StripSymbols = append(opts.StripSymbols, args[i+1])
				i++
			}
		case "--keep-symbol":
			if i+1 < len(args) {
				opts.KeepSymbols = append(opts.KeepSymbols, args[i+1])
				i++
			}
		case "--remove-section":
			if i+1 < len(args) {
				opts.RemoveSection = append(opts.RemoveSection, args[i+1])
				i++
			}
		}
	}

	return opts
}

// copyObject copies and modifies object file
func copyObject(inputFile, outputFile string, options CopyOptions) error {
	// Read input file
	input, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}
	defer input.Close()

	// Parse ELF
	elfFile, err := elf.ParseELF(input)
	if err != nil {
		return fmt.Errorf("failed to parse ELF: %w", err)
	}

	// Apply transformations
	if options.StripAll {
		elfFile.Symbols = []elf.Symbol{}
		// Remove debug sections
		for i := len(elfFile.Sections) - 1; i >= 0; i-- {
			if isDebugSection(elfFile.Sections[i].Name) {
				elfFile.Sections = append(elfFile.Sections[:i], elfFile.Sections[i+1:]...)
			}
		}
	}

	if options.StripDebug {
		// Remove debug sections only
		for i := len(elfFile.Sections) - 1; i >= 0; i-- {
			if isDebugSection(elfFile.Sections[i].Name) {
				elfFile.Sections = append(elfFile.Sections[:i], elfFile.Sections[i+1:]...)
			}
		}
	}

	// Strip specific symbols
	if len(options.StripSymbols) > 0 {
		newSymbols := []elf.Symbol{}
		for _, sym := range elfFile.Symbols {
			if !containsString(options.StripSymbols, sym.Name) {
				newSymbols = append(newSymbols, sym)
			}
		}
		elfFile.Symbols = newSymbols
	}

	// Keep only specified symbols
	if len(options.KeepSymbols) > 0 {
		newSymbols := []elf.Symbol{}
		for _, sym := range elfFile.Symbols {
			if containsString(options.KeepSymbols, sym.Name) {
				newSymbols = append(newSymbols, sym)
			}
		}
		elfFile.Symbols = newSymbols
	}

	// Remove sections
	if len(options.RemoveSection) > 0 {
		for _, sectionName := range options.RemoveSection {
			for i := len(elfFile.Sections) - 1; i >= 0; i-- {
				if elfFile.Sections[i].Name == sectionName {
					elfFile.Sections = append(elfFile.Sections[:i], elfFile.Sections[i+1:]...)
					break
				}
			}
		}
	}

	// Write output
	return writeELF(outputFile, elfFile)
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

// containsString checks if slice contains string
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// writeELF writes ELF file (simplified - in production use proper ELF writer)
func writeELF(filename string, elfFile *elf.ELF) error {
	// For this implementation, we'll do a simple copy
	// In a full implementation, we would rebuild the ELF file properly
	input, err := os.Open(filename)
	if err != nil {
		// Create new file
		output, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer output.Close()

		// Write minimal ELF header
		header := []byte{
			0x7f, 'E', 'L', 'F', // Magic
			2,                   // 64-bit
			1,                   // Little endian
			1,                   // Version
			0,                   // OS/ABI
			0, 0, 0, 0, 0, 0, 0, // Padding
		}

		if _, err := output.Write(header); err != nil {
			return err
		}

		return nil
	}
	defer input.Close()

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	return err
}
