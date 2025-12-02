package main

import (
	"fmt"
	"sync"
	"testing"
)

// Advanced Testing demonstrates advanced testing techniques

func main() {
	fmt.Println("Advanced testing examples")
	fmt.Println("Run: go test -v")
}

// TestTableDriven demonstrates table-driven tests
func TestTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		wantErr  bool
	}{
		{"positive number", 5, 25, false},
		{"zero", 0, 0, false},
		{"negative number", -3, 9, false},
		{"large number", 100, 10000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := square(tt.input)
			if result != tt.expected {
				t.Errorf("square(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// square squares a number
func square(n int) int {
	return n * n
}

// TestSubtests demonstrates subtests
func TestSubtests(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		result := add(2, 3)
		if result != 5 {
			t.Errorf("add(2, 3) = %d; expected 5", result)
		}
	})

	t.Run("subtraction", func(t *testing.T) {
		result := subtract(5, 3)
		if result != 2 {
			t.Errorf("subtract(5, 3) = %d; expected 2", result)
		}
	})

	t.Run("multiplication", func(t *testing.T) {
		result := multiply(4, 5)
		if result != 20 {
			t.Errorf("multiply(4, 5) = %d; expected 20", result)
		}
	})
}

// add performs addition
func add(a, b int) int {
	return a + b
}

// subtract performs subtraction
func subtract(a, b int) int {
	return a - b
}

// multiply performs multiplication
func multiply(a, b int) int {
	return a * b
}

// TestParallel demonstrates parallel testing
func TestParallel(t *testing.T) {
	t.Parallel()

	t.Run("test1", func(t *testing.T) {
		t.Parallel()
		// Test code
	})

	t.Run("test2", func(t *testing.T) {
		t.Parallel()
		// Test code
	})
}

// TestBenchmark demonstrates benchmarking
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		square(42)
	}
}

// TestBenchmarkComparison compares different implementations
func BenchmarkSquareComparison(b *testing.B) {
	b.Run("direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := 42 * 42
			_ = result
		}
	})

	b.Run("function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := square(42)
			_ = result
		}
	})
}

// TestFuzz demonstrates fuzzing (Go 1.18+)
func FuzzSquare(f *testing.F) {
	// Add seed corpus
	f.Add(5)
	f.Add(-3)
	f.Add(0)

	f.Fuzz(func(t *testing.T, n int) {
		result := square(n)

		// Properties to check
		if result < 0 {
			t.Errorf("square(%d) = %d; expected non-negative", n, result)
		}

		if n != 0 && result == 0 {
			t.Errorf("square(%d) = %d; expected non-zero for non-zero input", n, result)
		}
	})
}

// TestHelpers demonstrates test helpers
func TestHelpers(t *testing.T) {
	// Helper function
	assertEqual := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d; want %d", got, want)
		}
	}

	assertEqual(t, add(2, 3), 5)
	assertEqual(t, multiply(4, 5), 20)
}

// TestCleanup demonstrates cleanup functions
func TestCleanup(t *testing.T) {
	// Setup
	t.Cleanup(func() {
		// Cleanup code
		fmt.Println("Cleaning up test")
	})

	// Test code
}

// TestSkip demonstrates skipping tests
func TestSkip(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long test")
	}

	// Long-running test
}

// TestErrorCases demonstrates error case testing
func TestErrorCases(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr bool
	}{
		{"valid input", 10, false},
		{"zero input", 0, true},
		{"negative input", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePositive(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePositive(%d) error = %v; wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

// validatePositive validates positive number
func validatePositive(n int) error {
	if n <= 0 {
		return fmt.Errorf("number must be positive: %d", n)
	}
	return nil
}

// TestConcurrency demonstrates testing concurrent code
func TestConcurrency(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Concurrent increments
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != 100 {
		t.Errorf("counter = %d; expected 100", counter)
	}
}
