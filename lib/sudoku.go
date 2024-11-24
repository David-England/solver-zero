package lib

type Sudoku struct {
	Grid [9][9]int
}

// Represents a 3x3 subgrid of a sudoku.
type Sub struct {
	Grid [3][3]int
}

// Returns a 3x3 array containing the 3x3 subgrids of a sudoku.
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
