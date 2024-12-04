package pencil

import (
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
