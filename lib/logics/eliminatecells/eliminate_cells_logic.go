package eliminatecells

import (
	"solver-zero/lib"
)

type EliminateCellsLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *EliminateCellsLogic) RunStep() (bool, error) {
	isSuccessful := false

	for i := 0; i < 9; i++ {
		for num := 1; num <= 9; num++ {
			rowCandidates := logic.Sudoku.CandidateCellsInRow(i, num)
			colCandidates := logic.Sudoku.CandidateCellsInColumn(i, num)
			subCandidates := logic.Sudoku.CandidateCellsInSubgrid(i/3, i%3, num)

			setIfOneCandidate(rowCandidates, num, logic.Sudoku, &isSuccessful)
			setIfOneCandidate(colCandidates, num, logic.Sudoku, &isSuccessful)
			setIfOneCandidate(subCandidates, num, logic.Sudoku, &isSuccessful)
		}
	}

	logic.Sudoku.EliminateOptions()

	return isSuccessful, nil
}

func setIfOneCandidate(candidates []lib.Coords, num int, sud *lib.Sudoku, isSuccessful *bool) {
	if len(candidates) == 1 {
		cell := candidates[0]
		sud.Grid[cell.RowIndex][cell.ColumnIndex] = num
		*isSuccessful = true
	}
}
