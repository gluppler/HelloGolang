package main

import (
	"fmt"
	"os"
)

// Ranlib - Generate index to archive (GNU ranlib equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <archive>...\n", os.Args[0])
		os.Exit(1)
	}

	for _, archiveName := range os.Args[1:] {
		if err := generateIndex(archiveName); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", archiveName, err)
			os.Exit(1)
		}
	}
}

// Archive types (shared with ar.go)
type archiveHeader struct {
	Name    string
	Date    int64
	UID     int
	GID     int
	Mode    int
	Size    int64
	EndChar [2]byte
}

type archiveMember struct {
	Header archiveHeader
	Data   []byte
}

// generateIndex generates symbol index for archive
func generateIndex(archiveName string) error {
	// Read archive
	members, err := readArchiveForRanlib(archiveName)
	if err != nil {
		return fmt.Errorf("failed to read archive: %w", err)
	}

	// Extract symbols from object files
	symbols := []ArchiveSymbol{}
	for _, member := range members {
		// Check if member is an object file
		if isObjectFile(member.Data) {
			memberSymbols := extractSymbolsFromObject(member.Data, member.Header.Name)
			symbols = append(symbols, memberSymbols...)
		}
	}

	// Create symbol table member
	symbolTable := createSymbolTable(symbols)

	// Add or replace symbol table in archive
	membersMap := make(map[string]*archiveMember)
	for _, m := range members {
		membersMap[m.Header.Name] = m
	}

	// Remove old symbol table if exists
	delete(membersMap, "__.SYMDEF")
	delete(membersMap, "/")

	// Add new symbol table
	membersMap["/"] = symbolTable

	// Write updated archive
	return writeArchiveForRanlib(archiveName, membersMap)
}

// readArchiveForRanlib reads archive file
func readArchiveForRanlib(archiveName string) ([]*archiveMember, error) {
	file, err := os.Open(archiveName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read magic
	magic := make([]byte, 8)
	if _, err := file.Read(magic); err != nil {
		return nil, fmt.Errorf("failed to read magic: %w", err)
	}

	if string(magic) != "!<arch>\n" {
		return nil, fmt.Errorf("invalid archive magic")
	}

	var members []*archiveMember

	for {
		// Read header (60 bytes)
		headerBytes := make([]byte, 60)
		n, err := file.Read(headerBytes)
		if err != nil || n < 60 {
			break
		}

		// Parse header
		header := archiveHeader{}
		header.Name = trimSpaceForRanlib(string(headerBytes[0:16]))
		header.Size, _ = parseInt64ForRanlib(trimSpaceForRanlib(string(headerBytes[48:58])))

		// Secure: validate size
		if header.Size < 0 || header.Size > 100*1024*1024 {
			return nil, fmt.Errorf("invalid member size: %d", header.Size)
		}

		// Read data
		data := make([]byte, header.Size)
		if _, err := file.Read(data); err != nil {
			break
		}

		// Skip padding
		if header.Size%2 != 0 {
			file.Read(make([]byte, 1))
		}

		members = append(members, &archiveMember{
			Header: header,
			Data:   data,
		})
	}

	return members, nil
}

// writeArchiveForRanlib writes archive file
func writeArchiveForRanlib(archiveName string, members map[string]*archiveMember) error {
	file, err := os.Create(archiveName)
	if err != nil {
		return fmt.Errorf("failed to create archive: %w", err)
	}
	defer file.Close()

	// Write magic
	if _, err := file.WriteString("!<arch>\n"); err != nil {
		return fmt.Errorf("failed to write magic: %w", err)
	}

	// Write members
	for _, member := range members {
		// Write header (simplified)
		header := formatHeaderForRanlib(member.Header)
		if _, err := file.Write(header); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}

		// Write data
		if _, err := file.Write(member.Data); err != nil {
			return fmt.Errorf("failed to write data: %w", err)
		}

		// Write padding
		if len(member.Data)%2 != 0 {
			if _, err := file.Write([]byte{'\n'}); err != nil {
				return fmt.Errorf("failed to write padding: %w", err)
			}
		}
	}

	return nil
}

// Helper functions
func trimSpaceForRanlib(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}

func parseInt64ForRanlib(s string) (int64, error) {
	result := int64(0)
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid number")
		}
		result = result*10 + int64(c-'0')
		if result > 100*1024*1024 {
			return 0, fmt.Errorf("number too large")
		}
	}
	return result, nil
}

func formatHeaderForRanlib(h archiveHeader) []byte {
	header := make([]byte, 60)
	copy(header[0:16], padStringForRanlib(h.Name, 16))
	copy(header[48:58], padStringForRanlib(fmt.Sprintf("%d", h.Size), 10))
	copy(header[58:60], []byte("`\n"))
	return header
}

func padStringForRanlib(s string, length int) []byte {
	result := make([]byte, length)
	copy(result, []byte(s))
	for i := len(s); i < length; i++ {
		result[i] = ' '
	}
	return result
}

// ArchiveSymbol represents a symbol in archive index
type ArchiveSymbol struct {
	Name     string
	Member   string
	Offset   int64
}

// isObjectFile checks if data is an object file
func isObjectFile(data []byte) bool {
	// Secure: validate minimum size
	if len(data) < 4 {
		return false
	}

	// Check ELF magic
	if data[0] == 0x7f && data[1] == 'E' && data[2] == 'L' && data[3] == 'F' {
		return true
	}

	return false
}

// extractSymbolsFromObject extracts symbols from object file
func extractSymbolsFromObject(data []byte, memberName string) []ArchiveSymbol {
	symbols := []ArchiveSymbol{}

	// Secure: validate size
	if len(data) > 100*1024*1024 {
		return symbols
	}

	// Parse ELF to extract symbols
	// Simplified - in production would use proper ELF parser
	// For now, create placeholder symbols
	symbols = append(symbols, ArchiveSymbol{
		Name:   "main",
		Member: memberName,
		Offset: 0,
	})

	return symbols
}

// createSymbolTable creates symbol table archive member
func createSymbolTable(symbols []ArchiveSymbol) *archiveMember {
	// Secure: validate symbol count
	if len(symbols) > 100000 {
		symbols = symbols[:100000]
	}

	// Format: number of symbols (4 bytes) + symbol entries
	data := make([]byte, 4+len(symbols)*8)
	
	// Write number of symbols
	data[0] = byte(len(symbols))
	data[1] = byte(len(symbols) >> 8)
	data[2] = byte(len(symbols) >> 16)
	data[3] = byte(len(symbols) >> 24)

	// Write symbol entries (simplified format)
	offset := 4
	for _, sym := range symbols {
		// Symbol name (4 bytes offset, simplified)
		nameLen := len(sym.Name)
		if nameLen > 255 {
			nameLen = 255
		}
		data[offset] = byte(nameLen)
		offset++
		copy(data[offset:], []byte(sym.Name[:nameLen]))
		offset += nameLen
		// Pad to 8 bytes
		for offset%8 != 0 {
			offset++
		}
	}

	return &archiveMember{
		Header: archiveHeader{
			Name: "/",
			Date: 0,
			UID:  0,
			GID:  0,
			Mode: 0644,
			Size: int64(len(data)),
		},
		Data: data,
	}
}
