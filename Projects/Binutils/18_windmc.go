package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Windmc - Windows message compiler (GNU windmc equivalent)

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

	if err := compileMessages(inputFile, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// compileMessages compiles message file
func compileMessages(inputFile, outputFile string) error {
	// Secure: validate filenames
	if len(inputFile) > 255 || len(outputFile) > 255 {
		return fmt.Errorf("filename too long")
	}

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	// Secure: validate file size
	if len(data) > 10*1024*1024 {
		return fmt.Errorf("input file too large")
	}

	// Parse message definitions
	messages := parseMessageFile(string(data))

	// Generate binary message file
	binaryData := generateMessageBinary(messages)

	// Secure: validate output size
	if len(binaryData) > 50*1024*1024 {
		return fmt.Errorf("output file too large")
	}

	return os.WriteFile(outputFile, binaryData, 0644)
}

// Message represents a Windows message definition
type Message struct {
	ID      uint32
	Language uint16
	Text    string
}

// parseMessageFile parses .mc message file
func parseMessageFile(content string) []Message {
	messages := []Message{}
	lines := strings.Split(content, "\n")

	// Secure: limit number of lines
	if len(lines) > 100000 {
		lines = lines[:100000]
	}

	var currentID uint32 = 1
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		// Secure: validate line length
		if len(line) > 10000 {
			continue
		}

		// Parse message definition (simplified)
		if strings.HasPrefix(line, "MessageId=") {
			// Extract ID
			parts := strings.Split(line, "=")
			if len(parts) >= 2 {
				if id, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 32); err == nil {
					currentID = uint32(id)
				}
			}
		} else if strings.HasPrefix(line, "Language=") {
			// Language definition
			continue
		} else if line != "" && !strings.HasPrefix(line, ".") {
			// Message text
			messages = append(messages, Message{
				ID:       currentID,
				Language: 0x0409, // English (US)
				Text:     line,
			})
			currentID++
		}
	}

	// Secure: limit number of messages
	if len(messages) > 10000 {
		messages = messages[:10000]
	}

	return messages
}

// generateMessageBinary generates binary message file
func generateMessageBinary(messages []Message) []byte {
	// Binary format: header + message entries
	headerSize := 16
	entrySize := 8 + 256 // ID (4) + Language (2) + Padding (2) + Text (256 max)

	data := make([]byte, headerSize+len(messages)*entrySize)

	// Write header
	// Number of messages
	data[0] = byte(len(messages))
	data[1] = byte(len(messages) >> 8)
	data[2] = byte(len(messages) >> 16)
	data[3] = byte(len(messages) >> 24)

	// Write messages
	offset := headerSize
	for _, msg := range messages {
		// Message ID
		data[offset] = byte(msg.ID)
		data[offset+1] = byte(msg.ID >> 8)
		data[offset+2] = byte(msg.ID >> 16)
		data[offset+3] = byte(msg.ID >> 24)

		// Language
		data[offset+4] = byte(msg.Language)
		data[offset+5] = byte(msg.Language >> 8)

		// Text (truncate to 255 bytes)
		textBytes := []byte(msg.Text)
		if len(textBytes) > 255 {
			textBytes = textBytes[:255]
		}
		copy(data[offset+8:], textBytes)

		offset += entrySize
	}

	return data[:offset]
}
