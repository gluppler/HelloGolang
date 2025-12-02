package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Advanced Concurrency demonstrates advanced concurrency patterns and techniques

func main() {
	workerPoolAdvanced()
	rateLimitingAdvanced()
	circuitBreaker()
	semaphorePattern()
	barrierPattern()
	atomicOperations()
	runtimeControl()
	contextPropagation()
}

// workerPoolAdvanced demonstrates advanced worker pool with dynamic scaling
func workerPoolAdvanced() {
	type Job struct {
		ID     int
		Data   string
		Result chan string
	}

	jobs := make(chan Job, 100)
	results := make(chan string, 100)

	// Dynamic worker pool
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Simulate work
				result := fmt.Sprintf("Worker %d processed job %d: %s", workerID, job.ID, job.Data)
				results <- result
				job.Result <- result
			}
		}(i)
	}

	// Submit jobs
	go func() {
		defer close(jobs)
		for i := 0; i < 10; i++ {
			jobs <- Job{
				ID:     i,
				Data:   fmt.Sprintf("data-%d", i),
				Result: make(chan string, 1),
			}
		}
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for result := range results {
		fmt.Printf("Result: %s\n", result)
	}
}

// rateLimitingAdvanced demonstrates advanced rate limiting patterns
func rateLimitingAdvanced() {
	// Token bucket rate limiter
	type TokenBucket struct {
		tokens     int64
		capacity   int64
		refillRate time.Duration
		mu         sync.Mutex
		lastRefill time.Time
	}

	newTokenBucket := func(capacity int64, refillRate time.Duration) *TokenBucket {
		return &TokenBucket{
			tokens:     capacity,
			capacity:   capacity,
			refillRate: refillRate,
			lastRefill: time.Now(),
		}
	}

	takeToken := func(tb *TokenBucket) bool {
		tb.mu.Lock()
		defer tb.mu.Unlock()

		// Refill tokens
		now := time.Now()
		elapsed := now.Sub(tb.lastRefill)
		tokensToAdd := int64(elapsed / tb.refillRate)

		if tokensToAdd > 0 {
			tb.tokens = min(tb.tokens+tokensToAdd, tb.capacity)
			tb.lastRefill = now
		}

		if tb.tokens > 0 {
			tb.tokens--
			return true
		}
		return false
	}

	bucket := newTokenBucket(5, 100*time.Millisecond)

	// Test rate limiting
	for i := 0; i < 10; i++ {
		if takeToken(bucket) {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// min returns the minimum of two int64 values
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// circuitBreaker demonstrates circuit breaker pattern
func circuitBreaker() {
	type CircuitBreaker struct {
		failures     int64
		successes    int64
		state        int32 // 0: closed, 1: open, 2: half-open
		maxFailures  int64
		resetTimeout time.Duration
		lastFailTime time.Time
		mu           sync.RWMutex
	}

	cb := &CircuitBreaker{
		maxFailures:  5,
		resetTimeout: 1 * time.Second,
		state:        0, // closed
	}

	call := func() error {
		cb.mu.RLock()
		state := atomic.LoadInt32(&cb.state)
		cb.mu.RUnlock()

		if state == 1 { // open
			if time.Since(cb.lastFailTime) > cb.resetTimeout {
				atomic.StoreInt32(&cb.state, 2) // half-open
				fmt.Println("Circuit breaker: half-open")
			} else {
				return fmt.Errorf("circuit breaker is open")
			}
		}

		// Simulate operation
		err := simulateOperation()

		cb.mu.Lock()
		defer cb.mu.Unlock()

		if err != nil {
			atomic.AddInt64(&cb.failures, 1)
			cb.lastFailTime = time.Now()

			if atomic.LoadInt64(&cb.failures) >= cb.maxFailures {
				atomic.StoreInt32(&cb.state, 1) // open
				fmt.Println("Circuit breaker: opened")
			}
			return err
		}

		// Success
		atomic.StoreInt64(&cb.successes, 1)
		atomic.StoreInt64(&cb.failures, 0)
		atomic.StoreInt32(&cb.state, 0) // closed
		return nil
	}

	// Test circuit breaker
	for i := 0; i < 10; i++ {
		if err := call(); err != nil {
			fmt.Printf("Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// simulateOperation simulates an operation that might fail
func simulateOperation() error {
	// Simulate 30% failure rate
	if time.Now().UnixNano()%10 < 3 {
		return fmt.Errorf("operation failed")
	}
	return nil
}

// semaphorePattern demonstrates semaphore pattern for resource limiting
func semaphorePattern() {
	type Semaphore struct {
		ch chan struct{}
	}

	newSemaphore := func(limit int) *Semaphore {
		return &Semaphore{
			ch: make(chan struct{}, limit),
		}
	}

	acquire := func(s *Semaphore) {
		s.ch <- struct{}{}
	}

	release := func(s *Semaphore) {
		<-s.ch
	}

	sem := newSemaphore(3) // Allow 3 concurrent operations

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			acquire(sem)
			defer release(sem)

			fmt.Printf("Task %d: Acquired semaphore\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Task %d: Released semaphore\n", id)
		}(i)
	}

	wg.Wait()
}

// barrierPattern demonstrates barrier synchronization pattern
func barrierPattern() {
	type Barrier struct {
		count    int
		required int
		mu       sync.Mutex
		cond     *sync.Cond
	}

	newBarrier := func(required int) *Barrier {
		b := &Barrier{
			required: required,
		}
		b.cond = sync.NewCond(&b.mu)
		return b
	}

	wait := func(b *Barrier) {
		b.mu.Lock()
		b.count++
		count := b.count
		b.mu.Unlock()

		if count >= b.required {
			b.cond.Broadcast()
		} else {
			b.mu.Lock()
			for b.count < b.required {
				b.cond.Wait()
			}
			b.mu.Unlock()
		}
	}

	barrier := newBarrier(5)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Waiting at barrier\n", id)
			wait(barrier)
			fmt.Printf("Goroutine %d: Passed barrier\n", id)
		}(i)
	}

	wg.Wait()
}

// atomicOperations demonstrates atomic operations for lock-free programming
func atomicOperations() {
	var counter int64
	var wg sync.WaitGroup

	// Atomic increment
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic counter: %d\n", atomic.LoadInt64(&counter))

	// Atomic compare and swap
	var value int64 = 10
	old := atomic.LoadInt64(&value)
	newVal := int64(20)

	swapped := atomic.CompareAndSwapInt64(&value, old, newVal)
	fmt.Printf("CAS swapped: %t, value: %d\n", swapped, atomic.LoadInt64(&value))

	// Atomic value for complex types
	type Config struct {
		Host string
		Port int
	}

	var config atomic.Value
	config.Store(Config{Host: "localhost", Port: 8080})

	stored := config.Load().(Config)
	fmt.Printf("Stored config: %+v\n", stored)
}

// runtimeControl demonstrates runtime control and monitoring
func runtimeControl() {
	// Get runtime stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Runtime stats:\n")
	fmt.Printf("  Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("  CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("  Memory allocated: %d KB\n", m.Alloc/1024)
	fmt.Printf("  Total allocations: %d\n", m.Mallocs)

	// Force GC
	runtime.GC()

	// Set max procs
	oldProcs := runtime.GOMAXPROCS(4)
	fmt.Printf("  Old GOMAXPROCS: %d, New: 4\n", oldProcs)

	// Goexit from goroutine
	go func() {
		fmt.Println("Goroutine will exit")
		runtime.Goexit()
		fmt.Println("This won't print")
	}()

	time.Sleep(100 * time.Millisecond)
}

// contextPropagation demonstrates advanced context propagation
func contextPropagation() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Propagate context through layers
	processWithContext(ctx, "request-1")

	// Context with cancellation
	ctx2, cancel2 := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		processWithContext(ctx2, "request-2")
	}()

	time.Sleep(500 * time.Millisecond)
	cancel2()
	wg.Wait()
}

// processWithContext processes a request with context
func processWithContext(ctx context.Context, requestID string) {
	// Check context at start
	if ctx.Err() != nil {
		fmt.Printf("Request %s: Context already cancelled\n", requestID)
		return
	}

	// Simulate work with periodic context checks
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Request %s: Cancelled at step %d\n", requestID, i)
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("Request %s: Step %d completed\n", requestID, i)
		}
	}

	fmt.Printf("Request %s: Completed\n", requestID)
}
