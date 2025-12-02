package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context demonstrates context package for cancellation and timeouts

func main() {
	basicContext()
	contextWithTimeout()
	contextWithCancel()
	contextWithDeadline()
	contextWithValue()
	contextBestPractices()
}

// basicContext demonstrates basic context usage
func basicContext() {
	// Background context (root context)
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO context (should be replaced)
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	// Check if context is done
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	default:
		fmt.Println("Context is active")
	}
}

// contextWithTimeout demonstrates timeout context
func contextWithTimeout() {
	// Context with 1 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Simulate work
	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Printf("Work timed out: %v\n", ctx.Err())
	}
}

// contextWithCancel demonstrates cancellation context
func contextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	// Start goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine cancelled")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Cancel after 500ms
	time.Sleep(500 * time.Millisecond)
	cancel()

	time.Sleep(100 * time.Millisecond)
}

// contextWithDeadline demonstrates deadline context
func contextWithDeadline() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Printf("Deadline exceeded: %v\n", ctx.Err())
	}
}

// contextWithValue demonstrates value context
func contextWithValue() {
	// Add values to context
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "requestID", "req-456")

	// Retrieve values
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")

	fmt.Printf("User ID: %v\n", userID)
	fmt.Printf("Request ID: %v\n", requestID)

	// Type-safe value retrieval
	if uid, ok := userID.(int); ok {
		fmt.Printf("User ID (typed): %d\n", uid)
	}
}

// contextBestPractices demonstrates best practices
func contextBestPractices() {
	// Practice 1: Pass context as first parameter
	processRequest(context.Background(), "request-1")

	// Practice 2: Check context in loops
	processWithContextCheck(context.Background())

	// Practice 3: Use context for HTTP requests
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)

	// Practice 4: Don't store context in structs
	// Pass context explicitly to functions

	// Practice 5: Use context for cancellation propagation
	propagateCancellation()
}

// processRequest demonstrates context as first parameter
func processRequest(ctx context.Context, requestID string) {
	// Check if context is cancelled
	if ctx.Err() != nil {
		return
	}

	fmt.Printf("Processing request: %s\n", requestID)

	// Simulate work with context check
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("Request %s completed\n", requestID)
	case <-ctx.Done():
		fmt.Printf("Request %s cancelled\n", requestID)
	}
}

// processWithContextCheck demonstrates context checking in loops
func processWithContextCheck(ctx context.Context) {
	for i := 0; i < 10; i++ {
		// Check context before each iteration
		if ctx.Err() != nil {
			fmt.Println("Context cancelled, stopping")
			return
		}

		fmt.Printf("Processing item %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}
}

// propagateCancellation demonstrates cancellation propagation
func propagateCancellation() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id)
		}(i)
	}

	// Cancel after 500ms
	time.Sleep(500 * time.Millisecond)
	cancel()

	wg.Wait()
}

// worker demonstrates a worker that respects context
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelled\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// Context best practices:
// 1. Pass context as first parameter
// 2. Don't store context in structs
// 3. Always check ctx.Err() in loops
// 4. Use WithTimeout/WithDeadline for timeouts
// 5. Use WithCancel for cancellation
// 6. Use WithValue sparingly (for request-scoped data)
// 7. Never pass nil context (use context.Background())
// 8. Cancel contexts to prevent leaks
// 9. Context values should be immutable
// 10. Use context for cancellation propagation

func contextSummary() {
	fmt.Println("Context Summary:")
	fmt.Println("  - Cancellation propagation")
	fmt.Println("  - Timeout management")
	fmt.Println("  - Request-scoped values")
	fmt.Println("  - First parameter convention")
	fmt.Println("  - Essential for concurrent code")
}
