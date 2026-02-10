// Package sudoku implements a recursive backtracking solver for 9x9 Sudoku puzzles.
//
// The solver validates input, ensures unique solutions, and handles invalid or
// unsolvable puzzles. It uses a depth-first search with backtracking to explore
// the solution space efficiently.
//
// Usage:
//
//	package main
//
//	import "sudoku/sudoku"
//
//	func main() {
//		sudoku.Inpt()
//	}
//
// The program expects exactly 9 command-line arguments, each representing a row
// of the Sudoku grid. Empty cells are represented by '.' and filled cells by
// digits '1'-'9'.
//
// Author: Spiros Nikoloudakis
// Email: sp.nikoloudakis@gmail.com
// GitHub: @grsprs
// Year: 2026
package sudoku
