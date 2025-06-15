package main

import (
	"fmt"
	"net/http"
	"solver-zero/lib"
	"solver-zero/lib/logics/eliminatecells"
	"solver-zero/lib/logics/obvpairs"
	"solver-zero/lib/logics/obvsingles"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/sudoku", solveSudoku)

	router.Run("localhost:8081")
}

func solveSudoku(c *gin.Context) {
	grid := [9][9]int{}

	if err := c.BindJSON(&grid); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validateGrid(grid); err != nil {
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
	sud := lib.CreateSudoku(sudokuGrid)
	logics := []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: sud},
		&obvpairs.ObviousPairsLogic{Sudoku: sud},
	}
	solutionSteps = make([][9][9]int, 0)

	for isChanged := true; isChanged; {
		isChanged, runError = lib.RunStep(logics)

		if runError != nil {
			return nil, runError
		} else if isChanged && hasNewEntry(sud.Grid, solutionSteps) {
			solutionSteps = append(solutionSteps, sud.Grid)
		}
	}

	return
}

func hasNewEntry(latestGrid [9][9]int, existingSteps [][9][9]int) bool {
	countSteps := len(existingSteps)

	if countSteps < 1 {
		return true
	} else if latestGrid == existingSteps[countSteps-1] {
		return false
	} else {
		return true
	}
}

func validateGrid(sudokuGrid [9][9]int) error {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := sudokuGrid[i][j]

			if val < 0 || val > 9 {
				return fmt.Errorf("invalid entry: %v", val)
			}
		}
	}

	return nil
}
