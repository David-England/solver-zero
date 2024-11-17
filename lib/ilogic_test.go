package lib_test

import (
	"errors"
	"solver-zero/lib"
	"testing"
)

type stubLogic struct {
	returns bool
	errors  bool
}

func (stub stubLogic) RunStep() (bool, error) {
	if stub.errors {
		return stub.returns, errors.New("example error")
	} else {
		return stub.returns, nil
	}
}

func TestRunStepChanged(t *testing.T) {
	// Arrange
	returns := true
	errors := false
	logics := []lib.ILogic{stubLogic{returns, errors}}

	// Act
	isChanged, _ := lib.RunStep(logics)

	// Assert
	if !isChanged {
		t.Fatalf("returned %v but should be %v", isChanged, returns)
	}
}

func TestRunStepUnchanged(t *testing.T) {
	// Arrange
	returns := false
	errors := false
	logics := []lib.ILogic{stubLogic{returns, errors}}

	// Act
	isChanged, _ := lib.RunStep(logics)

	// Assert
	if isChanged {
		t.Fatalf("returned %v but should be %v", isChanged, returns)
	}
}

func TestRunStepErrors(t *testing.T) {
	// Arrange
	logics := []lib.ILogic{stubLogic{returns: false, errors: true}}

	// Act
	_, err := lib.RunStep(logics)

	// Assert
	if err == nil {
		t.Fatalf("didn't error but should")
	}
}

func TestFirstLogicReturns(t *testing.T) {
	// Arrange
	logics := []lib.ILogic{
		stubLogic{returns: true, errors: false},
		stubLogic{returns: false, errors: false},
	}

	// Act
	isChanged, _ := lib.RunStep(logics)

	// Assert
	if !isChanged {
		t.Fatalf("should return as changed")
	}
}

func TestSecondLogicErrors(t *testing.T) {
	// Arrange
	logics := []lib.ILogic{
		stubLogic{returns: false, errors: false},
		stubLogic{returns: false, errors: true},
	}

	// Act
	_, err := lib.RunStep(logics)

	// Assert
	if err == nil {
		t.Fatalf("didn't error but should")
	}
}
