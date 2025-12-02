package main

import (
	"fmt"
)

// Tree Algorithms - Comprehensive implementations of tree algorithms

func main() {
	demonstrateTreeAlgorithms()
}

func demonstrateTreeAlgorithms() {
	// Create sample BST
	root := &TreeNode{Value: 50}
	root.Insert(30)
	root.Insert(70)
	root.Insert(20)
	root.Insert(40)
	root.Insert(60)
	root.Insert(80)

	fmt.Println("Tree Traversals:")
	fmt.Print("Inorder: ")
	root.Inorder()
	fmt.Println()

	fmt.Print("Preorder: ")
	root.Preorder()
	fmt.Println()

	fmt.Print("Postorder: ")
	root.Postorder()
	fmt.Println()

	// Search
	found := root.Search(40)
	fmt.Printf("Search 40: %t\n", found)

	// Delete
	root = root.Delete(20)
	fmt.Print("After deleting 20, Inorder: ")
	root.Inorder()
	fmt.Println()
}

// TreeNode represents a node in binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a value into BST
// Time Complexity: O(h), Space Complexity: O(h)
func (n *TreeNode) Insert(value int) {
	// Secure: validate input
	if n == nil {
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
	// If value == n.Value, do nothing (no duplicates)
}

// Search searches for a value in BST
// Time Complexity: O(h), Space Complexity: O(h)
func (n *TreeNode) Search(value int) bool {
	// Secure: validate input
	if n == nil {
		return false
	}

	if value == n.Value {
		return true
	}

	if value < n.Value {
		return n.Left.Search(value)
	}

	return n.Right.Search(value)
}

// Delete deletes a value from BST
// Time Complexity: O(h), Space Complexity: O(h)
func (n *TreeNode) Delete(value int) *TreeNode {
	// Secure: validate input
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		// Node to delete found
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}

		// Node with two children
		minNode := n.Right.MinValue()
		n.Value = minNode.Value
		n.Right = n.Right.Delete(minNode.Value)
	}

	return n
}

// MinValue finds minimum value node
func (n *TreeNode) MinValue() *TreeNode {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Inorder performs inorder traversal
func (n *TreeNode) Inorder() {
	if n == nil {
		return
	}
	n.Left.Inorder()
	fmt.Printf("%d ", n.Value)
	n.Right.Inorder()
}

// Preorder performs preorder traversal
func (n *TreeNode) Preorder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	n.Left.Preorder()
	n.Right.Preorder()
}

// Postorder performs postorder traversal
func (n *TreeNode) Postorder() {
	if n == nil {
		return
	}
	n.Left.Postorder()
	n.Right.Postorder()
	fmt.Printf("%d ", n.Value)
}

// LevelOrder performs level-order traversal (BFS)
// Time Complexity: O(n), Space Complexity: O(n)
func (n *TreeNode) LevelOrder() []int {
	if n == nil {
		return nil
	}

	result := []int{}
	queue := []*TreeNode{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// Height calculates height of tree
// Time Complexity: O(n), Space Complexity: O(h)
func (n *TreeNode) Height() int {
	if n == nil {
		return -1
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsBalanced checks if tree is balanced
// Time Complexity: O(n), Space Complexity: O(h)
func (n *TreeNode) IsBalanced() bool {
	return n.checkBalance() != -1
}

// checkBalance returns height if balanced, -1 otherwise
func (n *TreeNode) checkBalance() int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.checkBalance()
	if leftHeight == -1 {
		return -1
	}

	rightHeight := n.Right.checkBalance()
	if rightHeight == -1 {
		return -1
	}

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Diameter calculates diameter of tree
// Time Complexity: O(n), Space Complexity: O(h)
func (n *TreeNode) Diameter() int {
	if n == nil {
		return 0
	}

	maxDiameter := 0
	n.diameterHelper(&maxDiameter)
	return maxDiameter
}

// diameterHelper calculates diameter recursively
func (n *TreeNode) diameterHelper(maxDiameter *int) int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.diameterHelper(maxDiameter)
	rightHeight := n.Right.diameterHelper(maxDiameter)

	currentDiameter := leftHeight + rightHeight + 1
	if currentDiameter > *maxDiameter {
		*maxDiameter = currentDiameter
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// LowestCommonAncestor finds LCA of two nodes
// Time Complexity: O(h), Space Complexity: O(h)
func (n *TreeNode) LowestCommonAncestor(p, q int) *TreeNode {
	if n == nil {
		return nil
	}

	if n.Value > p && n.Value > q {
		return n.Left.LowestCommonAncestor(p, q)
	}

	if n.Value < p && n.Value < q {
		return n.Right.LowestCommonAncestor(p, q)
	}

	return n
}

// ValidateBST checks if tree is valid BST
// Time Complexity: O(n), Space Complexity: O(h)
func (n *TreeNode) ValidateBST() bool {
	return n.validateBSTHelper(nil, nil)
}

// validateBSTHelper validates BST with min and max bounds
func (n *TreeNode) validateBSTHelper(min, max *int) bool {
	if n == nil {
		return true
	}

	if min != nil && n.Value <= *min {
		return false
	}

	if max != nil && n.Value >= *max {
		return false
	}

	return n.Left.validateBSTHelper(min, &n.Value) &&
		n.Right.validateBSTHelper(&n.Value, max)
}
