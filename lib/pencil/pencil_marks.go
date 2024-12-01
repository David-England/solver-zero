package pencil

import "solver-zero/lib"

type PencilMarks struct {
	// Dimensions: row, column, number - 1.
	cantBe [9][9][9]bool
}

func (pencilMarks *PencilMarks) CandidateNumbers(row, col int) []int {
	couldBe := make([]int, 0, 9)

	for num := 1; num <= 9; num++ {
		if !pencilMarks.cantBe[row][col][num-1] {
			couldBe = append(couldBe, num)
		}
	}

	return couldBe
}

func (pencilMarks *PencilMarks) EliminateOptions(sudoku *lib.Sudoku) {
	pencilMarks.cantBe = [9][9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.Grid[i][j] != 0 {
				pencilMarks.banRowForNumber(i, sudoku.Grid[i][j])
				pencilMarks.banColumnForNumber(j, sudoku.Grid[i][j])
				pencilMarks.banSubgridForNumber(i/3, j/3, sudoku.Grid[i][j])
				pencilMarks.banAllNumbersForCell(i, j)
			}
		}
	}
}

func (pencilMarks *PencilMarks) banRowForNumber(row int, num int) {
	for col := 0; col < 9; col++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}

func (pencilMarks *PencilMarks) banColumnForNumber(col int, num int) {
	for row := 0; row < 9; row++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}

func (pencilMarks *PencilMarks) banSubgridForNumber(subgridRow int, subgridColumn int, num int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pencilMarks.cantBe[3*subgridRow+i][3*subgridColumn+j][num-1] = true
		}
	}
}

func (pencilMarks *PencilMarks) banAllNumbersForCell(row, col int) {
	for num := 1; num <= 9; num++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}
