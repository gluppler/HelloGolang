package main

import (
	"fmt"
	"sort"
)

// Greedy Algorithms - Comprehensive implementations of greedy algorithms

func main() {
	demonstrateGreedyAlgorithms()
}

func demonstrateGreedyAlgorithms() {
	// Activity Selection
	activities := []Activity{
		{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
	}
	selected := ActivitySelection(activities)
	fmt.Println("Activity Selection:", selected)

	// Fractional Knapsack
	items := []Item{
		{60, 10}, {100, 20}, {120, 30},
	}
	capacity := 50
	profit := FractionalKnapsack(items, capacity)
	fmt.Printf("Fractional Knapsack (capacity %d): profit %.2f\n", capacity, profit)

	// Job Sequencing
	jobs := []Job{
		{'a', 2, 100}, {'b', 1, 19}, {'c', 2, 27}, {'d', 1, 25}, {'e', 3, 15},
	}
	sequence := JobSequencing(jobs)
	fmt.Println("Job Sequencing:", sequence)

	// Minimum Coin Change
	coins := []int{1, 5, 10, 25}
	amount := 67
	change := MinimumCoinChange(coins, amount)
	fmt.Printf("Minimum coins for %d: %v\n", amount, change)
}

// Activity represents an activity with start and finish time
type Activity struct {
	Start  int
	Finish int
}

// ActivitySelection selects maximum number of non-overlapping activities
// Time Complexity: O(n log n), Space Complexity: O(1)
func ActivitySelection(activities []Activity) []Activity {
	n := len(activities)

	// Secure: validate input
	if n == 0 {
		return nil
	}

	// Secure: validate activities
	for _, act := range activities {
		if act.Start < 0 || act.Finish < 0 || act.Start > act.Finish {
			return nil
		}
	}

	// Sort by finish time
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})

	selected := []Activity{activities[0]}
	lastFinish := activities[0].Finish

	for i := 1; i < n; i++ {
		// Secure: bounds checking
		if i < len(activities) {
			if activities[i].Start >= lastFinish {
				selected = append(selected, activities[i])
				lastFinish = activities[i].Finish
			}
		}
	}

	return selected
}

// Item represents an item with value and weight
type Item struct {
	Value  int
	Weight int
}

// FractionalKnapsack solves fractional knapsack problem
// Time Complexity: O(n log n), Space Complexity: O(1)
func FractionalKnapsack(items []Item, capacity int) float64 {
	n := len(items)

	// Secure: validate input
	if n == 0 || capacity <= 0 {
		return 0
	}

	// Secure: validate items
	for _, item := range items {
		if item.Value < 0 || item.Weight <= 0 {
			return 0
		}
	}

	// Sort by value/weight ratio (descending)
	sort.Slice(items, func(i, j int) bool {
		ratioI := float64(items[i].Value) / float64(items[i].Weight)
		ratioJ := float64(items[j].Value) / float64(items[j].Weight)
		return ratioI > ratioJ
	})

	totalValue := 0.0
	remainingCapacity := capacity

	for i := 0; i < n && remainingCapacity > 0; i++ {
		// Secure: bounds checking
		if i < len(items) {
			if items[i].Weight <= remainingCapacity {
				totalValue += float64(items[i].Value)
				remainingCapacity -= items[i].Weight
			} else {
				// Take fraction
				fraction := float64(remainingCapacity) / float64(items[i].Weight)
				totalValue += fraction * float64(items[i].Value)
				remainingCapacity = 0
			}
		}
	}

	return totalValue
}

// Job represents a job with id, deadline, and profit
type Job struct {
	ID       rune
	Deadline int
	Profit   int
}

// JobSequencing solves job sequencing problem
// Time Complexity: O(n²), Space Complexity: O(n)
func JobSequencing(jobs []Job) []rune {
	n := len(jobs)

	// Secure: validate input
	if n == 0 {
		return nil
	}

	// Secure: validate jobs
	for _, job := range jobs {
		if job.Deadline < 0 || job.Profit < 0 {
			return nil
		}
	}

	// Sort by profit (descending)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Profit > jobs[j].Profit
	})

	// Find maximum deadline
	maxDeadline := 0
	for _, job := range jobs {
		if job.Deadline > maxDeadline {
			maxDeadline = job.Deadline
		}
	}

	// Secure: validate maxDeadline
	if maxDeadline <= 0 {
		return nil
	}

	// Time slots
	timeSlot := make([]bool, maxDeadline)
	result := []rune{}

	for i := 0; i < n; i++ {
		// Secure: bounds checking
		if i < len(jobs) {
			// Find latest available slot
			for j := min(jobs[i].Deadline-1, maxDeadline-1); j >= 0; j-- {
				// Secure: bounds checking
				if j < len(timeSlot) && !timeSlot[j] {
					timeSlot[j] = true
					result = append(result, jobs[i].ID)
					break
				}
			}
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

// MinimumCoinChange finds minimum coins to make amount
// Time Complexity: O(n), Space Complexity: O(1)
func MinimumCoinChange(coins []int, amount int) []int {
	// Secure: validate input
	if amount < 0 || len(coins) == 0 {
		return nil
	}

	// Secure: validate coins
	for _, coin := range coins {
		if coin <= 0 {
			return nil
		}
	}

	// Sort coins in descending order
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	result := []int{}
	remaining := amount

	for i := 0; i < len(coins) && remaining > 0; i++ {
		// Secure: bounds checking
		if i < len(coins) && coins[i] > 0 {
			count := remaining / coins[i]
			if count > 0 {
				for j := 0; j < count; j++ {
					result = append(result, coins[i])
				}
				remaining %= coins[i]
			}
		}
	}

	// Secure: check if solution exists
	if remaining != 0 {
		return nil // No solution
	}

	return result
}

// HuffmanNode represents a node in Huffman tree
type HuffmanNode struct {
	Char   rune
	Freq   int
	Left   *HuffmanNode
	Right  *HuffmanNode
	IsLeaf bool
}

// HuffmanCoding implements Huffman coding algorithm
// Time Complexity: O(n log n), Space Complexity: O(n)
func HuffmanCoding(freq map[rune]int) *HuffmanNode {
	// Secure: validate input
	if len(freq) == 0 {
		return nil
	}

	// Secure: validate frequencies
	for char, f := range freq {
		if f < 0 {
			delete(freq, char)
		}
	}

	if len(freq) == 0 {
		return nil
	}

	// Create priority queue (simplified with slice)
	nodes := []*HuffmanNode{}
	for char, f := range freq {
		nodes = append(nodes, &HuffmanNode{
			Char:   char,
			Freq:   f,
			IsLeaf: true,
		})
	}

	// Build Huffman tree
	for len(nodes) > 1 {
		// Sort by frequency
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Freq < nodes[j].Freq
		})

		// Take two nodes with minimum frequency
		left := nodes[0]
		right := nodes[1]
		nodes = nodes[2:]

		// Create new internal node
		merged := &HuffmanNode{
			Freq:   left.Freq + right.Freq,
			Left:   left,
			Right:  right,
			IsLeaf: false,
		}

		nodes = append(nodes, merged)
	}

	if len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

// KruskalMST finds Minimum Spanning Tree using Kruskal's algorithm
// Time Complexity: O(E log E), Space Complexity: O(V)
func KruskalMST(edges []Edge, vertices int) []Edge {
	// Secure: validate input
	if vertices <= 0 || len(edges) == 0 {
		return nil
	}

	// Secure: validate edges
	for _, edge := range edges {
		if edge.From < 0 || edge.To < 0 || edge.From >= vertices || edge.To >= vertices {
			return nil
		}
		if edge.Weight < 0 {
			return nil
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Union-Find data structure
	parent := make([]int, vertices)
	for i := range parent {
		parent[i] = i
	}

	var findFunc func(int) int
	findFunc = func(x int) int {
		if parent[x] != x {
			parent[x] = findFunc(parent[x]) // Path compression
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := findFunc(x), findFunc(y)
		if px == py {
			return false // Cycle detected
		}
		parent[px] = py
		return true
	}

	result := []Edge{}

	for _, edge := range edges {
		if union(edge.From, edge.To) {
			result = append(result, edge)
			if len(result) == vertices-1 {
				break
			}
		}
	}

	return result
}

// Edge represents an edge with from, to, and weight
type Edge struct {
	From   int
	To     int
	Weight int
}
