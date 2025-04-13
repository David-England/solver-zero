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
	contents, readErr := os.ReadFile(filepath)

	if readErr != nil {
		return lib.Sudoku{}, readErr
	}

	grid, parseErr := parseCSV(string(contents))

	if parseErr != nil {
		return lib.Sudoku{}, parseErr
	}

	return *lib.CreateSudoku(grid), nil
}

func parseCSV(csv string) ([9][9]int, error) {
	lines := strings.Split(csv, "\r\n")
	grid := [9][9]int{}

	if len(lines) < 9 {
		return [9][9]int{}, errors.New("sudoku provided has fewer than 9 rows")
	}

	for i := 0; i < 9; i++ {
		var err error

		grid[i], err = parseLine(lines[i])

		if err != nil {
			return [9][9]int{}, err
		}
	}

	return grid, nil
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
