package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// Performance Optimization demonstrates optimization techniques

func main() {
	memoryOptimization()
	cpuOptimization()
	allocationOptimization()
	cachingPatterns()
	poolingPatterns()
	profilingTechniques()
}

// memoryOptimization demonstrates memory optimization techniques
func memoryOptimization() {
	// 1. Pre-allocate slices with known capacity
	preallocated := make([]int, 0, 1000) // length 0, capacity 1000
	fmt.Printf("Pre-allocated slice: len=%d, cap=%d\n", len(preallocated), cap(preallocated))

	// 2. Reuse slices
	var reusable []int
	for i := 0; i < 10; i++ {
		reusable = reusable[:0] // Reset length, keep capacity
		reusable = append(reusable, i)
	}

	// 3. Use struct instead of map for small key-value pairs
	type SmallStruct struct {
		Key   string
		Value int
	}

	structSize := unsafe.Sizeof(SmallStruct{})
	mapOverhead := 8 + 8 // approximate map overhead
	fmt.Printf("Struct size: %d bytes, Map overhead: ~%d bytes\n", structSize, mapOverhead)

	// 4. Avoid unnecessary allocations
	// Bad: creates new string
	bad := "prefix" + "suffix"
	_ = bad

	// Good: use strings.Builder for multiple concatenations
	var builder strings.Builder
	builder.WriteString("prefix")
	builder.WriteString("suffix")
	good := builder.String()
	_ = good
}

// cpuOptimization demonstrates CPU optimization techniques
func cpuOptimization() {
	// 1. Use bitwise operations where appropriate
	isEven := func(n int) bool {
		return n&1 == 0 // Faster than n%2 == 0
	}
	fmt.Printf("Is 4 even: %t\n", isEven(4))

	// 2. Avoid unnecessary function calls in loops
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	// Bad: function call in loop
	sum1 := 0
	for i := 0; i < len(numbers); i++ {
		sum1 += calculate(numbers[i]) // Function call overhead
	}

	// Good: inline calculation
	sum2 := 0
	for i := 0; i < len(numbers); i++ {
		sum2 += numbers[i] * 2 // Inline calculation
	}

	_ = sum1
	_ = sum2

	// 3. Use local variables to avoid repeated lookups
	type Config struct {
		Host string
		Port int
	}

	config := Config{Host: "localhost", Port: 8080}

	// Bad: repeated field access
	for i := 0; i < 1000; i++ {
		_ = config.Host // Repeated lookup
	}

	// Good: cache in local variable
	host := config.Host
	for i := 0; i < 1000; i++ {
		_ = host // Local variable access
	}
}

// calculate performs a calculation
func calculate(n int) int {
	return n * 2
}

// allocationOptimization demonstrates allocation optimization
func allocationOptimization() {
	// 1. Object pooling for frequently allocated objects
	type Buffer struct {
		data []byte
	}

	bufferPool := sync.Pool{
		New: func() interface{} {
			return &Buffer{
				data: make([]byte, 0, 1024),
			}
		},
	}

	// Get buffer from pool
	buf := bufferPool.Get().(*Buffer)
	buf.data = buf.data[:0] // Reset

	// Use buffer
	buf.data = append(buf.data, []byte("hello")...)

	// Return to pool
	bufferPool.Put(buf)

	// 2. Reuse byte slices
	var byteBuf []byte
	for i := 0; i < 10; i++ {
		byteBuf = byteBuf[:0] // Reset length
		byteBuf = append(byteBuf, byte('a'+i))
	}

	// 3. Avoid allocations in hot paths
	// Bad: allocates new slice
	processBad := func(data []int) {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		_ = dataCopy
	}

	// Good: work with original or use pool
	processGood := func(data []int) {
		// Work directly with data
		for i := range data {
			data[i] *= 2
		}
	}

	testData := []int{1, 2, 3, 4, 5}
	processBad(testData)
	processGood(testData)
}

// cachingPatterns demonstrates caching patterns
func cachingPatterns() {
	// 1. Simple in-memory cache with TTL
	type CacheEntry struct {
		Value     interface{}
		ExpiresAt time.Time
	}

	cache := make(map[string]*CacheEntry)
	var mu sync.RWMutex

	get := func(key string) (interface{}, bool) {
		mu.RLock()
		defer mu.RUnlock()

		entry, ok := cache[key]
		if !ok {
			return nil, false
		}

		if time.Now().After(entry.ExpiresAt) {
			return nil, false
		}

		return entry.Value, true
	}

	set := func(key string, value interface{}, ttl time.Duration) {
		mu.Lock()
		defer mu.Unlock()

		cache[key] = &CacheEntry{
			Value:     value,
			ExpiresAt: time.Now().Add(ttl),
		}
	}

	// Use cache
	set("key1", "value1", 5*time.Second)
	if val, ok := get("key1"); ok {
		fmt.Printf("Cache hit: %v\n", val)
	}

	// 2. LRU Cache implementation
	type LRUNode struct {
		Key   string
		Value interface{}
		Prev  *LRUNode
		Next  *LRUNode
	}

	type LRUCache struct {
		capacity int
		cache    map[string]*LRUNode
		head     *LRUNode
		tail     *LRUNode
		mu       sync.Mutex
	}

	newLRU := func(capacity int) *LRUCache {
		return &LRUCache{
			capacity: capacity,
			cache:    make(map[string]*LRUNode),
		}
	}

	_ = newLRU(10)
}

// poolingPatterns demonstrates object pooling patterns
func poolingPatterns() {
	// 1. Worker pool for goroutines
	type WorkerPool struct {
		workers   int
		taskQueue chan func()
		wg        sync.WaitGroup
	}

	newWorkerPool := func(workers int) *WorkerPool {
		wp := &WorkerPool{
			workers:   workers,
			taskQueue: make(chan func(), 100),
		}

		for i := 0; i < workers; i++ {
			wp.wg.Add(1)
			go func() {
				defer wp.wg.Done()
				for task := range wp.taskQueue {
					task()
				}
			}()
		}

		return wp
	}

	pool := newWorkerPool(5)

	// Submit tasks
	for i := 0; i < 10; i++ {
		taskID := i
		pool.taskQueue <- func() {
			fmt.Printf("Task %d executed\n", taskID)
		}
	}

	close(pool.taskQueue)
	pool.wg.Wait()

	// 2. Buffer pool
	bufferPool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 4096)
		},
	}

	// Get buffer
	buf := bufferPool.Get().([]byte)
	buf = buf[:0] // Reset

	// Use buffer
	buf = append(buf, []byte("data")...)

	// Return buffer
	bufferPool.Put(buf)
}

// profilingTechniques demonstrates profiling techniques
func profilingTechniques() {
	// 1. Memory profiling
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Memory stats:\n")
	fmt.Printf("  Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("  TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("  Sys: %d KB\n", m.Sys/1024)
	fmt.Printf("  NumGC: %d\n", m.NumGC)

	// 2. Force GC for testing
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("  After GC - Alloc: %d KB\n", m.Alloc/1024)

	// 3. CPU profiling (use pprof in production)
	// import _ "net/http/pprof"
	// go tool pprof http://localhost:6060/debug/pprof/profile

	// 4. Benchmark timing
	start := time.Now()
	// ... operation ...
	elapsed := time.Since(start)
	fmt.Printf("Operation took: %v\n", elapsed)

	// 5. Goroutine profiling
	fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
}
