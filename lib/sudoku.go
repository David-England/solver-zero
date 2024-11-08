package lib

type Sudoku struct {
	Grid [9][9]int
}

type Sub struct {
	Grid [3][3]int
}

func (sud Sudoku) GetSubs() [3][3]Sub {
	subs := [3][3]Sub{}
	return subs
}
