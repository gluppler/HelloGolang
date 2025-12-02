package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Standard Library demonstrates common standard library packages

func main() {
	fmtPackage()
	stringsPackage()
	bytesPackage()
	timePackage()
	jsonPackage()
	ioPackage()
	osPackage()
}

// fmtPackage demonstrates fmt package
func fmtPackage() {
	// Print functions
	fmt.Print("Print: ")
	fmt.Println("Hello")

	// Formatted print
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// String formatting
	s := fmt.Sprintf("Formatted: %s (%d)", name, age)
	fmt.Println(s)

	// Scanning (reading input)
	// fmt.Scan, fmt.Scanf, fmt.Scanln

	// Error formatting
	err := fmt.Errorf("error occurred: %s", "invalid input")
	fmt.Printf("Error: %v\n", err)
}

// stringsPackage demonstrates strings package
func stringsPackage() {
	text := "  Hello, World!  "

	// Manipulation
	trimmed := strings.TrimSpace(text)
	fmt.Printf("Trimmed: '%s'\n", trimmed)

	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)
	fmt.Printf("Upper: %s\n", upper)
	fmt.Printf("Lower: %s\n", lower)

	// Searching
	contains := strings.Contains(text, "World")
	fmt.Printf("Contains 'World': %t\n", contains)

	index := strings.Index(text, "World")
	fmt.Printf("Index of 'World': %d\n", index)

	// Splitting and joining
	parts := strings.Split("a,b,c", ",")
	fmt.Printf("Split: %v\n", parts)

	joined := strings.Join(parts, "-")
	fmt.Printf("Joined: %s\n", joined)

	// Replacement
	replaced := strings.Replace(text, "World", "Go", -1)
	fmt.Printf("Replaced: %s\n", replaced)

	// Prefix and suffix
	hasPrefix := strings.HasPrefix(text, "Hello")
	hasSuffix := strings.HasSuffix(text, "!")
	fmt.Printf("Has prefix 'Hello': %t\n", hasPrefix)
	fmt.Printf("Has suffix '!': %t\n", hasSuffix)

	// Builder for efficient string building
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	result := builder.String()
	fmt.Printf("Builder result: %s\n", result)
}

// bytesPackage demonstrates bytes package
func bytesPackage() {
	// Similar to strings but for []byte
	data := []byte("Hello, World!")

	// Contains
	contains := bytes.Contains(data, []byte("World"))
	fmt.Printf("Contains 'World': %t\n", contains)

	// Split
	parts := bytes.Split(data, []byte(","))
	fmt.Printf("Split into %d parts\n", len(parts))

	// Join
	joined := bytes.Join(parts, []byte("-"))
	fmt.Printf("Joined: %s\n", string(joined))

	// Buffer for efficient byte operations
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteByte(' ')
	buf.Write([]byte("World"))
	fmt.Printf("Buffer: %s\n", buf.String())

	// Compare
	data1 := []byte("abc")
	data2 := []byte("abc")
	equal := bytes.Equal(data1, data2)
	fmt.Printf("Bytes equal: %t\n", equal)
}

// timePackage demonstrates time package
func timePackage() {
	// Current time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	// Formatting
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Printf("Formatted: %s\n", formatted)

	// Parsing
	parsed, err := time.Parse("2006-01-02", "2024-01-15")
	if err == nil {
		fmt.Printf("Parsed: %v\n", parsed)
	}

	// Duration
	duration := 2 * time.Hour
	fmt.Printf("Duration: %v\n", duration)

	// Add time
	future := now.Add(24 * time.Hour)
	fmt.Printf("Future: %v\n", future)

	// Sleep
	fmt.Println("Sleeping for 100ms...")
	time.Sleep(100 * time.Millisecond)

	// Ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Timer
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	// Unix timestamp
	unix := now.Unix()
	fmt.Printf("Unix timestamp: %d\n", unix)

	// Time zones
	loc, _ := time.LoadLocation("America/New_York")
	nyTime := now.In(loc)
	fmt.Printf("NY time: %v\n", nyTime)
}

// jsonPackage demonstrates encoding/json package
func jsonPackage() {
	// Struct to JSON
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person := Person{Name: "Alice", Age: 30}

	jsonData, err := json.Marshal(person)
	if err == nil {
		fmt.Printf("JSON: %s\n", string(jsonData))
	}

	// Pretty print
	prettyJSON, _ := json.MarshalIndent(person, "", "  ")
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))

	// JSON to struct
	var decoded Person
	err = json.Unmarshal(jsonData, &decoded)
	if err == nil {
		fmt.Printf("Decoded: %+v\n", decoded)
	}

	// Map to JSON
	data := map[string]interface{}{
		"name": "Bob",
		"age":  25,
	}
	jsonMap, _ := json.Marshal(data)
	fmt.Printf("Map JSON: %s\n", string(jsonMap))
}

// ioPackage demonstrates io package
func ioPackage() {
	// Copy
	source := strings.NewReader("Hello, World!")
	dest := &bytes.Buffer{}

	n, err := io.Copy(dest, source)
	if err == nil {
		fmt.Printf("Copied %d bytes: %s\n", n, dest.String())
	}

	// ReadAll
	reader := strings.NewReader("Read all this")
	data, _ := io.ReadAll(reader)
	fmt.Printf("Read all: %s\n", string(data))

	// MultiReader
	r1 := strings.NewReader("First ")
	r2 := strings.NewReader("Second ")
	r3 := strings.NewReader("Third")
	multi := io.MultiReader(r1, r2, r3)
	all, _ := io.ReadAll(multi)
	fmt.Printf("MultiReader: %s\n", string(all))

	// TeeReader
	original := strings.NewReader("Tee this")
	var buf bytes.Buffer
	tee := io.TeeReader(original, &buf)
	teeData, _ := io.ReadAll(tee)
	fmt.Printf("TeeReader: %s (also written to buffer: %s)\n", string(teeData), buf.String())

	// LimitReader
	limit := io.LimitReader(strings.NewReader("This is a long string"), 10)
	limited, _ := io.ReadAll(limit)
	fmt.Printf("Limited to 10 bytes: %s\n", string(limited))
}

// osPackage demonstrates os package
func osPackage() {
	// Environment variables
	path := os.Getenv("PATH")
	fmt.Printf("PATH length: %d\n", len(path))

	// Set environment variable (for current process)
	os.Setenv("TEST_VAR", "test_value")
	value := os.Getenv("TEST_VAR")
	fmt.Printf("TEST_VAR: %s\n", value)

	// Working directory
	wd, _ := os.Getwd()
	fmt.Printf("Working directory: %s\n", wd)

	// File operations
	// Create file
	file, err := os.Create("temp.txt")
	if err == nil {
		file.WriteString("Hello, File!")
		file.Close()

		// Read file
		data, _ := os.ReadFile("temp.txt")
		fmt.Printf("File content: %s\n", string(data))

		// Remove file
		os.Remove("temp.txt")
	}

	// Process info
	pid := os.Getpid()
	fmt.Printf("Process ID: %d\n", pid)

	// Exit
	// os.Exit(0)  // Don't actually exit in example
}

// Additional standard library packages to explore:
// - net/http: HTTP client and server
// - net: Network I/O
// - context: Context for cancellation and timeouts
// - sync: Synchronization primitives
// - crypto: Cryptographic functions
// - encoding: Various encodings (base64, hex, etc.)
// - path/filepath: File path manipulation
// - flag: Command-line flag parsing
// - log: Logging
// - testing: Testing framework
// - sort: Sorting
// - math: Mathematical functions
// - regexp: Regular expressions
// - url: URL parsing

func standardLibrarySummary() {
	fmt.Println("Standard Library Summary:")
	fmt.Println("  - fmt: Formatting and I/O")
	fmt.Println("  - strings: String manipulation")
	fmt.Println("  - bytes: Byte slice operations")
	fmt.Println("  - time: Time and date operations")
	fmt.Println("  - encoding/json: JSON encoding/decoding")
	fmt.Println("  - io: I/O primitives")
	fmt.Println("  - os: OS interface")
	fmt.Println("  - net/http: HTTP client/server")
	fmt.Println("  - context: Context management")
	fmt.Println("  - sync: Synchronization")
	fmt.Println("  - And many more...")
}
