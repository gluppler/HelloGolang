package main

import (
	"fmt"
	"maps"
	"slices"
)

// Maps demonstrates map operations and patterns

func main() {
	mapBasics()
	mapOperations()
	mapIteration()
	mapSafety()
	mapPatterns()
}

// mapBasics demonstrates basic map operations
func mapBasics() {
	// Map declaration
	var m1 map[string]int
	fmt.Printf("Nil map: %v (nil: %t)\n", m1, m1 == nil)

	// Map literal
	m2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}
	fmt.Printf("Map literal: %v\n", m2)

	// Map using make
	m3 := make(map[string]int)
	m3["one"] = 1
	m3["two"] = 2
	fmt.Printf("Map with make: %v\n", m3)

	// Map with initial capacity hint
	m4 := make(map[string]int, 10)
	fmt.Printf("Pre-allocated map: %v (len: %d)\n", m4, len(m4))
}

// mapOperations demonstrates map operations
func mapOperations() {
	m := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}

	// Access
	fmt.Printf("Alice's age: %d\n", m["Alice"])

	// Secure: check existence before using zero value
	age, exists := m["David"]
	if exists {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Println("David not found in map")
	}

	// Modification
	m["Alice"] = 31
	fmt.Printf("Updated Alice's age: %d\n", m["Alice"])

	// Addition
	m["David"] = 28
	fmt.Printf("Added David: %v\n", m)

	// Deletion
	delete(m, "Bob")
	fmt.Printf("After deleting Bob: %v\n", m)

	// Length
	fmt.Printf("Map length: %d\n", len(m))
}

// mapIteration demonstrates iterating over maps
func mapIteration() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	// Iterate over key-value pairs
	fmt.Println("Key-value iteration:")
	for key, value := range m {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// Iterate over keys only
	fmt.Println("Keys only:")
	for key := range m {
		fmt.Printf("  %s\n", key)
	}

	// Get all keys as slice
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	fmt.Printf("Sorted keys: %v\n", keys)

	// Get all values as slice
	values := make([]int, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	fmt.Printf("Values: %v\n", values)
}

// mapSafety demonstrates safe map operations
func mapSafety() {
	// Secure: always check existence for critical operations
	m := map[string]int{
		"admin": 100,
		"user":  10,
	}

	// Safe access pattern
	role := "admin"
	if level, ok := m[role]; ok {
		fmt.Printf("Role '%s' has level %d\n", role, level)
	} else {
		fmt.Printf("Role '%s' not found\n", role)
	}

	// Safe deletion (delete is safe even if key doesn't exist)
	delete(m, "nonexistent") // No panic

	// Secure: prevent nil map panic
	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("Map is nil, initializing...")
		nilMap = make(map[string]int)
	}
	nilMap["key"] = 1 // Safe now

	// Concurrent access safety (maps are not thread-safe)
	// Use sync.Mutex or sync.RWMutex for concurrent access
	fmt.Println("Note: Maps are not thread-safe for concurrent writes")
}

// mapPatterns demonstrates common map patterns
func mapPatterns() {
	// Counting occurrences
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	fmt.Printf("Word counts: %v\n", counts)

	// Grouping
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 30},
		{"David", 25},
	}
	byAge := make(map[int][]string)
	for _, p := range people {
		byAge[p.Age] = append(byAge[p.Age], p.Name)
	}
	fmt.Printf("Grouped by age: %v\n", byAge)

	// Set simulation (map[string]bool or map[string]struct{})
	set1 := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	set2 := map[string]bool{
		"b": true,
		"c": true,
		"d": true,
	}

	// Check membership
	if set1["a"] {
		fmt.Println("'a' is in set1")
	}

	// Intersection
	intersection := make(map[string]bool)
	for key := range set1 {
		if set2[key] {
			intersection[key] = true
		}
	}
	fmt.Printf("Intersection: %v\n", intersection)

	// Union
	union := make(map[string]bool)
	maps.Copy(union, set1)
	maps.Copy(union, set2)
	fmt.Printf("Union: %v\n", union)

	// Map of maps
	matrix := map[string]map[string]int{
		"row1": {"col1": 1, "col2": 2},
		"row2": {"col1": 3, "col2": 4},
	}
	fmt.Printf("Map of maps: %v\n", matrix)

	// Secure: nested map access with safety checks
	if row, ok := matrix["row1"]; ok {
		if val, ok := row["col1"]; ok {
			fmt.Printf("matrix[row1][col1] = %d\n", val)
		}
	}
}

// mapComparison demonstrates map comparison (Go 1.21+)
func mapComparison() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}

	// Compare maps using maps.Equal
	equal := maps.Equal(m1, m2)
	fmt.Printf("m1 == m2: %t\n", equal)

	equal2 := maps.Equal(m1, m3)
	fmt.Printf("m1 == m3: %t\n", equal2)

	// Check if one map is a subset
	subset := map[string]int{"a": 1}
	isSubset := true
	for k, v := range subset {
		if val, ok := m1[k]; !ok || val != v {
			isSubset = false
			break
		}
	}
	fmt.Printf("subset is subset of m1: %t\n", isSubset)
}

// mapCloning demonstrates map cloning
func mapCloning() {
	original := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Shallow copy
	shallow := original
	shallow["a"] = 100
	fmt.Printf("Original after shallow copy: %v\n", original)

	// Reset
	original = map[string]int{"a": 1, "b": 2, "c": 3}

	// Deep copy using maps.Clone (Go 1.21+)
	cloned := maps.Clone(original)
	cloned["a"] = 200
	fmt.Printf("Original: %v, Cloned: %v\n", original, cloned)

	// Manual deep copy
	manual := make(map[string]int, len(original))
	for k, v := range original {
		manual[k] = v
	}
	manual["a"] = 300
	fmt.Printf("Original: %v, Manual copy: %v\n", original, manual)
}
