package main

import (
	"fmt"
	"math"
)

// Mathematical Algorithms - Comprehensive implementations of mathematical algorithms

func main() {
	demonstrateMathematicalAlgorithms()
}

func demonstrateMathematicalAlgorithms() {
	// GCD and LCM
	fmt.Printf("GCD(48, 18): %d\n", GCD(48, 18))
	fmt.Printf("LCM(12, 18): %d\n", LCM(12, 18))

	// Prime numbers
	fmt.Printf("Is 17 prime: %t\n", IsPrime(17))
	fmt.Printf("Primes up to 30: %v\n", SieveOfEratosthenes(30))

	// Factorial
	fmt.Printf("Factorial(5): %d\n", Factorial(5))

	// Power
	fmt.Printf("Power(2, 10): %d\n", Power(2, 10))

	// Fibonacci
	fmt.Printf("Fibonacci(10): %d\n", Fibonacci(10))

	// Permutations and Combinations
	fmt.Printf("Permutations(5, 3): %d\n", Permutations(5, 3))
	fmt.Printf("Combinations(5, 3): %d\n", Combinations(5, 3))
}

// GCD calculates Greatest Common Divisor using Euclidean algorithm
// Time Complexity: O(log(min(a, b))), Space Complexity: O(1)
func GCD(a, b int) int {
	// Secure: handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Secure: handle zero
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// GCDRecursive calculates GCD recursively
func GCDRecursive(a, b int) int {
	// Secure: handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Secure: handle zero
	if b == 0 {
		return a
	}

	return GCDRecursive(b, a%b)
}

// LCM calculates Least Common Multiple
// Time Complexity: O(log(min(a, b))), Space Complexity: O(1)
func LCM(a, b int) int {
	// Secure: handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Secure: handle zero
	if a == 0 || b == 0 {
		return 0
	}

	// Secure: prevent integer overflow
	gcd := GCD(a, b)
	if gcd == 0 {
		return 0
	}

	return (a / gcd) * b
}

// IsPrime checks if number is prime
// Time Complexity: O(√n), Space Complexity: O(1)
func IsPrime(n int) bool {
	// Secure: validate input
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Check up to √n
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// SieveOfEratosthenes finds all primes up to n
// Time Complexity: O(n log log n), Space Complexity: O(n)
func SieveOfEratosthenes(n int) []int {
	// Secure: validate input
	if n < 2 {
		return nil
	}

	// Create boolean array
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Sieve
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Collect primes
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// Factorial calculates factorial
// Time Complexity: O(n), Space Complexity: O(1)
func Factorial(n int) int {
	// Secure: validate input
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		// Secure: prevent integer overflow (simplified check)
		if result > math.MaxInt32/i {
			return 0 // Overflow
		}
		result *= i
	}

	return result
}

// FactorialRecursive calculates factorial recursively
func FactorialRecursive(n int) int {
	// Secure: validate input
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}

	return n * FactorialRecursive(n-1)
}

// Power calculates base^exponent
// Time Complexity: O(log n), Space Complexity: O(1)
func Power(base, exponent int) int {
	// Secure: validate input
	if exponent < 0 {
		return 0 // Handle negative exponent if needed
	}
	if exponent == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	if base == 1 {
		return 1
	}

	result := int64(1)
	base64 := int64(base)
	exp := int64(exponent)
	
	for exp > 0 {
		if exp%2 == 1 {
			// Secure: prevent integer overflow
			if result > math.MaxInt64/base64 {
				return 0 // Overflow
			}
			result *= base64
		}
		exp /= 2
		// Secure: prevent integer overflow
		if base64 > math.MaxInt64/base64 {
			return 0 // Overflow
		}
		base64 *= base64
	}
	
	return int(result)
}

// Fibonacci calculates nth Fibonacci number
func Fibonacci(n int) int {
	// Secure: validate input
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Permutations calculates nPr = n! / (n-r)!
// Time Complexity: O(n), Space Complexity: O(1)
func Permutations(n, r int) int {
	// Secure: validate input
	if n < 0 || r < 0 || r > n {
		return 0
	}

	result := 1
	for i := 0; i < r; i++ {
		// Secure: prevent integer overflow
		if result > math.MaxInt32/(n-i) {
			return 0 // Overflow
		}
		result *= (n - i)
	}

	return result
}

// Combinations calculates nCr = n! / (r! * (n-r)!)
// Time Complexity: O(min(r, n-r)), Space Complexity: O(1)
func Combinations(n, r int) int {
	// Secure: validate input
	if n < 0 || r < 0 || r > n {
		return 0
	}

	// Optimize: C(n, r) = C(n, n-r)
	if r > n-r {
		r = n - r
	}

	result := 1
	for i := 0; i < r; i++ {
		// Secure: prevent integer overflow
		if result > math.MaxInt32/(n-i) {
			return 0 // Overflow
		}
		result = result * (n - i) / (i + 1)
	}

	return result
}

// CatalanNumber calculates nth Catalan number
// Time Complexity: O(n), Space Complexity: O(1)
func CatalanNumber(n int) int {
	// Secure: validate input
	if n < 0 {
		return 0
	}

	// C(n) = (2n)! / ((n+1)! * n!)
	return Combinations(2*n, n) / (n + 1)
}

// ModularExponentiation calculates (base^exponent) % mod
// Time Complexity: O(log exponent), Space Complexity: O(1)
func ModularExponentiation(base, exponent, mod int) int {
	// Secure: validate input
	if mod <= 0 {
		return 0
	}
	if exponent < 0 {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := 1
	base = base % mod

	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		exponent /= 2
		base = (base * base) % mod
	}

	return result
}

// ExtendedGCD calculates GCD and coefficients using Extended Euclidean algorithm
// Returns: gcd, x, y such that ax + by = gcd(a, b)
func ExtendedGCD(a, b int) (int, int, int) {
	// Secure: handle zero
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := ExtendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1

	return gcd, x, y
}

// ModularInverse calculates modular multiplicative inverse
// Returns x such that (a * x) % m = 1
func ModularInverse(a, m int) int {
	// Secure: validate input
	if m <= 0 {
		return 0
	}

	gcd, x, _ := ExtendedGCD(a, m)
	if gcd != 1 {
		return 0 // Inverse doesn't exist
	}

	// Make x positive
	return (x%m + m) % m
}
