package lib

import (
	"testing"
)

func TestCreateSudoku(t *testing.T) {
	// Arrange
	num := 4
	grid := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, num, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Act
	sudoku := CreateSudoku(grid)

	// Assert
	if sudoku.Grid[3][3] != num {
		t.Fatalf("method CreateSudoku() failed to imbibe grid")
	}
	if !sudoku.pencilMarks.cantBe[3][4][num-1] {
		t.Fatalf("method CreateSudoku() failed to eliminate option in row of resolved entry")
	}
}

func TestGetSubs(t *testing.T) {
	// Arrange
	sudoku := Sudoku{
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
