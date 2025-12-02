package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Goroutines demonstrates goroutines, synchronization, and concurrency patterns

func main() {
	basicGoroutines()
	goroutineSynchronization()
	goroutineCommunication()
	goroutinePatterns()
	goroutineBestPractices()
}

// basicGoroutines demonstrates basic goroutine usage
func basicGoroutines() {
	// Simple goroutine
	go func() {
		fmt.Println("Goroutine 1: Hello from goroutine")
	}()

	// Named function as goroutine
	go sayHello("Alice")

	// Give goroutines time to execute
	time.Sleep(100 * time.Millisecond)

	// Multiple goroutines
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: running\n", id)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

// sayHello prints a greeting
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// goroutineSynchronization demonstrates synchronization primitives
func goroutineSynchronization() {
	// WaitGroup for waiting goroutines
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d: starting\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d: done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")

	// Mutex for protecting shared data
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter)

	// RWMutex for read-write locks
	var rwmu sync.RWMutex
	data := make(map[string]int)

	// Writers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwmu.Lock()
			data[fmt.Sprintf("key%d", id)] = id
			rwmu.Unlock()
		}(i)
	}

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwmu.RLock()
			_ = len(data)
			rwmu.RUnlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Data map: %v\n", data)
}

// goroutineCommunication demonstrates communication between goroutines
func goroutineCommunication() {
	// Channel basics
	ch := make(chan string)

	go func() {
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}()

	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
	}

	// Buffered channel
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	close(buffered)

	for val := range buffered {
		fmt.Printf("Buffered value: %d\n", val)
	}

	// Channel directions
	sender := make(chan<- int)   // send-only
	receiver := make(<-chan int) // receive-only
	_ = sender
	_ = receiver
}

// goroutinePatterns demonstrates common goroutine patterns
func goroutinePatterns() {
	// Pattern 1: Worker pool
	workerPool()

	// Pattern 2: Pipeline
	pipeline()

	// Pattern 3: Fan-out/Fan-in
	fanOutFanIn()

	// Pattern 4: Timeout pattern
	timeoutPattern()

	// Pattern 5: Rate limiting
	rateLimiting()
}

// workerPool demonstrates worker pool pattern
func workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	var wg sync.WaitGroup
	numWorkers := 3

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// pipeline demonstrates pipeline pattern
func pipeline() {
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()

	// Stage 3: Print results
	for sq := range squares {
		fmt.Printf("Squared: %d\n", sq)
	}
}

// fanOutFanIn demonstrates fan-out/fan-in pattern
func fanOutFanIn() {
	input := make(chan int)

	// Fan-out: multiple workers
	output1 := make(chan int)
	output2 := make(chan int)

	go func() {
		defer close(output1)
		for n := range input {
			output1 <- n * 2
		}
	}()

	go func() {
		defer close(output2)
		for n := range input {
			output2 <- n * 3
		}
	}()

	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()

	// Fan-in: collect results
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for val := range output1 {
			fmt.Printf("Output1: %d\n", val)
		}
	}()

	go func() {
		defer wg.Done()
		for val := range output2 {
			fmt.Printf("Output2: %d\n", val)
		}
	}()

	wg.Wait()
}

// timeoutPattern demonstrates timeout pattern
func timeoutPattern() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
}

// rateLimiting demonstrates rate limiting
func rateLimiting() {
	requests := make(chan int, 5)

	// Rate limiter: allow 1 request per 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Printf("Processing request %d\n", req)
	}
}

// goroutineBestPractices demonstrates best practices
func goroutineBestPractices() {
	// Practice 1: Always use WaitGroup or channels to wait
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Properly synchronized goroutine")
	}()
	wg.Wait()

	// Practice 2: Use context for cancellation
	ctx := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		select {
		case <-ctx:
			fmt.Println("Goroutine cancelled")
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Goroutine completed")
		}
	}()

	close(ctx)
	<-done

	// Practice 3: Avoid goroutine leaks
	leakPrevention()

	// Practice 4: Use sync.Once for one-time initialization
	var once sync.Once
	initFunc := func() {
		fmt.Println("Initialized once")
	}

	for i := 0; i < 3; i++ {
		go once.Do(initFunc)
	}
	time.Sleep(100 * time.Millisecond)

	// Practice 5: Check GOMAXPROCS
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
}

// leakPrevention demonstrates preventing goroutine leaks
func leakPrevention() {
	// Secure: always ensure goroutines can exit
	done := make(chan struct{})
	result := make(chan int)

	go func() {
		defer close(result)
		select {
		case <-time.After(100 * time.Millisecond):
			result <- 42
		case <-done:
			return
		}
	}()

	// Cancel if needed
	close(done)

	// Try to receive with timeout
	select {
	case val := <-result:
		fmt.Printf("Got result: %d\n", val)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("No result received")
	}
}
