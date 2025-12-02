package elf

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ELF represents an ELF file structure
type ELF struct {
	Class       string
	Data        string
	Version     uint32
	OSABI       string
	Type        string
	Machine     string
	Entry       uint64
	Header      ELFHeader
	Sections    []Section
	Segments    []Segment
	Symbols     []Symbol
	StringTable []byte
}

// ELFHeader represents ELF file header
type ELFHeader struct {
	Magic      [4]byte
	Class      byte
	Data       byte
	Version    byte
	OSABI      byte
	ABIVersion byte
	Padding    [7]byte
	Type       uint16
	Machine    uint16
	Version32  uint32
	Entry64    uint64
	PhOff64    uint64
	ShOff64    uint64
	Flags      uint32
	EhSize     uint16
	PhentSize  uint16
	PhNum      uint16
	ShentSize  uint16
	ShNum      uint16
	ShStrndx   uint16
}

// Section represents an ELF section
type Section struct {
	Name      string
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	AddrAlign uint64
	EntSize   uint64
	Data      []byte
}

// Segment represents an ELF program segment
type Segment struct {
	Type   uint32
	Flags  uint32
	Offset uint64
	VAddr  uint64
	PAddr  uint64
	FileSz uint64
	MemSz  uint64
	Align  uint64
}

// Symbol represents an ELF symbol
type Symbol struct {
	Name    string
	Value   uint64
	Size    uint64
	Info    byte
	Other   byte
	Shndx   uint16
	Type    string
	Binding string
}

// ParseELF parses an ELF file
func ParseELF(r io.ReadSeeker) (*ELF, error) {
	elf := &ELF{}

	// Read ELF header
	var header ELFHeader
	if err := binary.Read(r, binary.LittleEndian, &header.Magic); err != nil {
		return nil, fmt.Errorf("failed to read magic: %w", err)
	}

	// Validate ELF magic
	if header.Magic[0] != 0x7f || header.Magic[1] != 'E' || header.Magic[2] != 'L' || header.Magic[3] != 'F' {
		return nil, fmt.Errorf("invalid ELF magic")
	}

	// Read rest of header
	if _, err := r.Seek(4, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &header.Class); err != nil {
		return nil, fmt.Errorf("failed to read class: %w", err)
	}

	// Determine endianness
	var endian binary.ByteOrder = binary.LittleEndian
	if header.Data == 2 {
		endian = binary.BigEndian
	}

	// Read remaining header fields
	if _, err := r.Seek(4, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek: %w", err)
	}

	headerBytes := make([]byte, 64)
	if _, err := r.Read(headerBytes); err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	// Parse class
	if header.Class == 1 {
		elf.Class = "ELF32"
		// Parse 32-bit header
		if len(headerBytes) >= 52 {
			elf.Type = GetELFType(endian.Uint16(headerBytes[16:18]))
			elf.Machine = GetMachine(endian.Uint16(headerBytes[18:20]))
			elf.Version = endian.Uint32(headerBytes[20:24])
			elf.Entry = uint64(endian.Uint32(headerBytes[24:28]))
		}
	} else if header.Class == 2 {
		elf.Class = "ELF64"
		// Parse 64-bit header
		if len(headerBytes) >= 64 {
			elf.Type = GetELFType(endian.Uint16(headerBytes[16:18]))
			elf.Machine = GetMachine(endian.Uint16(headerBytes[18:20]))
			elf.Version = endian.Uint32(headerBytes[20:24])
			elf.Entry = endian.Uint64(headerBytes[24:32])
		}
	} else {
		return nil, fmt.Errorf("invalid ELF class: %d", header.Class)
	}

	// Parse data encoding
	if header.Data == 1 {
		elf.Data = "Little Endian"
	} else if header.Data == 2 {
		elf.Data = "Big Endian"
	} else {
		return nil, fmt.Errorf("invalid data encoding: %d", header.Data)
	}

	// Parse OS/ABI
	elf.OSABI = GetOSABI(headerBytes[7])

	elf.Header = header

	// Parse sections
	if err := parseSections(r, elf, endian); err != nil {
		return nil, fmt.Errorf("failed to parse sections: %w", err)
	}

	// Parse symbols
	if err := parseSymbols(r, elf, endian); err != nil {
		// Symbols are optional, so we don't fail
		_ = err
	}

	return elf, nil
}

// parseSections parses ELF sections
func parseSections(r io.ReadSeeker, elf *ELF, endian binary.ByteOrder) error {
	// Secure: validate section header offset
	if elf.Header.ShOff64 == 0 {
		return nil // No sections
	}

	if _, err := r.Seek(int64(elf.Header.ShOff64), io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to sections: %w", err)
	}

	// Secure: validate section count
	if elf.Header.ShNum == 0 || elf.Header.ShNum > 10000 {
		return fmt.Errorf("invalid section count: %d", elf.Header.ShNum)
	}

	elf.Sections = make([]Section, elf.Header.ShNum)

	// Read section headers
	for i := uint16(0); i < elf.Header.ShNum; i++ {
		var section Section

		if elf.Class == "ELF32" {
			// 32-bit section header (40 bytes)
			shBytes := make([]byte, 40)
			if _, err := r.Read(shBytes); err != nil {
				return fmt.Errorf("failed to read section header: %w", err)
			}

			section.Name = fmt.Sprintf("section_%d", i)
			section.Type = endian.Uint32(shBytes[4:8])
			section.Flags = uint64(endian.Uint32(shBytes[8:12]))
			section.Addr = uint64(endian.Uint32(shBytes[12:16]))
			section.Offset = uint64(endian.Uint32(shBytes[16:20]))
			section.Size = uint64(endian.Uint32(shBytes[20:24]))
			section.Link = endian.Uint32(shBytes[24:28])
			section.Info = endian.Uint32(shBytes[28:32])
			section.AddrAlign = uint64(endian.Uint32(shBytes[32:36]))
			section.EntSize = uint64(endian.Uint32(shBytes[36:40]))
		} else {
			// 64-bit section header (64 bytes)
			shBytes := make([]byte, 64)
			if _, err := r.Read(shBytes); err != nil {
				return fmt.Errorf("failed to read section header: %w", err)
			}

			section.Name = fmt.Sprintf("section_%d", i)
			section.Type = endian.Uint32(shBytes[4:8])
			section.Flags = endian.Uint64(shBytes[8:16])
			section.Addr = endian.Uint64(shBytes[16:24])
			section.Offset = endian.Uint64(shBytes[24:32])
			section.Size = endian.Uint64(shBytes[32:40])
			section.Link = endian.Uint32(shBytes[40:44])
			section.Info = endian.Uint32(shBytes[44:48])
			section.AddrAlign = endian.Uint64(shBytes[48:56])
			section.EntSize = endian.Uint64(shBytes[56:64])
		}

		elf.Sections[i] = section
	}

	// Read string table for section names
	if elf.Header.ShStrndx < elf.Header.ShNum {
		strSection := elf.Sections[elf.Header.ShStrndx]
		if strSection.Offset > 0 && strSection.Size > 0 {
			// Secure: validate string table size
			if strSection.Size > 100*1024*1024 { // 100MB limit
				return fmt.Errorf("string table too large: %d", strSection.Size)
			}

			if _, err := r.Seek(int64(strSection.Offset), io.SeekStart); err != nil {
				return fmt.Errorf("failed to seek to string table: %w", err)
			}

			elf.StringTable = make([]byte, strSection.Size)
			if _, err := r.Read(elf.StringTable); err != nil && err != io.EOF {
				return fmt.Errorf("failed to read string table: %w", err)
			}

			// Update section names
			for i := range elf.Sections {
				nameOffset := endian.Uint32([]byte{
					elf.Sections[i].Name[0],
					elf.Sections[i].Name[1],
					elf.Sections[i].Name[2],
					elf.Sections[i].Name[3],
				})
				if nameOffset < uint32(len(elf.StringTable)) {
					elf.Sections[i].Name = ReadCString(elf.StringTable[nameOffset:])
				}
			}
		}
	}

	return nil
}

// parseSymbols parses ELF symbols
func parseSymbols(r io.ReadSeeker, elf *ELF, endian binary.ByteOrder) error {
	// Find .symtab section
	var symtabSection *Section
	for i := range elf.Sections {
		if elf.Sections[i].Type == 2 { // SHT_SYMTAB
			symtabSection = &elf.Sections[i]
			break
		}
	}

	if symtabSection == nil {
		return nil // No symbol table
	}

	// Secure: validate symbol table size
	if symtabSection.Size == 0 || symtabSection.EntSize == 0 {
		return nil
	}

	if symtabSection.Size > 100*1024*1024 { // 100MB limit
		return fmt.Errorf("symbol table too large: %d", symtabSection.Size)
	}

	if _, err := r.Seek(int64(symtabSection.Offset), io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to symbols: %w", err)
	}

	numSymbols := symtabSection.Size / symtabSection.EntSize
	// Secure: limit number of symbols
	if numSymbols > 100000 {
		numSymbols = 100000
	}

	elf.Symbols = make([]Symbol, numSymbols)

	// Find string table for symbol names
	var strtabSection *Section
	for i := range elf.Sections {
		if elf.Sections[i].Type == 3 && elf.Sections[i].Link == uint32(symtabSection.Info) { // SHT_STRTAB
			strtabSection = &elf.Sections[i]
			break
		}
	}

	var strtab []byte
	if strtabSection != nil && strtabSection.Size > 0 {
		// Secure: validate string table size
		if strtabSection.Size > 100*1024*1024 {
			return fmt.Errorf("string table too large: %d", strtabSection.Size)
		}

		if _, err := r.Seek(int64(strtabSection.Offset), io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek to string table: %w", err)
		}

		strtab = make([]byte, strtabSection.Size)
		if _, err := r.Read(strtab); err != nil && err != io.EOF {
			return fmt.Errorf("failed to read string table: %w", err)
		}
	}

	// Read symbols
	for i := uint64(0); i < numSymbols; i++ {
		var symbol Symbol

		if elf.Class == "ELF32" {
			// 32-bit symbol (16 bytes)
			symBytes := make([]byte, 16)
			if _, err := r.Read(symBytes); err != nil {
				break
			}

			nameIdx := endian.Uint32(symBytes[0:4])
			if nameIdx < uint32(len(strtab)) {
				symbol.Name = ReadCString(strtab[nameIdx:])
			}
			symbol.Value = uint64(endian.Uint32(symBytes[4:8]))
			symbol.Size = uint64(endian.Uint32(symBytes[8:12]))
			symbol.Info = symBytes[12]
			symbol.Other = symBytes[13]
			symbol.Shndx = endian.Uint16(symBytes[14:16])
		} else {
			// 64-bit symbol (24 bytes)
			symBytes := make([]byte, 24)
			if _, err := r.Read(symBytes); err != nil {
				break
			}

			nameIdx := endian.Uint32(symBytes[0:4])
			if nameIdx < uint32(len(strtab)) {
				symbol.Name = ReadCString(strtab[nameIdx:])
			}
			symbol.Info = symBytes[4]
			symbol.Other = symBytes[5]
			symbol.Shndx = endian.Uint16(symBytes[6:8])
			symbol.Value = endian.Uint64(symBytes[8:16])
			symbol.Size = endian.Uint64(symBytes[16:24])
		}

		symbol.Type = GetSymbolType(symbol.Info & 0x0f)
		symbol.Binding = GetSymbolBinding((symbol.Info >> 4) & 0x0f)

		elf.Symbols[i] = symbol
	}

	return nil
}

// ReadCString reads a null-terminated string
func ReadCString(data []byte) string {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}

// GetELFType returns ELF type name
func GetELFType(t uint16) string {
	types := map[uint16]string{
		0: "ET_NONE",
		1: "ET_REL",
		2: "ET_EXEC",
		3: "ET_DYN",
		4: "ET_CORE",
	}
	if name, ok := types[t]; ok {
		return name
	}
	return fmt.Sprintf("ET_UNKNOWN(%d)", t)
}

// GetMachine returns machine type name
func GetMachine(m uint16) string {
	machines := map[uint16]string{
		0x00: "EM_NONE",
		0x02: "EM_SPARC",
		0x03: "EM_386",
		0x08: "EM_MIPS",
		0x14: "EM_PPC",
		0x28: "EM_ARM",
		0x3E: "EM_X86_64",
		0xB7: "EM_AARCH64",
	}
	if name, ok := machines[m]; ok {
		return name
	}
	return fmt.Sprintf("EM_UNKNOWN(0x%02x)", m)
}

// GetOSABI returns OS/ABI name
func GetOSABI(abi byte) string {
	abis := map[byte]string{
		0x00: "ELFOSABI_NONE",
		0x01: "ELFOSABI_HPUX",
		0x02: "ELFOSABI_NETBSD",
		0x03: "ELFOSABI_LINUX",
		0x06: "ELFOSABI_SOLARIS",
		0x07: "ELFOSABI_AIX",
		0x08: "ELFOSABI_IRIX",
		0x09: "ELFOSABI_FREEBSD",
		0x0C: "ELFOSABI_OPENBSD",
	}
	if name, ok := abis[abi]; ok {
		return name
	}
	return fmt.Sprintf("ELFOSABI_UNKNOWN(0x%02x)", abi)
}

// GetSymbolType returns symbol type name
func GetSymbolType(t byte) string {
	types := map[byte]string{
		0:  "STT_NOTYPE",
		1:  "STT_OBJECT",
		2:  "STT_FUNC",
		3:  "STT_SECTION",
		4:  "STT_FILE",
		10: "STT_COMMON",
		13: "STT_TLS",
	}
	if name, ok := types[t]; ok {
		return name
	}
	return fmt.Sprintf("STT_UNKNOWN(%d)", t)
}

// GetSymbolBinding returns symbol binding name
func GetSymbolBinding(b byte) string {
	bindings := map[byte]string{
		0: "STB_LOCAL",
		1: "STB_GLOBAL",
		2: "STB_WEAK",
	}
	if name, ok := bindings[b]; ok {
		return name
	}
	return fmt.Sprintf("STB_UNKNOWN(%d)", b)
}
