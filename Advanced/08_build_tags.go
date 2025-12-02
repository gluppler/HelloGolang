package main

import (
	"fmt"
	"runtime"
)

// Build Tags demonstrates build tags and conditional compilation

func main() {
	buildTagsExample()
	platformSpecific()
	featureFlags()
}

// buildTagsExample demonstrates basic build tags
// Build tags are comments at the top of files:
// //go:build tag_name
// or
// // +build tag_name

func buildTagsExample() {
	fmt.Println("Build tags example")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)

	// Example build tags:
	// //go:build linux
	// //go:build windows
	// //go:build darwin
	// //go:build !windows
	// //go:build linux || darwin
	// //go:build debug
	// //go:build !production
}

// platformSpecific demonstrates platform-specific code
func platformSpecific() {
	// Use build tags for platform-specific implementations
	// File: file_linux.go
	// //go:build linux
	// package main
	// func getOS() string { return "linux" }

	// File: file_windows.go
	// //go:build windows
	// package main
	// func getOS() string { return "windows" }

	// File: file_darwin.go
	// //go:build darwin
	// package main
	// func getOS() string { return "darwin" }

	fmt.Printf("Current OS: %s\n", runtime.GOOS)
}

// featureFlags demonstrates feature flags with build tags
func featureFlags() {
	// Enable features with build tags
	// Build with: go build -tags debug
	// File: debug.go
	// //go:build debug
	// package main
	// const DebugMode = true

	// File: release.go
	// //go:build !debug
	// package main
	// const DebugMode = false

	fmt.Println("Feature flags example")
	fmt.Println("Use build tags to enable/disable features")
}

// Build tag examples:
// 1. Platform-specific: //go:build linux
// 2. Architecture-specific: //go:build amd64
// 3. Feature flags: //go:build debug
// 4. Test files: //go:build test
// 5. Integration tests: //go:build integration
// 6. Exclude from build: //go:build ignore

// Usage:
// go build -tags debug
// go build -tags "linux debug"
// go test -tags integration
