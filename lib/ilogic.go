package lib

type ILogic interface {
	RunStep(Sudoku) (Sudoku, bool)
}

func RunStep(logics []ILogic, inputSudoku Sudoku) (Sudoku, bool) {
	for _, logic := range logics {
		sudoku, isChanged := logic.RunStep(inputSudoku)

		if isChanged {
			return sudoku, true
		} else {
			inputSudoku = sudoku
		}
	}

	return inputSudoku, false
}
