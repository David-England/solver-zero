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
		isSuccessful = resolveRow(i, &pencilMarks, logic.Sudoku) || isSuccessful
		isSuccessful = resolveCol(i, &pencilMarks, logic.Sudoku) || isSuccessful
		isSuccessful = resolveSub(i/3, i%3, &pencilMarks, logic.Sudoku) || isSuccessful
	}

	return isSuccessful, nil
}

func resolveRow(row int, pencilMarks *pencil.PencilMarks, sud *lib.Sudoku) (isSuccessful bool) {
	for num := 1; num <= 9; num++ {
		candidateCells := pencilMarks.CandidateCellsInRow(row, num)

		if len(candidateCells) == 1 {
			cell := candidateCells[0]
			sud.Grid[cell.RowIndex][cell.ColumnIndex] = num
			isSuccessful = true
		}
	}

	return
}

func resolveCol(col int, pencilMarks *pencil.PencilMarks, sud *lib.Sudoku) (isSuccessful bool) {
	for num := 1; num <= 9; num++ {
		candidateCells := pencilMarks.CandidateCellsInColumn(col, num)

		if len(candidateCells) == 1 {
			cell := candidateCells[0]
			sud.Grid[cell.RowIndex][cell.ColumnIndex] = num
			isSuccessful = true
		}
	}

	return
}

func resolveSub(subgridRow, subgridCol int, pencilMarks *pencil.PencilMarks,
	sud *lib.Sudoku) (isSuccessful bool) {
	for num := 1; num <= 9; num++ {
		candidateCells := pencilMarks.CandidateCellsInSubgrid(subgridRow, subgridCol, num)

		if len(candidateCells) == 1 {
			cell := candidateCells[0]
			sud.Grid[cell.RowIndex][cell.ColumnIndex] = num
			isSuccessful = true
		}
	}

	return
}
