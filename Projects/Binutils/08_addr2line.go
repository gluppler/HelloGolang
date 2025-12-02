package main

import (
	"fmt"
	"os"
	"strconv"

	"hellogolang/Projects/Binutils/elf"
)

// Addr2line - Convert addresses to file/line (GNU addr2line equivalent)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s -e <executable> <address>...\n", os.Args[0])
		os.Exit(1)
	}

	executable := ""
	addresses := []uint64{}

	// Parse arguments
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-e" && i+1 < len(os.Args) {
			executable = os.Args[i+1]
			i++
		} else {
			// Parse address
			addr, err := parseAddress(os.Args[i])
			if err == nil {
				addresses = append(addresses, addr)
			}
		}
	}

	if executable == "" {
		fmt.Fprintf(os.Stderr, "Error: no executable specified\n")
		os.Exit(1)
	}

	if len(addresses) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no addresses specified\n")
		os.Exit(1)
	}

	file, err := os.Open(executable)
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

	for _, addr := range addresses {
		file, line := addrToLine(elfFile, addr)
		if file != "" {
			fmt.Printf("%s:%d\n", file, line)
		} else {
			fmt.Printf("??:0\n")
		}
	}
}

// parseAddress parses address string
func parseAddress(s string) (uint64, error) {
	// Remove 0x prefix if present
	if len(s) > 2 && s[0:2] == "0x" {
		s = s[2:]
	}

	// Secure: validate length
	if len(s) > 16 {
		return 0, fmt.Errorf("address too long")
	}

	val, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// addrToLine converts address to file:line
func addrToLine(elfFile *elf.ELF, addr uint64) (string, int) {
	// Find section containing address
	for _, section := range elfFile.Sections {
		if addr >= section.Addr && addr < section.Addr+section.Size {
			// In a full implementation, we would parse DWARF debug info
			// For now, return section name as file
			return section.Name, 0
		}
	}

	return "", 0
}
