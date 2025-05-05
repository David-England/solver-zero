package obvpairs_test

import (
	"solver-zero/lib"
	"solver-zero/lib/logics/obvpairs"
	"testing"
)

func TestObviousPairRow(t *testing.T) {
	// Arrange
	sudoku := lib.CreateSudoku(
		[9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{3, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 3},
			{4, 0, 0, 0, 0, 0, 0, 0, 9},
			{5, 0, 0, 0, 0, 0, 0, 0, 8},
			{6, 0, 0, 0, 0, 0, 0, 0, 7},
			{7, 0, 0, 0, 0, 0, 0, 0, 6},
			{8, 0, 0, 0, 0, 0, 0, 0, 5},
			{9, 0, 0, 0, 0, 0, 0, 0, 4},
		})
	logic := obvpairs.ObviousPairsLogic{Sudoku: sudoku}

	// Act
	logic.RunStep()

	// Assert
	if len(sudoku.CandidateCellsInRow(0, 1)) != 2 {
		t.Fatalf("failed to perform obvious pair cancellation in row")
	}
}

func TestObviousPairColumn(t *testing.T) {
	// Arrange
	sudoku := lib.CreateSudoku(
		[9][9]int{
			{0, 3, 0, 4, 5, 6, 7, 8, 9},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 3, 9, 8, 7, 6, 5, 4},
		})
	logic := obvpairs.ObviousPairsLogic{Sudoku: sudoku}

	// Act
	logic.RunStep()

	// Assert
	if len(sudoku.CandidateCellsInColumn(0, 1)) != 2 {
		t.Fatalf("failed to perform obvious pair cancellation in column")
	}
}

func TestObviousPairSubgrid(t *testing.T) {
	// Arrange
	sudoku := lib.CreateSudoku(
		[9][9]int{
			{4, 5, 6, 0, 0, 0, 7, 8, 9},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 3, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		})
	logic := obvpairs.ObviousPairsLogic{Sudoku: sudoku}

	// Act
	logic.RunStep()

	// Assert
	if len(sudoku.CandidateCellsInSubgrid(0, 1, 1)) != 2 {
		t.Fatalf("failed to perform obvious pair cancellation in subgrid")
	}
}
