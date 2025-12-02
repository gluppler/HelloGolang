package main

import (
	"fmt"
	"os"
	"strings"

	"hellogolang/Projects/Binutils/elf"
)

// As - Assembler (GNU as equivalent - simplified)

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

	if err := assembleFile(inputFile, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// assembleFile assembles source file to object file
func assembleFile(inputFile, outputFile string) error {
	// Read assembly source
	source, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	// Secure: validate file size
	if len(source) > 10*1024*1024 { // 10MB limit
		return fmt.Errorf("input file too large")
	}

	// Parse assembly (simplified - full implementation would parse all instructions)
	instructions := parseAssembly(string(source))

	// Generate object file
	object := generateObjectFile(instructions)

	// Write object file
	return writeObjectFile(outputFile, object)
}

// Instruction represents an assembly instruction
type Instruction struct {
	Opcode string
	Operands []string
	Label   string
}

// parseAssembly parses assembly source
func parseAssembly(source string) []Instruction {
	lines := strings.Split(source, "\n")
	instructions := []Instruction{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// Secure: validate line length
		if len(line) > 10000 {
			continue
		}

		// Parse label
		var label string
		if idx := strings.Index(line, ":"); idx != -1 {
			label = strings.TrimSpace(line[:idx])
			line = strings.TrimSpace(line[idx+1:])
		}

		// Parse instruction
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		opcode := parts[0]
		operands := []string{}
		if len(parts) > 1 {
			operands = parts[1:]
		}

		instructions = append(instructions, Instruction{
			Opcode:   opcode,
			Operands: operands,
			Label:    label,
		})
	}

	return instructions
}

// generateObjectFile generates object file from instructions
func generateObjectFile(instructions []Instruction) *elf.ELF {
	// Create minimal ELF object file
	elfFile := &elf.ELF{
		Class:   "ELF64",
		Data:    "Little Endian",
		Version: 1,
		OSABI:   "ELFOSABI_LINUX",
		Type:    "ET_REL",
		Machine: "EM_X86_64",
		Entry:   0,
		Sections: []elf.Section{
			{
				Name:   ".text",
				Type:   1, // PROGBITS
				Flags:  6, // ALLOC | EXECINSTR
				Addr:   0,
				Offset: 0,
				Size:   uint64(len(instructions) * 4), // Approximate
				Data:   generateCode(instructions),
			},
		},
		Symbols: generateSymbols(instructions),
	}

	return elfFile
}

// generateCode generates machine code from instructions
func generateCode(instructions []Instruction) []byte {
	// Simplified code generation
	// In production, would properly encode each instruction
	code := make([]byte, 0, len(instructions)*4)

	for _, inst := range instructions {
		// Simple encoding (placeholder)
		switch inst.Opcode {
		case "ret":
			code = append(code, 0xc3)
		case "nop":
			code = append(code, 0x90)
		default:
			// Default: NOP
			code = append(code, 0x90)
		}
	}

	return code
}

// generateSymbols generates symbols from instructions
func generateSymbols(instructions []Instruction) []elf.Symbol {
	symbols := []elf.Symbol{}

	for i, inst := range instructions {
		if inst.Label != "" {
			symbols = append(symbols, elf.Symbol{
				Name:    inst.Label,
				Value:   uint64(i * 4),
				Size:    0,
				Info:    2, // STT_FUNC
				Type:    "STT_FUNC",
				Binding: "STB_LOCAL",
			})
		}
	}

	return symbols
}

// writeObjectFile writes object file
func writeObjectFile(filename string, elfFile *elf.ELF) error {
	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	// Write ELF header (simplified - in production use proper ELF writer)
	// For now, just write section data
	// Write section data
	for _, section := range elfFile.Sections {
		if len(section.Data) > 0 {
			if _, err := output.Write(section.Data); err != nil {
				return err
			}
		}
	}

	return nil
}
