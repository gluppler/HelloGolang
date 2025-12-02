package main

import (
	"fmt"
	"io"
	"strings"
)

// Interfaces demonstrates interface types, implementation, and polymorphism

func main() {
	basicInterfaces()
	interfaceComposition()
	emptyInterface()
	typeAssertions()
	typeSwitches()
	interfaceBestPractices()
}

// Shape defines a geometric shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates rectangle area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates rectangle perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

// Area calculates circle area
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates circle perimeter
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// basicInterfaces demonstrates basic interface usage
func basicInterfaces() {
	var s Shape

	rect := Rectangle{Width: 5, Height: 3}
	s = rect
	fmt.Printf("Rectangle area: %.2f, perimeter: %.2f\n", s.Area(), s.Perimeter())

	circle := Circle{Radius: 4}
	s = circle
	fmt.Printf("Circle area: %.2f, perimeter: %.2f\n", s.Area(), s.Perimeter())

	// Interface slice
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 4},
	}

	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}

// Reader and Writer interfaces for composition example
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

// ReadWriter is a composed interface
type ReadWriter interface {
	Reader
	Writer
}

// StringBuffer implements ReadWriter
type StringBuffer struct {
	buf strings.Builder
}

// Read reads from buffer
func (sb *StringBuffer) Read(p []byte) (int, error) {
	data := sb.buf.String()
	if len(data) == 0 {
		return 0, io.EOF
	}
	n := copy(p, []byte(data))
	return n, nil
}

// Write writes to buffer
func (sb *StringBuffer) Write(p []byte) (int, error) {
	n, err := sb.buf.Write(p)
	return n, err
}

// interfaceComposition demonstrates interface composition
func interfaceComposition() {
	buf := &StringBuffer{}
	buf.Write([]byte("Hello, World!"))

	data := make([]byte, 20)
	n, _ := buf.Read(data)
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
}

// emptyInterface demonstrates empty interface usage
func emptyInterface() {
	// Empty interface can hold any type
	var i interface{}

	i = 42
	fmt.Printf("int: %v (type: %T)\n", i, i)

	i = "hello"
	fmt.Printf("string: %v (type: %T)\n", i, i)

	i = []int{1, 2, 3}
	fmt.Printf("slice: %v (type: %T)\n", i, i)

	// Slice of empty interface
	values := []interface{}{1, "two", 3.0, true}
	for _, v := range values {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

// typeAssertions demonstrates type assertions
func typeAssertions() {
	var i interface{} = "hello"

	// Type assertion
	s, ok := i.(string)
	if ok {
		fmt.Printf("String value: %s\n", s)
	}

	// Type assertion without ok check (panics if wrong type)
	i = 42
	if s, ok := i.(string); ok {
		fmt.Printf("String: %s\n", s)
	} else {
		fmt.Printf("Not a string, value: %v\n", i)
	}

	// Multiple type assertions
	processValue(42)
	processValue("hello")
	processValue([]int{1, 2, 3})
}

// processValue processes different value types
func processValue(v interface{}) {
	if s, ok := v.(string); ok {
		fmt.Printf("Processing string: %s\n", s)
	} else if n, ok := v.(int); ok {
		fmt.Printf("Processing int: %d\n", n)
	} else if sl, ok := v.([]int); ok {
		fmt.Printf("Processing slice: %v\n", sl)
	} else {
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// typeSwitches demonstrates type switches
func typeSwitches() {
	var i interface{} = 42

	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case []int:
		fmt.Printf("Slice: %v\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// Type switch with multiple cases
	processWithTypeSwitch(42)
	processWithTypeSwitch("hello")
	processWithTypeSwitch(true)
	processWithTypeSwitch([]int{1, 2, 3})
}

// processWithTypeSwitch processes value using type switch
func processWithTypeSwitch(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("Int: %d\n", val)
	case string:
		fmt.Printf("String: %s\n", val)
	case bool:
		fmt.Printf("Bool: %t\n", val)
	case []int:
		fmt.Printf("Int slice: %v\n", val)
	default:
		fmt.Printf("Other: %v (type: %T)\n", val, val)
	}
}

// interfaceBestPractices demonstrates interface best practices
func interfaceBestPractices() {
	// Accept interfaces, return structs
	reader := strings.NewReader("Hello, World!")
	processReader(reader)

	// Interface with method sets
	var w FileWriter = &StringWriter{}
	w.Write([]byte("Test"))
}

// FileWriter interface (renamed to avoid conflict)
type FileWriter interface {
	Write([]byte) (int, error)
}

// StringWriter implements FileWriter
type StringWriter struct {
	buf strings.Builder
}

// Write writes bytes to buffer
func (sw *StringWriter) Write(p []byte) (int, error) {
	return sw.buf.Write(p)
}

// processReader processes a Reader
func processReader(r io.Reader) {
	data := make([]byte, 12)
	n, err := r.Read(data)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
}

// Comparable interface (built-in)
func comparableInterface() {
	// Types that implement == and != are comparable
	var a, b int = 5, 5
	fmt.Printf("a == b: %t\n", a == b)

	var s1, s2 string = "hello", "hello"
	fmt.Printf("s1 == s2: %t\n", s1 == s2)

	// Slices and maps are not comparable
	// sl1 := []int{1, 2, 3}
	// sl2 := []int{1, 2, 3}
	// sl1 == sl2  // compile error
}

// Stringer interface (from fmt package)
type PersonExample struct {
	Name string
	Age  int
}

// String implements fmt.Stringer
func (p PersonExample) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func stringerExample() {
	p := PersonExample{Name: "Alice", Age: 30}
	fmt.Println(p) // Uses String() method
}
