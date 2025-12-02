package main

import (
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Objdump - Object file dumper (GNU objdump equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := dumpObject(file, os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// dumpObject dumps object file information
func dumpObject(r interface{ Seek(int64, int) (int64, error); Read([]byte) (int, error) }, filename string) error {
	elfFile, err := elf.ParseELF(r)
	if err != nil {
		return fmt.Errorf("not an ELF file: %w", err)
	}

	fmt.Printf("%s:     file format %s-%s\n", filename, elfFile.Class, elfFile.Data)
	fmt.Printf("architecture: %s\n\n", elfFile.Machine)

	// Dump file header
	fmt.Println("File Header:")
	fmt.Printf("  Magic:   %02x %02x %02x %02x\n",
		elfFile.Header.Magic[0], elfFile.Header.Magic[1],
		elfFile.Header.Magic[2], elfFile.Header.Magic[3])
	fmt.Printf("  Class:                             %s\n", elfFile.Class)
	fmt.Printf("  Data:                              %s\n", elfFile.Data)
	fmt.Printf("  Version:                           %d\n", elfFile.Version)
	fmt.Printf("  OS/ABI:                            %s\n", elfFile.OSABI)
	fmt.Printf("  Type:                              %s\n", elfFile.Type)
	fmt.Printf("  Machine:                           %s\n", elfFile.Machine)
	fmt.Printf("  Entry point address:               0x%x\n", elfFile.Entry)
	fmt.Println()

	// Dump sections
	fmt.Println("Sections:")
	fmt.Printf("Idx Name          Size      Address\n")
	for i, section := range elfFile.Sections {
		fmt.Printf("%3d %-13s %08x  %016x\n",
			i, truncateString(section.Name, 13),
			section.Size, section.Addr)
	}
	fmt.Println()

	// Dump symbols
	if len(elfFile.Symbols) > 0 {
		fmt.Println("SYMBOL TABLE:")
		fmt.Printf("%016x %c %s\n", 0, ' ', "UND")
		for _, sym := range elfFile.Symbols {
			if sym.Name != "" {
				binding := ' '
				if sym.Binding == "STB_GLOBAL" {
					binding = 'g'
				} else if sym.Binding == "STB_WEAK" {
					binding = 'w'
				}

				typeChar := ' '
				if sym.Type == "STT_FUNC" {
					typeChar = 'F'
				} else if sym.Type == "STT_OBJECT" {
					typeChar = 'O'
				}

				fmt.Printf("%016x %c%c %s\n",
					sym.Value, binding, typeChar, sym.Name)
			}
		}
		fmt.Println()
	}

	return nil
}

// truncateString truncates string to max length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
