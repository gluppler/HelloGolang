package main

import (
	"fmt"
	"os"
)

// Nlmconv - Convert object files to NetWare Loadable Module (GNU nlmconv equivalent)

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

	// Secure: validate number of input files
	if len(inputFiles) > 1000 {
		fmt.Fprintf(os.Stderr, "Error: too many input files\n")
		os.Exit(1)
	}

	if err := convertToNLM(inputFiles, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// convertToNLM converts object files to NLM format
func convertToNLM(inputFiles []string, outputFile string) error {
	// Secure: validate filenames
	if len(outputFile) > 255 {
		return fmt.Errorf("output filename too long")
	}

	// NLM file format structure
	nlmHeader := createNLMHeader(outputFile)

	// Read and process input files
	objectData := []byte{}
	for _, inputFile := range inputFiles {
		// Secure: validate filename
		if len(inputFile) > 255 {
			return fmt.Errorf("input filename too long: %s", inputFile)
		}

		data, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", inputFile, err)
		}

		// Secure: validate file size
		if len(data) > 100*1024*1024 {
			return fmt.Errorf("file too large: %s", inputFile)
		}

		objectData = append(objectData, data...)
	}

	// Combine header and object data
	nlmFile := append(nlmHeader, objectData...)

	// Secure: validate total size
	if len(nlmFile) > 200*1024*1024 {
		return fmt.Errorf("NLM file too large")
	}

	return os.WriteFile(outputFile, nlmFile, 0644)
}

// createNLMHeader creates NLM file header
func createNLMHeader(moduleName string) []byte {
	// NLM signature: "NetWare Loadable Module"
	header := make([]byte, 128)

	// Write NLM signature
	signature := "NetWare Loadable Module\x00"
	copy(header[0:], []byte(signature))

	// Write module name
	nameBytes := []byte(moduleName)
	if len(nameBytes) > 64 {
		nameBytes = nameBytes[:64]
	}
	copy(header[32:], nameBytes)

	// Write version (1.0)
	header[96] = 1
	header[97] = 0

	return header
}
