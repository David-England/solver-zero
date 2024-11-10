package lib

type ILogic interface {
	SetSudoku(*Sudoku)
	RunStep() bool
}

func RunStep(logics []ILogic) bool {
	for _, logic := range logics {
		isChanged := logic.RunStep()

		if isChanged {
			return true
		}
	}

	return false
}
