package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Ld - Linker (GNU ld equivalent - simplified)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s -o <output> <input>...\n", os.Args[0])
		os.Exit(1)
	}

	outputFile := ""
	inputFiles := []string{}

	// Parse arguments
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-o" && i+1 < len(os.Args) {
			outputFile = os.Args[i+1]
			i++
		} else {
			inputFiles = append(inputFiles, os.Args[i])
		}
	}

	if outputFile == "" {
		fmt.Fprintf(os.Stderr, "Error: no output file specified\n")
		os.Exit(1)
	}

	if len(inputFiles) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no input files specified\n")
		os.Exit(1)
	}

	if err := linkFiles(inputFiles, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// linkFiles links multiple object files
func linkFiles(inputFiles []string, outputFile string) error {
	// Secure: validate number of input files
	if len(inputFiles) > 1000 {
		return fmt.Errorf("too many input files")
	}

	// Parse all input files
	objects := []*elf.ELF{}
	allSymbols := []elf.Symbol{}
	allSections := []elf.Section{}

	for _, filename := range inputFiles {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", filename, err)
		}

		elfFile, err := elf.ParseELF(file)
		file.Close()

		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", filename, err)
		}

		objects = append(objects, elfFile)
		allSymbols = append(allSymbols, elfFile.Symbols...)
		allSections = append(allSections, elfFile.Sections...)
	}

	// Resolve symbols
	resolvedSymbols := resolveSymbols(allSymbols)

	// Merge sections
	mergedSections := mergeSections(allSections)

	// Create output ELF
	outputELF := &elf.ELF{
		Class:    objects[0].Class,
		Data:     objects[0].Data,
		Version:  1,
		OSABI:    objects[0].OSABI,
		Type:     "ET_EXEC",
		Machine:  objects[0].Machine,
		Entry:    findEntryPoint(objects),
		Sections: mergedSections,
		Symbols:  resolvedSymbols,
	}

	// Write output
	return writeLinkedELF(outputFile, outputELF)
}

// resolveSymbols resolves symbol references
func resolveSymbols(symbols []elf.Symbol) []elf.Symbol {
	symbolMap := make(map[string]*elf.Symbol)
	resolved := []elf.Symbol{}

	for i := range symbols {
		sym := &symbols[i]
		if existing, exists := symbolMap[sym.Name]; exists {
			// Resolve conflict (prefer defined over undefined)
			if sym.Type != "STT_NOTYPE" && existing.Type == "STT_NOTYPE" {
				symbolMap[sym.Name] = sym
			}
		} else {
			symbolMap[sym.Name] = sym
		}
	}

	for _, sym := range symbolMap {
		resolved = append(resolved, *sym)
	}

	return resolved
}

// mergeSections merges sections from multiple objects
func mergeSections(sections []elf.Section) []elf.Section {
	sectionMap := make(map[string]*elf.Section)
	merged := []elf.Section{}

	for i := range sections {
		section := &sections[i]
		if existing, exists := sectionMap[section.Name]; exists {
			// Merge section data
			existing.Data = append(existing.Data, section.Data...)
			existing.Size += section.Size
		} else {
			sectionMap[section.Name] = section
		}
	}

	for _, section := range sectionMap {
		merged = append(merged, *section)
	}

	return merged
}

// findEntryPoint finds entry point from objects
func findEntryPoint(objects []*elf.ELF) uint64 {
	for _, obj := range objects {
		if obj.Entry != 0 {
			return obj.Entry
		}
	}
	return 0
}

// writeLinkedELF writes linked ELF file
func writeLinkedELF(filename string, elfFile *elf.ELF) error {
	// In a full implementation, we would properly construct ELF file
	// For now, create a minimal executable
	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	// Write ELF header
	header := createELFHeader(elfFile)
	if _, err := output.Write(header); err != nil {
		return err
	}

	// Write sections
	for _, section := range elfFile.Sections {
		if len(section.Data) > 0 {
			// Secure: validate section size
			if len(section.Data) > 100*1024*1024 {
				return fmt.Errorf("section too large: %s", section.Name)
			}

			if _, err := output.Write(section.Data); err != nil {
				return err
			}
		}
	}

	return nil
}

// createELFHeader creates ELF header bytes
func createELFHeader(elfFile *elf.ELF) []byte {
	header := make([]byte, 64)

	// Magic
	header[0] = 0x7f
	header[1] = 'E'
	header[2] = 'L'
	header[3] = 'F'

	// Class (64-bit)
	header[4] = 2

	// Data (little endian)
	header[5] = 1

	// Version
	header[6] = 1

	// OS/ABI
	header[7] = 0

	// Type (executable)
	binary.LittleEndian.PutUint16(header[16:18], 2)

	// Machine (x86_64)
	binary.LittleEndian.PutUint16(header[18:20], 0x3e)

	// Version
	binary.LittleEndian.PutUint32(header[20:24], 1)

	// Entry point
	binary.LittleEndian.PutUint64(header[24:32], elfFile.Entry)

	// Program header offset
	binary.LittleEndian.PutUint64(header[32:40], 64)

	// Section header offset
	binary.LittleEndian.PutUint64(header[40:48], 64+uint64(len(elfFile.Sections))*64)

	// Flags
	binary.LittleEndian.PutUint32(header[48:52], 0)

	// Header size
	binary.LittleEndian.PutUint16(header[52:54], 64)

	// Program header size
	binary.LittleEndian.PutUint16(header[54:56], 56)

	// Number of program headers
	binary.LittleEndian.PutUint16(header[56:58], 1)

	// Section header size
	binary.LittleEndian.PutUint16(header[58:60], 64)

	// Number of section headers
	binary.LittleEndian.PutUint16(header[60:62], uint16(len(elfFile.Sections)))

	// String table index
	binary.LittleEndian.PutUint16(header[62:64], 0)

	return header
}
