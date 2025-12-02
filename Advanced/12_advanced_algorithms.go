package main

import (
	"fmt"
	"sort"
)

// Advanced Algorithms demonstrates advanced algorithm implementations

func main() {
	sortingAlgorithms()
	searchingAlgorithms()
	dynamicProgramming()
	greedyAlgorithms()
	graphAlgorithms()
}

// sortingAlgorithms demonstrates sorting algorithms
func sortingAlgorithms() {
	// Quick sort
	var quicksort func([]int) []int
	quicksort = func(arr []int) []int {
		if len(arr) <= 1 {
			return arr
		}

		pivot := arr[len(arr)/2]
		var left, middle, right []int

		for _, v := range arr {
			if v < pivot {
				left = append(left, v)
			} else if v == pivot {
				middle = append(middle, v)
			} else {
				right = append(right, v)
			}
		}

		left = quicksort(left)
		right = quicksort(right)

		return append(append(left, middle...), right...)
	}

	data := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := quicksort(data)
	fmt.Printf("Quick sorted: %v\n", sorted)

	// Merge function
	merge := func(left, right []int) []int {
		result := make([]int, 0, len(left)+len(right))
		i, j := 0, 0

		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}

		result = append(result, left[i:]...)
		result = append(result, right[j:]...)

		return result
	}

	// Merge sort
	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		if len(arr) <= 1 {
			return arr
		}

		mid := len(arr) / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])

		return merge(left, right)
	}

	data2 := []int{64, 34, 25, 12, 22, 11, 90}
	sorted2 := mergeSort(data2)
	fmt.Printf("Merge sorted: %v\n", sorted2)
}

// searchingAlgorithms demonstrates searching algorithms
func searchingAlgorithms() {
	// Binary search
	binarySearch := func(arr []int, target int) int {
		left, right := 0, len(arr)-1

		for left <= right {
			mid := left + (right-left)/2

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

	sorted := []int{1, 3, 5, 7, 9, 11, 13, 15}
	index := binarySearch(sorted, 7)
	fmt.Printf("Binary search for 7: index %d\n", index)

	// Linear search
	linearSearch := func(arr []int, target int) int {
		for i, v := range arr {
			if v == target {
				return i
			}
		}
		return -1
	}

	index2 := linearSearch(sorted, 9)
	fmt.Printf("Linear search for 9: index %d\n", index2)
}

// dynamicProgramming demonstrates dynamic programming
func dynamicProgramming() {
	// Fibonacci with memoization
	fibMemo := make(map[int]int)
	
	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}

		if val, ok := fibMemo[n]; ok {
			return val
		}

		fibMemo[n] = fibonacci(n-1) + fibonacci(n-2)
		return fibMemo[n]
	}

	fmt.Printf("Fibonacci(10): %d\n", fibonacci(10))

	// Longest Common Subsequence
	lcs := func(s1, s2 string) int {
		m, n := len(s1), len(s2)
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}

		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}

		return dp[m][n]
	}

	result := lcs("ABCDGH", "AEDFHR")
	fmt.Printf("LCS of 'ABCDGH' and 'AEDFHR': %d\n", result)
}

// max returns maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// greedyAlgorithms demonstrates greedy algorithms
func greedyAlgorithms() {
	// Activity selection problem
	type Activity struct {
		Start int
		End   int
	}

	activitySelection := func(activities []Activity) []Activity {
		// Sort by end time
		sort.Slice(activities, func(i, j int) bool {
			return activities[i].End < activities[j].End
		})

		selected := []Activity{activities[0]}
		lastEnd := activities[0].End

		for i := 1; i < len(activities); i++ {
			if activities[i].Start >= lastEnd {
				selected = append(selected, activities[i])
				lastEnd = activities[i].End
			}
		}

		return selected
	}

	activities := []Activity{
		{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
	}
	selected := activitySelection(activities)
	fmt.Printf("Selected activities: %v\n", selected)
}

// graphAlgorithms demonstrates graph algorithms
func graphAlgorithms() {
	// Dijkstra's algorithm (simplified)
	type Edge struct {
		To     int
		Weight int
	}

	dijkstra := func(graph map[int][]Edge, start int) map[int]int {
		dist := make(map[int]int)
		visited := make(map[int]bool)

		// Initialize distances
		for node := range graph {
			dist[node] = 1 << 31 // Large number
		}
		dist[start] = 0

		// Simplified implementation
		for len(visited) < len(graph) {
			// Find unvisited node with minimum distance
			minNode := -1
			minDist := 1 << 31

			for node, d := range dist {
				if !visited[node] && d < minDist {
					minDist = d
					minNode = node
				}
			}

			if minNode == -1 {
				break
			}

			visited[minNode] = true

			// Update distances to neighbors
			for _, edge := range graph[minNode] {
				if !visited[edge.To] {
					newDist := dist[minNode] + edge.Weight
					if newDist < dist[edge.To] {
						dist[edge.To] = newDist
					}
				}
			}
		}

		return dist
	}

	graph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}

	distances := dijkstra(graph, 0)
	fmt.Printf("Shortest distances from node 0: %v\n", distances)
}
