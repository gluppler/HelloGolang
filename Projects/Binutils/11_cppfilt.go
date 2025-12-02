package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// C++filt - Demangle C++ symbols (GNU c++filt equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <mangled-name>...\n", os.Args[0])
		os.Exit(1)
	}

	for _, name := range os.Args[1:] {
		demangled := demangle(name)
		fmt.Println(demangled)
	}
}

// demangle demangles C++ symbol name
func demangle(mangled string) string {
	// Secure: validate input length
	if len(mangled) > 10000 {
		return mangled
	}

	// Remove leading underscore if present
	if len(mangled) > 0 && mangled[0] == '_' {
		mangled = mangled[1:]
	}

	// Check if it's a mangled name
	if !strings.HasPrefix(mangled, "_Z") {
		return mangled // Not mangled
	}

	// Simple demangling (full implementation would parse Itanium ABI)
	// This is a simplified version
	if strings.HasPrefix(mangled, "_ZN") {
		// Mangled class member or namespace
		return demangleClassMember(mangled)
	}

	if strings.HasPrefix(mangled, "_Z") {
		// Mangled function
		return demangleFunction(mangled)
	}

	return mangled
}

// demangleFunction demangles function name
func demangleFunction(mangled string) string {
	// Secure: validate format
	if len(mangled) < 3 {
		return mangled
	}

	// Skip _Z
	mangled = mangled[2:]

	// Simple pattern matching
	// Full implementation would properly parse Itanium ABI mangling
	re := regexp.MustCompile(`^(\d+)([a-zA-Z_][a-zA-Z0-9_]*)`)
	matches := re.FindStringSubmatch(mangled)

	if len(matches) >= 3 {
		nameLen := parseIntSafe(matches[1])
		if nameLen > 0 && nameLen < len(matches[2]) {
			name := matches[2][:nameLen]
			// Secure: validate name
			if isValidIdentifier(name) {
				return name
			}
		}
	}

	return "_Z" + mangled
}

// demangleClassMember demangles class member name
func demangleClassMember(mangled string) string {
	// Secure: validate format
	if len(mangled) < 4 {
		return mangled
	}

	// Skip _ZN
	mangled = mangled[3:]

	// Simple pattern matching for class::member
	parts := []string{}

	for len(mangled) > 0 {
		re := regexp.MustCompile(`^(\d+)([a-zA-Z_][a-zA-Z0-9_]*)`)
		matches := re.FindStringSubmatch(mangled)

		if len(matches) >= 3 {
			nameLen := parseIntSafe(matches[1])
			if nameLen > 0 && nameLen <= len(matches[2]) {
				name := matches[2][:nameLen]
				if isValidIdentifier(name) {
					parts = append(parts, name)
					mangled = mangled[len(matches[0]):]
					continue
				}
			}
		}

		break
	}

	if len(parts) > 0 {
		return strings.Join(parts, "::")
	}

	return "_ZN" + mangled
}

// parseIntSafe safely parses integer
func parseIntSafe(s string) int {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
		// Secure: prevent overflow
		if result > 10000 {
			return 0
		}
	}
	return result
}

// isValidIdentifier checks if string is valid identifier
func isValidIdentifier(s string) bool {
	if len(s) == 0 || len(s) > 1000 {
		return false
	}

	// Check first character
	if !((s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z') || s[0] == '_') {
		return false
	}

	// Check remaining characters
	for i := 1; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}

	return true
}
