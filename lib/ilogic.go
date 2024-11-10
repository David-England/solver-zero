package lib

type ILogic interface {
	RunStep(*Sudoku) bool
}

func RunStep(logics []ILogic, inputSudoku *Sudoku) bool {
	for _, logic := range logics {
		isChanged := logic.RunStep(inputSudoku)

		if isChanged {
			return true
		}
	}

	return false
}
