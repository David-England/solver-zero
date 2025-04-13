package lib

type Sudoku struct {
	Grid        [9][9]int
	PencilMarks PencilMarks
}

type Sub struct {
	Grid [3][3]int
}

type Coords struct {
	RowIndex    int
	ColumnIndex int
}

func (suduko Sudoku) GetSubs() (subs [3][3]Sub) {
	for subRow := 0; subRow < 3; subRow++ {
		for subCol := 0; subCol < 3; subCol++ {
			for cellRow := 0; cellRow < 3; cellRow++ {
				for cellCol := 0; cellCol < 3; cellCol++ {
					subs[subRow][subCol].Grid[cellRow][cellCol] =
						suduko.Grid[3*subRow+cellRow][3*subCol+cellCol]
				}
			}
		}
	}

	return
}

func (sudoku *Sudoku) CandidateNumbers(row, col int) []int {
	return sudoku.PencilMarks.CandidateNumbers(row, col)
}

func (sudoku *Sudoku) CandidateCellsInRow(row, num int) []Coords {
	return sudoku.PencilMarks.CandidateCellsInRow(row, num)
}

func (sudoku *Sudoku) CandidateCellsInColumn(col, num int) []Coords {
	return sudoku.PencilMarks.CandidateCellsInColumn(col, num)
}

func (sudoku *Sudoku) CandidateCellsInSubgrid(subgridRow, subgridCol, num int) []Coords {
	return sudoku.PencilMarks.CandidateCellsInSubgrid(subgridRow, subgridCol, num)
}

func (sudoku *Sudoku) EliminateOptions() {
	sudoku.PencilMarks.EliminateOptions(sudoku)
}
