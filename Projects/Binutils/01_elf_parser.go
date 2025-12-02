package main

import (
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// ELF Parser - Core ELF file parsing functionality

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <elf-file>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	elfFile, err := elf.ParseELF(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing ELF: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ELF File: %s\n", os.Args[1])
	fmt.Printf("Class: %s\n", elfFile.Class)
	fmt.Printf("Data: %s\n", elfFile.Data)
	fmt.Printf("Version: %d\n", elfFile.Version)
	fmt.Printf("OS/ABI: %s\n", elfFile.OSABI)
	fmt.Printf("Type: %s\n", elfFile.Type)
	fmt.Printf("Machine: %s\n", elfFile.Machine)
	fmt.Printf("Entry Point: 0x%x\n", elfFile.Entry)
}

