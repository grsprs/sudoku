# Sudoku Solver - Documentation

**Developer:** Spiros Nikoloudakis  
**Year:** 2026

## Description

This program solves 9x9 Sudoku puzzles using a backtracking algorithm. It takes 9 command-line arguments, each representing a row of the Sudoku grid, where dots (`.`) represent empty cells and digits (`1-9`) represent filled cells.

## Project Structure

```
sudoku/
├── go.mod                  # Go module definition
├── main.go                 # Entry point
├── README.md              # This file
└── sudoku/
    ├── io.go              # Input validation and output formatting
    ├── io_test.go         # Tests for input validation
    ├── process.go         # Sudoku solving logic
    └── process_test.go    # Tests for solving logic
```

## Algorithm

The solver uses a **backtracking algorithm** with the following approach:

1. **Input Validation**: Validates that exactly 9 arguments are provided, each containing exactly 9 characters (digits 1-9 or dots)
2. **Recursive Solving**: Iterates through empty cells and tries placing numbers 1-9
3. **Validation**: For each placement, checks if the number is valid according to Sudoku rules:
   - Number must not exist in the same row
   - Number must not exist in the same column
   - Number must not exist in the same 3x3 box
4. **Backtracking**: If a placement leads to no solution, backtracks and tries the next number
5. **Solution Detection**: Tracks the number of solutions found and stops after finding the first valid solution

## Usage

### Running the Program

```bash
go run . "row1" "row2" "row3" "row4" "row5" "row6" "row7" "row8" "row9"
```

### Example

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

### Expected Output

The program displays:
1. Input Sudoku grid
2. Solved Sudoku grid (if solvable)
3. "Error" message for invalid inputs or unsolvable puzzles

## Input Format

- Each argument represents one row of the Sudoku grid
- Use `.` for empty cells
- Use digits `1-9` for filled cells
- Each row must contain exactly 9 characters

## Error Handling

The program prints "Error" and exits for:
- Wrong number of arguments (not exactly 9)
- Invalid row length (not exactly 9 characters)
- Invalid characters (not `.` or `1-9`)
- Unsolvable puzzles
- Puzzles with multiple solutions

## Testing

Run tests with:

```bash
go test ./...
```

Tests cover:
- Valid input processing
- Invalid argument count
- Invalid characters
- Invalid string length
- Number placement validation
- Row, column, and box validation

## Implementation Details

### Key Functions

- **Inpt()**: Validates command-line input and initiates solving
- **Process()**: Main solving function using backtracking
- **PutNum()**: Validates if a number can be placed at a given position
- **solsud**: Stores the first valid solution found
- **sum**: Tracks the number of solutions to ensure uniqueness

### Variables

- **pin**: Current state of the Sudoku grid during solving
- **solsud**: Stores the final solution
- **sum**: Counter for number of solutions found

## Requirements

- Go 1.21 or higher
- No external dependencies

## Notes

- The solver ensures only one valid solution exists
- If multiple solutions are found, the program returns an error
- The backtracking algorithm guarantees finding a solution if one exists
