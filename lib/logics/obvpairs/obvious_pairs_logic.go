package obvpairs

import (
	"slices"
	"solver-zero/lib"
)

type ObviousPairsLogic struct {
	Sudoku         *lib.Sudoku
	pairsInRow     []pairInRow
	pairsInColumn  []pairInColumn
	pairsInSubgrid []pairInSubgrid
}

type pairInRow struct {
	nums  [2]int
	row   int
	cells [2]lib.Coords
}

type pairInColumn struct {
	nums   [2]int
	column int
	cells  [2]lib.Coords
}

type pairInSubgrid struct {
	nums          [2]int
	subgridRow    int
	subgridColumn int
	cells         [2]lib.Coords
}

func (logic *ObviousPairsLogic) RunStep() (bool, error) {
	isSuccessful := false
	logic.pairsInRow = make([]pairInRow, 0)
	logic.pairsInColumn = make([]pairInColumn, 0)
	logic.pairsInSubgrid = make([]pairInSubgrid, 0)

	logic.identifyPairs()
	logic.banOutsidePairs(&isSuccessful)

	return isSuccessful, nil
}

func (logic *ObviousPairsLogic) identifyPairs() {
	for row := 0; row < 9; row++ {
		logic.findPairsRow(row)
	}
	for col := 0; col < 9; col++ {
		logic.findPairsColumn(col)
	}
	for subg := 0; subg < 9; subg++ {
		logic.findPairsSubgrid(subg/3, subg%3)
	}
}

func (logic *ObviousPairsLogic) banOutsidePairs(hasBanned *bool) {
	for _, pair := range logic.pairsInRow {
		cols := [2]int{pair.cells[0].ColumnIndex, pair.cells[1].ColumnIndex}
		banRowExcept(pair.nums, cols, pair.row, logic.Sudoku, hasBanned)
	}
	for _, pair := range logic.pairsInColumn {
		rows := [2]int{pair.cells[0].RowIndex, pair.cells[1].RowIndex}
		banColumnExcept(pair.nums, rows, pair.column, logic.Sudoku, hasBanned)
	}
	for _, pair := range logic.pairsInSubgrid {
		banSubgridExcept(pair.nums, pair.cells, pair.subgridRow, pair.subgridColumn, logic.Sudoku,
			hasBanned)
	}
}

func (logic *ObviousPairsLogic) findPairsRow(row int) {
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

func (logic *ObviousPairsLogic) findPairsColumn(col int) {
	cellsWith2 := make(map[int][2]int)

	for row := 0; row < 9; row++ {
		if cand := logic.Sudoku.CandidateNumbers(row, col); len(cand) == 2 {
			candSortedArr := sort(([2]int)(cand))

			for k, v := range cellsWith2 {
				if candSortedArr == v {

					obviousPair := pairInColumn{
						nums:   candSortedArr,
						column: col,
						cells: [2]lib.Coords{
							{RowIndex: k, ColumnIndex: col},
							{RowIndex: row, ColumnIndex: col},
						},
					}

					logic.pairsInColumn = append(logic.pairsInColumn, obviousPair)
					break
				}
			}

			cellsWith2[row] = candSortedArr
		}
	}
}

func (logic *ObviousPairsLogic) findPairsSubgrid(subgridRow, subgridColumn int) {
	cellsWith2 := make(map[int][2]int)

	for i := 0; i < 9; i++ {
		row := 3*subgridRow + i/3
		col := 3*subgridColumn + i%3

		if cand := logic.Sudoku.CandidateNumbers(row, col); len(cand) == 2 {
			candSortedArr := sort(([2]int)(cand))

			for k, v := range cellsWith2 {
				if candSortedArr == v {

					obviousPair := pairInSubgrid{
						nums:          candSortedArr,
						subgridRow:    subgridRow,
						subgridColumn: subgridColumn,
						cells: [2]lib.Coords{
							{RowIndex: 3*subgridRow + k/3, ColumnIndex: 3*subgridColumn + k%3},
							{RowIndex: row, ColumnIndex: col},
						},
					}

					logic.pairsInSubgrid = append(logic.pairsInSubgrid, obviousPair)
					break
				}
			}

			cellsWith2[i] = candSortedArr
		}
	}
}

func banRowExcept(numsToBan, exceptCols [2]int, row int, sudoku *lib.Sudoku, hasBanned *bool) {
	for col := 0; col < 9; col++ {
		if !(col == exceptCols[0] || col == exceptCols[1]) {
			preExistCandidates := sudoku.CandidateNumbers(row, col)

			for _, num := range numsToBan {
				if slices.Contains(preExistCandidates, num) {
					sudoku.Ban(num, lib.Coords{RowIndex: row, ColumnIndex: col})
					*hasBanned = true
				}
			}
		}
	}
}

func banColumnExcept(numsToBan, exceptRows [2]int, col int, sudoku *lib.Sudoku, hasBanned *bool) {
	for row := 0; row < 9; row++ {
		if !(row == exceptRows[0] || row == exceptRows[1]) {
			preExistCandidates := sudoku.CandidateNumbers(row, col)

			for _, num := range numsToBan {
				if slices.Contains(preExistCandidates, num) {
					sudoku.Ban(num, lib.Coords{RowIndex: row, ColumnIndex: col})
					*hasBanned = true
				}
			}
		}
	}
}

func banSubgridExcept(numsToBan [2]int, exceptCells [2]lib.Coords, subgridRow, subgridColumn int,
	sudoku *lib.Sudoku, hasBanned *bool) {
	for row := 3 * subgridRow; row < 3*(subgridRow+1); row++ {
		for col := 3 * subgridColumn; col < 3*(subgridColumn+1); col++ {
			cell := lib.Coords{RowIndex: row, ColumnIndex: col}

			if !(cell == exceptCells[0] || cell == exceptCells[1]) {
				preExistCandidates := sudoku.CandidateNumbers(row, col)

				for _, num := range numsToBan {
					if slices.Contains(preExistCandidates, num) {
						sudoku.Ban(num, cell)
						*hasBanned = true
					}
				}
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
