# Contributing Examples – Go (Golang)

This document provides **Go-specific examples** of how to apply the 10 contributing principles from `CONTRIBUTING.md`.
It is designed for developers working on **cross-platform Go projects** (Linux, macOS, Windows), at both small and enterprise scale.

---

## 1. Clarity Over Cleverness

Write code that is **readable first**, optimized later.

```go
// ❌ Bad: Too clever, unclear
result := make([]int, 0)
for i := 0; i < 10; i++ {
    result = append(result, i*i)
}

// ✅ Good: Explicit, self-explanatory
func SquaresUpTo(n int) []int {
    squares := make([]int, 0, n)
    for i := 0; i < n; i++ {
        squares = append(squares, i*i)
    }
    return squares
}
```

> ✅ Anyone reading the good version understands intent instantly.

---

## 2. Cross-Platform Awareness

Always test on **Linux + macOS + Windows**. Avoid platform-locked calls.

```go
// ❌ Bad: Hardcoded path (only works on Unix-like systems)
f, _ := os.Open("/tmp/data.txt")

// ✅ Good: Cross-platform temp directory
f, _ := os.Open(filepath.Join(os.TempDir(), "data.txt"))
```

> ✅ Uses `os.TempDir()` and `filepath.Join`, which work everywhere.

---

## 3. Minimal Dependencies

Prefer Go standard library. Only import external libraries if truly needed.

```go
// ❌ Bad: Pulling in a third-party JSON lib for simple decoding
import "github.com/some/jsonlib"

// ✅ Good: Standard library is sufficient
import "encoding/json"

type Config struct {
    Port int `json:"port"`
}

func LoadConfig(r io.Reader) (*Config, error) {
    var cfg Config
    err := json.NewDecoder(r).Decode(&cfg)
    return &cfg, err
}
```

> ✅ The stdlib is production-ready; don’t overcomplicate.

---

## 4. Strong Error Handling

Always check and propagate errors — never silently ignore them.

```go
// ❌ Bad: Ignored error
data, _ := os.ReadFile("config.json")

// ✅ Good: Explicit error handling
data, err := os.ReadFile("config.json")
if err != nil {
    return fmt.Errorf("failed to read config: %w", err)
}
```

> ✅ Explicit error propagation makes debugging in prod much easier.

---

## 5. Testing & Validation

Every function that can fail should have a **unit test**.

```go
// file: mathutils.go
func Add(a, b int) int {
    return a + b
}

// file: mathutils_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Fatalf("expected 5, got %d", result)
    }
}
```

> ✅ Go’s `testing` package is built-in; no excuses for missing tests.

---

## 6. Documentation First

Every **package, function, and type** should be documented.

```go
// Package storage provides a simple wrapper over file operations.
package storage

// Save writes data into a given file path. It overwrites existing data.
func Save(path string, data []byte) error {
    return os.WriteFile(path, data, 0644)
}
```

> ✅ Godoc comments double as API docs and onboarding guides.

---

## 7. Security by Default

Never trust input, always sanitize.

```go
// ❌ Bad: Directly running user input
cmd := exec.Command("sh", "-c", userInput)
cmd.Run()

// ✅ Good: Explicitly validated input
allowed := map[string]bool{"ls": true, "pwd": true}
if !allowed[userInput] {
    return fmt.Errorf("command not allowed")
}
cmd := exec.Command(userInput)
cmd.Run()
```

> ✅ Prevents command injection vulnerabilities.

---

## 8. Consistency in Style

Follow `gofmt` and idiomatic Go style.

```go
// ❌ Bad: Non-idiomatic, inconsistent names
func compute_value(x int) int { return x * x }

// ✅ Good: Idiomatic Go naming
func ComputeValue(x int) int { return x * x }
```

> ✅ Always run `gofmt` / `goimports`. CI should enforce this.

---

## 9. Performance Mindfulness

Don’t over-optimize prematurely, but avoid obvious inefficiencies.

```go
// ❌ Bad: Repeated string concatenation
s := ""
for i := 0; i < 1000; i++ {
    s += "data"
}

// ✅ Good: Use strings.Builder
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("data")
}
s := b.String()
```

> ✅ Efficient and still clear; good balance for enterprise use.

---

## 10. Collaboration & Review

Write code assuming **others will read, review, and maintain it**.

```go
// ❌ Bad: No context
func DoIt() {}

// ✅ Good: Clear naming + comment
// ProcessOrder validates and processes an incoming order.
func ProcessOrder(orderID string) error {
    // TODO: Add payment validation
    return nil
}
```

> ✅ Reviewers see intent immediately, making collaboration smoother.

---

# Final Notes

* Run `go test ./...` before pushing.
* Ensure `go fmt ./...` passes.
* Keep PRs **small, reviewed, and tested**.

---
