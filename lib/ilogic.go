package lib

// Represents a particular logical heuristic for solving a sudoku.
type ILogic interface {
	// Runs one logical operation, returning true if any cells were resolved.
	RunStep() (bool, error)
}

// Runs sequentially over a slice of ILogics, returning true on the first that resolves a cell, or
// false if none are able.
func RunStep(logics []ILogic) (bool, error) {
	for _, logic := range logics {
		isChanged, err := logic.RunStep()

		if err != nil {
			return isChanged, err
		}

		if isChanged {
			return true, nil
		}
	}

	return false, nil
}
