# Sudoku Solver

**Developer:** Spiros Nikoloudakis  
**Year:** 2026

## Description

A command-line Sudoku solver written in Go that uses recursive backtracking algorithm to solve 9x9 Sudoku puzzles. The program validates input, ensures unique solutions, and handles invalid or unsolvable puzzles.

## Features

- ✓ Recursive backtracking algorithm
- ✓ Input validation (format, length, characters)
- ✓ Unique solution verification
- ✓ Error handling for invalid inputs
- ✓ Comprehensive test coverage

## Installation

```bash
git clone <repository-url>
cd sudoku
go mod download
```

## Usage

### Instructions

- Create a program that resolves a sudoku.

- A valid sudoku has only one possible solution.

Make sure you submit all the necessary files to run the program.

### Running the Program

```bash
go run . "row1" "row2" "row3" "row4" "row5" "row6" "row7" "row8" "row9"
```

Each argument represents one row of the Sudoku grid:
- Use `.` for empty cells
- Use digits `1-9` for filled cells
- Each row must contain exactly 9 characters

### Example 1:

Example of output for a valid sudoku :

```console
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
$
```

### Example 2:

Examples of expected outputs for invalid inputs or sudokus :

```console
$ go run . 1 2 3 4 | cat -e
Error$
$ go run . | cat -e
Error$
$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
$
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Project Structure

```
sudoku/
├── go.mod              # Go module definition
├── main.go             # Entry point
├── LICENSE             # MIT License
├── README.md           # This file
├── .gitignore          # Git ignore rules
├── DOCUMENTATION.md    # Detailed documentation
├── PROJECT_ANALYSIS.md # Algorithm analysis
└── sudoku/
    ├── io.go           # Input/output handling
    ├── io_test.go      # Input validation tests
    ├── process.go      # Solving logic
    └── process_test.go # Solving logic tests
```

## Algorithm

The solver uses **recursive backtracking**:

1. Find an empty cell
2. Try placing numbers 1-9
3. For each number, validate against Sudoku rules:
   - Number must not exist in the same row
   - Number must not exist in the same column
   - Number must not exist in the same 3x3 box
4. If valid, place the number and recursively solve the rest
5. If no solution found, backtrack and try the next number
6. Ensure only one unique solution exists

## Requirements

- Go 1.21 or higher
- No external dependencies

## License

MIT License - see [LICENSE](LICENSE) file for details

## Author

**Spiros Nikoloudakis** - 2026
