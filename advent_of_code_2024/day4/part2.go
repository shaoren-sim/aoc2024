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
	verticalStrings := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		verticalStrings[i] = verticalStringsBuilder[i].String()
	}

	return verticalStrings
}

func findIndices(inputString string, substring string) []int {
	count := strings.Count(inputString, substring)
	if count == 0 {
		return []int{}
	}
	inds := make([]int, count)

	i := 0
	n := -1
	for {
		n = strings.Index(inputString, substring)
		if n == -1 {
			break
		}
		if i == 0 {
			inds[i] = n
		} else {
			inds[i] = n + inds[i-1] + 1
		}
		inputString = inputString[n+1:]
		i += 1
	}
	return inds
}

func findIndicesForwardReverse(inputString string, substring string) []int {
	revSubstring := ReverseWord(substring)
	forwardInds := findIndices(inputString, substring)
	reverseInds := findIndices(inputString, revSubstring)
	return append(forwardInds, reverseInds...)
}

func SearchX(upperDiagonal []string, lowerDiagonal []string, xWord string) int {
	// Ensure that diagonals are the same length
	if len(upperDiagonal) != len(lowerDiagonal) {
		panic("Mismatched diagonal arrays.")
	}

	count := 0

	for i := 0; i < len(upperDiagonal); i++ {
		for j := i; j < len(lowerDiagonal); j++ {
			upperDiagInds := findIndicesForwardReverse(upperDiagonal[i], xWord)
			lowerDiagInds := findIndicesForwardReverse(lowerDiagonal[j], xWord)

			// If either of the diagonals has no matches (len=0), break.
			if len(upperDiagInds) == 0 {
				continue
			}
			if len(lowerDiagInds) == 0 {
				continue
			}
			for _, upperDiagInd := range upperDiagInds {
				for _, lowerDiagInd := range lowerDiagInds {
					if upperDiagInd == lowerDiagInd {
						fmt.Println(upperDiagonal[i], lowerDiagonal[j])
						fmt.Println(upperDiagInds, lowerDiagInds)
						count += 1
					}
				}
			}

		}

		// // For the lower diagonal, count both forward and backwards occurences.
		// lowerDiagInds := findIndices(lowerDiagonal[i], xWord)
		// lowerDiagRevInds := findIndices(lowerDiagonal[i], revXWord)
		// if lowerDiagRevCount > lowerDiagCount {
		// 	lowerDiagCount = lowerDiagRevCount
		// }
		//
		// if SearchForWord(upperDiagonal[i], xWord) == 1 || SearchForWord(upperDiagonal[i], revXWord) == 1 {
		// 	if SearchForWord(lowerDiagonal[i], xWord) == 1 || SearchForWord(lowerDiagonal[i], revXWord) == 1 {
		// 		fmt.Println("Match found")
		// 		count += 1
		// 	}
		// }
	}

	return count
}

func CountXOccurences(lines []string, xWord string) int {
	verticalStrings := buildVerticalStrings(lines)
	upperDiagonal, lowerDiagonal := BuildDiagonalStrings(verticalStrings, lines)
	// fmt.Println(upperDiagonal)
	// fmt.Println(lowerDiagonal)

	count := SearchX(upperDiagonal, lowerDiagonal, xWord)
	fmt.Println("Final count", count)

	return count
}

func testXWordSearch() {
	testDiagString := `XMAS
XMAS
XMAS
XMAS`

	testLines, err := StringToLines(testDiagString)
	if err != nil {
		panic("Error parsing test string")
	}
	if CountXOccurences(testLines, "MAS") != 2 {
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
	if CountXOccurences(testLines, "MAS") != 9 {
		panic("Failed test case.")
	}

}

func testFindIndices() {
	testString := "SMASAMSAM"
	indices := findIndicesForwardReverse(testString, "MAS")

	truth := []int{1, 3, 6}

	if len(truth) != len(indices) {
		panic("Failed testFindIndices, incorrect number of hits.")
	}
	for i := 0; i < len(truth); i++ {
		if truth[i] != indices[i] {
			panic("Failed testFindIndices, incorrect index.")
		}
	}

	testString = "SMASAMSAMSMASAMSAM"
	indices = findIndicesForwardReverse(testString, "MAS")

	// Since we just blind join, order is upper-then-lower
	truth = []int{1, 10, 3, 6, 12, 15}

	if len(truth) != len(indices) {
		panic("Failed testFindIndices, incorrect number of hits.")
	}
	for i := 0; i < len(truth); i++ {
		if truth[i] != indices[i] {
			panic("Failed testFindIndices, incorrect index.")
		}
	}
}

func MainPart2() {
	testFindIndices()
	testXWordSearch()

	// parsedLines := GetInputLines()
	// fmt.Println(len(parsedLines), "lines from input.")
	// answer := CountXOccurences(parsedLines, "MAS")
	fmt.Printf("Answer Part 2: %d\n", 0)
}
