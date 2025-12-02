package main

import (
	"fmt"
	"math"
)

// Searching Algorithms - Comprehensive implementations of all major searching algorithms

func main() {
	demonstrateSearchingAlgorithms()
}

func demonstrateSearchingAlgorithms() {
	sortedArr := []int{2, 5, 8, 12, 16, 23, 38, 45, 67, 78, 89, 95}
	unsortedArr := []int{64, 34, 25, 12, 22, 11, 90}

	target := 23

	fmt.Println("Sorted array:", sortedArr)
	fmt.Println("Target:", target)

	// Linear Search
	index := LinearSearch(unsortedArr, target)
	fmt.Printf("Linear Search (unsorted): index %d\n", index)

	// Binary Search
	index = BinarySearch(sortedArr, target)
	fmt.Printf("Binary Search: index %d\n", index)

	// Interpolation Search
	index = InterpolationSearch(sortedArr, target)
	fmt.Printf("Interpolation Search: index %d\n", index)

	// Exponential Search
	index = ExponentialSearch(sortedArr, target)
	fmt.Printf("Exponential Search: index %d\n", index)

	// Jump Search
	index = JumpSearch(sortedArr, target)
	fmt.Printf("Jump Search: index %d\n", index)

	// Ternary Search
	index = TernarySearch(sortedArr, target)
	fmt.Printf("Ternary Search: index %d\n", index)
}

// LinearSearch searches for target in array using linear search
// Time Complexity: O(n), Space Complexity: O(1)
func LinearSearch(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// BinarySearch searches for target in sorted array using binary search
// Time Complexity: O(log n), Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		// Secure: prevent integer overflow
		mid := left + (right-left)/2

		// Secure: bounds checking
		if mid < 0 || mid >= len(arr) {
			break
		}

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive searches recursively
func BinarySearchRecursive(arr []int, target, left, right int) int {
	// Secure: bounds checking
	if left < 0 || right >= len(arr) || left > right {
		return -1
	}

	mid := left + (right-left)/2

	// Secure: bounds checking
	if mid < 0 || mid >= len(arr) {
		return -1
	}

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}

	return BinarySearchRecursive(arr, target, mid+1, right)
}

// InterpolationSearch searches using interpolation search
// Time Complexity: O(log log n) average, O(n) worst, Space Complexity: O(1)
func InterpolationSearch(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right && target >= arr[left] && target <= arr[right] {
		// Secure: prevent division by zero
		if arr[right] == arr[left] {
			if arr[left] == target {
				return left
			}
			return -1
		}

		// Secure: prevent integer overflow
		pos := left + ((right-left)*(target-arr[left]))/(arr[right]-arr[left])

		// Secure: bounds checking
		if pos < 0 || pos >= len(arr) {
			break
		}

		if arr[pos] == target {
			return pos
		}

		if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}

// ExponentialSearch searches using exponential search
// Time Complexity: O(log n), Space Complexity: O(1)
func ExponentialSearch(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	// If target is at first position
	if arr[0] == target {
		return 0
	}

	// Find range for binary search
	i := 1
	for i < len(arr) && arr[i] <= target {
		i *= 2
	}

	// Secure: bounds checking
	right := i
	if right >= len(arr) {
		right = len(arr) - 1
	}

	// Binary search in the found range
	return BinarySearchRecursive(arr, target, i/2, right)
}

// JumpSearch searches using jump search
// Time Complexity: O(√n), Space Complexity: O(1)
func JumpSearch(arr []int, target int) int {
	// Secure: bounds checking
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Finding block size to be jumped
	step := int(math.Sqrt(float64(n)))

	// Finding the block where element is present
	prev := 0
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	// Doing a linear search for target in block
	for arr[prev] < target {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	// Secure: bounds checking
	if prev < n && arr[prev] == target {
		return prev
	}

	return -1
}

// min returns minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TernarySearch searches using ternary search
// Time Complexity: O(log₃ n), Space Complexity: O(1)
func TernarySearch(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	return ternarySearchHelper(arr, target, 0, len(arr)-1)
}

// ternarySearchHelper is recursive helper for ternary search
func ternarySearchHelper(arr []int, target, left, right int) int {
	// Secure: bounds checking
	if left < 0 || right >= len(arr) || left > right {
		return -1
	}

	// Divide array into three parts
	mid1 := left + (right-left)/3
	mid2 := right - (right-left)/3

	// Secure: bounds checking
	if mid1 < 0 || mid1 >= len(arr) || mid2 < 0 || mid2 >= len(arr) {
		return -1
	}

	if arr[mid1] == target {
		return mid1
	}

	if arr[mid2] == target {
		return mid2
	}

	if target < arr[mid1] {
		return ternarySearchHelper(arr, target, left, mid1-1)
	} else if target > arr[mid2] {
		return ternarySearchHelper(arr, target, mid2+1, right)
	} else {
		return ternarySearchHelper(arr, target, mid1+1, mid2-1)
	}
}

// FindFirst finds first occurrence of target
func FindFirst(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		// Secure: bounds checking
		if mid < 0 || mid >= len(arr) {
			break
		}

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// FindLast finds last occurrence of target
func FindLast(arr []int, target int) int {
	// Secure: bounds checking
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		// Secure: bounds checking
		if mid < 0 || mid >= len(arr) {
			break
		}

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// CountOccurrences counts occurrences of target in sorted array
func CountOccurrences(arr []int, target int) int {
	first := FindFirst(arr, target)
	if first == -1 {
		return 0
	}

	last := FindLast(arr, target)
	return last - first + 1
}
