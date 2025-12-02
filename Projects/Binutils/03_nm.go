package main

import (
	"fmt"
	"os"
	"sort"

	"hellogolang/Projects/Binutils/elf"
)

// Nm - List symbols from object files (GNU nm equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file> [options]\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	elfFile, err := elf.ParseELF(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := listSymbols(elfFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// listSymbols lists symbols in nm format
func listSymbols(elfFile *elf.ELF) error {
	if len(elfFile.Symbols) == 0 {
		return nil
	}

	// Sort symbols by value
	symbols := make([]elf.Symbol, len(elfFile.Symbols))
	copy(symbols, elfFile.Symbols)

	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i].Value < symbols[j].Value
	})

	// Print symbols
	for _, sym := range symbols {
		if sym.Name == "" {
			continue
		}

		// Determine symbol type character
		typeChar := byte('?')
		switch sym.Type {
		case "STT_NOTYPE":
			typeChar = 'U' // Undefined
		case "STT_OBJECT":
			typeChar = 'D' // Data
		case "STT_FUNC":
			typeChar = 'T' // Text (code)
		case "STT_SECTION":
			typeChar = 'S' // Section
		case "STT_FILE":
			typeChar = 'A' // Absolute
		}

		// Determine binding
		if sym.Binding == "STB_GLOBAL" {
			typeChar = toUpper(typeChar)
		} else if sym.Binding == "STB_WEAK" {
			if typeChar == 'U' {
				typeChar = 'w'
			} else {
				typeChar = toLower(typeChar)
			}
		}

		// Print symbol
		if sym.Value == 0 && sym.Type == "STT_NOTYPE" {
			fmt.Printf("                 %c %s\n", typeChar, sym.Name)
		} else {
			if elfFile.Class == "ELF64" {
				fmt.Printf("%016x %c %s\n", sym.Value, typeChar, sym.Name)
			} else {
				fmt.Printf("%08x %c %s\n", uint32(sym.Value), typeChar, sym.Name)
			}
		}
	}

	return nil
}

// toUpper converts character to uppercase
func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

// toLower converts character to lowercase
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
