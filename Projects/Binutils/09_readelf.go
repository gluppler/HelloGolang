package main

import (
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Readelf - Display information about ELF files (GNU readelf equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <option> <elf-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options: -h (header), -S (sections), -s (symbols), -l (segments)\n")
		os.Exit(1)
	}

	option := "-h"
	filename := ""

	if len(os.Args) == 2 {
		filename = os.Args[1]
	} else {
		option = os.Args[1]
		filename = os.Args[2]
	}

	file, err := os.Open(filename)
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

	switch option {
	case "-h", "--file-header":
		showFileHeader(elfFile, filename)
	case "-S", "--section-headers":
		showSectionHeaders(elfFile)
	case "-s", "--symbols":
		showSymbols(elfFile)
	case "-l", "--program-headers":
		showProgramHeaders(elfFile)
	case "-a", "--all":
		showAll(elfFile, filename)
	default:
		showFileHeader(elfFile, filename)
	}
}

// truncateString truncates string to max length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// showFileHeader shows ELF file header
func showFileHeader(elfFile *elf.ELF, filename string) {
	fmt.Printf("ELF Header:\n")
	fmt.Printf("  Magic:   ")
	for i := 0; i < 16; i++ {
		if i < 4 {
			fmt.Printf(" %02x", elfFile.Header.Magic[i])
		} else {
			fmt.Printf(" %02x", 0)
		}
	}
	fmt.Println()
	fmt.Printf("  Class:                             %s\n", elfFile.Class)
	fmt.Printf("  Data:                              %s\n", elfFile.Data)
	fmt.Printf("  Version:                           %d (current)\n", elfFile.Version)
	fmt.Printf("  OS/ABI:                            %s\n", elfFile.OSABI)
	fmt.Printf("  ABI Version:                      0\n")
	fmt.Printf("  Type:                              %s\n", elfFile.Type)
	fmt.Printf("  Machine:                           %s\n", elfFile.Machine)
	fmt.Printf("  Version:                           0x%x\n", elfFile.Version)
	fmt.Printf("  Entry point address:               0x%x\n", elfFile.Entry)
	fmt.Printf("  Start of program headers:          %d (bytes into file)\n", elfFile.Header.PhOff64)
	fmt.Printf("  Start of section headers:          %d (bytes into file)\n", elfFile.Header.ShOff64)
	fmt.Printf("  Flags:                             0x%x\n", elfFile.Header.Flags)
	fmt.Printf("  Size of this header:               %d (bytes)\n", elfFile.Header.EhSize)
	fmt.Printf("  Size of program headers:           %d (bytes)\n", elfFile.Header.PhentSize)
	fmt.Printf("  Number of program headers:         %d\n", elfFile.Header.PhNum)
	fmt.Printf("  Size of section headers:           %d (bytes)\n", elfFile.Header.ShentSize)
	fmt.Printf("  Number of section headers:         %d\n", elfFile.Header.ShNum)
	fmt.Printf("  Section header string table index: %d\n", elfFile.Header.ShStrndx)
}

// showSectionHeaders shows section headers
func showSectionHeaders(elfFile *elf.ELF) {
	fmt.Printf("There are %d section headers, starting at offset 0x%x:\n\n",
		len(elfFile.Sections), elfFile.Header.ShOff64)
	fmt.Printf("Section Headers:\n")
	fmt.Printf("  [Nr] Name              Type            Address          Off    Size   ES Flg Lk Inf Al\n")

	for i, section := range elfFile.Sections {
		typeName := getSectionType(section.Type)
		fmt.Printf("  [%2d] %-17s %-15s %016x %06x %06x %02x %3s %2d %3d %2d\n",
			i, truncateString(section.Name, 17), typeName,
			section.Addr, section.Offset, section.Size,
			section.EntSize, getSectionFlags(section.Flags),
			section.Link, section.Info, section.AddrAlign)
	}
}

// showSymbols shows symbol table
func showSymbols(elfFile *elf.ELF) {
	if len(elfFile.Symbols) == 0 {
		fmt.Println("No symbol table found")
		return
	}

	fmt.Printf("Symbol table '.symtab' contains %d entries:\n", len(elfFile.Symbols))
	fmt.Printf("   Num:    Value          Size Type    Bind   Vis      Ndx Name\n")

	for i, sym := range elfFile.Symbols {
		ndx := "UND"
		if sym.Shndx != 0 {
			ndx = fmt.Sprintf("%3d", sym.Shndx)
		}

		fmt.Printf("%6d: %016x %5d %-7s %-6s DEFAULT %3s %s\n",
			i, sym.Value, sym.Size, sym.Type, sym.Binding, ndx, sym.Name)
	}
}

// showProgramHeaders shows program headers
func showProgramHeaders(elfFile *elf.ELF) {
	fmt.Printf("Elf file type is %s\n", elfFile.Type)
	fmt.Printf("Entry point 0x%x\n", elfFile.Entry)
	fmt.Printf("There are %d program headers, starting at offset %d\n\n",
		elfFile.Header.PhNum, elfFile.Header.PhOff64)

	if len(elfFile.Segments) > 0 {
		fmt.Printf("Program Headers:\n")
		fmt.Printf("  Type           Offset             VirtAddr           PhysAddr\n")
		fmt.Printf("                 FileSiz            MemSiz              Flags  Align\n")

		for _, seg := range elfFile.Segments {
			typeName := getSegmentType(seg.Type)
			flags := getSegmentFlags(seg.Flags)
			fmt.Printf("  %-14s 0x%016x 0x%016x 0x%016x\n",
				typeName, seg.Offset, seg.VAddr, seg.PAddr)
			fmt.Printf("                 0x%016x 0x%016x %3s     0x%x\n",
				seg.FileSz, seg.MemSz, flags, seg.Align)
		}
	} else {
		fmt.Printf("No program headers found\n")
	}
}

// showAll shows all information
func showAll(elfFile *elf.ELF, filename string) {
	showFileHeader(elfFile, filename)
	fmt.Println()
	showSectionHeaders(elfFile)
	fmt.Println()
	showSymbols(elfFile)
	fmt.Println()
	showProgramHeaders(elfFile)
}

// getSectionType returns section type name
func getSectionType(t uint32) string {
	types := map[uint32]string{
		0:  "NULL",
		1:  "PROGBITS",
		2:  "SYMTAB",
		3:  "STRTAB",
		4:  "RELA",
		5:  "HASH",
		6:  "DYNAMIC",
		7:  "NOTE",
		8:  "NOBITS",
		9:  "REL",
		10: "SHLIB",
		11: "DYNSYM",
	}
	if name, ok := types[t]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN(%d)", t)
}

// getSectionFlags returns section flags string
func getSectionFlags(flags uint64) string {
	result := ""
	if flags&0x1 != 0 {
		result += "W"
	}
	if flags&0x2 != 0 {
		result += "A"
	}
	if flags&0x4 != 0 {
		result += "X"
	}
	if result == "" {
		result = "   "
	}
	return result
}

// getSegmentType returns segment type name
func getSegmentType(t uint32) string {
	types := map[uint32]string{
		0: "NULL",
		1: "LOAD",
		2: "DYNAMIC",
		3: "INTERP",
		4: "NOTE",
	}
	if name, ok := types[t]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN(%d)", t)
}

// getSegmentFlags returns segment flags string
func getSegmentFlags(flags uint32) string {
	result := ""
	if flags&0x1 != 0 {
		result += "E"
	}
	if flags&0x2 != 0 {
		result += "W"
	}
	if flags&0x4 != 0 {
		result += "R"
	}
	return result
}
