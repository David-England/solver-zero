package lib

type ILogic interface {
	RunStep() (bool, error)
}

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
