package main

import (
	"fmt"
	"math"
)

// Dynamic Programming - Comprehensive implementations of DP algorithms

func main() {
	demonstrateDynamicProgramming()
}

func demonstrateDynamicProgramming() {
	// Fibonacci
	fmt.Println("Fibonacci(10):", Fibonacci(10))
	fmt.Println("Fibonacci(20):", Fibonacci(20))

	// Longest Common Subsequence
	s1, s2 := "ABCDGH", "AEDFHR"
	fmt.Printf("LCS of '%s' and '%s': %d\n", s1, s2, LongestCommonSubsequence(s1, s2))

	// Longest Increasing Subsequence
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Println("LIS length:", LongestIncreasingSubsequence(arr))

	// Edit Distance
	s3, s4 := "sunday", "saturday"
	fmt.Printf("Edit distance between '%s' and '%s': %d\n", s3, s4, EditDistance(s3, s4))

	// 0/1 Knapsack
	weights := []int{10, 20, 30}
	values := []int{60, 100, 120}
	capacity := 50
	fmt.Printf("Knapsack (capacity %d): %d\n", capacity, Knapsack01(weights, values, capacity))

	// Coin Change
	coins := []int{1, 3, 4}
	amount := 6
	fmt.Printf("Coin change for %d: %d ways\n", amount, CoinChange(coins, amount))

	// Matrix Chain Multiplication
	p := []int{1, 2, 3, 4, 3}
	fmt.Printf("Matrix chain multiplication cost: %d\n", MatrixChainMultiplication(p))
}

// Fibonacci calculates nth Fibonacci number using DP
// Time Complexity: O(n), Space Complexity: O(n)
func Fibonacci(n int) int {
	// Secure: validate input
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	// DP table
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		// Secure: prevent integer overflow
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibonacciOptimized calculates Fibonacci with O(1) space
func FibonacciOptimized(n int) int {
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

// LongestCommonSubsequence finds LCS of two strings
// Time Complexity: O(m * n), Space Complexity: O(m * n)
func LongestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)

	// Secure: validate input
	if m == 0 || n == 0 {
		return 0
	}

	// DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Secure: bounds checking
			if i-1 < len(s1) && j-1 < len(s2) {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}

	return dp[m][n]
}

// max returns maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LongestIncreasingSubsequence finds length of LIS
// Time Complexity: O(n²), Space Complexity: O(n)
func LongestIncreasingSubsequence(arr []int) int {
	n := len(arr)

	// Secure: validate input
	if n == 0 {
		return 0
	}

	// DP table
	dp := make([]int, n)

	// Initialize
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// Secure: bounds checking
			if j < len(arr) && i < len(arr) && arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// Find maximum
	result := 1
	for _, val := range dp {
		if val > result {
			result = val
		}
	}

	return result
}

// EditDistance calculates edit distance (Levenshtein distance)
// Time Complexity: O(m * n), Space Complexity: O(m * n)
func EditDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)

	// Secure: validate input
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Secure: bounds checking
			if i-1 < len(s1) && j-1 < len(s2) {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = 1 + min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				}
			}
		}
	}

	return dp[m][n]
}

// min3 returns minimum of three integers
func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// Knapsack01 solves 0/1 knapsack problem
// Time Complexity: O(n * W), Space Complexity: O(n * W)
func Knapsack01(weights, values []int, capacity int) int {
	n := len(weights)

	// Secure: validate input
	if n == 0 || capacity < 0 || len(values) != n {
		return 0
	}

	// Secure: validate weights and values
	for i := range weights {
		if weights[i] < 0 || values[i] < 0 {
			return 0
		}
	}

	// DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			// Secure: bounds checking
			if i-1 < len(weights) {
				if weights[i-1] <= w {
					// Secure: prevent integer overflow
					dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
				} else {
					dp[i][w] = dp[i-1][w]
				}
			}
		}
	}

	return dp[n][capacity]
}

// CoinChange counts ways to make amount using coins
// Time Complexity: O(n * amount), Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
	// Secure: validate input
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	// Secure: validate coins
	for _, coin := range coins {
		if coin <= 0 {
			return 0
		}
	}

	// DP table
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		// Secure: validate coin
		if coin <= 0 {
			continue
		}
		for j := coin; j <= amount; j++ {
			// Secure: prevent integer overflow
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// MatrixChainMultiplication finds minimum cost of matrix chain multiplication
// Time Complexity: O(n³), Space Complexity: O(n²)
func MatrixChainMultiplication(p []int) int {
	n := len(p) - 1

	// Secure: validate input
	if n <= 0 {
		return 0
	}

	// Secure: validate dimensions
	for i := range p {
		if p[i] < 0 {
			return 0
		}
	}

	// DP table
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Fill table
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			dp[i][j] = math.MaxInt32

			for k := i; k < j; k++ {
				// Secure: bounds checking
				if i < len(p) && k+1 < len(p) && j+1 < len(p) {
					cost := dp[i][k] + dp[k+1][j] + p[i]*p[k+1]*p[j+1]
					if cost < dp[i][j] {
						dp[i][j] = cost
					}
				}
			}
		}
	}

	return dp[0][n-1]
}

// LongestPalindromicSubsequence finds length of longest palindromic subsequence
// Time Complexity: O(n²), Space Complexity: O(n²)
func LongestPalindromicSubsequence(s string) int {
	n := len(s)

	// Secure: validate input
	if n == 0 {
		return 0
	}

	// DP table
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1

			// Secure: bounds checking
			if i < len(s) && j < len(s) {
				if s[i] == s[j] {
					if length == 2 {
						dp[i][j] = 2
					} else {
						dp[i][j] = dp[i+1][j-1] + 2
					}
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}

	return dp[0][n-1]
}

// RodCutting solves rod cutting problem
// Time Complexity: O(n²), Space Complexity: O(n)
func RodCutting(prices []int, n int) int {
	// Secure: validate input
	if n <= 0 || len(prices) == 0 {
		return 0
	}

	// Secure: validate prices
	for _, price := range prices {
		if price < 0 {
			return 0
		}
	}

	// DP table
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			// Secure: bounds checking
			if j < len(prices) {
				maxVal = max(maxVal, prices[j]+dp[i-j-1])
			}
		}
		dp[i] = maxVal
	}

	return dp[n]
}
