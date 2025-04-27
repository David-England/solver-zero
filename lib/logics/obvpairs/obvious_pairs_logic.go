package obvpairs

import "solver-zero/lib"

type ObviousPairsLogic struct {
	Sudoku *lib.Sudoku
}

func (logic *ObviousPairsLogic) RunStep() (bool, error) {
	return false, nil
}
