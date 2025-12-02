package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Advanced Channels demonstrates advanced channel patterns and techniques

func main() {
	channelPipelines()
	channelOrPattern()
	channelMergePattern()
	channelBroadcast()
	channelTimeout()
	channelBackpressure()
	channelCancellation()
}

// channelPipelines demonstrates complex pipeline patterns
func channelPipelines() {
	// Multi-stage pipeline
	numbers := generateNumbers(10)
	squared := square(numbers)
	doubled := double(squared)
	filtered := filterEven(doubled)

	fmt.Println("Pipeline results:")
	for val := range filtered {
		fmt.Printf("  %d\n", val)
	}
}

// generateNumbers generates numbers
func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			out <- i
		}
	}()
	return out
}

// square squares numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// double doubles numbers
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// filterEven filters even numbers
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

// channelOrPattern demonstrates "or" pattern for multiple channels
func channelOrPattern() {
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		out := make(chan interface{})
		var wg sync.WaitGroup

		for _, ch := range channels {
			wg.Add(1)
			go func(c <-chan interface{}) {
				defer wg.Done()
				for val := range c {
					select {
					case out <- val:
					case <-out:
						return
					}
				}
			}(ch)
		}

		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
		close(ch1)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from ch2"
		close(ch2)
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch3 <- "from ch3"
		close(ch3)
	}()

	merged := or(ch1, ch2, ch3)
	fmt.Println("Or pattern results:")
	for val := range merged {
		fmt.Printf("  Received: %v\n", val)
	}
}

// channelMergePattern demonstrates merging multiple channels
func channelMergePattern() {
	merge := func(channels ...<-chan int) <-chan int {
		out := make(chan int)
		var wg sync.WaitGroup

		output := func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	ch1 := generateNumbers(5)
	ch2 := generateNumbers(5)
	ch3 := generateNumbers(5)

	merged := merge(ch1, ch2, ch3)

	fmt.Println("Merged channel results:")
	count := 0
	for val := range merged {
		fmt.Printf("  %d\n", val)
		count++
		if count >= 15 {
			break
		}
	}
}

// channelBroadcast demonstrates broadcast pattern
func channelBroadcast() {
	type Broadcaster struct {
		mu        sync.RWMutex
		listeners []chan string
	}

	broadcaster := &Broadcaster{
		listeners: make([]chan string, 0),
	}

	subscribe := func(b *Broadcaster) <-chan string {
		b.mu.Lock()
		defer b.mu.Unlock()

		ch := make(chan string, 1)
		b.listeners = append(b.listeners, ch)
		return ch
	}

	broadcast := func(b *Broadcaster, message string) {
		b.mu.RLock()
		defer b.mu.RUnlock()

		for _, listener := range b.listeners {
			select {
			case listener <- message:
			default:
				// Skip if listener is full
			}
		}
	}

	// Subscribe listeners
	ch1 := subscribe(broadcaster)
	ch2 := subscribe(broadcaster)
	ch3 := subscribe(broadcaster)

	// Broadcast message
	broadcast(broadcaster, "Hello, subscribers!")

	// Receive from listeners
	fmt.Println("Broadcast results:")
	fmt.Printf("  Listener 1: %s\n", <-ch1)
	fmt.Printf("  Listener 2: %s\n", <-ch2)
	fmt.Printf("  Listener 3: %s\n", <-ch3)
}

// channelTimeout demonstrates timeout patterns
func channelTimeout() {
	// Operation with timeout
	doWork := func() <-chan string {
		result := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			result <- "work completed"
		}()
		return result
	}

	select {
	case result := <-doWork():
		fmt.Printf("Work result: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Work timed out")
	}

	// Context-based timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	workCh := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		workCh <- "work done"
	}()

	select {
	case result := <-workCh:
		fmt.Printf("Context work result: %s\n", result)
	case <-ctx.Done():
		fmt.Printf("Context work cancelled: %v\n", ctx.Err())
	}
}

// channelBackpressure demonstrates backpressure handling
func channelBackpressure() {
	// Producer with backpressure
	producer := func(out chan<- int) {
		defer close(out)
		for i := 0; i < 100; i++ {
			select {
			case out <- i:
				// Successfully sent
			case <-time.After(10 * time.Millisecond):
				// Backpressure: channel full, skip or retry
				fmt.Printf("Backpressure at %d\n", i)
			}
		}
	}

	// Slow consumer
	consumer := func(in <-chan int) {
		for val := range in {
			fmt.Printf("Consumed: %d\n", val)
			time.Sleep(50 * time.Millisecond) // Slow processing
		}
	}

	ch := make(chan int, 5) // Small buffer creates backpressure

	go producer(ch)
	consumer(ch)
}

// channelCancellation demonstrates cancellation patterns
func channelCancellation() {
	// Cancellable generator
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; ; i++ {
				select {
				case out <- i:
				case <-ctx.Done():
					fmt.Println("Generator cancelled")
					return
				}
			}
		}()
		return out
	}

	ctx, cancel := context.WithCancel(context.Background())

	numbers := generator(ctx)

	// Consume for a bit
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	fmt.Println("Cancellable generator:")
	count := 0
	for val := range numbers {
		fmt.Printf("  %d\n", val)
		count++
		if count >= 10 {
			break
		}
	}

	// Graceful shutdown pattern
	shutdown := make(chan struct{})
	work := make(chan int, 10)

	// Worker
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case job := <-work:
				fmt.Printf("Processing job: %d\n", job)
			case <-shutdown:
				fmt.Println("Worker shutting down")
				return
			}
		}
	}()

	// Send some work
	for i := 0; i < 5; i++ {
		work <- i
	}

	// Shutdown
	close(shutdown)
	wg.Wait()
}
