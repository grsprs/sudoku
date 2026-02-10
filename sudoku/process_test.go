package sudoku

import "testing"

// TestPutNumValidation tests the validation logic for placing numbers
func TestPutNumValidation(t *testing.T) {
	pin = make([][]int, 9)
	for i := 0; i < 9; i++ {
		pin[i] = make([]int, 9)
	}

	// Test row validation
	pin[0][0] = 5
	if PutNum(0, 1, 5) {
		t.Error("PutNum should return false when number exists in row")
	}

	// Test column validation
	pin[1][0] = 0
	pin[0][0] = 0
	pin[1][1] = 7
	if PutNum(0, 1, 7) {
		t.Error("PutNum should return false when number exists in column")
	}

	// Test 3x3 box validation
	pin[1][1] = 0
	pin[2][2] = 9
	if PutNum(0, 0, 9) {
		t.Error("PutNum should return false when number exists in 3x3 box")
	}

	// Test valid placement
	if !PutNum(0, 0, 1) {
		t.Error("PutNum should return true for valid placement")
	}
}
