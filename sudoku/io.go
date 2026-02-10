// Package sudoku implements a recursive backtracking solver for 9x9 Sudoku puzzles.
//
// Author: Spiros Nikoloudakis
// Email: sp.nikoloudakis@gmail.com
// GitHub: @grsprs
// Year: 2026
package sudoku

import (
	"fmt"
	"os"
)

// pin stores the current state of the sudoku grid during solving
var pin [][]int

// solsud stores the final solved sudoku grid
var solsud [][]int

// sum tracks the number of solutions found to ensure uniqueness
var sum int

// Inpt validates command-line arguments, processes input, and initiates solving
func Inpt() {
	// Initialize solution counter and allocate memory for grids
	sum = 0
	pin = make([][]int, 9)
	solsud = make([][]int, 9)
	fmt.Println()
	
	// Validate argument count (program name + 9 rows)
	if len(os.Args) != 10 {
		fmt.Println("Error Length of Arguments")
		return
	}
	
	fmt.Println(" Input Sudoku")
	fmt.Println("--------------")
	
	// Process each row of the sudoku grid
	for i := 1; i <= 9; i++ {
		pin[i-1] = make([]int, 9)
		solsud[i-1] = make([]int, 9)
		
		// Validate row length must be exactly 9 characters
		if len(os.Args[i]) != 9 {
			fmt.Println("Error Length of number")
			return
		}

		// Process each cell in the row
		for j := 0; j < 9; j++ {
			// Validate character is either '.' or digit '1'-'9'
			if !((os.Args[i][j] == '.') || (os.Args[i][j] >= '1' && os.Args[i][j] <= '9')) {
				fmt.Println("Error no <.> or number")
				return
			} else {
				// Convert '.' to 0, digits to their numeric value
				if os.Args[i][j] == '.' {
					pin[i-1][j] = 0
				} else {
					pin[i-1][j] = int(os.Args[i][j] - '0')
				}
				fmt.Print(pin[i-1][j], "  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	
	// Attempt to solve the sudoku puzzle
	if Process() {
		// Only print solution if exactly one solution was found
		if sum == 1 {
			fmt.Println(" Solve Sudoku")
			fmt.Println("--------------")
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Print(solsud[i][j])
					fmt.Print("  ")
				}
				fmt.Println()
			}
		}
	}
}
