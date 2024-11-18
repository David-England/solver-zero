package lib_test

import (
	"solver-zero/lib"
	"testing"
)

func TestGetSubs(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{4, 4, 4, 0, 0, 0, 0, 0, 0},
			{5, 5, 5, 0, 0, 0, 0, 0, 0},
			{6, 6, 6, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Act
	subgrids := sudoku.GetSubs()

	// Assert
	if subgrids[1][0].Grid != [3][3]int{{4, 4, 4}, {5, 5, 5}, {6, 6, 6}} {
		t.Fatalf("method GetSubs() returned wrong subgrid")
	}
}
