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
	sud := lib.Sudoku{}

	if err := c.BindJSON(&sud); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if solutionSteps, err := runSolver(sud); err != nil {
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
