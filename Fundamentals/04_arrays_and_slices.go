package main

import (
	"fmt"
	"slices"
)

// Arrays and Slices demonstrates array and slice operations

func main() {
	arrays()
	slicesBasic()
	sliceOperations()
	sliceAppending()
	sliceCopying()
	sliceFiltering()
	multidimensionalSlices()
}

// arrays demonstrates fixed-size arrays
func arrays() {
	// Array declaration
	var arr1 [5]int
	fmt.Printf("Zero-initialized array: %v\n", arr1)

	// Array initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Initialized array: %v\n", arr2)

	// Array with inferred size
	arr3 := [...]int{10, 20, 30}
	fmt.Printf("Inferred size array: %v (length: %d)\n", arr3, len(arr3))

	// Array access and modification
	arr2[0] = 100
	fmt.Printf("Modified array: %v\n", arr2)

	// Array iteration
	fmt.Println("Array iteration:")
	for i, v := range arr2 {
		fmt.Printf("  [%d] = %d\n", i, v)
	}

	// Array comparison (arrays are comparable)
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Arrays equal: %t\n", arr2 == arr4)
}

// slicesBasic demonstrates basic slice operations
func slicesBasic() {
	// Slice declaration
	var sl1 []int
	fmt.Printf("Nil slice: %v (nil: %t)\n", sl1, sl1 == nil)

	// Slice literal
	sl2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v (length: %d, capacity: %d)\n", sl2, len(sl2), cap(sl2))

	// Slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	sl3 := arr[1:4] // [20, 30, 40]
	fmt.Printf("Slice from array: %v\n", sl3)

	// Slice using make
	sl4 := make([]int, 5)     // length 5, capacity 5
	sl5 := make([]int, 5, 10) // length 5, capacity 10
	fmt.Printf("make([]int, 5): %v (len: %d, cap: %d)\n", sl4, len(sl4), cap(sl4))
	fmt.Printf("make([]int, 5, 10): %v (len: %d, cap: %d)\n", sl5, len(sl5), cap(sl5))

	// Slice access
	fmt.Printf("sl2[0] = %d\n", sl2[0])
	fmt.Printf("sl2[1:3] = %v\n", sl2[1:3])
	fmt.Printf("sl2[:3] = %v\n", sl2[:3])
	fmt.Printf("sl2[2:] = %v\n", sl2[2:])
	fmt.Printf("sl2[:] = %v\n", sl2[:])
}

// sliceOperations demonstrates slice operations
func sliceOperations() {
	sl := []int{1, 2, 3, 4, 5}

	// Length and capacity
	fmt.Printf("Slice: %v\n", sl)
	fmt.Printf("Length: %d, Capacity: %d\n", len(sl), cap(sl))

	// Safe access with bounds checking
	if len(sl) > 0 {
		fmt.Printf("First element: %d\n", sl[0])
		fmt.Printf("Last element: %d\n", sl[len(sl)-1])
	}

	// Secure: bounds checking before access
	index := 10
	if index >= 0 && index < len(sl) {
		fmt.Printf("sl[%d] = %d\n", index, sl[index])
	} else {
		fmt.Printf("Index %d out of bounds\n", index)
	}

	// Slice operations using slices package (Go 1.21+)
	slices.Sort(sl)
	fmt.Printf("Sorted: %v\n", sl)

	contains := slices.Contains(sl, 3)
	fmt.Printf("Contains 3: %t\n", contains)

	indexOf := slices.Index(sl, 3)
	fmt.Printf("Index of 3: %d\n", indexOf)
}

// sliceAppending demonstrates appending to slices
func sliceAppending() {
	// Start with empty slice
	var sl []int
	fmt.Printf("Initial: %v (len: %d, cap: %d)\n", sl, len(sl), cap(sl))

	// Append single element
	sl = append(sl, 1)
	fmt.Printf("After append(1): %v (len: %d, cap: %d)\n", sl, len(sl), cap(sl))

	// Append multiple elements
	sl = append(sl, 2, 3, 4)
	fmt.Printf("After append(2,3,4): %v (len: %d, cap: %d)\n", sl, len(sl), cap(sl))

	// Append another slice
	sl2 := []int{5, 6, 7}
	sl = append(sl, sl2...)
	fmt.Printf("After append(sl2...): %v (len: %d, cap: %d)\n", sl, len(sl), cap(sl))

	// Pre-allocate with known capacity for better performance
	sl3 := make([]int, 0, 10) // length 0, capacity 10
	for i := 0; i < 5; i++ {
		sl3 = append(sl3, i)
	}
	fmt.Printf("Pre-allocated: %v (len: %d, cap: %d)\n", sl3, len(sl3), cap(sl3))
}

// sliceCopying demonstrates copying slices
func sliceCopying() {
	source := []int{1, 2, 3, 4, 5}

	// Shallow copy (both point to same underlying array)
	shallow := source
	shallow[0] = 100
	fmt.Printf("Source after shallow copy modification: %v\n", source)

	// Reset
	source = []int{1, 2, 3, 4, 5}

	// Deep copy using append
	deep1 := append([]int(nil), source...)
	deep1[0] = 200
	fmt.Printf("Source: %v, Deep copy: %v\n", source, deep1)

	// Deep copy using make and copy
	deep2 := make([]int, len(source))
	copy(deep2, source)
	deep2[0] = 300
	fmt.Printf("Source: %v, Deep copy: %v\n", source, deep2)

	// Copy with different sizes
	small := make([]int, 3)
	copied := copy(small, source)
	fmt.Printf("Copied %d elements to smaller slice: %v\n", copied, small)

	// Copy using slices.Clone (Go 1.21+)
	cloned := slices.Clone(source)
	cloned[0] = 400
	fmt.Printf("Source: %v, Cloned: %v\n", source, cloned)
}

// sliceFiltering demonstrates filtering and transforming slices
func sliceFiltering() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	evens := filterEven(numbers)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Evens: %v\n", evens)

	// Filter using function
	isOdd := func(n int) bool { return n%2 != 0 }
	odds := filter(numbers, isOdd)
	fmt.Printf("Odds: %v\n", odds)

	// Map transformation
	doubled := mapSlice(numbers, func(n int) int { return n * 2 })
	fmt.Printf("Doubled: %v\n", doubled)
}

// filterEven filters even numbers from a slice
func filterEven(numbers []int) []int {
	var result []int
	for _, n := range numbers {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

// filter filters a slice using a predicate function
func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range numbers {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// mapSlice transforms each element using a function
func mapSlice(numbers []int, transform func(int) int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = transform(n)
	}
	return result
}

// multidimensionalSlices demonstrates multi-dimensional slices
func multidimensionalSlices() {
	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Matrix:")
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// Dynamic 2D slice
	dynamic := make([][]int, 3)
	for i := range dynamic {
		dynamic[i] = make([]int, 3)
		for j := range dynamic[i] {
			dynamic[i][j] = i*3 + j + 1
		}
	}
	fmt.Println("Dynamic 2D Matrix:")
	for i, row := range dynamic {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// Jagged slice (different row lengths)
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	fmt.Println("Jagged slice:")
	for i, row := range jagged {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
}
