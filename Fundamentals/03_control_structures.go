package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Control Structures demonstrates if/else, switch, for loops, and control flow

func main() {
	ifElseStatements()
	switchStatements()
	forLoops()
	rangeLoops()
	breakAndContinue()
	gotoStatement()
	deferStatement()
}

// ifElseStatements demonstrates if/else and if initialization
func ifElseStatements() {
	x := 10

	// Basic if
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// if-else
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	// if-else if-else
	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	}

	// if with initialization (common pattern)
	if err := validateNumber(x); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Number is valid")
	}

	// Secure: proper bounds checking
	if x < 0 || x > 100 {
		fmt.Println("x is out of valid range")
	}
}

// validateNumber validates a number (returns error if invalid)
func validateNumber(n int) error {
	if n < 0 {
		return fmt.Errorf("number cannot be negative")
	}
	return nil
}

// switchStatements demonstrates switch statements
func switchStatements() {
	day := "Monday"

	// Basic switch
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("End of work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Midweek")
	}

	// Switch with initialization
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// Switch on type (type switch)
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// Switch with fallthrough
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}

// forLoops demonstrates different for loop patterns
func forLoops() {
	// Traditional for loop
	fmt.Println("Traditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-style loop
	fmt.Println("While-style loop:")
	j := 0
	for j < 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("Infinite loop with break:")
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()

	// Loop with continue
	fmt.Println("Loop with continue (skip even numbers):")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// rangeLoops demonstrates range loops
func rangeLoops() {
	// Range over slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Range over slice:")
	for index, value := range numbers {
		fmt.Printf("  [%d] = %d\n", index, value)
	}

	// Range over slice (index only)
	fmt.Println("Range over slice (index only):")
	for i := range numbers {
		fmt.Printf("  [%d] = %d\n", i, numbers[i])
	}

	// Range over slice (value only)
	fmt.Println("Range over slice (value only):")
	for _, value := range numbers {
		fmt.Printf("  %d\n", value)
	}

	// Range over map
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}
	fmt.Println("Range over map:")
	for name, age := range ages {
		fmt.Printf("  %s: %d\n", name, age)
	}

	// Range over string (runes)
	text := "Hello"
	fmt.Println("Range over string:")
	for i, char := range text {
		fmt.Printf("  [%d] = %c (U+%04X)\n", i, char, char)
	}

	// Range over channel
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println("Range over channel:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}
}

// breakAndContinue demonstrates break and continue with labels
func breakAndContinue() {
	fmt.Println("Nested loops with labeled break:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Printf("  [%d,%d] ", i, j)
		}
		fmt.Println()
	}

	fmt.Println("Nested loops with labeled continue:")
OuterLoop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue OuterLoop2
			}
			fmt.Printf("  [%d,%d] ", i, j)
		}
		fmt.Println()
	}
}

// gotoStatement demonstrates goto (use sparingly)
func gotoStatement() {
	fmt.Println("Goto example:")
	i := 0

Start:
	if i < 5 {
		fmt.Printf("  %d ", i)
		i++
		goto Start
	}
	fmt.Println()

	// Secure: goto can be useful for error handling
	if err := riskyOperation(); err != nil {
		goto ErrorHandler
	}
	fmt.Println("Operation succeeded")
	return

ErrorHandler:
	fmt.Println("Error occurred, cleaning up...")
}

// riskyOperation simulates an operation that might fail
func riskyOperation() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return fmt.Errorf("operation failed")
	}
	return nil
}

// deferStatement demonstrates defer for cleanup
func deferStatement() {
	fmt.Println("Defer examples:")

	// Basic defer
	defer fmt.Println("This runs last")
	fmt.Println("This runs first")

	// Multiple defers (LIFO order)
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	fmt.Println("Main execution")

	// Defer with function
	defer func() {
		fmt.Println("Deferred function")
	}()

	// Defer with arguments evaluation
	value := 10
	defer fmt.Printf("Deferred value: %d\n", value)
	value = 20
	fmt.Printf("Current value: %d\n", value)

	// Defer for resource cleanup
	resourceCleanup()
}

// resourceCleanup demonstrates defer for resource management
func resourceCleanup() {
	fmt.Println("Resource cleanup example:")

	// Simulate opening a resource
	fmt.Println("Opening resource...")
	defer func() {
		fmt.Println("Closing resource...")
	}()

	fmt.Println("Using resource...")
}
