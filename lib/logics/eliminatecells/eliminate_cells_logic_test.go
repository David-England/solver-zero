package eliminatecells_test

import (
	"solver-zero/lib"
	"solver-zero/lib/logics/eliminatecells"
	"testing"
)

func TestPureElimination(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 4, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 4, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{4, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 4, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	logic := eliminatecells.EliminateCellsLogic{&sudoku}

	// Act
	isChanged, _ := logic.RunStep()

	// Assert
	if !isChanged {
		t.Fatalf("should have changed but claiming to have not")
	}
	if sudoku.Grid[2][2] != 4 {
		t.Fatalf("subgrid failed to resolve")
	}
}

func TestResolveTrivialRow(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 2, 3, 0, 5, 6, 7, 8, 9},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	logic := eliminatecells.EliminateCellsLogic{&sudoku}

	// Act
	isChanged, _ := logic.RunStep()

	// Assert
	if !isChanged {
		t.Fatalf("should have changed but claiming to have not")
	}
	if sudoku.Grid[3][3] != 4 {
		t.Fatalf("row failed to resolve")
	}
}

func TestResolveTrivialColumn(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 3, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 0, 0, 0, 0},
			{0, 0, 0, 6, 0, 0, 0, 0, 0},
			{0, 0, 0, 7, 0, 0, 0, 0, 0},
			{0, 0, 0, 8, 0, 0, 0, 0, 0},
			{0, 0, 0, 9, 0, 0, 0, 0, 0},
		},
	}
	logic := eliminatecells.EliminateCellsLogic{&sudoku}

	// Act
	isChanged, _ := logic.RunStep()

	// Assert
	if !isChanged {
		t.Fatalf("should have changed but claiming to have not")
	}
	if sudoku.Grid[3][3] != 4 {
		t.Fatalf("column failed to resolve")
	}
}

func TestResolveTrivialSubgrid(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 2, 3},
			{0, 0, 0, 0, 0, 0, 0, 5, 6},
			{0, 0, 0, 0, 0, 0, 7, 8, 9},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	logic := eliminatecells.EliminateCellsLogic{&sudoku}

	// Act
	isChanged, _ := logic.RunStep()

	// Assert
	if !isChanged {
		t.Fatalf("should have changed but claiming to have not")
	}
	if sudoku.Grid[4][6] != 4 {
		t.Fatalf("subgrid failed to resolve")
	}
}

func TestResolveNothing(t *testing.T) {
	// Arrange
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	logic := eliminatecells.EliminateCellsLogic{&sudoku}

	// Act
	isChanged, _ := logic.RunStep()

	// Assert
	if isChanged {
		t.Fatalf("claiming to have changed when shouldn't have")
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.Grid[i][j] != 0 {
				t.Fatalf("cell (%v, %v) is %v but should be 0", i, j, sudoku.Grid[i][j])
			}
		}
	}
}
