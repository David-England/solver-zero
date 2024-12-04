package eliminatecells

import (
	"solver-zero/lib"
	"solver-zero/lib/pencil"
)

type EliminateCellsLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *EliminateCellsLogic) RunStep() (bool, error) {
	isSuccessful := false
	pencilMarks := pencil.PencilMarks{}

	pencilMarks.EliminateOptions(logic.Sudoku)

	for i := 0; i < 9; i++ {
		for num := 1; num <= 9; num++ {
			rowCandidates := pencilMarks.CandidateCellsInRow(i, num)
			colCandidates := pencilMarks.CandidateCellsInColumn(i, num)
			subCandidates := pencilMarks.CandidateCellsInSubgrid(i/3, i%3, num)

			isSuccessful = setIfOneCandidate(rowCandidates, num, logic.Sudoku) || isSuccessful
			isSuccessful = setIfOneCandidate(colCandidates, num, logic.Sudoku) || isSuccessful
			isSuccessful = setIfOneCandidate(subCandidates, num, logic.Sudoku) || isSuccessful
		}
	}

	return isSuccessful, nil
}

func setIfOneCandidate(candidates []lib.Coords, num int, sud *lib.Sudoku) (isSuccessful bool) {
	if len(candidates) == 1 {
		cell := candidates[0]
		sud.Grid[cell.RowIndex][cell.ColumnIndex] = num
		isSuccessful = true
	}

	return
}
