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
				isAttemptSuccessful, err := attemptCell(logic.Sudoku, i, j)

				isSuccessful = isSuccessful || isAttemptSuccessful

				if err != nil {
					return isSuccessful, err
				}
			}
		}
	}

	logic.Sudoku.EliminateOptions()

	return isSuccessful, nil
}

func attemptCell(sudoku *lib.Sudoku, row, col int) (bool, error) {
	couldBe := sudoku.CandidateNumbers(row, col)

	if len(couldBe) == 1 {
		sudoku.Grid[row][col] = couldBe[0]
		return true, nil
	} else if len(couldBe) == 0 {
		return false, fmt.Errorf("no number works for cell (%v, %v)", row+1, col+1)
	} else {
		return false, nil
	}
}
