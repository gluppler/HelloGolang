package main

import (
	"fmt"
)

// Backtracking Algorithms - Comprehensive implementations of backtracking algorithms

func main() {
	demonstrateBacktrackingAlgorithms()
}

func demonstrateBacktrackingAlgorithms() {
	// N-Queens
	n := 4
	solutions := SolveNQueens(n)
	fmt.Printf("N-Queens (%d): %d solutions\n", n, len(solutions))
	for i, sol := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		printBoard(sol, n)
	}

	// Sudoku Solver
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("\nSudoku Solver:")
	if SolveSudoku(board) {
		printSudoku(board)
	} else {
		fmt.Println("No solution exists")
	}

	// Subset Sum
	arr := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	fmt.Printf("\nSubset Sum: array %v, target %d\n", arr, sum)
	if SubsetSum(arr, sum) {
		fmt.Println("Subset with given sum exists")
	} else {
		fmt.Println("No subset with given sum")
	}

	// Permutations
	nums := []int{1, 2, 3}
	perms := Permute(nums)
	fmt.Printf("\nPermutations of %v: %v\n", nums, perms)
}

// SolveNQueens solves N-Queens problem
// Time Complexity: O(N!), Space Complexity: O(N)
func SolveNQueens(n int) [][]int {
	// Secure: validate input
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return [][]int{{0}}
	}
	if n < 4 && n > 1 {
		return nil // No solution for n=2,3
	}

	solutions := [][]int{}
	board := make([]int, n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// Valid solution found
			solution := make([]int, n)
			copy(solution, board)
			solutions = append(solutions, solution)
			return
		}

		for col := 0; col < n; col++ {
			if isValidQueen(board, row, col) {
				board[row] = col
				backtrack(row + 1)
			}
		}
	}

	backtrack(0)
	return solutions
}

// isValidQueen checks if queen can be placed at (row, col)
func isValidQueen(board []int, row, col int) bool {
	// Secure: bounds checking
	if row < 0 || row >= len(board) || col < 0 {
		return false
	}

	for i := 0; i < row; i++ {
		// Secure: bounds checking
		if i >= len(board) {
			break
		}
		// Check same column
		if board[i] == col {
			return false
		}
		// Check diagonals
		if abs(board[i]-col) == abs(i-row) {
			return false
		}
	}

	return true
}

// abs returns absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// printBoard prints chess board
func printBoard(board []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Secure: bounds checking
			if i < len(board) && board[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// SolveSudoku solves Sudoku puzzle
// Time Complexity: O(9^(empty cells)), Space Complexity: O(1)
func SolveSudoku(board [][]int) bool {
	// Secure: validate board
	if len(board) != 9 {
		return false
	}
	for i := range board {
		if len(board[i]) != 9 {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Secure: bounds checking
			if row >= len(board) || col >= len(board[row]) {
				continue
			}

			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidSudoku(board, row, col, num) {
						board[row][col] = num

						if SolveSudoku(board) {
							return true
						}

						board[row][col] = 0 // Backtrack
					}
				}
				return false
			}
		}
	}

	return true
}

// isValidSudoku checks if number can be placed at (row, col)
func isValidSudoku(board [][]int, row, col, num int) bool {
	// Secure: validate input
	if row < 0 || row >= 9 || col < 0 || col >= 9 || num < 1 || num > 9 {
		return false
	}

	// Secure: bounds checking
	if row >= len(board) || col >= len(board[row]) {
		return false
	}

	// Check row
	for c := 0; c < 9; c++ {
		if c < len(board[row]) && board[row][c] == num {
			return false
		}
	}

	// Check column
	for r := 0; r < 9; r++ {
		if r < len(board) && col < len(board[r]) && board[r][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow, startCol := (row/3)*3, (col/3)*3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if r < len(board) && c < len(board[r]) && board[r][c] == num {
				return false
			}
		}
	}

	return true
}

// printSudoku prints Sudoku board
func printSudoku(board [][]int) {
	// Secure: validate board
	if len(board) != 9 {
		return
	}

	for i := 0; i < 9; i++ {
		// Secure: bounds checking
		if i >= len(board) {
			break
		}
		if len(board[i]) != 9 {
			continue
		}

		for j := 0; j < 9; j++ {
			// Secure: bounds checking
			if j < len(board[i]) {
				fmt.Printf("%d ", board[i][j])
			}
			if (j+1)%3 == 0 && j < 8 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 && i < 8 {
			fmt.Println("------+-------+------")
		}
	}
}

// SubsetSum checks if subset with given sum exists
// Time Complexity: O(2^n), Space Complexity: O(n)
func SubsetSum(arr []int, sum int) bool {
	// Secure: validate input
	if sum < 0 {
		return false
	}
	if len(arr) == 0 {
		return sum == 0
	}

	// Secure: validate array elements
	for _, val := range arr {
		if val < 0 {
			return false
		}
	}

	var backtrack func(index, currentSum int) bool
	backtrack = func(index, currentSum int) bool {
		// Secure: bounds checking
		if index < 0 || index > len(arr) {
			return false
		}

		if currentSum == sum {
			return true
		}

		if index >= len(arr) || currentSum > sum {
			return false
		}

		// Include current element
		if backtrack(index+1, currentSum+arr[index]) {
			return true
		}

		// Exclude current element
		return backtrack(index+1, currentSum)
	}

	return backtrack(0, 0)
}

// Permute generates all permutations
// Time Complexity: O(n! * n), Space Complexity: O(n)
func Permute(nums []int) [][]int {
	// Secure: validate input
	if len(nums) == 0 {
		return [][]int{}
	}

	result := [][]int{}

	var backtrack func(current []int, used []bool)
	backtrack = func(current []int, used []bool) {
		// Secure: bounds checking
		if len(current) == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// Secure: bounds checking
			if i >= len(used) {
				break
			}

			if !used[i] {
				used[i] = true
				current = append(current, nums[i])
				backtrack(current, used)
				current = current[:len(current)-1]
				used[i] = false
			}
		}
	}

	backtrack([]int{}, make([]bool, len(nums)))
	return result
}

// CombinationSum finds all combinations that sum to target
// Time Complexity: O(2^n), Space Complexity: O(target)
func CombinationSum(candidates []int, target int) [][]int {
	// Secure: validate input
	if target < 0 || len(candidates) == 0 {
		return nil
	}

	// Secure: validate candidates
	for _, val := range candidates {
		if val <= 0 {
			return nil
		}
	}

	result := [][]int{}

	var backtrack func(index int, current []int, sum int)
	backtrack = func(index int, current []int, sum int) {
		// Secure: bounds checking
		if index < 0 || index > len(candidates) {
			return
		}

		if sum == target {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		if index >= len(candidates) || sum > target {
			return
		}

		// Include current candidate
		current = append(current, candidates[index])
		backtrack(index, current, sum+candidates[index])
		current = current[:len(current)-1]

		// Skip current candidate
		backtrack(index+1, current, sum)
	}

	backtrack(0, []int{}, 0)
	return result
}

// WordSearch checks if word exists in 2D board
// Time Complexity: O(m * n * 4^L), Space Complexity: O(L)
func WordSearch(board [][]byte, word string) bool {
	// Secure: validate input
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	// Secure: validate board dimensions
	for i := range board {
		if len(board[i]) != n {
			return false
		}
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		// Secure: bounds checking
		if index >= len(word) {
			return true
		}
		if row < 0 || row >= m || col < 0 || col >= n {
			return false
		}
		if row >= len(board) || col >= len(board[row]) {
			return false
		}
		if row >= len(visited) || col >= len(visited[row]) {
			return false
		}
		if visited[row][col] {
			return false
		}
		if index >= len(word) || board[row][col] != word[index] {
			return false
		}

		visited[row][col] = true

		// Check all directions
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			if backtrack(row+dir[0], col+dir[1], index+1) {
				return true
			}
		}

		visited[row][col] = false
		return false
	}

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
