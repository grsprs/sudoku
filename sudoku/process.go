// Package sudoku implements a recursive backtracking solver for 9x9 Sudoku puzzles.
//
// Author: Spiros Nikoloudakis
// Email: sp.nikoloudakis@gmail.com
// GitHub: @grsprs
// Year: 2026
package sudoku

// Process solves the sudoku puzzle using recursive backtracking
// Returns true if a valid solution exists, false otherwise
func Process() bool {
	// Iterate through all cells to find empty ones
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Found an empty cell (represented by 0)
			if pin[row][col] == 0 {
				// Try placing numbers 1-9
				for num := 1; num <= 9; num++ {
					// Check if number is valid at this position
					if PutNum(row, col, num) {
						// Place the number
						pin[row][col] = num
						// Recursively attempt to solve the rest
						if !Process() {
							return false
						}
						// Backtrack: remove the number and try next (triggered when Process returns true)
						pin[row][col] = 0
					}
				}
				// No valid number found for this cell, trigger backtracking
				return true
			}
		}
	}
	
	// All cells filled - found a complete solution
	sum++
	// Store the first solution found
	if sum == 1 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				solsud[i][j] = pin[i][j]
			}
		}
	}

	// Reject puzzles with multiple solutions
	if sum > 1 {
		return false
	}
	return true
}

// PutNum validates if a number can be placed at the given position
// Checks row, column, and 3x3 box constraints
func PutNum(row int, col int, num int) bool {
	// Check if number exists in the same row or column
	for i := 0; i < 9; i++ {
		if (pin[row][i] == num) || (pin[i][col] == num) {
			return false
		}
	}
	
	// Calculate starting position of the 3x3 box
	srow := (row / 3) * 3
	scol := (col / 3) * 3
	
	// Check if number exists in the 3x3 box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if pin[i+srow][j+scol] == num {
				return false
			}
		}
	}
	return true
}
