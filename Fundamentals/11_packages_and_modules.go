package main

import (
	"fmt"
	"math"
	"strings"
)

// Packages and Modules demonstrates package organization and module usage

func main() {
	packageBasics()
	packageOrganization()
	visibilityRules()
	importAliases()
	blankImports()
}

// packageBasics demonstrates basic package concepts
func packageBasics() {
	// Using standard library packages
	fmt.Println("Using fmt package")

	// Math operations
	result := math.Sqrt(16)
	fmt.Printf("Square root of 16: %.2f\n", result)

	// String operations
	text := "Hello, World!"
	upper := strings.ToUpper(text)
	fmt.Printf("Uppercase: %s\n", upper)

	// Package functions
	fmt.Printf("Custom function: %d\n", addNumbers(10, 20))
}

// addNumbers is a package-level function
func addNumbers(a, b int) int {
	return a + b
}

// packageOrganization demonstrates package organization
func packageOrganization() {
	// Package structure:
	// - Each directory is a package
	// - Files in same package share the same package name
	// - Package name should match directory name (except main)

	fmt.Println("Package organization:")
	fmt.Println("  - One package per directory")
	fmt.Println("  - Package name matches directory")
	fmt.Println("  - main package is executable")
	fmt.Println("  - Other packages are libraries")
}

// visibilityRules demonstrates visibility rules
func visibilityRules() {
	// Exported (public): starts with uppercase
	// Unexported (private): starts with lowercase

	// This function is exported (visible outside package)
	ExportedFunction()

	// This function is unexported (only visible in package)
	unexportedFunction()

	// Example struct
	p := Person{
		Name: "Alice", // exported field
		age:  30,      // unexported field
	}

	fmt.Printf("Person: %+v\n", p)
	p.SetAge(31) // exported method
	fmt.Printf("Updated age: %d\n", p.GetAge())
}

// ExportedFunction is visible outside the package
func ExportedFunction() {
	fmt.Println("This function is exported")
}

// unexportedFunction is only visible within the package
func unexportedFunction() {
	fmt.Println("This function is unexported")
}

// Person demonstrates visibility in structs
type Person struct {
	Name string // exported
	age  int    // unexported
}

// SetAge sets the age (exported method)
func (p *Person) SetAge(age int) {
	if age < 0 || age > 150 {
		return
	}
	p.age = age
}

// GetAge returns the age (exported method)
func (p *Person) GetAge() int {
	return p.age
}

// importAliases demonstrates import aliases
func importAliases() {
	// Standard import
	// import "fmt"

	// Aliased import
	// import f "fmt"
	// f.Println("Hello")

	// Dot import (not recommended)
	// import . "fmt"
	// Println("Hello")

	fmt.Println("Use import aliases to avoid naming conflicts")
}

// blankImports demonstrates blank imports
func blankImports() {
	// Blank import for side effects
	// import _ "package/name"

	// Used when package has init() functions that need to run
	// but you don't use any exported identifiers

	fmt.Println("Blank imports are used for side effects")
}

// Package-level variables
var (
	packageVar1 = "package variable 1"
	packageVar2 = 42
)

// Package-level constants
const (
	PackageConst1 = "package constant 1"
	PackageConst2 = 100
)

// init function runs automatically when package is imported
func init() {
	fmt.Println("Package initialization")
}

// ModuleExample demonstrates Go modules
func ModuleExample() {
	// Go modules:
	// - go.mod file defines module
	// - go.sum file tracks checksums
	// - Versioning with semantic versioning
	// - Dependency management

	fmt.Println("Go modules provide dependency management")
	fmt.Println("  - go mod init <module-name>")
	fmt.Println("  - go mod tidy")
	fmt.Println("  - go mod download")
	fmt.Println("  - go get <package>")
}

// Best practices for packages
func packageBestPractices() {
	// 1. Keep packages focused and cohesive
	// 2. Use clear, descriptive names
	// 3. Minimize exported API surface
	// 4. Document exported functions/types
	// 5. Group related functionality
	// 6. Avoid circular dependencies
	// 7. Use internal packages for private APIs

	fmt.Println("Package best practices:")
	fmt.Println("  - Single responsibility")
	fmt.Println("  - Clear naming")
	fmt.Println("  - Minimal public API")
	fmt.Println("  - Good documentation")
}
