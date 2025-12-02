package main

import (
	"fmt"
)

// String Algorithms - Comprehensive implementations of string algorithms

func main() {
	demonstrateStringAlgorithms()
}

func demonstrateStringAlgorithms() {
	text := "ABABDABACDABABCABCABABC"
	pattern := "ABABCABAB"

	// KMP Algorithm
	matches := KMPSearch(text, pattern)
	fmt.Printf("KMP Search: pattern '%s' found at indices: %v\n", pattern, matches)

	// Rabin-Karp Algorithm
	matches2 := RabinKarpSearch(text, pattern)
	fmt.Printf("Rabin-Karp Search: pattern '%s' found at indices: %v\n", pattern, matches2)

	// Boyer-Moore Algorithm
	matches3 := BoyerMooreSearch(text, pattern)
	fmt.Printf("Boyer-Moore Search: pattern '%s' found at indices: %v\n", pattern, matches3)

	// Longest Common Substring
	s1, s2 := "ABCDGH", "ACDGHR"
	lcs := LongestCommonSubstring(s1, s2)
	fmt.Printf("Longest Common Substring of '%s' and '%s': %s\n", s1, s2, lcs)

	// Longest Palindromic Substring
	s := "forgeeksskeegfor"
	lps := LongestPalindromicSubstring(s)
	fmt.Printf("Longest Palindromic Substring of '%s': %s\n", s, lps)
}

// KMPSearch finds all occurrences of pattern in text using KMP algorithm
// Time Complexity: O(n + m), Space Complexity: O(m)
func KMPSearch(text, pattern string) []int {
	n, m := len(text), len(pattern)

	// Secure: validate input
	if m == 0 {
		return nil
	}
	if n < m {
		return nil
	}

	// Build failure function (LPS array)
	lps := buildLPS(pattern)

	result := []int{}
	i, j := 0, 0

	for i < n {
		// Secure: bounds checking
		if i < len(text) && j < len(pattern) {
			if text[i] == pattern[j] {
				i++
				j++
			}

			if j == m {
				// Pattern found
				result = append(result, i-j)
				j = lps[j-1]
			} else if i < n && text[i] != pattern[j] {
				if j != 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		} else {
			break
		}
	}

	return result
}

// buildLPS builds the longest proper prefix which is also suffix
func buildLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)

	length := 0
	i := 1

	for i < m {
		// Secure: bounds checking
		if i < len(pattern) && length < len(pattern) {
			if pattern[i] == pattern[length] {
				length++
				lps[i] = length
				i++
			} else {
				if length != 0 {
					length = lps[length-1]
				} else {
					lps[i] = 0
					i++
				}
			}
		} else {
			break
		}
	}

	return lps
}

// RabinKarpSearch finds all occurrences using Rabin-Karp algorithm
// Time Complexity: O(n + m) average, O(n * m) worst, Space Complexity: O(1)
func RabinKarpSearch(text, pattern string) []int {
	n, m := len(text), len(pattern)

	// Secure: validate input
	if m == 0 {
		return nil
	}
	if n < m {
		return nil
	}

	const base = 256
	const mod = 101

	// Calculate hash of pattern and first window of text
	patternHash := 0
	textHash := 0
	h := 1

	// Calculate h = base^(m-1) % mod
	for i := 0; i < m-1; i++ {
		h = (h * base) % mod
	}

	// Calculate initial hashes
	for i := 0; i < m; i++ {
		// Secure: bounds checking
		if i < len(pattern) {
			patternHash = (base*patternHash + int(pattern[i])) % mod
		}
		if i < len(text) {
			textHash = (base*textHash + int(text[i])) % mod
		}
	}

	result := []int{}

	// Slide pattern over text
	for i := 0; i <= n-m; i++ {
		// Secure: bounds checking
		if i < 0 || i+m > len(text) {
			break
		}

		// Check hash values
		if patternHash == textHash {
			// Check characters one by one
			match := true
			for j := 0; j < m; j++ {
				// Secure: bounds checking
				if i+j >= len(text) || j >= len(pattern) || text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				result = append(result, i)
			}
		}

		// Calculate hash for next window
		if i < n-m {
			// Secure: bounds checking
			if i < len(text) && i+m < len(text) {
				textHash = (base*(textHash-int(text[i])*h) + int(text[i+m])) % mod
				// Handle negative hash
				if textHash < 0 {
					textHash = textHash + mod
				}
			}
		}
	}

	return result
}

// BoyerMooreSearch finds all occurrences using Boyer-Moore algorithm
// Time Complexity: O(n * m) worst, O(n/m) best, Space Complexity: O(m)
func BoyerMooreSearch(text, pattern string) []int {
	n, m := len(text), len(pattern)

	// Secure: validate input
	if m == 0 {
		return nil
	}
	if n < m {
		return nil
	}

	// Build bad character table
	badChar := make(map[byte]int)
	for i := 0; i < m; i++ {
		// Secure: bounds checking
		if i < len(pattern) {
			badChar[pattern[i]] = i
		}
	}

	result := []int{}
	s := 0

	for s <= n-m {
		// Secure: bounds checking
		if s < 0 || s+m > len(text) {
			break
		}

		j := m - 1

		// Match pattern from right to left
		for j >= 0 && s+j < len(text) && j < len(pattern) && text[s+j] == pattern[j] {
			j--
		}

		if j < 0 {
			// Pattern found
			result = append(result, s)
			// Shift pattern
			if s+m < n {
				// Secure: bounds checking
				if s+m < len(text) {
					if lastOcc, ok := badChar[text[s+m]]; ok {
						s += m - lastOcc
					} else {
						s += m + 1
					}
				} else {
					s++
				}
			} else {
				s++
			}
		} else {
			// Shift pattern
			if s+j < len(text) {
				if lastOcc, ok := badChar[text[s+j]]; ok {
					shift := j - lastOcc
					if shift > 0 {
						s += shift
					} else {
						s++
					}
				} else {
					s += j + 1
				}
			} else {
				s++
			}
		}
	}

	return result
}

// LongestCommonSubstring finds longest common substring
// Time Complexity: O(m * n), Space Complexity: O(m * n)
func LongestCommonSubstring(s1, s2 string) string {
	m, n := len(s1), len(s2)

	// Secure: validate input
	if m == 0 || n == 0 {
		return ""
	}

	// DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0
	endIndex := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Secure: bounds checking
			if i-1 < len(s1) && j-1 < len(s2) {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
					if dp[i][j] > maxLength {
						maxLength = dp[i][j]
						endIndex = i - 1
					}
				} else {
					dp[i][j] = 0
				}
			}
		}
	}

	if maxLength == 0 {
		return ""
	}

	// Secure: bounds checking
	startIndex := endIndex - maxLength + 1
	if startIndex < 0 || startIndex >= len(s1) {
		return ""
	}

	return s1[startIndex : endIndex+1]
}

// LongestPalindromicSubstring finds longest palindromic substring
// Time Complexity: O(n²), Space Complexity: O(1)
func LongestPalindromicSubstring(s string) string {
	n := len(s)

	// Secure: validate input
	if n == 0 {
		return ""
	}

	start, maxLen := 0, 1

	// Expand around center
	expandAroundCenter := func(left, right int) int {
		// Secure: bounds checking
		for left >= 0 && right < n && left < len(s) && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	for i := 0; i < n; i++ {
		// Odd length palindromes
		len1 := expandAroundCenter(i, i)
		// Even length palindromes
		len2 := expandAroundCenter(i, i+1)

		length := max(len1, len2)
		if length > maxLen {
			maxLen = length
			start = i - (length-1)/2
		}
	}

	// Secure: bounds checking
	if start < 0 || start+maxLen > len(s) {
		return ""
	}

	return s[start : start+maxLen]
}

// max returns maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ZAlgorithm finds all occurrences using Z-algorithm
// Time Complexity: O(n + m), Space Complexity: O(n + m)
func ZAlgorithm(text, pattern string) []int {
	concat := pattern + "$" + text
	n := len(concat)

	// Secure: validate input
	if n == 0 {
		return nil
	}

	z := make([]int, n)
	l, r := 0, 0

	for i := 1; i < n; i++ {
		// Secure: bounds checking
		if i >= len(concat) {
			break
		}

		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}

		// Expand
		for i+z[i] < n && concat[z[i]] == concat[i+z[i]] {
			z[i]++
		}

		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	result := []int{}
	patternLen := len(pattern)

	for i := patternLen + 1; i < n; i++ {
		// Secure: bounds checking
		if i < len(z) && z[i] == patternLen {
			result = append(result, i-patternLen-1)
		}
	}

	return result
}

// min returns minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
