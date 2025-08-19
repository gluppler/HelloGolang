# Contributing Examples (Go)

This document shows how to apply our coding principles when working with Go projects.
Use these examples as a guide to keep contributions clean, scalable, and easy to maintain.

---

## 1. Functional Principles

Prefer pure functions that avoid side effects and are easy to test.

```go
// Bad: relies on global state
var counter int

func Increment() int {
    counter++
    return counter
}

// Good: pure function with explicit input/output
func Increment(n int) int {
    return n + 1
}
```

---

## 2. Readability First

Code should be easy to follow for any contributor.

```go
// Bad: cryptic and unclear
func prc(i int) int { return i<<1 + 3 }

// Good: descriptive names, simple logic
func ProcessValue(value int) int {
    return (value * 2) + 3
}
```

---

## 3. Unit Tests Everywhere

Every feature should have test coverage.

```go
// file: mathutil.go
func Square(n int) int {
    return n * n
}

// file: mathutil_test.go
package main

import "testing"

func TestSquare(t *testing.T) {
    got := Square(4)
    want := 16

    if got != want {
        t.Errorf("Square(4) = %d; want %d", got, want)
    }
}
```

---

## 4. Refactor Continuously

Keep code modular and maintainable.

```go
// Bad: tightly coupled
func HandleUser(name string, age int) string {
    return fmt.Sprintf("%s is %d years old", name, age)
}

// Good: separate concerns
type User struct {
    Name string
    Age  int
}

func (u User) Info() string {
    return fmt.Sprintf("%s is %d years old", u.Name, u.Age)
}
```

---

## 5. Error Handling

Always check and propagate errors properly.

```go
// Bad: ignoring errors
data, _ := os.ReadFile("file.txt")

// Good: explicit error handling
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatalf("failed to read file: %v", err)
}
```

---

## 6. Idiomatic Go

Follow Go conventions: use short variable names in small scopes, longer names in wider scopes.

```go
// Good practice
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

userName := "gabe"
fmt.Println(userName)
```

---

## 7. Keep Functions Small

Each function should do one thing well.

```go
// Bad: mixed logic
func SaveUser(u User) {
    // validate
    if u.Name == "" { return }
    // save to db
    db.Save(u)
    // send email
    email.Send(u.Email)
}

// Good: single responsibility
func ValidateUser(u User) bool { return u.Name != "" }
func SaveToDB(u User) { db.Save(u) }
func NotifyUser(u User) { email.Send(u.Email) }
```

---

## 8. Documentation & Comments

Explain intent, not obvious code.

```go
// Bad: redundant comment
// Add one to n
func Increment(n int) int { return n + 1 }

// Good: describes purpose
// Increment returns the next integer in sequence.
func Increment(n int) int { return n + 1 }
```

---

## 9. Avoid Premature Optimization

Write clear code first, optimize only if needed.

```go
// Bad: micro-optimized, unreadable
func FastAdd(nums []int) (sum int) {
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }
    return
}

// Good: clear and efficient enough
func Sum(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

---

## 10. Consistency

Follow the same patterns everywhere to reduce friction.
Use `gofmt` and `golangci-lint` before committing.

---

👉 These are the **Go-specific patterns** you should follow in every contribution.

---
