package main

import (
	"fmt"
	"log"
	"os"
	"solver-zero/lib"
	"solver-zero/lib/logics/eliminatenumbers"
)

func main() {
	sud, importErr := importSudoku(os.Args[1])

	if importErr != nil {
		log.Fatal(importErr)
	}

	var logics []lib.ILogic = []lib.ILogic{
		&eliminatenumbers.EliminateNumbersLogic{Sudoku: &sud},
	}

	fmt.Println(sud.Grid)
	if runErr := run(logics); runErr != nil {
		log.Fatal(runErr)
	}
	fmt.Println(sud.Grid)
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
