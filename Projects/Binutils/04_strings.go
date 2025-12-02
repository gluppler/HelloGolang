package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Strings - Print printable strings in files (GNU strings equivalent)

func main() {
	minLen := 4
	args := os.Args[1:]

	// Parse arguments
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-n length] <file>...\n", os.Args[0])
		os.Exit(1)
	}

	files := []string{}
	for i := 0; i < len(args); i++ {
		if args[i] == "-n" && i+1 < len(args) {
			// Secure: validate length
			if n, err := parseInt(args[i+1]); err == nil && n > 0 && n <= 100 {
				minLen = n
			}
			i++
		} else {
			files = append(files, args[i])
		}
	}

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No files specified\n")
		os.Exit(1)
	}

	for _, filename := range files {
		if err := extractStrings(filename, minLen); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", filename, err)
		}
	}
}

// extractStrings extracts printable strings from file
func extractStrings(filename string, minLen int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Secure: limit file size
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	// Secure: prevent reading extremely large files
	maxSize := int64(100 * 1024 * 1024) // 100MB
	if stat.Size() > maxSize {
		return fmt.Errorf("file too large: %d bytes", stat.Size())
	}

	reader := bufio.NewReader(file)
	currentString := []byte{}

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			if len(currentString) >= minLen {
				fmt.Printf("%s\n", string(currentString))
			}
			break
		}

		// Check if byte is printable ASCII
		if b >= 32 && b <= 126 {
			currentString = append(currentString, b)
		} else {
			if len(currentString) >= minLen {
				fmt.Printf("%s\n", string(currentString))
			}
			currentString = currentString[:0]
		}

		// Secure: limit string length
		if len(currentString) > 10000 {
			currentString = currentString[:0]
		}
	}

	return nil
}

// parseInt safely parses integer
func parseInt(s string) (int, error) {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid number")
		}
		result = result*10 + int(c-'0')
		// Secure: prevent integer overflow
		if result > 100 {
			return 0, fmt.Errorf("number too large")
		}
	}
	return result, nil
}
