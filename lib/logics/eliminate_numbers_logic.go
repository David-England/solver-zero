package logics

import "solver-zero/lib"

type EliminateNumbersLogic struct {
	sudoku *lib.Sudoku
}

func (logic *EliminateNumbersLogic) SetSudoku(sudoku *lib.Sudoku) {
	logic.sudoku = sudoku
}

func (logic *EliminateNumbersLogic) RunStep() bool {
	cantBe := [9][9][9]bool{}
	isChanged := false

	eliminateOptions(&cantBe, logic.sudoku)
	isChanged = setResolvedCells(&cantBe, logic.sudoku)

	return isChanged
}

func eliminateOptions(cantBe *[9][9][9]bool, sudoku *lib.Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.Grid[i][j] != 0 {
				banRowForNumber(cantBe, i, sudoku.Grid[i][j])
				banColumnForNumber(cantBe, j, sudoku.Grid[i][j])
				banSubgridForNumber(cantBe, i/3, j/3, sudoku.Grid[i][j])
			}
		}
	}
}

func setResolvedCells(cantBe *[9][9][9]bool, sudoku *lib.Sudoku) (isSuccessful bool) {
	isSuccessful = false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.Grid[i][j] == 0 {
				couldBe := make([]int, 9)

				for num := 1; num <= 9; num++ {
					if !cantBe[i][j][num-1] {
						couldBe = append(couldBe, num)
					}
				}

				if len(couldBe) == 1 {
					sudoku.Grid[i][j] = couldBe[0]
					isSuccessful = true
				}
			}
		}
	}

	return
}

func banRowForNumber(cantBe *[9][9][9]bool, row int, num int) {
	for col := 0; col < 9; col++ {
		cantBe[row][col][num-1] = true
	}
}

func banColumnForNumber(cantBe *[9][9][9]bool, col int, num int) {
	for row := 0; row < 9; row++ {
		cantBe[row][col][num-1] = true
	}
}

func banSubgridForNumber(cantBe *[9][9][9]bool, subgridRow int, subgridColumn int, num int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cantBe[3*subgridRow+i][3*subgridColumn+j][num-1] = true
		}
	}
}
