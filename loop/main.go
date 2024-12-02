package main

import (
	"fmt"
	"log"
	"os"
	"solver-zero/lib"
	"solver-zero/lib/logics/eliminatecells"
	"solver-zero/lib/logics/obvsingles"
)

func main() {
	sud, importErr := importSudoku(os.Args[1])

	if importErr != nil {
		log.Fatal(importErr)
	}

	var logics []lib.ILogic = []lib.ILogic{
		&obvsingles.ObviousSinglesLogic{Sudoku: &sud},
		&eliminatecells.EliminateCellsLogic{Sudoku: &sud},
	}

	printSudoku(sud)
	if runErr := run(logics); runErr != nil {
		log.Fatal(runErr)
	}
	printSudoku(sud)
}

func run(logics []lib.ILogic) (runError error) {
	for isChanged := true; isChanged; {
		isChanged, runError = lib.RunStep(logics)

		if runError != nil {
			return
		}
	}

	return
}

func printSudoku(sudoku lib.Sudoku) {
	fmt.Println("-------------------------------")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			printRow(sudoku.Grid[3*i+j])
			fmt.Println()
		}

		fmt.Println("-------------------------------")
	}
}

func printRow(row [9]int) {
	fmt.Print("|")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if val := row[3*i+j]; val != 0 {
				fmt.Printf(" %v ", val)
			} else {
				fmt.Print("   ")
			}
		}

		fmt.Print("|")
	}
}
