package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Advanced Error Handling demonstrates advanced error handling patterns

func main() {
	errorWrapping()
	errorChains()
	customErrorTypes()
	errorRecovery()
	errorLogging()
	errorMetrics()
}

// errorWrapping demonstrates error wrapping and unwrapping
func errorWrapping() {
	// Wrap errors with context
	originalErr := errors.New("database connection failed")

	wrappedErr := fmt.Errorf("failed to fetch user: %w", originalErr)
	doubleWrapped := fmt.Errorf("user service error: %w", wrappedErr)

	fmt.Printf("Original error: %v\n", originalErr)
	fmt.Printf("Wrapped error: %v\n", wrappedErr)
	fmt.Printf("Double wrapped: %v\n", doubleWrapped)

	// Unwrap errors
	if unwrapped := errors.Unwrap(doubleWrapped); unwrapped != nil {
		fmt.Printf("Unwrapped: %v\n", unwrapped)
	}

	// Check if error is specific type
	if errors.Is(doubleWrapped, originalErr) {
		fmt.Println("Error chain contains original error")
	}

	// Extract specific error type
	var dbErr *DatabaseError
	if errors.As(doubleWrapped, &dbErr) {
		fmt.Printf("Extracted DatabaseError: %v\n", dbErr)
	}
}

// DatabaseError represents a database error
type DatabaseError struct {
	Code    string
	Message string
	Err     error
}

// Error implements error interface
func (e *DatabaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error [%s]: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("database error [%s]: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// errorChains demonstrates error chain traversal
func errorChains() {
	// Build error chain
	dbErr := &DatabaseError{
		Code:    "CONN_001",
		Message: "connection timeout",
		Err:     errors.New("network unreachable"),
	}
	
	var err error = dbErr
	err = fmt.Errorf("query failed: %w", err)
	err = fmt.Errorf("service unavailable: %w", err)
	
	// Traverse error chain
	fmt.Println("Error chain:")
	current := err
	depth := 0
	for current != nil {
		fmt.Printf("  [%d] %v\n", depth, current)
		current = errors.Unwrap(current)
		depth++
		if depth > 10 { // Prevent infinite loops
			break
		}
	}
	
	// Find specific error in chain
	if errors.Is(err, dbErr) {
		fmt.Println("Found CONN_001 in error chain")
	}
}

// customErrorTypes demonstrates custom error types
func customErrorTypes() {
	// Validation error
	validationErr := &ValidationError{
		Field:   "email",
		Message: "invalid format",
		Code:    "VAL_001",
	}

	fmt.Printf("Validation error: %v\n", validationErr)
	fmt.Printf("Error code: %s\n", validationErr.Code)

	// Business logic error
	businessErr := &BusinessError{
		Operation: "transfer",
		Reason:    "insufficient funds",
		Code:      "BIZ_001",
	}

	fmt.Printf("Business error: %v\n", businessErr)

	// Temporary error
	tempErr := &TemporaryError{
		Message:    "service temporarily unavailable",
		RetryAfter: 30,
	}

	if tempErr.Temporary() {
		fmt.Printf("Temporary error, retry after %d seconds\n", tempErr.RetryAfter)
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
	Code    string
}

// Error implements error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error [%s] on field '%s': %s", e.Code, e.Field, e.Message)
}

// BusinessError represents a business logic error
type BusinessError struct {
	Operation string
	Reason    string
	Code      string
}

// Error implements error interface
func (e *BusinessError) Error() string {
	return fmt.Sprintf("business error [%s] in operation '%s': %s", e.Code, e.Operation, e.Reason)
}

// TemporaryError represents a temporary error
type TemporaryError struct {
	Message    string
	RetryAfter int // seconds
}

// Error implements error interface
func (e *TemporaryError) Error() string {
	return e.Message
}

// Temporary indicates if error is temporary
func (e *TemporaryError) Temporary() bool {
	return true
}

// errorRecovery demonstrates error recovery patterns
func errorRecovery() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)

			// Log stack trace
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			fmt.Printf("Stack trace:\n%s\n", buf[:n])
		}
	}()

	// Safe operation with recovery
	safeOperation := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic recovered: %v", r)
			}
		}()

		// Risky operation
		riskyCode()
		return nil
	}

	if err := safeOperation(); err != nil {
		fmt.Printf("Operation error: %v\n", err)
	}
}

// riskyCode may panic
func riskyCode() {
	panic("something went wrong")
}

// errorLogging demonstrates structured error logging
func errorLogging() {
	// Structured logging
	logError := func(err error, context map[string]interface{}) {
		var fields []string
		fields = append(fields, fmt.Sprintf("error=%q", err.Error()))

		for k, v := range context {
			fields = append(fields, fmt.Sprintf("%s=%v", k, v))
		}

		log.Printf("[ERROR] %s", strings.Join(fields, " "))
	}

	err := errors.New("operation failed")
	logError(err, map[string]interface{}{
		"user_id":   123,
		"operation": "transfer",
		"amount":    100.0,
		"timestamp": "2024-01-15T10:00:00Z",
	})

	// Error with stack trace
	logErrorWithStack := func(err error) {
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		log.Printf("[ERROR] %v\nStack:\n%s", err, buf[:n])
	}

	logErrorWithStack(errors.New("critical error"))
}

// errorMetrics demonstrates error metrics and monitoring
func errorMetrics() {
	type ErrorMetrics struct {
		counts map[string]int64
		mu     sync.Mutex
	}

	metrics := &ErrorMetrics{
		counts: make(map[string]int64),
	}

	recordError := func(metrics *ErrorMetrics, err error) {
		metrics.mu.Lock()
		defer metrics.mu.Unlock()

		errorType := fmt.Sprintf("%T", err)
		metrics.counts[errorType]++
	}

	// Record some errors
	recordError(metrics, &ValidationError{Field: "email", Message: "invalid"})
	recordError(metrics, &BusinessError{Operation: "transfer", Reason: "insufficient funds"})
	recordError(metrics, &ValidationError{Field: "password", Message: "too short"})

	// Get metrics
	metrics.mu.Lock()
	fmt.Println("Error metrics:")
	for errType, count := range metrics.counts {
		fmt.Printf("  %s: %d\n", errType, count)
	}
	metrics.mu.Unlock()
}

// Advanced error handling patterns
func advancedErrorPatterns() {
	// Pattern 1: Error aggregation
	type ErrorList []error

	aggregateErrors := func(errs []error) error {
		if len(errs) == 0 {
			return nil
		}
		if len(errs) == 1 {
			return errs[0]
		}

		var messages []string
		for _, err := range errs {
			messages = append(messages, err.Error())
		}
		return fmt.Errorf("multiple errors: %s", strings.Join(messages, "; "))
	}

	errs := []error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}

	aggErr := aggregateErrors(errs)
	fmt.Printf("Aggregated error: %v\n", aggErr)

	// Pattern 2: Error retry with exponential backoff
	retryWithBackoff := func(operation func() error, maxRetries int) error {
		for i := 0; i < maxRetries; i++ {
			err := operation()
			if err == nil {
				return nil
			}

			// Check if error is retryable
			if !isRetryable(err) {
				return err
			}

			// Exponential backoff
			backoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
			time.Sleep(backoff)
		}
		return fmt.Errorf("operation failed after %d retries", maxRetries)
	}

	_ = retryWithBackoff(func() error {
		return errors.New("temporary failure")
	}, 3)
}

// isRetryable checks if error is retryable
func isRetryable(err error) bool {
	var tempErr *TemporaryError
	return errors.As(err, &tempErr)
}
