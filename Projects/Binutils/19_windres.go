package main

import (
	"fmt"
	"os"
	"strings"
)

// Windres - Windows resource compiler (GNU windres equivalent)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s -o <output> <input>\n", os.Args[0])
		os.Exit(1)
	}

	outputFile := ""
	inputFile := ""

	// Parse arguments
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-o" && i+1 < len(os.Args) {
			outputFile = os.Args[i+1]
			i++
		} else if !strings.HasPrefix(os.Args[i], "-") {
			inputFile = os.Args[i]
		}
	}

	if outputFile == "" {
		fmt.Fprintf(os.Stderr, "Error: no output file specified\n")
		os.Exit(1)
	}

	if inputFile == "" {
		fmt.Fprintf(os.Stderr, "Error: no input file specified\n")
		os.Exit(1)
	}

	if err := compileResource(inputFile, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// compileResource compiles resource file
func compileResource(inputFile, outputFile string) error {
	// Secure: validate filenames
	if len(inputFile) > 255 || len(outputFile) > 255 {
		return fmt.Errorf("filename too long")
	}

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	// Secure: validate file size
	if len(data) > 10*1024*1024 {
		return fmt.Errorf("input file too large")
	}

	// Parse resource definitions
	resources := parseResourceFile(string(data))

	// Generate COFF resource object file
	coffData := generateCOFFResource(resources)

	// Secure: validate output size
	if len(coffData) > 50*1024*1024 {
		return fmt.Errorf("output file too large")
	}

	return os.WriteFile(outputFile, coffData, 0644)
}

// Resource represents a Windows resource
type Resource struct {
	Type     string
	Name     string
	Language uint16
	Data     []byte
}

// parseResourceFile parses .rc resource file
func parseResourceFile(content string) []Resource {
	resources := []Resource{}
	lines := strings.Split(content, "\n")

	// Secure: limit number of lines
	if len(lines) > 100000 {
		lines = lines[:100000]
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") {
			continue
		}

		// Secure: validate line length
		if len(line) > 10000 {
			continue
		}

		// Parse resource definition (simplified)
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			resources = append(resources, Resource{
				Type:     parts[0],
				Name:     parts[1],
				Language: 0x0409, // English (US)
				Data:     []byte(strings.Join(parts[2:], " ")),
			})
		}
	}

	// Secure: limit number of resources
	if len(resources) > 10000 {
		resources = resources[:10000]
	}

	return resources
}

// generateCOFFResource generates COFF format resource object
func generateCOFFResource(resources []Resource) []byte {
	// COFF header (20 bytes)
	header := make([]byte, 20)

	// Machine type (x86)
	header[0] = 0x4c
	header[1] = 0x01

	// Number of sections
	header[2] = 1
	header[3] = 0

	// Timestamp
	// (simplified - would use actual timestamp)

	// Pointer to symbol table
	// (simplified)

	// Number of symbols
	// (simplified)

	// Optional header size
	header[16] = 0
	header[17] = 0

	// Characteristics
	header[18] = 0x01 // RELOCS_STRIPPED
	header[19] = 0x00

	// Section header for .rsrc section
	sectionHeader := make([]byte, 40)
	copy(sectionHeader[0:8], []byte(".rsrc\x00\x00\x00"))

	// Section data
	sectionData := []byte{}
	for _, res := range resources {
		// Resource entry (simplified)
		entry := make([]byte, 16)
		copy(entry[0:4], []byte(res.Type))
		copy(entry[4:8], []byte(res.Name))
		entry[8] = byte(res.Language)
		entry[9] = byte(res.Language >> 8)
		
		// Secure: limit resource data size
		resData := res.Data
		if len(resData) > 65535 {
			resData = resData[:65535]
		}
		entry[10] = byte(len(resData))
		entry[11] = byte(len(resData) >> 8)
		
		sectionData = append(sectionData, entry...)
		sectionData = append(sectionData, resData...)
	}

	// Update section header with size
	sectionHeader[16] = byte(len(sectionData))
	sectionHeader[17] = byte(len(sectionData) >> 8)
	sectionHeader[18] = byte(len(sectionData) >> 16)
	sectionHeader[19] = byte(len(sectionData) >> 24)

	// Combine all parts
	result := append(header, sectionHeader...)
	result = append(result, sectionData...)

	return result
}
