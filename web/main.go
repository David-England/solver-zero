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
	sud := lib.Sudoku{}

	if err := c.BindJSON(&sud); err != nil {
		return
	}

	solutionSteps, _ := runSolver(sud)

	c.IndentedJSON(http.StatusOK, solutionSteps)
}

func runSolver(sud lib.Sudoku) ([]lib.Sudoku, error) {
	var logics []lib.ILogic = []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: &sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: &sud},
	}

	lib.RunStep(logics)
	solutionSteps := append(make([]lib.Sudoku, 0), sud)

	return solutionSteps, nil
}
