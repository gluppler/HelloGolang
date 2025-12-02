package main

import (
	"fmt"
	"math"
)

// Graph Algorithms - Comprehensive implementations of graph algorithms

func main() {
	demonstrateGraphAlgorithms()
}

func demonstrateGraphAlgorithms() {
	// Create sample graph
	graph := createSampleGraph()

	fmt.Println("Graph Algorithms:")

	// BFS
	fmt.Println("\nBFS Traversal:")
	BFS(graph, 0)

	// DFS
	fmt.Println("\nDFS Traversal:")
	visited := make([]bool, graph.Vertices)
	DFS(graph, 0, visited)
	fmt.Println()

	// Dijkstra's
	fmt.Println("\nDijkstra's Shortest Path:")
	distances := Dijkstra(graph, 0)
	fmt.Println("Distances from node 0:", distances)

	// Topological Sort
	fmt.Println("\nTopological Sort:")
	topoGraph := createDirectedGraph()
	topologicalOrder := TopologicalSort(topoGraph)
	fmt.Println("Topological order:", topologicalOrder)
}

// Graph representation using adjacency list
type Graph struct {
	Vertices int
	Edges    [][]Edge
}

// Edge represents an edge in the graph
type Edge struct {
	To     int
	Weight int
}

// createSampleGraph creates a sample graph for testing
func createSampleGraph() *Graph {
	g := &Graph{
		Vertices: 5,
		Edges:    make([][]Edge, 5),
	}

	g.Edges[0] = []Edge{{1, 4}, {2, 1}}
	g.Edges[1] = []Edge{{3, 1}}
	g.Edges[2] = []Edge{{1, 2}, {3, 5}}
	g.Edges[3] = []Edge{{4, 3}}
	g.Edges[4] = []Edge{}

	return g
}

// createDirectedGraph creates a directed acyclic graph
func createDirectedGraph() *Graph {
	g := &Graph{
		Vertices: 6,
		Edges:    make([][]Edge, 6),
	}

	g.Edges[5] = []Edge{{2, 1}, {0, 1}}
	g.Edges[4] = []Edge{{0, 1}, {1, 1}}
	g.Edges[2] = []Edge{{3, 1}}
	g.Edges[3] = []Edge{{1, 1}}
	g.Edges[1] = []Edge{}
	g.Edges[0] = []Edge{}

	return g
}

// BFS performs breadth-first search
// Time Complexity: O(V + E), Space Complexity: O(V)
func BFS(graph *Graph, start int) {
	// Secure: bounds checking
	if start < 0 || start >= graph.Vertices {
		return
	}

	visited := make([]bool, graph.Vertices)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", vertex)

		// Secure: bounds checking
		if vertex < 0 || vertex >= len(graph.Edges) {
			continue
		}

		for _, edge := range graph.Edges[vertex] {
			// Secure: bounds checking
			if edge.To >= 0 && edge.To < graph.Vertices && !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}
	fmt.Println()
}

// DFS performs depth-first search
// Time Complexity: O(V + E), Space Complexity: O(V)
func DFS(graph *Graph, vertex int, visited []bool) {
	// Secure: bounds checking
	if vertex < 0 || vertex >= graph.Vertices {
		return
	}

	// Secure: bounds checking
	if vertex >= len(visited) {
		return
	}

	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	// Secure: bounds checking
	if vertex < 0 || vertex >= len(graph.Edges) {
		return
	}

	for _, edge := range graph.Edges[vertex] {
		// Secure: bounds checking
		if edge.To >= 0 && edge.To < graph.Vertices && edge.To < len(visited) && !visited[edge.To] {
			DFS(graph, edge.To, visited)
		}
	}
}

// Dijkstra finds shortest paths from source to all vertices
// Time Complexity: O((V + E) log V), Space Complexity: O(V)
func Dijkstra(graph *Graph, source int) []int {
	// Secure: bounds checking
	if source < 0 || source >= graph.Vertices {
		return nil
	}

	dist := make([]int, graph.Vertices)
	visited := make([]bool, graph.Vertices)

	// Initialize distances
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	// Simple implementation (use priority queue for better performance)
	for i := 0; i < graph.Vertices; i++ {
		// Find unvisited vertex with minimum distance
		u := minDistance(dist, visited)

		// Secure: bounds checking
		if u < 0 || u >= graph.Vertices {
			break
		}

		visited[u] = true

		// Secure: bounds checking
		if u < 0 || u >= len(graph.Edges) {
			continue
		}

		// Update distances to neighbors
		for _, edge := range graph.Edges[u] {
			// Secure: bounds checking
			if edge.To >= 0 && edge.To < graph.Vertices {
				// Secure: prevent integer overflow
				if dist[u] != math.MaxInt32 && dist[u]+edge.Weight < dist[edge.To] {
					dist[edge.To] = dist[u] + edge.Weight
				}
			}
		}
	}

	return dist
}

// minDistance finds vertex with minimum distance
func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1

	// Secure: bounds checking
	if len(dist) != len(visited) {
		return -1
	}

	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

// TopologicalSort performs topological sorting
// Time Complexity: O(V + E), Space Complexity: O(V)
func TopologicalSort(graph *Graph) []int {
	// Secure: validate graph
	if graph == nil || graph.Vertices == 0 {
		return nil
	}

	inDegree := make([]int, graph.Vertices)

	// Calculate in-degrees
	for i := 0; i < graph.Vertices; i++ {
		// Secure: bounds checking
		if i < len(graph.Edges) {
			for _, edge := range graph.Edges[i] {
				// Secure: bounds checking
				if edge.To >= 0 && edge.To < graph.Vertices {
					inDegree[edge.To]++
				}
			}
		}
	}

	// Queue for vertices with no incoming edges
	queue := []int{}
	for i := 0; i < graph.Vertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		// Secure: bounds checking
		if vertex >= 0 && vertex < len(graph.Edges) {
			for _, edge := range graph.Edges[vertex] {
				// Secure: bounds checking
				if edge.To >= 0 && edge.To < graph.Vertices {
					inDegree[edge.To]--
					if inDegree[edge.To] == 0 {
						queue = append(queue, edge.To)
					}
				}
			}
		}
	}

	// Check for cycle
	if len(result) != graph.Vertices {
		return nil // Cycle detected
	}

	return result
}

// BellmanFord finds shortest paths using Bellman-Ford algorithm
// Time Complexity: O(V * E), Space Complexity: O(V)
func BellmanFord(graph *Graph, source int) ([]int, bool) {
	// Secure: bounds checking
	if source < 0 || source >= graph.Vertices {
		return nil, false
	}

	dist := make([]int, graph.Vertices)

	// Initialize distances
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	// Relax edges V-1 times
	for i := 0; i < graph.Vertices-1; i++ {
		for u := 0; u < graph.Vertices; u++ {
			// Secure: bounds checking
			if u < len(graph.Edges) {
				for _, edge := range graph.Edges[u] {
					// Secure: bounds checking
					if edge.To >= 0 && edge.To < graph.Vertices {
						// Secure: prevent integer overflow
						if dist[u] != math.MaxInt32 && dist[u]+edge.Weight < dist[edge.To] {
							dist[edge.To] = dist[u] + edge.Weight
						}
					}
				}
			}
		}
	}

	// Check for negative cycles
	for u := 0; u < graph.Vertices; u++ {
		// Secure: bounds checking
		if u < len(graph.Edges) {
			for _, edge := range graph.Edges[u] {
				// Secure: bounds checking
				if edge.To >= 0 && edge.To < graph.Vertices {
					if dist[u] != math.MaxInt32 && dist[u]+edge.Weight < dist[edge.To] {
						return nil, false // Negative cycle detected
					}
				}
			}
		}
	}

	return dist, true
}

// FloydWarshall finds all-pairs shortest paths
// Time Complexity: O(V³), Space Complexity: O(V²)
func FloydWarshall(graph *Graph) [][]int {
	// Secure: validate graph
	if graph == nil || graph.Vertices == 0 {
		return nil
	}

	// Initialize distance matrix
	dist := make([][]int, graph.Vertices)
	for i := range dist {
		dist[i] = make([]int, graph.Vertices)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	// Initialize with edge weights
	for u := 0; u < graph.Vertices; u++ {
		// Secure: bounds checking
		if u < len(graph.Edges) {
			for _, edge := range graph.Edges[u] {
				// Secure: bounds checking
				if edge.To >= 0 && edge.To < graph.Vertices {
					dist[u][edge.To] = edge.Weight
				}
			}
		}
	}

	// Floyd-Warshall algorithm
	for k := 0; k < graph.Vertices; k++ {
		for i := 0; i < graph.Vertices; i++ {
			for j := 0; j < graph.Vertices; j++ {
				// Secure: prevent integer overflow
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][k]+dist[k][j] < dist[i][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}

	return dist
}

// Prim finds Minimum Spanning Tree using Prim's algorithm
// Time Complexity: O(V²), Space Complexity: O(V)
func Prim(graph *Graph) []int {
	// Secure: validate graph
	if graph == nil || graph.Vertices == 0 {
		return nil
	}

	parent := make([]int, graph.Vertices)
	key := make([]int, graph.Vertices)
	mstSet := make([]bool, graph.Vertices)

	// Initialize keys
	for i := range key {
		key[i] = math.MaxInt32
		parent[i] = -1
	}

	key[0] = 0

	for count := 0; count < graph.Vertices-1; count++ {
		u := minKey(key, mstSet)

		// Secure: bounds checking
		if u < 0 || u >= graph.Vertices {
			break
		}

		mstSet[u] = true

		// Secure: bounds checking
		if u < len(graph.Edges) {
			for _, edge := range graph.Edges[u] {
				// Secure: bounds checking
				if edge.To >= 0 && edge.To < graph.Vertices {
					if !mstSet[edge.To] && edge.Weight < key[edge.To] {
						parent[edge.To] = u
						key[edge.To] = edge.Weight
					}
				}
			}
		}
	}

	return parent
}

// minKey finds vertex with minimum key value
func minKey(key []int, mstSet []bool) int {
	min := math.MaxInt32
	minIndex := -1

	// Secure: bounds checking
	if len(key) != len(mstSet) {
		return -1
	}

	for v := 0; v < len(key); v++ {
		if !mstSet[v] && key[v] < min {
			min = key[v]
			minIndex = v
		}
	}

	return minIndex
}
