package main

import (
	"fmt"
	"os"
	"solver-zero/lib"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello universe.")
	fmt.Println(importSudoku(os.Args[1]).Grid)
}

func importSudoku(filepath string) lib.Sudoku {
	out, _ := os.ReadFile(filepath)
	return parseCSV(string(out))
}

func parseCSV(csv string) (suduko lib.Sudoku) {
	lines := strings.Split(csv, "\n")

	for i := 0; i < 9; i++ {
		cells := strings.Split(lines[i], ",")

		for j := 0; j < 9; j++ {
			if len(cells[j]) > 0 {
				suduko.Grid[i][j], _ = strconv.Atoi(cells[j])
			}
		}
	}

	return suduko
}
