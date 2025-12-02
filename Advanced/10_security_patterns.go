package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

// Security Patterns demonstrates secure coding patterns

func main() {
	secureRandom()
	secureComparison()
	inputValidation()
	sqlInjectionPrevention()
	xssPrevention()
	secureStorage()
	rateLimitingSecurity()
}

// secureRandom demonstrates secure random number generation
func secureRandom() {
	// Secure: Use crypto/rand for cryptographic operations
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Printf("Error generating random bytes: %v\n", err)
		return
	}

	// Encode to base64 for storage/transmission
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	fmt.Printf("Secure random string: %s\n", randomString)

	// Generate secure token
	token := generateSecureToken(32)
	fmt.Printf("Secure token: %s\n", token)
}

// generateSecureToken generates a secure random token
func generateSecureToken(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

// secureComparison demonstrates constant-time comparison
func secureComparison() {
	// Secure: Use constant-time comparison for secrets
	secret1 := []byte("my-secret-key")
	secret2 := []byte("my-secret-key")
	secret3 := []byte("wrong-secret")

	// Constant-time comparison (prevents timing attacks)
	equal1 := subtle.ConstantTimeCompare(secret1, secret2) == 1
	equal2 := subtle.ConstantTimeCompare(secret1, secret3) == 1

	fmt.Printf("Secret1 == Secret2: %t\n", equal1)
	fmt.Printf("Secret1 == Secret3: %t\n", equal2)

	// Insecure: Regular comparison (vulnerable to timing attacks)
	// if secret1 == secret2 { ... } // DON'T USE FOR SECRETS
}

// inputValidation demonstrates input validation
func inputValidation() {
	// Secure: Always validate input
	validateInput := func(input string) error {
		// Check length
		if len(input) == 0 {
			return fmt.Errorf("input cannot be empty")
		}
		if len(input) > 100 {
			return fmt.Errorf("input too long")
		}

		// Check for dangerous characters
		for _, char := range input {
			if char < 32 || char > 126 {
				return fmt.Errorf("input contains invalid characters")
			}
		}

		return nil
	}

	// Test validation
	inputs := []string{
		"valid input",
		"",
		"a" + string(make([]byte, 200)), // Too long
		"valid\x00null",                 // Contains null byte
	}

	for _, input := range inputs {
		if err := validateInput(input); err != nil {
			fmt.Printf("Validation failed for '%s': %v\n", input, err)
		} else {
			fmt.Printf("Validation passed for '%s'\n", input)
		}
	}
}

// sqlInjectionPrevention demonstrates SQL injection prevention
func sqlInjectionPrevention() {
	// Secure: Use parameterized queries
	// Example with database/sql:
	// db.Query("SELECT * FROM users WHERE id = ?", userID)
	// db.Query("SELECT * FROM users WHERE name = ? AND email = ?", name, email)

	// Insecure: String concatenation (VULNERABLE)
	// query := "SELECT * FROM users WHERE name = '" + userName + "'"
	// This is vulnerable to SQL injection

	fmt.Println("SQL Injection Prevention:")
	fmt.Println("  - Always use parameterized queries")
	fmt.Println("  - Never concatenate user input into SQL")
	fmt.Println("  - Validate and sanitize input")
	fmt.Println("  - Use prepared statements")
}

// xssPrevention demonstrates XSS prevention
func xssPrevention() {
	// Secure: Escape output
	escapeHTML := func(s string) string {
		// In production, use html/template or html.EscapeString
		escaped := ""
		for _, char := range s {
			switch char {
			case '<':
				escaped += "&lt;"
			case '>':
				escaped += "&gt;"
			case '&':
				escaped += "&amp;"
			case '"':
				escaped += "&quot;"
			case '\'':
				escaped += "&#39;"
			default:
				escaped += string(char)
			}
		}
		return escaped
	}

	userInput := "<script>alert('XSS')</script>"
	escaped := escapeHTML(userInput)
	fmt.Printf("Original: %s\n", userInput)
	fmt.Printf("Escaped: %s\n", escaped)

	fmt.Println("XSS Prevention:")
	fmt.Println("  - Escape all user input before output")
	fmt.Println("  - Use html/template for HTML output")
	fmt.Println("  - Validate and sanitize input")
	fmt.Println("  - Use Content Security Policy (CSP)")
}

// secureStorage demonstrates secure storage patterns
func secureStorage() {
	// Secure: Never store plaintext passwords
	// Use bcrypt, argon2, or similar

	type User struct {
		ID           int
		Username     string
		PasswordHash string // Never store plaintext
		Salt         string
	}

	// Example password hashing (use crypto/bcrypt in production)
	hashPassword := func(password string) (hash string, salt string, err error) {
		// Generate salt
		saltBytes := make([]byte, 16)
		_, err = rand.Read(saltBytes)
		if err != nil {
			return "", "", err
		}
		salt = base64.URLEncoding.EncodeToString(saltBytes)

		// In production, use bcrypt.CompareHashAndPassword
		// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		hash = "hashed_password_here" // Placeholder
		return hash, salt, nil
	}

	password := "user-password"
	hash, salt, err := hashPassword(password)
	if err == nil {
		fmt.Printf("Password hashed: %s\n", hash)
		fmt.Printf("Salt: %s\n", salt)
	}

	fmt.Println("Secure Storage:")
	fmt.Println("  - Hash passwords with bcrypt or argon2")
	fmt.Println("  - Use unique salts per password")
	fmt.Println("  - Never store sensitive data in plaintext")
	fmt.Println("  - Encrypt sensitive data at rest")
}

// rateLimitingSecurity demonstrates rate limiting for security
func rateLimitingSecurity() {
	// Rate limiting to prevent brute force attacks
	type RateLimiter struct {
		attempts    map[string][]time.Time
		maxAttempts int
		window      time.Duration
		mu          sync.Mutex
	}

	limiter := &RateLimiter{
		attempts:    make(map[string][]time.Time),
		maxAttempts: 5,
		window:      15 * time.Minute,
	}

	checkRateLimit := func(rl *RateLimiter, key string) bool {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()
		attempts := rl.attempts[key]

		// Remove old attempts outside window
		validAttempts := []time.Time{}
		for _, attempt := range attempts {
			if now.Sub(attempt) < rl.window {
				validAttempts = append(validAttempts, attempt)
			}
		}

		if len(validAttempts) >= rl.maxAttempts {
			return false // Rate limited
		}

		// Record new attempt
		validAttempts = append(validAttempts, now)
		rl.attempts[key] = validAttempts

		return true // Allowed
	}

	// Test rate limiting
	key := "user@example.com"
	for i := 0; i < 7; i++ {
		allowed := checkRateLimit(limiter, key)
		if allowed {
			fmt.Printf("Attempt %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Attempt %d: Rate limited\n", i+1)
		}
	}
}
