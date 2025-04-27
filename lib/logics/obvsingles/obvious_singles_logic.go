package obvsingles

import (
	"fmt"
	"solver-zero/lib"
)

type ObviousSinglesLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *ObviousSinglesLogic) RunStep() (bool, error) {
	isSuccessful := false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if logic.Sudoku.Grid[i][j] == 0 {
				couldBe := logic.Sudoku.CandidateNumbers(i, j)

				if len(couldBe) == 1 {
					resolveCell(logic.Sudoku, i, j, couldBe[0])
					isSuccessful = true
				} else if len(couldBe) == 0 {
					return isSuccessful, fmt.Errorf("no number works for cell (%v, %v)", i+1, j+1)
				}
			}
		}
	}

	return isSuccessful, nil
}

func resolveCell(sudoku *lib.Sudoku, row, col, num int) {
	sudoku.Grid[row][col] = num
	sudoku.EliminateOptionsForCell(row, col, num)
}
