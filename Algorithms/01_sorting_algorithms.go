package main

import (
	"fmt"
)

// Sorting Algorithms - Comprehensive implementations of all major sorting algorithms

func main() {
	demonstrateSortingAlgorithms()
}

func demonstrateSortingAlgorithms() {
	data := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 88}

	fmt.Println("Original array:", data)

	// Bubble Sort
	bubbleData := make([]int, len(data))
	copy(bubbleData, data)
	BubbleSort(bubbleData)
	fmt.Println("Bubble Sort:", bubbleData)

	// Selection Sort
	selectionData := make([]int, len(data))
	copy(selectionData, data)
	SelectionSort(selectionData)
	fmt.Println("Selection Sort:", selectionData)

	// Insertion Sort
	insertionData := make([]int, len(data))
	copy(insertionData, data)
	InsertionSort(insertionData)
	fmt.Println("Insertion Sort:", insertionData)

	// Merge Sort
	mergeData := make([]int, len(data))
	copy(mergeData, data)
	MergeSort(mergeData)
	fmt.Println("Merge Sort:", mergeData)

	// Quick Sort
	quickData := make([]int, len(data))
	copy(quickData, data)
	QuickSort(quickData)
	fmt.Println("Quick Sort:", quickData)

	// Heap Sort
	heapData := make([]int, len(data))
	copy(heapData, data)
	HeapSort(heapData)
	fmt.Println("Heap Sort:", heapData)

	// Counting Sort
	countingData := []int{4, 2, 2, 8, 3, 3, 1}
	CountingSort(countingData, 10)
	fmt.Println("Counting Sort:", countingData)

	// Radix Sort
	radixData := []int{170, 45, 75, 90, 802, 24, 2, 66}
	RadixSort(radixData)
	fmt.Println("Radix Sort:", radixData)

	// Bucket Sort
	bucketData := []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	BucketSort(bucketData)
	fmt.Println("Bucket Sort:", bucketData)

	// Shell Sort
	shellData := make([]int, len(data))
	copy(shellData, data)
	ShellSort(shellData)
	fmt.Println("Shell Sort:", shellData)
}

// BubbleSort sorts array using bubble sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Optimized: break if no swaps occurred
		if !swapped {
			break
		}
	}
}

// SelectionSort sorts array using selection sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func SelectionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// InsertionSort sorts array using insertion sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func InsertionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Secure: bounds checking
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// MergeSort sorts array using merge sort algorithm
// Time Complexity: O(n log n), Space Complexity: O(n)
func MergeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	mergeSortHelper(arr, 0, n-1)
}

// mergeSortHelper is the recursive helper for merge sort
func mergeSortHelper(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		mergeSortHelper(arr, left, mid)
		mergeSortHelper(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

// merge merges two sorted subarrays
func merge(arr []int, left, mid, right int) {
	// Secure: bounds checking
	if left < 0 || mid < left || right < mid || right >= len(arr) {
		return
	}

	n1 := mid - left + 1
	n2 := right - mid

	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	copy(leftArr, arr[left:left+n1])
	copy(rightArr, arr[mid+1:mid+1+n2])

	i, j, k := 0, 0, left

	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// QuickSort sorts array using quick sort algorithm
// Time Complexity: O(n log n) average, O(n²) worst, Space Complexity: O(log n)
func QuickSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	quickSortHelper(arr, 0, n-1)
}

// quickSortHelper is the recursive helper for quick sort
func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

// partition partitions the array and returns pivot index
func partition(arr []int, low, high int) int {
	// Secure: bounds checking
	if low < 0 || high >= len(arr) || low > high {
		return low
	}

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// HeapSort sorts array using heap sort algorithm
// Time Complexity: O(n log n), Space Complexity: O(1)
func HeapSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

// heapify maintains heap property
func heapify(arr []int, n, i int) {
	// Secure: bounds checking
	if i < 0 || i >= n || n > len(arr) {
		return
	}

	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// CountingSort sorts array using counting sort algorithm
// Time Complexity: O(n + k), Space Complexity: O(k)
// k is the range of input
func CountingSort(arr []int, maxVal int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Secure: validate maxVal
	if maxVal < 0 {
		return
	}

	// Count array
	count := make([]int, maxVal+1)

	// Count occurrences
	for i := 0; i < n; i++ {
		// Secure: bounds checking
		if arr[i] < 0 || arr[i] > maxVal {
			continue
		}
		count[arr[i]]++
	}

	// Modify count array to store positions
	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	// Output array
	output := make([]int, n)

	// Build output array
	for i := n - 1; i >= 0; i-- {
		// Secure: bounds checking
		if arr[i] < 0 || arr[i] > maxVal {
			continue
		}
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	// Copy back to original array
	copy(arr, output)
}

// RadixSort sorts array using radix sort algorithm
// Time Complexity: O(d * (n + k)), Space Complexity: O(n + k)
// d is number of digits, k is the base
func RadixSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find maximum number to know number of digits
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// Secure: handle negative numbers
	if max < 0 {
		return
	}

	// Do counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp)
	}
}

// countingSortByDigit performs counting sort by digit
func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// Count occurrences
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		if index >= 0 && index < 10 {
			count[index]++
		}
	}

	// Change count to position
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build output array
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		if index >= 0 && index < 10 {
			output[count[index]-1] = arr[i]
			count[index]--
		}
	}

	// Copy output to original array
	copy(arr, output)
}

// BucketSort sorts array using bucket sort algorithm
// Time Complexity: O(n + k) average, Space Complexity: O(n)
func BucketSort(arr []float64) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Create buckets
	buckets := make([][]float64, n)

	// Secure: validate input
	for i := 0; i < n; i++ {
		if arr[i] < 0 || arr[i] > 1 {
			// Bucket sort typically works for numbers in range [0, 1)
			// For other ranges, normalize first
			continue
		}
		index := int(float64(n) * arr[i])
		if index >= n {
			index = n - 1
		}
		buckets[index] = append(buckets[index], arr[i])
	}

	// Sort individual buckets
	for i := 0; i < n; i++ {
		InsertionSortFloat(buckets[i])
	}

	// Concatenate buckets
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
}

// InsertionSortFloat sorts float array using insertion sort
func InsertionSortFloat(arr []float64) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// ShellSort sorts array using shell sort algorithm
// Time Complexity: O(n²) worst case, Space Complexity: O(1)
func ShellSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Start with large gap, then reduce
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do gapped insertion sort
		for i := gap; i < n; i++ {
			temp := arr[i]
			var j int

			// Secure: bounds checking
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = temp
		}
	}
}
