package main

import (
	"solver-zero/lib"
	"strings"
	"testing"
)

func TestParseLineOf4s(t *testing.T) {
	// Arrange
	line := "4,4,4,4,4,4,4,4,4"

	// Act
	actual, err := parseLine(line)

	// Assert
	if err != nil {
		t.Fatalf("parsing line of 4s errored: %v", err.Error())
	} else if actual != [9]int{4, 4, 4, 4, 4, 4, 4, 4, 4} {
		t.Fatalf("should be all 4s, but is %v", actual)
	}
}

func TestParseEmptyLine(t *testing.T) {
	// Arrange
	line := ",,,,,,,,"

	// Act
	actual, err := parseLine(line)

	// Assert
	if err != nil {
		t.Fatalf("parsing empty line errored: %v", err.Error())
	} else if actual != [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0} {
		t.Fatalf("should be all 0s, but is %v", actual)
	}
}

func TestParseTooShortLine(t *testing.T) {
	// Arrange
	line := "1,2,3,4,5,6,7,8"

	// Act
	_, err := parseLine(line)

	// Assert
	if err == nil {
		t.Fatalf("parsing too short line didn't error but should have")
	}
}

func TestParseTooBigInt(t *testing.T) {
	// Arrange
	line := ",,,,10,,,,"

	// Act
	_, err := parseLine(line)

	// Assert
	if err == nil {
		t.Fatalf("parsing too big int didn't error but should have")
	}
}

func TestParseBadInt(t *testing.T) {
	// Arrange
	line := ",,,,A,,,,"

	// Act
	_, err := parseLine(line)

	// Assert
	if err == nil {
		t.Fatalf("parsing non-integer value didn't error but should have")
	}
}

func TestCSV(t *testing.T) {
	// Arrange
	csv := createGoodCSV()
	expected := createGoodSudoku()

	// Act
	actual, err := parseCSV(csv)

	// Assert
	if err != nil {
		t.Fatalf("parsing CSV failed: %v", err.Error())
	} else if actual != expected {
		t.Fatalf("should be %v, but was %v", expected, actual)
	}
}

func TestParseTooFewLines(t *testing.T) {
	// Arrange
	csv := "1,,,,,,,,\r\n,2,,,,,,,\r\n,,3,,,,,,\r\n" +
		",,,4,,,,,\r\n,,,,5,,,,\r\n,,,,,6,,,\r\n" +
		",,,,,,7,,\r\n,,,,,,,8,"

	// Act
	_, err := parseCSV(csv)

	// Assert
	if err == nil {
		t.Fatalf("parsing CSV with eight lines didn't error but should have")
	}
}

func TestParseCSVWithBadLine(t *testing.T) {
	// Arrange
	csv := "A,B,C,D,E,F,G,H,I\r\n,2,,,,,,,\r\n,,3,,,,,,\r\n" +
		",,,4,,,,,\r\n,,,,5,,,,\r\n,,,,,6,,,\r\n" +
		",,,,,,7,,\r\n,,,,,,,8,\r\n,,,,,,,,9"

	// Act
	_, err := parseCSV(csv)

	// Assert
	if err == nil {
		t.Fatalf("parsing CSV containing non-integers didn't error but should have")
	}
}

func TestParseNewlineEOL(t *testing.T) {
	// Arrange
	csv := strings.ReplaceAll(createGoodCSV(), "\r\n", "\n")

	// Act
	_, err := parseCSV(csv)

	// Assert
	if err == nil {
		t.Fatalf("parsing CSV with newline-only EOL sequence didn't error but should have")
	}
}

func TestParseNewlineReturnEOL(t *testing.T) {
	// Arrange
	csv := strings.ReplaceAll(createGoodCSV(), "\r\n", "\n\r")

	// Act
	_, err := parseCSV(csv)

	// Assert
	if err == nil {
		t.Fatalf("parsing CSV with newline-return EOL sequence didn't error but should have")
	}
}

func createGoodCSV() string {
	return "1,,,,,,,,\r\n,2,,,,,,,\r\n,,3,,,,,,\r\n" +
		",,,4,,,,,\r\n,,,,5,,,,\r\n,,,,,6,,,\r\n" +
		",,,,,,7,,\r\n,,,,,,,8,\r\n,,,,,,,,9"
}

func createGoodSudoku() lib.Sudoku {
	return lib.Sudoku{
		Grid: [9][9]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 3, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 4, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 5, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 6, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 7, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 8, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 9},
		},
	}
}
