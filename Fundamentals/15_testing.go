package main

import (
	"fmt"
	"testing"
)

// Testing demonstrates Go testing framework and best practices

// Example test function (would be in a _test.go file)
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("add(2, 3) = %d; expected %d", result, expected)
	}
}

// add is a simple function to test
func add(a, b int) int {
	return a + b
}

// TestSubtract demonstrates table-driven tests
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zero", 5, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("subtract(%d, %d) = %d; expected %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// subtract performs subtraction
func subtract(a, b int) int {
	return a - b
}

// TestDivide demonstrates error testing
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"fractional result", 7, 2, 3.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

// divide performs division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// BenchmarkAdd demonstrates benchmarking
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(10, 20)
	}
}

// ExampleAdd demonstrates example tests (appear in documentation)
func ExampleAdd() {
	result := add(3, 4)
	fmt.Println(result)
	// Output: 7
}

// TestMain demonstrates test setup/teardown
func TestMain(m *testing.M) {
	// Setup
	fmt.Println("Setting up tests...")

	// Run tests
	_ = m.Run() // In actual test file, use: os.Exit(m.Run())

	// Teardown
	fmt.Println("Cleaning up tests...")
}

// TestParallel demonstrates parallel testing
func TestParallel(t *testing.T) {
	t.Parallel()

	// This test can run in parallel with other tests
	t.Run("subtest1", func(t *testing.T) {
		t.Parallel()
		// Test code
	})

	t.Run("subtest2", func(t *testing.T) {
		t.Parallel()
		// Test code
	})
}

// TestSkip demonstrates skipping tests
func TestSkip(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test")
	}

	// Long-running test code
}

// TestCleanup demonstrates cleanup functions
func TestCleanup(t *testing.T) {
	// Setup
	t.Cleanup(func() {
		// Cleanup code runs after test
		fmt.Println("Cleaning up...")
	})

	// Test code
}

// Testing best practices:
// 1. Use table-driven tests for multiple cases
// 2. Test both success and error cases
// 3. Use descriptive test names
// 4. Keep tests simple and focused
// 5. Use subtests for organization
// 6. Benchmark performance-critical code
// 7. Use test helpers for common setup
// 8. Test edge cases and boundaries
// 9. Use testify or similar for assertions (optional)
// 10. Keep tests fast and independent

func testingSummary() {
	fmt.Println("Testing Summary:")
	fmt.Println("  - Test functions: TestXxx")
	fmt.Println("  - Benchmark functions: BenchmarkXxx")
	fmt.Println("  - Example functions: ExampleXxx")
	fmt.Println("  - Run: go test")
	fmt.Println("  - Run with coverage: go test -cover")
	fmt.Println("  - Run benchmarks: go test -bench=.")
	fmt.Println("  - Run specific test: go test -run TestName")
}

func main() {
	// This file demonstrates Go testing concepts
	// To run actual tests, use: go test
	// This main function is for demonstration purposes only
	
	fmt.Println("Go Testing Framework Demonstration")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("This file demonstrates:")
	fmt.Println("  - Test functions (TestXxx)")
	fmt.Println("  - Table-driven tests")
	fmt.Println("  - Error testing")
	fmt.Println("  - Benchmarking")
	fmt.Println("  - Example tests")
	fmt.Println("  - Test setup/teardown")
	fmt.Println("  - Parallel testing")
	fmt.Println()
	fmt.Println("To run tests, use: go test")
	fmt.Println()
	
	// Demonstrate functions
	fmt.Println("Function demonstrations:")
	fmt.Printf("  add(2, 3) = %d\n", add(2, 3))
	fmt.Printf("  subtract(5, 3) = %d\n", subtract(5, 3))
	
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("  divide(10, 2) error: %v\n", err)
	} else {
		fmt.Printf("  divide(10, 2) = %.1f\n", result)
	}
	
	testingSummary()
}
