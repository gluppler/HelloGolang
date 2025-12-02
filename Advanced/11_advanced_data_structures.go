package main

import (
	"fmt"
)

// Advanced Data Structures demonstrates advanced data structure implementations

func main() {
	linkedList()
	binaryTree()
	heap()
	trie()
	graph()
}

// linkedList demonstrates linked list implementation
func linkedList() {
	type Node struct {
		Value int
		Next  *Node
	}

	type LinkedList struct {
		Head *Node
	}

	append := func(list *LinkedList, value int) {
		newNode := &Node{Value: value}

		if list.Head == nil {
			list.Head = newNode
			return
		}

		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	prepend := func(list *LinkedList, value int) {
		newNode := &Node{Value: value, Next: list.Head}
		list.Head = newNode
	}

	display := func(list *LinkedList) {
		current := list.Head
		for current != nil {
			fmt.Printf("%d -> ", current.Value)
			current = current.Next
		}
		fmt.Println("nil")
	}

	list := &LinkedList{}
	append(list, 1)
	append(list, 2)
	append(list, 3)
	prepend(list, 0)

	fmt.Println("Linked List:")
	display(list)
}

// TreeNode represents a binary tree node
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// insert inserts a value into BST
func insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value}
	}
	
	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}
	
	return root
}

// inOrderTraversal performs inorder traversal
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Value)
	inOrderTraversal(root.Right)
}

// binaryTree demonstrates binary tree implementation
func binaryTree() {
	var root *TreeNode
	root = insert(root, 5)
	root = insert(root, 3)
	root = insert(root, 7)
	root = insert(root, 1)
	root = insert(root, 4)
	
	fmt.Println("Binary Tree (in-order):")
	inOrderTraversal(root)
	fmt.Println()
}

// heap demonstrates heap implementation
func heap() {
	type MinHeap struct {
		data []int
	}

	parent := func(i int) int {
		return (i - 1) / 2
	}

	left := func(i int) int {
		return 2*i + 1
	}

	right := func(i int) int {
		return 2*i + 2
	}

	heapifyUp := func(h *MinHeap, i int) {
		for i > 0 && h.data[parent(i)] > h.data[i] {
			h.data[parent(i)], h.data[i] = h.data[i], h.data[parent(i)]
			i = parent(i)
		}
	}

	var heapifyDown func(*MinHeap, int)
	heapifyDown = func(h *MinHeap, i int) {
		smallest := i
		l := left(i)
		r := right(i)

		if l < len(h.data) && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < len(h.data) && h.data[r] < h.data[smallest] {
			smallest = r
		}

		if smallest != i {
			h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
			heapifyDown(h, smallest)
		}
	}

	insert := func(h *MinHeap, value int) {
		h.data = append(h.data, value)
		heapifyUp(h, len(h.data)-1)
	}

	extractMin := func(h *MinHeap) (int, bool) {
		if len(h.data) == 0 {
			return 0, false
		}

		min := h.data[0]
		h.data[0] = h.data[len(h.data)-1]
		h.data = h.data[:len(h.data)-1]

		if len(h.data) > 0 {
			heapifyDown(h, 0)
		}

		return min, true
	}

	heap := &MinHeap{}
	insert(heap, 5)
	insert(heap, 3)
	insert(heap, 7)
	insert(heap, 1)
	insert(heap, 4)

	fmt.Println("Min Heap:")
	for {
		if min, ok := extractMin(heap); ok {
			fmt.Printf("%d ", min)
		} else {
			break
		}
	}
	fmt.Println()
}

// trie demonstrates trie (prefix tree) implementation
func trie() {
	type TrieNode struct {
		children map[rune]*TrieNode
		isEnd    bool
	}

	type Trie struct {
		root *TrieNode
	}

	newTrie := func() *Trie {
		return &Trie{
			root: &TrieNode{
				children: make(map[rune]*TrieNode),
			},
		}
	}

	insert := func(t *Trie, word string) {
		current := t.root
		for _, char := range word {
			if current.children[char] == nil {
				current.children[char] = &TrieNode{
					children: make(map[rune]*TrieNode),
				}
			}
			current = current.children[char]
		}
		current.isEnd = true
	}

	search := func(t *Trie, word string) bool {
		current := t.root
		for _, char := range word {
			if current.children[char] == nil {
				return false
			}
			current = current.children[char]
		}
		return current.isEnd
	}

	trie := newTrie()
	insert(trie, "hello")
	insert(trie, "world")
	insert(trie, "help")

	fmt.Println("Trie:")
	fmt.Printf("  Search 'hello': %t\n", search(trie, "hello"))
	fmt.Printf("  Search 'help': %t\n", search(trie, "help"))
	fmt.Printf("  Search 'hel': %t\n", search(trie, "hel"))
}

// graph demonstrates graph implementation
func graph() {
	type Graph struct {
		vertices map[int][]int
	}

	newGraph := func() *Graph {
		return &Graph{
			vertices: make(map[int][]int),
		}
	}

	addEdge := func(g *Graph, from, to int) {
		g.vertices[from] = append(g.vertices[from], to)
		// For undirected graph, also add reverse edge
		g.vertices[to] = append(g.vertices[to], from)
	}

	var dfs func(*Graph, int, map[int]bool)
	dfs = func(g *Graph, start int, visited map[int]bool) {
		visited[start] = true
		fmt.Printf("%d ", start)

		for _, neighbor := range g.vertices[start] {
			if !visited[neighbor] {
				dfs(g, neighbor, visited)
			}
		}
	}

	graph := newGraph()
	addEdge(graph, 0, 1)
	addEdge(graph, 0, 2)
	addEdge(graph, 1, 3)
	addEdge(graph, 2, 4)

	fmt.Println("Graph DFS:")
	visited := make(map[int]bool)
	dfs(graph, 0, visited)
	fmt.Println()
}
