package main

import (
	"errors"
	"fmt"
	"os"
	"solver-zero/lib"
	"strconv"
	"strings"
)

func importSudoku(filepath string) (lib.Sudoku, error) {
	contents, err := os.ReadFile(filepath)

	if err != nil {
		return lib.Sudoku{}, err
	}

	return parseCSV(string(contents))
}

func parseCSV(csv string) (lib.Sudoku, error) {
	lines := strings.Split(csv, "\r\n")
	var sudoku lib.Sudoku

	if len(lines) < 9 {
		return lib.Sudoku{}, errors.New("sudoku provided has fewer than 9 rows")
	}

	for i := 0; i < 9; i++ {
		var err error

		sudoku.Grid[i], err = parseLine(lines[i])

		if err != nil {
			return lib.Sudoku{}, err
		}
	}

	return sudoku, nil
}

func parseLine(line string) ([9]int, error) {
	cells := strings.Split(line, ",")
	row := [9]int{}

	if len(cells) < 9 {
		return [9]int{}, errors.New("one of the lines has fewer than 9 cells")
	}

	for j := 0; j < 9; j++ {
		if len(cells[j]) > 0 {
			x, parseErr := strconv.Atoi(cells[j])

			if parseErr != nil {
				return [9]int{}, parseErr
			} else if x < 1 || x > 9 {
				return [9]int{}, fmt.Errorf("value not in [1,9]: %v", x)
			} else {
				row[j] = x
			}
		}
	}

	return row, nil
}
