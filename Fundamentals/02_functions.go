package main

import (
	"fmt"
	"math"
)

// Functions demonstrates function declarations, parameters, and return values

func main() {
	basicFunction()
	multipleReturns()
	namedReturns()
	variadicFunction()
	firstClassFunctions()
	closures()
	recursion()
}

// basicFunction demonstrates a simple function
func basicFunction() {
	result := add(10, 20)
	fmt.Printf("add(10, 20) = %d\n", result)
}

// add performs addition of two integers
func add(a, b int) int {
	return a + b
}

// multipleReturns demonstrates functions with multiple return values
func multipleReturns() {
	quotient, remainder := divide(17, 5)
	fmt.Printf("divide(17, 5) = quotient: %d, remainder: %d\n", quotient, remainder)

	// Ignoring one return value
	_, rem := divide(20, 3)
	fmt.Printf("divide(20, 3) remainder: %d\n", rem)
}

// divide returns quotient and remainder
func divide(a, b int) (int, int) {
	if b == 0 {
		// Secure: prevent division by zero
		return 0, 0
	}
	return a / b, a % b
}

// namedReturns demonstrates named return values
func namedReturns() {
	sum, product := calculate(5, 7)
	fmt.Printf("calculate(5, 7) = sum: %d, product: %d\n", sum, product)
}

// calculate uses named return values
func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// variadicFunction demonstrates variadic functions
func variadicFunction() {
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("sumAll(1,2,3,4,5) = %d\n", sum)

	numbers := []int{10, 20, 30}
	sum2 := sumAll(numbers...)
	fmt.Printf("sumAll([]int{10,20,30}...) = %d\n", sum2)
}

// sumAll accepts variable number of arguments
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// firstClassFunctions demonstrates functions as first-class citizens
func firstClassFunctions() {
	// Function as a variable
	var operation func(int, int) int
	operation = multiply

	result := operation(5, 6)
	fmt.Printf("operation(5, 6) = %d\n", result)

	// Function as parameter
	result2 := applyOperation(10, 4, multiply)
	fmt.Printf("applyOperation(10, 4, multiply) = %d\n", result2)

	result3 := applyOperation(10, 4, subtract)
	fmt.Printf("applyOperation(10, 4, subtract) = %d\n", result3)
}

// multiply multiplies two integers
func multiply(a, b int) int {
	return a * b
}

// subtract subtracts b from a
func subtract(a, b int) int {
	return a - b
}

// applyOperation applies a function to two integers
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// closures demonstrates closure functions
func closures() {
	// Counter closure
	counter := makeCounter()
	fmt.Printf("Counter: %d, %d, %d\n", counter(), counter(), counter())

	// Multiplier closure
	multiplyBy := makeMultiplier(5)
	fmt.Printf("multiplyBy(5)(3) = %d\n", multiplyBy(3))
	fmt.Printf("multiplyBy(5)(7) = %d\n", multiplyBy(7))
}

// makeCounter returns a closure that increments a counter
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeMultiplier returns a closure that multiplies by a factor
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// recursion demonstrates recursive functions
func recursion() {
	factorial := calculateFactorial(5)
	fmt.Printf("factorial(5) = %d\n", factorial)

	fib := fibonacci(10)
	fmt.Printf("fibonacci(10) = %d\n", fib)
}

// calculateFactorial calculates factorial recursively
func calculateFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

// fibonacci calculates nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// advancedFunction demonstrates advanced function features
func advancedFunction() {
	// Function that returns a function
	powerOf := func(exp float64) func(float64) float64 {
		return func(base float64) float64 {
			return math.Pow(base, exp)
		}
	}

	square := powerOf(2)
	cube := powerOf(3)

	fmt.Printf("square(5) = %.2f\n", square(5))
	fmt.Printf("cube(3) = %.2f\n", cube(3))
}
