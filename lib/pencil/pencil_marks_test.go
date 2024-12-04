package pencil

import (
	"fmt"
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
	sudoku := getEmptySudoku()

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

func TestCandidateNumbers(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateNumbers(3, 3)

	// Assert
	assertContains(candidates, 4, "4 should be candidate, but was not", t)
	assertLength(candidates, 1,
		"should be exactly 1 candidate number in row with 8 solved cells", t)
}

func TestCandidatesInRow(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	row := 3
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInRow(row, 4)

	// Assert
	cell := lib.Coords{RowIndex: row, ColumnIndex: 3}
	assertContains(candidates, cell, fmt.Sprintf("%v should be candidate, but was not", cell), t)
	assertLength(candidates, 1, "should be exactly 1 candidate cell in row with 8 solved cells", t)
}

func TestCandidatesInColumn(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	col := 3
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
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInColumn(col, 4)

	// Assert
	cell := lib.Coords{RowIndex: 3, ColumnIndex: col}
	assertContains(candidates, cell, fmt.Sprintf("%v should be candidate, but was not", cell), t)
	assertLength(candidates, 1,
		"should be exactly 1 candidate cell in column with 8 solved cells", t)
}

func TestCandidatesInSubgrid(t *testing.T) {
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
			{0, 0, 0, 0, 0, 0, 1, 2, 3},
			{0, 0, 0, 0, 0, 0, 0, 5, 6},
			{0, 0, 0, 0, 0, 0, 7, 8, 9},
		},
	}
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInSubgrid(subRow, subCol, 4)

	// Assert
	cell := lib.Coords{RowIndex: 7, ColumnIndex: 6}
	assertContains(candidates, cell, fmt.Sprintf("%v should be candidate, but was not", cell), t)
	assertLength(candidates, 1,
		"should be exactly 1 candidate cell in subgrid with 8 solved cells", t)
}

func TestCandidateNumbersForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	sudoku := getEmptySudoku()
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateNumbers(0, 0)

	// Assert
	for num := 1; num <= 9; num++ {
		assertContains(candidates, num, fmt.Sprintf(
			"every number should be a candidate for an empty sudoku, but %v was not", num), t)
	}
	assertLength(candidates, 9, fmt.Sprintf(
		"should be 9 candidate numbers for empty sudoku, but was %v", len(candidates)), t)
}

func TestCandidatesInRowForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	row := 0
	sudoku := getEmptySudoku()
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInRow(row, 4)

	// Assert
	for i := 0; i < 9; i++ {
		cell := lib.Coords{RowIndex: row, ColumnIndex: i}
		assertContains(candidates, cell, fmt.Sprintf(
			"every cell should be a candidate for an empty row, but %v was not", cell), t)
	}
	assertLength(candidates, 9, fmt.Sprintf(
		"should be 9 candidate cells for empty row, but was %v", len(candidates)), t)
}

func TestCandidatesInColForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	col := 0
	sudoku := getEmptySudoku()
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInColumn(col, 4)

	// Assert
	for i := 0; i < 9; i++ {
		cell := lib.Coords{RowIndex: i, ColumnIndex: col}
		assertContains(candidates, cell, fmt.Sprintf(
			"every cell should be a candidate for an empty column, but %v was not", cell), t)
	}
	assertLength(candidates, 9, fmt.Sprintf(
		"should be 9 candidate cells for empty column, but was %v", len(candidates)), t)
}

func TestCandidatesInSubForEmpty(t *testing.T) {
	// Arrange
	sut := PencilMarks{}
	subRow, subCol := 2, 2
	sudoku := getEmptySudoku()
	sut.EliminateOptions(&sudoku)

	// Act
	candidates := sut.CandidateCellsInSubgrid(subRow, subCol, 4)

	// Assert
	for row := 3 * subRow; row < 3*(subRow+1); row++ {
		for col := 3 * subCol; col < 3*(subCol+1); col++ {
			cell := lib.Coords{RowIndex: row, ColumnIndex: col}
			assertContains(candidates, cell, fmt.Sprintf(
				"every cell should be a candidate for an empty subgrid, but %v was not", cell), t)
		}
	}
	assertLength(candidates, 9, fmt.Sprintf(
		"should be 9 candidate cells for empty subgrid, but was %v", len(candidates)), t)
}

func getEmptySudoku() lib.Sudoku {
	return lib.Sudoku{
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
}

func assertContains[T comparable](container []T, item T, failMessage string, t *testing.T) {
	if !slices.Contains(container, item) {
		t.Fatal(failMessage)
	}
}

func assertLength[T comparable](container []T, expected int, failMessage string, t *testing.T) {
	if len(container) != expected {
		t.Fatal(failMessage)
	}
}
