package main

import (
	"fmt"
	"strings"
)

func buildVerticalStrings(lines []string) []string {
	verticalStringsBuilder := make([]strings.Builder, len(lines[0]))
	for i := 0; i < len(lines); i++ {
		verticalStringsBuilder[i] = strings.Builder{}
	}

	for _, line := range lines {
		for rowInd, char := range line {
			// Construct vertical strings
			_, err := verticalStringsBuilder[rowInd].WriteRune(char)
			if err != nil {
				panic("Error during building vertical string")
			}
		}
	}

	// Build strings from the builder.
	verticalStrings := make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		verticalStrings[i] = verticalStringsBuilder[i].String()
	}

	return verticalStrings
}

// 3x3 Search heuristic
func heuristic(searchSlice [][]string, xWordParts []string, revXWordParts []string) bool {
	// Check if the characters match our search condition.

	// Upper-left to lower-right check
	startRow := searchSlice[0]
	if startRow[0] == xWordParts[0] {
		// fmt.Println("Enter condition 0")
		for partInd, char := range xWordParts {
			subSlice := searchSlice[partInd]
			if subSlice[partInd] != char {
				return false
			}
		}
		// fmt.Println("Clear cond 0")
	} else if startRow[0] == revXWordParts[0] {
		// fmt.Println("Enter condition 1")
		for partInd, char := range revXWordParts {
			subSlice := searchSlice[partInd]
			if subSlice[partInd] != char {
				return false
			}
		}
		// fmt.Println("Clear cond 1")
	} else {
		// Return false if not a match.
		// Technically redundant since the trigger for this function is a match.
		return false
	}

	// Lower-left to upper-right check
	startRow = searchSlice[len(searchSlice)-1]
	if startRow[0] == xWordParts[0] {
		// fmt.Println("Enter condition 3")
		for partInd, char := range xWordParts {
			subSlice := searchSlice[len(searchSlice)-1-partInd]
			if subSlice[partInd] != char {
				return false
			}
		}
		// fmt.Println("Pass condition 3")
	} else if startRow[0] == revXWordParts[0] {
		// fmt.Println("Enter condition 4")
		for partInd, char := range revXWordParts {
			subSlice := searchSlice[len(searchSlice)-1-partInd]
			if subSlice[partInd] != char {
				return false
			}
		}
		// fmt.Println("Pass condition 4")
	} else {
		// Return false if not a match.
		return false
	}
	return true
}

func SearchX(matrix [][]string, xWord string) int {
	count := 0

	revXWord := ReverseWord(xWord)

	xWordParts := strings.Split(xWord, "")
	revXWordParts := strings.Split(revXWord, "")

	rows := len(matrix)
	cols := len(matrix[0])

	searchDistance := len(xWordParts)
	// Find the occurences of the pattern.
	for i := 0; i <= rows-searchDistance; i++ {
		row := matrix[i]
		for j := 0; j <= cols-searchDistance; j++ {
			char := string(row[j])
			// Character match
			if char == xWordParts[0] || char == revXWordParts[0] {
				// Make a copy here to not affect the underlying slice.
				searchSlice := append(matrix[:0:0], matrix[i:i+searchDistance]...)
				for x := range searchSlice {
					subSlice := searchSlice[x]
					searchSlice[x] = subSlice[j : j+searchDistance]
				}

				// Perform heuristic test.
				if heuristic(searchSlice, xWordParts, revXWordParts) {
					count += 1
				}
			}
		}
	}

	return count
}

func linesToMatrix(lines []string) [][]string {
	rows := len(lines)
	cols := len(lines[0])
	matrix := make([][]string, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]string, cols)
		stringParts := strings.Split(lines[i], "")
		for j := 0; j < cols; j++ {
			matrix[i][j] = stringParts[j]
		}
	}

	return matrix
}

func CountXOccurences(lines []string, xWord string) int {
	matrix := linesToMatrix(lines)
	count := SearchX(matrix, xWord)
	// fmt.Println("Final count", count)

	return count
}

func testXWordSearch() {
	testString := `00000
11111
22222
33333
44444`
	testLines, err := StringToLines(testString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "123") != 3 {
		panic("Failed test case.")
	}

	testDiagString := `XMAS
XMAS
XMAS
XMAS`

	testLines, err = StringToLines(testDiagString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "MAS") != 2 {
		panic("Failed test case.")
	}

	testString = `0000000000
1111111111
2222222222
3333333333
4444444444
5555555555`
	testLines, err = StringToLines(testString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "MAS") != 0 {
		panic("Failed test case.")
	}

	testString = `0M0S000000
11A11MSMS1
2M2S2MAA22
33A3ASMSM3
4M4S4M4444
5555555555
S6S6S6S6S6
7A7A7A7A77
M8M8M8M8M8
9999999999`
	testLines, err = StringToLines(testString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "MAS") != 9 {
		panic("Failed test case.")
	}

	testString = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
	testLines, err = StringToLines(testString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "MAS") != 9 {
		panic("Failed test case.")
	}

	testString = `MMMSXXMASM
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
	if CountXOccurences(testLines, "MAS") != 9 {
		panic("Failed test case.")
	}
}

func MainPart2() {
	testXWordSearch()

	parsedLines := GetInputLines()
	answer := CountXOccurences(parsedLines, "MAS")
	fmt.Printf("Answer Part 2: %d\n", answer)
}
