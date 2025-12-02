package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Ar - Archive utility (GNU ar equivalent)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <operation> <archive> [files...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Operations: r (replace), t (table), x (extract), d (delete)\n")
		os.Exit(1)
	}

	operation := os.Args[1]
	archiveName := os.Args[2]

	switch operation {
	case "r":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Error: no files specified\n")
			os.Exit(1)
		}
		if err := createArchive(archiveName, os.Args[3:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "t":
		if err := listArchive(archiveName); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "x":
		files := []string{}
		if len(os.Args) > 3 {
			files = os.Args[3:]
		}
		if err := extractArchive(archiveName, files); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "d":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Error: no files specified\n")
			os.Exit(1)
		}
		if err := deleteFromArchive(archiveName, os.Args[3:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown operation: %s\n", operation)
		os.Exit(1)
	}
}

// ArchiveHeader represents an archive member header
type ArchiveHeader struct {
	Name    string
	Date    int64
	UID     int
	GID     int
	Mode    int
	Size    int64
	EndChar [2]byte
}

// createArchive creates or updates an archive
func createArchive(archiveName string, files []string) error {
	// Read existing archive if it exists
	members := map[string]*ArchiveMember{}
	if _, err := os.Stat(archiveName); err == nil {
		existing, err := readArchive(archiveName)
		if err == nil {
			for _, m := range existing {
				members[m.Header.Name] = m
			}
		}
	}

	// Add or replace files
	for _, filename := range files {
		// Secure: validate filename
		if len(filename) > 255 {
			return fmt.Errorf("filename too long: %s", filename)
		}

		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", filename, err)
		}

		// Secure: limit file size
		if len(data) > 100*1024*1024 { // 100MB
			return fmt.Errorf("file too large: %s", filename)
		}

		stat, err := os.Stat(filename)
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", filename, err)
		}

		members[filename] = &ArchiveMember{
			Header: ArchiveHeader{
				Name: filename,
				Date: stat.ModTime().Unix(),
				UID:  0,
				GID:  0,
				Mode: 0644,
				Size: int64(len(data)),
			},
			Data: data,
		}
	}

	// Write archive
	return writeArchive(archiveName, members)
}

// ArchiveMember represents an archive member
type ArchiveMember struct {
	Header ArchiveHeader
	Data   []byte
}

// readArchive reads an archive file
func readArchive(archiveName string) ([]*ArchiveMember, error) {
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

	var members []*ArchiveMember

	for {
		// Read header (60 bytes)
		headerBytes := make([]byte, 60)
		n, err := file.Read(headerBytes)
		if err == io.EOF || n < 60 {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read header: %w", err)
		}

		// Parse header
		header := ArchiveHeader{}
		header.Name = trimSpace(string(headerBytes[0:16]))
		header.Date, _ = parseInt64(trimSpace(string(headerBytes[16:28])))
		header.UID, _ = parseIntSafe(trimSpace(string(headerBytes[28:34])))
		header.GID, _ = parseIntSafe(trimSpace(string(headerBytes[34:40])))
		header.Mode, _ = parseIntOctal(trimSpace(string(headerBytes[40:48])))
		header.Size, _ = parseInt64(trimSpace(string(headerBytes[48:58])))
		copy(header.EndChar[:], headerBytes[58:60])

		// Secure: validate size
		if header.Size < 0 || header.Size > 100*1024*1024 {
			return nil, fmt.Errorf("invalid member size: %d", header.Size)
		}

		// Read data
		data := make([]byte, header.Size)
		if _, err := io.ReadFull(file, data); err != nil {
			return nil, fmt.Errorf("failed to read member data: %w", err)
		}

		// Skip padding (even byte boundary)
		if header.Size%2 != 0 {
			file.Read(make([]byte, 1))
		}

		members = append(members, &ArchiveMember{
			Header: header,
			Data:   data,
		})
	}

	return members, nil
}

// writeArchive writes an archive file
func writeArchive(archiveName string, members map[string]*ArchiveMember) error {
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
		// Write header
		header := formatHeader(member.Header)
		if _, err := file.Write(header); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}

		// Write data
		if _, err := file.Write(member.Data); err != nil {
			return fmt.Errorf("failed to write data: %w", err)
		}

		// Write padding if needed
		if len(member.Data)%2 != 0 {
			if _, err := file.Write([]byte{'\n'}); err != nil {
				return fmt.Errorf("failed to write padding: %w", err)
			}
		}
	}

	return nil
}

// formatHeader formats archive header
func formatHeader(h ArchiveHeader) []byte {
	header := make([]byte, 60)
	copy(header[0:16], padString(h.Name, 16))
	copy(header[16:28], padString(fmt.Sprintf("%d", h.Date), 12))
	copy(header[28:34], padString(fmt.Sprintf("%d", h.UID), 6))
	copy(header[34:40], padString(fmt.Sprintf("%d", h.GID), 6))
	copy(header[40:48], padString(fmt.Sprintf("%o", h.Mode), 8))
	copy(header[48:58], padString(fmt.Sprintf("%d", h.Size), 10))
	copy(header[58:60], []byte("`\n"))
	return header
}

// padString pads string to specified length
func padString(s string, length int) []byte {
	result := make([]byte, length)
	copy(result, []byte(s))
	for i := len(s); i < length; i++ {
		result[i] = ' '
	}
	return result
}

// trimSpace trims whitespace
func trimSpace(s string) string {
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

// parseInt64 safely parses int64
func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// parseIntSafe safely parses integer
func parseIntSafe(s string) (int, error) {
	return strconv.Atoi(s)
}

// parseIntOctal safely parses octal integer
func parseIntOctal(s string) (int, error) {
	return strconv.Atoi(s)
}

// listArchive lists archive contents
func listArchive(archiveName string) error {
	members, err := readArchive(archiveName)
	if err != nil {
		return err
	}

	for _, member := range members {
		fmt.Println(member.Header.Name)
	}

	return nil
}

// extractArchive extracts files from archive
func extractArchive(archiveName string, files []string) error {
	members, err := readArchive(archiveName)
	if err != nil {
		return err
	}

	extractAll := len(files) == 0

	for _, member := range members {
		if extractAll || contains(files, member.Header.Name) {
			// Secure: validate filename
			if len(member.Header.Name) == 0 || len(member.Header.Name) > 255 {
				continue
			}

			// Secure: prevent path traversal
			if containsPathTraversal(member.Header.Name) {
				continue
			}

			if err := os.WriteFile(member.Header.Name, member.Data, os.FileMode(member.Header.Mode)); err != nil {
				return fmt.Errorf("failed to write %s: %w", member.Header.Name, err)
			}
		}
	}

	return nil
}

// deleteFromArchive deletes files from archive
func deleteFromArchive(archiveName string, files []string) error {
	members, err := readArchive(archiveName)
	if err != nil {
		return err
	}

	// Create map of files to delete
	deleteMap := make(map[string]bool)
	for _, f := range files {
		deleteMap[f] = true
	}

	// Filter out deleted members
	newMembers := make(map[string]*ArchiveMember)
	for _, member := range members {
		if !deleteMap[member.Header.Name] {
			newMembers[member.Header.Name] = member
		}
	}

	return writeArchive(archiveName, newMembers)
}

// contains checks if slice contains string
func contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// containsPathTraversal checks for path traversal attacks
func containsPathTraversal(path string) bool {
	// Secure: check for dangerous patterns
	dangerous := []string{"../", "..\\", "/", "\\"}
	for _, pattern := range dangerous {
		if bytes.Contains([]byte(path), []byte(pattern)) {
			return true
		}
	}
	return false
}
