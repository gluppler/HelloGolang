package main

import (
	"fmt"
)

// Generics demonstrates generic types and functions (Go 1.18+)

func main() {
	genericFunctions()
	genericTypes()
	typeConstraints()
	comparableConstraint()
	orderedConstraint()
	genericInterfaces()
	genericBestPractices()
}

// genericFunctions demonstrates generic functions
func genericFunctions() {
	// Generic function with type parameter
	result1 := identity(42)
	fmt.Printf("identity(42) = %v (type: %T)\n", result1, result1)

	result2 := identity("hello")
	fmt.Printf("identity(\"hello\") = %v (type: %T)\n", result2, result2)

	// Generic function with multiple type parameters
	pair := Pair[int, string]{First: 1, Second: "one"}
	fmt.Printf("Pair: %+v\n", pair)

	// Generic function with constraints
	sum := add(10, 20)
	fmt.Printf("add(10, 20) = %d\n", sum)

	// Generic slice operations
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapSlice(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Doubled: %v\n", doubled)

	filtered := filterSlice(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Filtered evens: %v\n", filtered)
}

// identity is a generic identity function
func identity[T any](value T) T {
	return value
}

// Pair is a generic pair type
type Pair[T, U any] struct {
	First  T
	Second U
}

// Numeric constraint for numeric types
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// add adds two numbers of the same numeric type
func add[T Numeric](a, b T) T {
	return a + b
}

// mapSlice applies a function to each element
func mapSlice[T any](slice []T, fn func(T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// filterSlice filters elements based on predicate
func filterSlice[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// genericTypes demonstrates generic types
func genericTypes() {
	// Generic stack
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Printf("Stack top: %d\n", stack.Peek())
	fmt.Printf("Stack pop: %d\n", stack.Pop())
	fmt.Printf("Stack size: %d\n", stack.Size())

	// Generic map wrapper
	cache := NewCache[string, int]()
	cache.Set("one", 1)
	cache.Set("two", 2)

	val, ok := cache.Get("one")
	if ok {
		fmt.Printf("Cache get: %d\n", val)
	}
}

// Stack is a generic stack implementation
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

// Size returns the number of items
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Cache is a generic cache implementation
type Cache[K comparable, V any] struct {
	data map[K]V
}

// NewCache creates a new cache
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

// Set stores a value
func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

// Get retrieves a value
func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return value, ok
}

// typeConstraints demonstrates type constraints
func typeConstraints() {
	// Using built-in constraints
	fmt.Printf("Max(10, 20) = %d\n", maxValue(10, 20))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", maxValue(3.14, 2.71))

	// Custom constraint
	fmt.Printf("Sum(1, 2, 3) = %d\n", sumNumbers(1, 2, 3))
	fmt.Printf("Sum(1.5, 2.5, 3.5) = %.2f\n", sumNumbers(1.5, 2.5, 3.5))
}

// Ordered constraint for types that support comparison
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// maxValue returns the maximum of two values
func maxValue[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// sumNumbers sums numeric values
func sumNumbers[T Numeric](values ...T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// comparableConstraint demonstrates comparable constraint
func comparableConstraint() {
	// Comparable types can use == and !=
	fmt.Printf("Equal(1, 1) = %t\n", equal(1, 1))
	fmt.Printf("Equal(\"a\", \"b\") = %t\n", equal("a", "b"))

	// Find in slice
	numbers := []int{1, 2, 3, 4, 5}
	found := contains(numbers, 3)
	fmt.Printf("Contains 3: %t\n", found)
}

// equal checks if two comparable values are equal
func equal[T comparable](a, b T) bool {
	return a == b
}

// contains checks if slice contains value
func contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// orderedConstraint demonstrates ordered constraint
func orderedConstraint() {
	// Ordered types support <, <=, >, >=
	numbers := []int{5, 2, 8, 1, 9}
	min := findMin(numbers)
	max := findMax(numbers)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Sort
	sorted := sortSlice(numbers)
	fmt.Printf("Sorted: %v\n", sorted)
}

// findMin finds minimum value
func findMin[T Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// findMax finds maximum value
func findMax[T Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// sortSlice sorts a slice (simple bubble sort for demo)
func sortSlice[T Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// genericInterfaces demonstrates generic interfaces
func genericInterfaces() {
	// Generic interface
	var container Container[int] = &IntContainer{data: []int{1, 2, 3}}
	fmt.Printf("Container size: %d\n", container.Size())
	fmt.Printf("Container get(1): %d\n", container.Get(1))
}

// Container is a generic interface
type Container[T any] interface {
	Size() int
	Get(index int) T
	Set(index int, value T)
}

// IntContainer implements Container[int]
type IntContainer struct {
	data []int
}

// Size returns the size
func (c *IntContainer) Size() int {
	return len(c.data)
}

// Get retrieves an element
func (c *IntContainer) Get(index int) int {
	if index < 0 || index >= len(c.data) {
		var zero int
		return zero
	}
	return c.data[index]
}

// Set sets an element
func (c *IntContainer) Set(index int, value int) {
	if index >= 0 && index < len(c.data) {
		c.data[index] = value
	}
}

// genericBestPractices demonstrates best practices
func genericBestPractices() {
	// 1. Use generics when you have type-safe operations
	// 2. Prefer interfaces when behavior matters more than types
	// 3. Use constraints appropriately
	// 4. Keep generic code readable
	// 5. Document type parameters

	fmt.Println("Generic best practices:")
	fmt.Println("  - Use for type-safe operations")
	fmt.Println("  - Appropriate constraints")
	fmt.Println("  - Clear documentation")
	fmt.Println("  - Readable code")
}
