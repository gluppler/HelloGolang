# Contributing Examples — Go

This document illustrates **practical coding principles** for contributing to this repository. All contributors should follow these practices to ensure **readability, scalability, and maintainability**, in the spirit of production-grade projects like `cURL`, `Docker`, and `Kubernetes`.

---

## 1. Functional Thinking in Go

Go is imperative, but functional principles help keep code predictable.

```go
// Bad: Mutates global state
var counter int

func Increment() {
    counter++
}

// Good: Pure function returning a new value
func Increment(n int) int {
    return n + 1
}
```

---

## 2. Readable & Idiomatic Go

Follow `Effective Go` and format everything with `gofmt`.

```go
// Bad: inconsistent style
func add(a int,b int)int{ return a+b }

// Good: idiomatic Go
func add(a, b int) int {
    return a + b
}
```

---

## 3. Concurrency Patterns

Use goroutines and channels responsibly, with `context` for cancellation.

```go
// Worker pool example
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
}
```

---

## 4. Error Handling

Always return explicit errors. Wrap with context.

```go
// Bad: panic used for normal error cases
file, err := os.Open("config.json")
if err != nil {
    panic("file not found")
}

// Good: return errors up the stack
file, err := os.Open("config.json")
if err != nil {
    return fmt.Errorf("open config.json: %w", err)
}
```

---

## 5. Testing & Benchmarks

Always write **table-driven tests**. Use benchmarks to check performance.

```go
// add_test.go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b int
        want int
    }{
        {1, 2, 3},
        {10, 20, 30},
    }

    for _, tt := range tests {
        got := add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(10, 20)
    }
}
```

---

## 6. Refactoring Principles

Prefer small packages and dependency injection.

```go
// Bad: Hard dependency
type Service struct{}

func (s Service) SaveToDB() {}

// Good: Use interfaces
type Repository interface {
    Save(data string) error
}

type Service struct {
    Repo Repository
}
```

---

## 7. Maintainable APIs

Keep APIs versioned and stable.

```go
// v1/user.go
package user

type User struct {
    Name string
    Age  int
}

// v2/user.go (non-breaking addition)
package user

type User struct {
    Name string
    Age  int
    Email string // new field
}
```

---

## 8. Performance & Memory

Profile with `pprof`, avoid allocations when possible.

```go
// Bad: Causes unnecessary allocations
func joinStrings(list []string) string {
    result := ""
    for _, s := range list {
        result += s
    }
    return result
}

// Good: Use strings.Builder
func joinStrings(list []string) string {
    var b strings.Builder
    for _, s := range list {
        b.WriteString(s)
    }
    return b.String()
}
```

---

## 9. Modularity & Scalability

Use Go modules and keep packages focused.

```bash
go mod init github.com/yourname/project
```

Structure:

```
project/
├── cmd/         # CLI entrypoints
├── pkg/         # Shared libraries
├── internal/    # Private code
└── tests/       # Integration tests
```

---

## 10. CI/CD & Style Enforcement

Automate linting and testing with GitHub Actions.

```yaml
# .github/workflows/go.yml
name: Go CI

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: 1.22
      - run: go build ./...
      - run: go test ./... -v
      - run: golangci-lint run
```

---

✅ Following these principles ensures that the project remains **scalable, idiomatic, and production-ready** for long-term maintainability.

---
