package eliminatecells

import "solver-zero/lib"

type EliminateCellsLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *EliminateCellsLogic) RunStep() (bool, error) {
	return false, nil
}
