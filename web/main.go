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

	if solutionSteps, err := runSolver(grid); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, solutionSteps)
	}
}

func runSolver(sudokuGrid [9][9]int) (solutionSteps [][9][9]int, runError error) {
	sud := lib.Sudoku{Grid: sudokuGrid}
	logics := []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: &sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: &sud},
	}
	solutionSteps = make([][9][9]int, 0)

	for isChanged := true; isChanged; {
		isChanged, runError = lib.RunStep(logics)

		if runError != nil {
			return nil, runError
		} else if isChanged {
			solutionSteps = append(solutionSteps, sud.Grid)
		}
	}

	return
}
