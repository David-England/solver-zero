package main

import (
	"errors"
	"net/http"
	"solver-zero/lib"
	"solver-zero/lib/logics/eliminatecells"
	"solver-zero/lib/logics/obvsingles"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sudoku", solveSudoku)

	router.Run("localhost:8081")
}

func solveSudoku(c *gin.Context) {
	grid := [9][9]int{}

	if err := c.BindJSON(&grid); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	solutionSteps, err := runSolver(lib.Sudoku{Grid: grid})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, solutionSteps)
	}
}

func runSolver(sud lib.Sudoku) ([]lib.Sudoku, error) {
	var logics []lib.ILogic = []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: &sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: &sud},
	}

	if isSuccessful, err := lib.RunStep(logics); err != nil {
		return nil, err
	} else if !isSuccessful {
		return nil, errors.New("sudoku too hard")
	} else {
		return append(make([]lib.Sudoku, 0), sud), nil
	}
}
