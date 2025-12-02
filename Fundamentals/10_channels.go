package main

import (
	"fmt"
	"sync"
	"time"
)

// Channels demonstrates channel operations, patterns, and best practices

func main() {
	channelBasics()
	channelDirections()
	channelSelect()
	channelPatterns()
	channelSafety()
}

// channelBasics demonstrates basic channel operations
func channelBasics() {
	// Unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Printf("Received: %d\n", value)

	// Buffered channel
	buffered := make(chan string, 3)
	buffered <- "first"
	buffered <- "second"
	buffered <- "third"
	close(buffered)

	for msg := range buffered {
		fmt.Printf("Buffered: %s\n", msg)
	}

	// Channel with capacity
	capCh := make(chan int, 5)
	fmt.Printf("Channel capacity: %d\n", cap(capCh))
	fmt.Printf("Channel length: %d\n", len(capCh))
}

// channelDirections demonstrates channel direction restrictions
func channelDirections() {
	// Bidirectional channel
	ch := make(chan int)

	// Send-only channel
	var sendOnly chan<- int = ch

	// Receive-only channel
	var receiveOnly <-chan int = ch

	go func() {
		sendOnly <- 100
		close(sendOnly)
	}()

	value := <-receiveOnly
	fmt.Printf("Received from directional channel: %d\n", value)

	// Function with channel directions
	sendToChannel(ch)
	receiveFromChannel(ch)
}

// sendToChannel accepts a send-only channel
func sendToChannel(ch chan<- int) {
	ch <- 200
}

// receiveFromChannel accepts a receive-only channel
func receiveFromChannel(ch <-chan int) {
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

// channelSelect demonstrates select statement
func channelSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send to channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Select on multiple channels
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	}

	// Select with default (non-blocking)
	select {
	case msg := <-ch1:
		fmt.Printf("Got: %s\n", msg)
	default:
		fmt.Println("No message ready")
	}

	// Select with timeout
	select {
	case msg := <-ch1:
		fmt.Printf("Got: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

	// Select for sending
	ch3 := make(chan int, 1)
	select {
	case ch3 <- 42:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Channel full")
	}
}

// channelPatterns demonstrates common channel patterns
func channelPatterns() {
	// Pattern 1: Generator
	numbers := numberGenerator(5)
	for n := range numbers {
		fmt.Printf("Generated: %d\n", n)
	}

	// Pattern 2: Multiplexing (fan-in)
	ch1 := numberGenerator(3)
	ch2 := numberGenerator(3)
	multiplexed := fanIn(ch1, ch2)

	for i := 0; i < 6; i++ {
		fmt.Printf("Multiplexed: %d\n", <-multiplexed)
	}

	// Pattern 3: Demultiplexing (fan-out)
	input := numberGenerator(10)
	output1 := make(chan int)
	output2 := make(chan int)

	go fanOut(input, output1, output2)

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

	// Pattern 4: Pipeline
	pipeline()

	// Pattern 5: Request-Response
	requestResponse()
}

// numberGenerator generates numbers up to n
func numberGenerator(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

// fanIn combines multiple channels into one
func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}

// fanOut distributes input to multiple outputs
func fanOut(input <-chan int, output1, output2 chan<- int) {
	defer close(output1)
	defer close(output2)

	for val := range input {
		select {
		case output1 <- val:
		case output2 <- val:
		}
	}
}

// pipeline demonstrates a processing pipeline
func pipeline() {
	// Stage 1: Generate
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Square
	squares := make(chan int)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()

	// Stage 3: Double
	doubled := make(chan int)
	go func() {
		defer close(doubled)
		for sq := range squares {
			doubled <- sq * 2
		}
	}()

	// Collect results
	for result := range doubled {
		fmt.Printf("Pipeline result: %d\n", result)
	}
}

// requestResponse demonstrates request-response pattern
func requestResponse() {
	type Request struct {
		ID   int
		Data string
	}

	type Response struct {
		ID     int
		Result string
	}

	requests := make(chan Request)
	responses := make(chan Response)

	// Worker
	go func() {
		for req := range requests {
			responses <- Response{
				ID:     req.ID,
				Result: fmt.Sprintf("Processed: %s", req.Data),
			}
		}
	}()

	// Send requests
	go func() {
		defer close(requests)
		for i := 1; i <= 3; i++ {
			requests <- Request{
				ID:   i,
				Data: fmt.Sprintf("data%d", i),
			}
		}
	}()

	// Receive responses
	for i := 0; i < 3; i++ {
		resp := <-responses
		fmt.Printf("Response: %+v\n", resp)
	}
}

// channelSafety demonstrates safe channel operations
func channelSafety() {
	// Secure: Always check if channel is closed
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}

	// Secure: Prevent sending on closed channel
	safeSend()

	// Secure: Prevent receiving from nil channel
	safeReceive()

	// Secure: Use context for cancellation
	cancellableOperation()
}

// safeSend demonstrates safe sending
func safeSend() {
	ch := make(chan int, 1)
	closed := false

	send := func(value int) bool {
		if closed {
			return false
		}
		select {
		case ch <- value:
			return true
		default:
			return false
		}
	}

	closeChannel := func() {
		if !closed {
			close(ch)
			closed = true
		}
	}

	send(42)
	closeChannel()

	if !send(43) {
		fmt.Println("Cannot send on closed channel")
	}
}

// safeReceive demonstrates safe receiving
func safeReceive() {
	var ch chan int = nil

	// Secure: nil channel receive blocks forever, use select with default
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("Channel is nil, skipping receive")
	}
}

// cancellableOperation demonstrates cancellation pattern
func cancellableOperation() {
	done := make(chan struct{})
	data := make(chan int)

	// Producer
	go func() {
		defer close(data)
		for i := 0; i < 10; i++ {
			select {
			case data <- i:
			case <-done:
				fmt.Println("Producer cancelled")
				return
			}
		}
	}()

	// Consumer with cancellation
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(done)
	}()

	// Process data
	for value := range data {
		fmt.Printf("Processing: %d\n", value)
	}
}

// channelBestPractices demonstrates best practices
func channelBestPractices() {
	// Practice 1: Close channels to signal completion
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Printf("Value: %d\n", val)
	}

	// Practice 2: Use buffered channels when appropriate
	buffered := make(chan string, 10)
	for i := 0; i < 5; i++ {
		buffered <- fmt.Sprintf("item%d", i)
	}
	close(buffered)

	// Practice 3: Use select for non-blocking operations
	nonBlocking()

	// Practice 4: Use context for timeouts and cancellation
	withTimeout()
}

// nonBlocking demonstrates non-blocking operations
func nonBlocking() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Channel full, operation would block")
	}

	select {
	case val := <-ch:
		fmt.Printf("Received: %d\n", val)
	default:
		fmt.Println("No value available")
	}
}

// withTimeout demonstrates timeout pattern
func withTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Printf("Got result: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
}
