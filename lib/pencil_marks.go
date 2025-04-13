package lib

type pencilMarks struct {
	// Dimensions: row, column, number - 1.
	cantBe [9][9][9]bool
}

func (pencilMarks *pencilMarks) candidateNumbers(row, col int) []int {
	couldBe := make([]int, 0, 9)

	for num := 1; num <= 9; num++ {
		if !pencilMarks.cantBe[row][col][num-1] {
			couldBe = append(couldBe, num)
		}
	}

	return couldBe
}

func (pencilMarks *pencilMarks) candidateCellsInRow(row, num int) []Coords {
	couldBe := make([]Coords, 0, 9)

	for col := 0; col < 9; col++ {
		if !pencilMarks.cantBe[row][col][num-1] {
			couldBe = append(couldBe, Coords{RowIndex: row, ColumnIndex: col})
		}
	}

	return couldBe
}

func (pencilMarks *pencilMarks) candidateCellsInColumn(col, num int) []Coords {
	couldBe := make([]Coords, 0, 9)

	for row := 0; row < 9; row++ {
		if !pencilMarks.cantBe[row][col][num-1] {
			couldBe = append(couldBe, Coords{RowIndex: row, ColumnIndex: col})
		}
	}

	return couldBe
}

func (pencilMarks *pencilMarks) candidateCellsInSubgrid(
	subgridRow, subgridCol, num int) []Coords {
	couldBe := make([]Coords, 0, 9)

	for i := 3 * subgridRow; i < 3*(subgridRow+1); i++ {
		for j := 3 * subgridCol; j < 3*(subgridCol+1); j++ {
			if !pencilMarks.cantBe[i][j][num-1] {
				couldBe = append(couldBe, Coords{RowIndex: i, ColumnIndex: j})
			}
		}
	}

	return couldBe
}

func (pencilMarks *pencilMarks) eliminateOptions(sudoku *Sudoku) {
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

func (pencilMarks *pencilMarks) banRowForNumber(row int, num int) {
	for col := 0; col < 9; col++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}

func (pencilMarks *pencilMarks) banColumnForNumber(col int, num int) {
	for row := 0; row < 9; row++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}

func (pencilMarks *pencilMarks) banSubgridForNumber(subgridRow int, subgridColumn int, num int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pencilMarks.cantBe[3*subgridRow+i][3*subgridColumn+j][num-1] = true
		}
	}
}

func (pencilMarks *pencilMarks) banAllNumbersForCell(row, col int) {
	for num := 1; num <= 9; num++ {
		pencilMarks.cantBe[row][col][num-1] = true
	}
}
