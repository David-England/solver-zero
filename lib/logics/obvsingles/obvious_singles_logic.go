package obvsingles

import (
	"fmt"
	"solver-zero/lib"
	"solver-zero/lib/pencil"
)

type ObviousSinglesLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *ObviousSinglesLogic) RunStep() (bool, error) {
	pencilMarks := pencil.PencilMarksMatrix{}

	pencilMarks.EliminateOptions(logic.Sudoku)
	isChanged, err := setResolvedCells(&pencilMarks, logic.Sudoku)

	return isChanged, err
}

func setResolvedCells(pencilMarks *pencil.PencilMarksMatrix, sudoku *lib.Sudoku) (bool, error) {
	isSuccessful := false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.Grid[i][j] == 0 {
				couldBe := make([]int, 0, 9)

				for num := 1; num <= 9; num++ {
					if !pencilMarks.CantBe[i][j][num-1] {
						couldBe = append(couldBe, num)
					}
				}

				if len(couldBe) == 1 {
					sudoku.Grid[i][j] = couldBe[0]
					isSuccessful = true
				} else if len(couldBe) == 0 {
					return isSuccessful, fmt.Errorf("no number works for cell (%v, %v)", i+1, j+1)
				}
			}
		}
	}

	return isSuccessful, nil
}
