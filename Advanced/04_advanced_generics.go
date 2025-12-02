package main

import (
	"fmt"
	"sync"
)

// Advanced Generics demonstrates advanced generic patterns and techniques

func main() {
	genericConstraints()
	genericTypeSets()
	genericInterfaces()
	genericMethods()
	genericReflection()
	genericPerformance()
}

// Numeric constraint for numeric types
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// genericConstraints demonstrates advanced constraint patterns
func genericConstraints() {
	// Sum function with numeric constraint
	fmt.Printf("Sum of ints: %d\n", sumNumeric(1, 2, 3, 4, 5))
	fmt.Printf("Sum of floats: %.2f\n", sumNumeric(1.5, 2.5, 3.5))
	
	// Find function with comparable constraint
	numbers := []int{10, 20, 30, 40, 50}
	index := findComparable(numbers, 30)
	fmt.Printf("Found 30 at index: %d\n", index)
}

// sumNumeric sums numeric values
func sumNumeric[T Numeric](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// findComparable finds value in slice
func findComparable[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Stringer interface for type sets example
type Stringer interface {
	String() string
}

// Int type for Stringer example
type Int int

// String implements Stringer for Int
func (i Int) String() string {
	return fmt.Sprintf("Int(%d)", i)
}

// Float type for Stringer example
type Float float64

// String implements Stringer for Float
func (f Float) String() string {
	return fmt.Sprintf("Float(%.2f)", f)
}

// Number interface for union type constraint
type Number interface {
	~int | ~float64
}

// printStringer prints a Stringer value
func printStringer[T Stringer](value T) {
	fmt.Println(value.String())
}

// doubleNumber doubles a number
func doubleNumber[T Number](value T) T {
	return value * 2
}

// genericTypeSets demonstrates type sets and union types
func genericTypeSets() {
	// Function that works with Stringer
	printStringer(Int(42))
	printStringer(Float(3.14))
	
	// Function with union constraint
	fmt.Printf("Double int: %d\n", doubleNumber(21))
	fmt.Printf("Double float: %.2f\n", doubleNumber(10.5))
}

// Container is a generic container interface
type Container[T any] interface {
	Add(item T)
	Get(index int) T
	Size() int
}

// SliceContainer is a generic slice container
type SliceContainer[T any] struct {
	items []T
}

// NewSliceContainer creates a new slice container
func NewSliceContainer[T any]() *SliceContainer[T] {
	return &SliceContainer[T]{
		items: make([]T, 0),
	}
}

// Add adds an item to the container
func (c *SliceContainer[T]) Add(item T) {
	c.items = append(c.items, item)
}

// Get retrieves an item by index
func (c *SliceContainer[T]) Get(index int) T {
	if index < 0 || index >= len(c.items) {
		var zero T
		return zero
	}
	return c.items[index]
}

// Size returns the number of items
func (c *SliceContainer[T]) Size() int {
	return len(c.items)
}

// genericInterfaces demonstrates generic interfaces
func genericInterfaces() {
	// Use generic container
	container := NewSliceContainer[string]()
	container.Add("hello")
	container.Add("world")
	fmt.Printf("Container size: %d\n", container.Size())
	fmt.Printf("Container[0]: %s\n", container.Get(0))
}

// Stack is a generic stack implementation
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// genericMethods demonstrates generic methods
func genericMethods() {
	// Use generic stack
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	if val, ok := stack.Peek(); ok {
		fmt.Printf("Stack peek: %d\n", val)
	}
	
	if val, ok := stack.Pop(); ok {
		fmt.Printf("Stack pop: %d\n", val)
	}
}

// genericReflection demonstrates generics with reflection
func genericReflection() {
	numbers := []int{1, 2, 3, 4, 5}
	
	// Map: square each number
	squared := mapSliceGeneric(numbers, func(n int) int {
		return n * n
	})
	fmt.Printf("Squared: %v\n", squared)
	
	// Filter: even numbers
	evens := filterSliceGeneric(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Evens: %v\n", evens)
	
	// Reduce: sum
	sum := reduceGeneric(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("Sum: %d\n", sum)
}

// mapSliceGeneric maps slice elements
func mapSliceGeneric[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// filterSliceGeneric filters slice elements
func filterSliceGeneric[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// reduceGeneric reduces slice to single value
func reduceGeneric[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// genericPerformance demonstrates performance considerations
func genericPerformance() {
	// Generics are compiled to concrete types
	// No runtime overhead compared to interfaces
	
	// Example: Generic binary search
	sortedInts := []int{1, 3, 5, 7, 9, 11, 13}
	index := binarySearchGeneric(sortedInts, 7)
	fmt.Printf("Found 7 at index: %d\n", index)
	
	// Generic cache with type safety
	cache := NewCache[string, int]()
	cache.Set("one", 1)
	cache.Set("two", 2)
	
	if val, ok := cache.Get("one"); ok {
		fmt.Printf("Cache value: %d\n", val)
	}
}

// binarySearchGeneric performs binary search
func binarySearchGeneric[T interface{ ~int | ~float64 }](slice []T, target T) int {
	left, right := 0, len(slice)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

// Cache is a generic cache implementation
type Cache[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewCache creates a new cache
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

// Get retrieves a value from cache
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Set stores a value in cache
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

