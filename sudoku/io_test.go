package sudoku

import (
	"os"
	"testing"
)

// TestValidInput tests valid sudoku input
func TestValidInput(t *testing.T) {
	os.Args = []string{"sudoku", ".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	Inpt()
}

// TestInvalidArgumentCount tests error handling for wrong number of arguments
func TestInvalidArgumentCount(t *testing.T) {
	os.Args = []string{"sudoku", "1", "2", "3"}
	Inpt()
}

// TestInvalidCharacters tests error handling for invalid characters
func TestInvalidCharacters(t *testing.T) {
	os.Args = []string{"sudoku", "not", "a", "sudoku", "puzzle", "here", "test", "data", "more", "args"}
	Inpt()
}

// TestInvalidLength tests error handling for wrong string length
func TestInvalidLength(t *testing.T) {
	os.Args = []string{"sudoku", "53..8294.", "8..34...5", "3542761..", "..6.3...4", "9....162.", ".9...3.78", "7438.9...", "..5..43.1"}
	Inpt()
}
