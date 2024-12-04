package main

import (
	"fmt"
	"slices"
	"strings"
)

func ReverseWord(word string) string {
	// Reverse the word
	characters := strings.Split(word, "")
	slices.Reverse(characters)
	return strings.Join(characters, "")
}

func BuildDiagonalStrings(rows []string, cols []string) ([]string, []string) {
	// Diagonal strings building.
	// To calculate the number of diagonals in 1 direction
	rowLength := len(rows[0])
	colLength := len(cols[0])
	numDiagonalsPerDirection := rowLength + colLength - 1
	// Horizontal >= Vertical
	upperDiagonalStringsBuilder := make([]strings.Builder, numDiagonalsPerDirection)
	for i := 0; i < numDiagonalsPerDirection; i++ {
		upperDiagonalStringsBuilder[i] = strings.Builder{}
	}
	lowerDiagonalStringsBuilder := make([]strings.Builder, numDiagonalsPerDirection)
	for i := 0; i < numDiagonalsPerDirection; i++ {
		lowerDiagonalStringsBuilder[i] = strings.Builder{}
	}
	// Check to see whether the vertical or horizontal is longer.
	// Upper-left to Lower-right
	for colInd := range cols {
		for rowInd := 0; rowInd < len(rows); rowInd++ {
			row := rows[rowInd]

			// Upper-left starting diagonals
			diagInd := rowInd + colInd
			upperDiagonalStringsBuilder[diagInd].WriteByte(row[colInd])

			// Upper-right starting diagonals
			diagInd = rowInd - colInd + colLength - 1
			lowerDiagonalStringsBuilder[diagInd].WriteByte(row[colInd])
		}
	}

	// Build strings from the builder.
	upperDiagonalStrings := make([]string, numDiagonalsPerDirection)
	for i := 0; i < numDiagonalsPerDirection; i++ {
		upperDiagonalStrings[i] = upperDiagonalStringsBuilder[i].String()
	}

	lowerDiagonalStrings := make([]string, numDiagonalsPerDirection)
	for i := 0; i < numDiagonalsPerDirection; i++ {
		lowerDiagonalStrings[i] = lowerDiagonalStringsBuilder[i].String()
	}

	return upperDiagonalStrings, lowerDiagonalStrings
}

func SearchAll(inputLines []string, searchWord string) int {
	// Store count
	count := 0

	// Create the vertical string constructors.
	// Check assumption here that all columns have the same length.
	lineLength := len(inputLines[0])
	for _, line := range inputLines[1:] {
		if len(line) != lineLength {
			panic("line lengths do not match.")
		}
	}
	verticalStringsBuilder := make([]strings.Builder, len(inputLines[0]))
	for i := 0; i < len(inputLines); i++ {
		verticalStringsBuilder[i] = strings.Builder{}
	}

	for _, line := range inputLines {
		// Horizontal Search
		count += SearchForWord(line, searchWord)
		count += SearchForWord(line, ReverseWord(searchWord))

		for rowInd, char := range line {
			// Construct vertical strings
			_, err := verticalStringsBuilder[rowInd].WriteRune(char)
			if err != nil {
				panic("Error during building vertical string")
			}
		}
	}

	// Vertical search
	// Build strings from the builder.
	verticalStrings := make([]string, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		verticalStrings[i] = verticalStringsBuilder[i].String()
	}

	// Count vertical words
	for _, verticalString := range verticalStrings {
		count += SearchForWord(verticalString, searchWord)
		count += SearchForWord(verticalString, ReverseWord(searchWord))
	}

	// Count diagonal words
	upperDiagonals, lowerDiagonals := BuildDiagonalStrings(verticalStrings, inputLines)
	for _, diagonalString := range upperDiagonals {
		count += SearchForWord(diagonalString, searchWord)
		count += SearchForWord(diagonalString, ReverseWord(searchWord))
	}
	for _, diagonalString := range lowerDiagonals {
		count += SearchForWord(diagonalString, searchWord)
		count += SearchForWord(diagonalString, ReverseWord(searchWord))
	}

	return count
}

func SearchForWord(inputString string, searchWord string) int {
	// fmt.Println("Found", strings.Count(inputString, searchWord), "in", inputString)
	return strings.Count(inputString, searchWord)
}

func CountOccurences(inputLines []string, searchWord string) int {
	// Store count.
	count := 0

	// Operate on the forward and backwards word.
	count += SearchAll(inputLines, searchWord)
	// fmt.Println("Final count", count)

	return count
}

func testWordSearch() {
	testDiagString := `XMAS
XMAS
XMAS
XMAS`

	testLines, err := StringToLines(testDiagString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountOccurences(testLines, "XMAS") != 6 {
		panic("Failed test case.")
	}

	testString := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	testLines, err = StringToLines(testString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountOccurences(testLines, "XMAS") != 18 {
		panic("Failed test case.")
	}

}

func MainPart1() {
	testWordSearch()

	parsedLines := GetInputLines()
	fmt.Println(len(parsedLines), "lines from input.")
	answer := CountOccurences(parsedLines, "XMAS")
	fmt.Printf("Answer Part 1: %d\n", answer)
}
