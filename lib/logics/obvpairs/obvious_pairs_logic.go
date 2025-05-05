package obvpairs

import "solver-zero/lib"

type ObviousPairsLogic struct {
	Sudoku     *lib.Sudoku
	pairsInRow []pairInRow
}

type pairInRow struct {
	nums  [2]int
	row   int
	cells [2]lib.Coords
}

func (logic *ObviousPairsLogic) RunStep() (bool, error) {
	logic.pairsInRow = make([]pairInRow, 0)

	for row := 0; row < 9; row++ {
		logic.resolveRow(row)
	}

	for _, pair := range logic.pairsInRow {
		cols := [2]int{pair.cells[0].ColumnIndex, pair.cells[1].ColumnIndex}
		banRowExcept(pair.nums, cols, pair.row, logic.Sudoku)
	}

	return false, nil
}

func (logic *ObviousPairsLogic) resolveRow(row int) {
	cellsWith2 := make(map[int][2]int)

	for col := 0; col < 9; col++ {
		if cand := logic.Sudoku.CandidateNumbers(row, col); len(cand) == 2 {
			candSortedArr := sort(([2]int)(cand))

			for k, v := range cellsWith2 {
				if candSortedArr == v {

					obviousPair := pairInRow{
						nums: candSortedArr,
						row:  row,
						cells: [2]lib.Coords{
							{RowIndex: row, ColumnIndex: k},
							{RowIndex: row, ColumnIndex: col},
						},
					}

					logic.pairsInRow = append(logic.pairsInRow, obviousPair)
					break
				}
			}

			cellsWith2[col] = candSortedArr
		}
	}
}

func banRowExcept(numsToBan, exceptCols [2]int, row int, sudoku *lib.Sudoku) {
	for col := 0; col < 9; col++ {
		if !(col == exceptCols[0] || col == exceptCols[1]) {
			for _, num := range numsToBan {
				sudoku.Ban(num, lib.Coords{RowIndex: row, ColumnIndex: col})
			}
		}
	}
}

func sort(x [2]int) [2]int {
	if x[0] > x[1] {
		return [2]int{x[1], x[0]}
	} else {
		// Ignore x[0] == x[1] case for now
		return x
	}
}
