package main

import (
	"fmt"
	"os"

	"hellogolang/Projects/Binutils/elf"
)

// Size - List section sizes (GNU size equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>...\n", os.Args[0])
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		if err := showSize(filename); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", filename, err)
		}
	}
}

// showSize shows section sizes
func showSize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	elfFile, err := elf.ParseELF(file)
	if err != nil {
		return fmt.Errorf("not an ELF file: %w", err)
	}

	var textSize, dataSize, bssSize uint64

	for _, section := range elfFile.Sections {
		switch section.Name {
		case ".text", ".init", ".fini":
			textSize += section.Size
		case ".data", ".rodata", ".sdata":
			dataSize += section.Size
		case ".bss", ".sbss":
			bssSize += section.Size
		}
	}

	total := textSize + dataSize + bssSize

	fmt.Printf("   text    data     bss     dec     hex filename\n")
	fmt.Printf("%7d %7d %7d %7d %7x %s\n",
		textSize, dataSize, bssSize, total, total, filename)

	return nil
}
