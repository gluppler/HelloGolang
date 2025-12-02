package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Error Handling demonstrates error handling patterns and best practices

func main() {
	basicErrors()
	customErrors()
	errorWrapping()
	errorChecking()
	errorPatterns()
	panicAndRecover()
}

// basicErrors demonstrates basic error handling
func basicErrors() {
	// Creating errors
	err1 := errors.New("something went wrong")
	fmt.Printf("Error: %v\n", err1)

	err2 := fmt.Errorf("operation failed: %s", "invalid input")
	fmt.Printf("Formatted error: %v\n", err2)

	// Error checking
	if err := divideFloat(10, 2); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Division successful")
	}

	if err := divideFloat(10, 0); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// divideFloat performs division with error handling (renamed to avoid conflict)
func divideFloat(a, b float64) error {
	if b == 0 {
		return errors.New("division by zero")
	}
	result := a / b
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	return nil
}

// customErrors demonstrates custom error types
func customErrors() {
	// Custom error type
	err := &ValidationError{
		Field:   "email",
		Message: "invalid email format",
	}
	fmt.Printf("Validation error: %v\n", err)

	// Using custom error
	if err := validateUser("", "invalid-email"); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// validateUser validates user input
func validateUser(username, email string) error {
	if username == "" {
		return &ValidationError{
			Field:   "username",
			Message: "username cannot be empty",
		}
	}
	if email == "" {
		return &ValidationError{
			Field:   "email",
			Message: "email cannot be empty",
		}
	}
	return nil
}

// errorWrapping demonstrates error wrapping (Go 1.13+)
func errorWrapping() {
	// Wrap errors
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("operation failed: %w", originalErr)

	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// Unwrap errors
	if unwrapped := errors.Unwrap(wrappedErr); unwrapped != nil {
		fmt.Printf("Unwrapped: %v\n", unwrapped)
	}

	// Check if error is specific type
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("Error is the original error")
	}

	// Check if error is of specific type
	var validationErr *ValidationError
	if errors.As(wrappedErr, &validationErr) {
		fmt.Printf("Error is ValidationError: %v\n", validationErr)
	}
}

// errorChecking demonstrates comprehensive error checking
func errorChecking() {
	// Secure: always check errors
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Multiple error checks
	if err := step1(); err != nil {
		fmt.Printf("Step 1 failed: %v\n", err)
		return
	}

	if err := step2(); err != nil {
		fmt.Printf("Step 2 failed: %v\n", err)
		return
	}

	fmt.Println("All steps completed successfully")
}

// step1 simulates a step that might fail
func step1() error {
	return nil
}

// step2 simulates a step that might fail
func step2() error {
	return nil
}

// errorPatterns demonstrates common error handling patterns
func errorPatterns() {
	// Pattern 1: Early return
	result, err := parseAndProcess("123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %d\n", result)

	// Pattern 2: Error aggregation
	var errs []error
	if err := validateField1(""); err != nil {
		errs = append(errs, err)
	}
	if err := validateField2(""); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		fmt.Printf("Multiple errors: %v\n", errs)
	}

	// Pattern 3: Retry with error
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		if err := retryableOperation(); err == nil {
			fmt.Println("Operation succeeded")
			break
		}
		if i == maxRetries-1 {
			fmt.Printf("Operation failed after %d retries\n", maxRetries)
		}
	}
}

// parseAndProcess parses and processes a string
func parseAndProcess(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s': %w", s, err)
	}

	if num < 0 {
		return 0, fmt.Errorf("number must be non-negative: %d", num)
	}

	return num * 2, nil
}

// validateField1 validates field 1
func validateField1(value string) error {
	if value == "" {
		return errors.New("field1 cannot be empty")
	}
	return nil
}

// validateField2 validates field 2
func validateField2(value string) error {
	if value == "" {
		return errors.New("field2 cannot be empty")
	}
	return nil
}

// retryableOperation simulates an operation that might fail
func retryableOperation() error {
	return errors.New("temporary failure")
}

// panicAndRecover demonstrates panic and recover
func panicAndRecover() {
	// Secure: use recover to handle panics gracefully
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	// This will panic but be recovered
	riskyFunction(true)

	fmt.Println("This line executes after recovery")
}

// riskyFunction may panic
func riskyFunction(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong")
	}
	fmt.Println("Function completed normally")
}

// safeDivision demonstrates safe division with panic recovery
func safeDivision(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}

	result = a / b
	return result, nil
}

// errorContext demonstrates adding context to errors
func errorContext() {
	// Add context at each level
	if err := processUser(""); err != nil {
		fmt.Printf("Failed to process user: %v\n", err)
	}
}

// processUser processes user data
func processUser(username string) error {
	if err := validateUsername(username); err != nil {
		return fmt.Errorf("processing user failed: %w", err)
	}
	return nil
}

// validateUsername validates username
func validateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}

// ErrorList demonstrates collecting multiple errors
type ErrorList []error

// Error implements error interface
func (el ErrorList) Error() string {
	if len(el) == 0 {
		return "no errors"
	}
	if len(el) == 1 {
		return el[0].Error()
	}
	return fmt.Sprintf("%d errors: %v", len(el), []error(el))
}

// Add adds an error to the list
func (el *ErrorList) Add(err error) {
	if err != nil {
		*el = append(*el, err)
	}
}

func errorListExample() {
	var errs ErrorList

	errs.Add(validateField1(""))
	errs.Add(validateField2(""))

	if len(errs) > 0 {
		fmt.Printf("Errors: %v\n", errs)
	}
}
