package main

import (
	"fmt"
	"unsafe"
)

// Variables and Types demonstrates Go's type system and variable declarations

func main() {
	// Basic variable declarations
	varExplicit()
	varInferred()
	varShort()
	constants()
	pointers()
	typeConversions()
	zeroValues()
}

// varExplicit demonstrates explicit variable declaration with type
func varExplicit() {
	var name string = "Go"
	var age int = 14
	var isActive bool = true

	fmt.Printf("Explicit: name=%s, age=%d, isActive=%t\n", name, age, isActive)
}

// varInferred demonstrates type inference
func varInferred() {
	var name = "Go"     // string inferred
	var age = 14        // int inferred
	var isActive = true // bool inferred

	fmt.Printf("Inferred: name=%s, age=%d, isActive=%t\n", name, age, isActive)
}

// varShort demonstrates short variable declaration (most common)
func varShort() {
	name := "Go"
	age := 14
	isActive := true

	fmt.Printf("Short: name=%s, age=%d, isActive=%t\n", name, age, isActive)
}

// constants demonstrates constant declarations
func constants() {
	const pi = 3.14159
	const e = 2.71828
	const (
		statusOK    = 200
		statusError = 500
	)

	// Typed constants
	const maxInt int64 = 9223372036854775807

	fmt.Printf("Constants: pi=%.5f, e=%.5f, statusOK=%d\n", pi, e, statusOK)
	fmt.Printf("Max int64: %d\n", maxInt)
}

// pointers demonstrates pointer usage
func pointers() {
	x := 42
	p := &x // p is a pointer to x

	fmt.Printf("Value: %d, Pointer: %p, Dereferenced: %d\n", x, p, *p)

	*p = 100 // Modify value through pointer
	fmt.Printf("After modification: x=%d\n", x)
}

// typeConversions demonstrates type conversions
func typeConversions() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Type conversions: int=%d, float64=%.2f, uint=%d\n", i, f, u)

	// String conversion
	var s string = fmt.Sprintf("%d", i)
	fmt.Printf("String conversion: %s\n", s)
}

// zeroValues demonstrates zero values for different types
func zeroValues() {
	var (
		i   int
		f   float64
		b   bool
		s   string
		ptr *int
		sl  []int
		m   map[string]int
	)

	fmt.Printf("Zero values:\n")
	fmt.Printf("  int: %d\n", i)
	fmt.Printf("  float64: %g\n", f)
	fmt.Printf("  bool: %t\n", b)
	fmt.Printf("  string: %q\n", s)
	fmt.Printf("  pointer: %v\n", ptr)
	fmt.Printf("  slice: %v (nil: %t)\n", sl, sl == nil)
	fmt.Printf("  map: %v (nil: %t)\n", m, m == nil)
}

// typeSizes demonstrates size of different types
func typeSizes() {
	fmt.Printf("Type sizes:\n")
	fmt.Printf("  int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("  int8: %d bytes\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("  int16: %d bytes\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("  int32: %d bytes\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("  int64: %d bytes\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("  float32: %d bytes\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("  float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("  bool: %d bytes\n", unsafe.Sizeof(bool(false)))
	fmt.Printf("  string: %d bytes (header)\n", unsafe.Sizeof(""))
}
