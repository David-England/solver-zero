package pencil

import (
	"slices"
	"solver-zero/lib"
	"testing"
)

func TestBanRow(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	num := 4
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, num, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Act
	sut.EliminateOptions(&sudoku)

	// Assert
	row := 3
	for col := 0; col < 9; col++ {
		if !sut.cantBe[row][col][num-1] {
			t.Fatalf("should ban row %v, but (%v, %v) was not banned", row, row, col)
		}
	}
}

func TestBanColumn(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	num := 4
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, num, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Act
	sut.EliminateOptions(&sudoku)

	// Assert
	col := 3
	for row := 0; row < 9; row++ {
		if !sut.cantBe[row][col][num-1] {
			t.Fatalf("should ban column %v, but (%v, %v) was not banned", col, row, col)
		}
	}
}

func TestBanSubgrid(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	num := 4
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, num, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Act
	sut.EliminateOptions(&sudoku)

	// Assert
	for row := 3; row < 6; row++ {
		for col := 3; col < 6; col++ {
			if !sut.cantBe[row][col][num-1] {
				t.Fatalf("should ban central subgrid, but (%v, %v) was not banned", row, col)
			}
		}
	}
}

func TestBanAllNumbers(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	num := 4
	sudoku := lib.Sudoku{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, num, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Act
	sut.EliminateOptions(&sudoku)

	// Assert
	row, col := 3, 3
	for num := 1; num <= 9; num++ {
		if !sut.cantBe[row][col][num-1] {
			t.Fatalf("should ban all numbers for solved cell (%v, %v), but %v was not",
				row, col, num)
		}
	}
}

func TestBanNothingForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
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

	// Act
	sut.EliminateOptions(&sudoku)

	// Assert
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for num := 1; num <= 9; num++ {
				if sut.cantBe[i][j][num-1] {
					t.Fatalf("shouldn't eliminate any for empty sudoku,"+
						" but eliminated %v for (%v, %v)", num, i, j)
				}
			}
		}
	}
}

func TestCandidateNumbersForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateNumbers(0, 0)

	// Assert
	for num := 1; num <= 9; num++ {
		if !slices.Contains(candidates, num) {
			t.Fatalf("every number should be a candidate for an empty sudoku, but %v was not", num)
		}
	}
	if len(candidates) != 9 {
		t.Fatalf("should be 9 candidate numbers for empty sudoku, but was %v", len(candidates))
	}
}

func TestCandidatesInRowForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	row := 0
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInRow(row, 4)

	// Assert
	for i := 0; i < 9; i++ {
		cell := lib.Coords{RowIndex: row, ColumnIndex: i}
		if !slices.Contains(candidates, cell) {
			t.Fatalf("every cell should be a candidate for an empty row, but %v was not", cell)
		}
	}
	if len(candidates) != 9 {
		t.Fatalf("should be 9 candidate cells for empty row, but was %v", len(candidates))
	}
}

func TestCandidatesInColForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	col := 0
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInColumn(col, 4)

	// Assert
	for i := 0; i < 9; i++ {
		cell := lib.Coords{RowIndex: i, ColumnIndex: col}
		if !slices.Contains(candidates, cell) {
			t.Fatalf("every cell should be a candidate for an empty column, but %v was not", cell)
		}
	}
	if len(candidates) != 9 {
		t.Fatalf("should be 9 candidate cells for empty column, but was %v", len(candidates))
	}
}

func TestCandidatesInSubForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	subRow, subCol := 2, 2
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInSubgrid(subRow, subCol, 4)

	// Assert
	for row := 3 * subRow; row < 3*(subRow+1); row++ {
		for col := 3 * subCol; col < 3*(subCol+1); col++ {
			cell := lib.Coords{RowIndex: row, ColumnIndex: col}
			if !slices.Contains(candidates, cell) {
				t.Fatalf("every cell should be a candidate for an empty subgrid,"+
					"but %v was not", cell)
			}
		}
	}
	if len(candidates) != 9 {
		t.Fatalf("should be 9 candidate cells for empty subgrid, but was %v", len(candidates))
	}
}
