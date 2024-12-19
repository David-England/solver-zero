package main

import (
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

func runSolver(sud lib.Sudoku) (solutionSteps []lib.Sudoku, runError error) {
	var logics []lib.ILogic = []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: &sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: &sud},
	}
	solutionSteps = make([]lib.Sudoku, 0)

	for isChanged := true; isChanged; {
		isChanged, runError = lib.RunStep(logics)

		if runError != nil {
			return nil, runError
		} else if isChanged {
			solutionSteps = append(solutionSteps, sud)
		}
	}

	return
}
