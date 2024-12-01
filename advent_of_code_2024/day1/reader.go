package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return lines, nil
}

func SplitLinesBySeparator(lines []string, separator string) [][]string {
	var result [][]string

	// Iterating over each line in the slice.
	for _, line := range lines {
		parts := strings.Split(line, separator)

		result = append(result, parts)
	}

	return result
}

func ParseCSVLikeFile(rawFile string, separator string) [][]string {
	lines, err := StringToLines(rawFile)
	if err != nil {
		log.Fatal(err)
	}

	return SplitLinesBySeparator(lines, separator)
}

func ExtractColumns(parsedFile [][]string) [][]string {
	// If rows is empty, return an empty slice
	if len(parsedFile) == 0 {
		return nil
	}

	// Get number of columns
	numCols := len(parsedFile[0])

	cols := make([][]string, numCols)

	for colIdx := 0; colIdx < numCols; colIdx++ {
		var col []string

		for _, row := range parsedFile {
			col = append(col, row[colIdx])
		}

		cols[colIdx] = col
	}

	return cols
}

func GetNumericalColumns() ([]int, []int) {
	// Download input by emulating curl command.
	var inputURL string = "https://adventofcode.com/2024/day/1/input"
	var sessionCookie string = GetCookieFromEnvVar()

	input := DownloadInputWithCookie(inputURL, sessionCookie)

	// Here, lines are in the format of [][]string
	parsedLines := ParseCSVLikeFile(input, "   ")
	columns := ExtractColumns(parsedLines)

	// break up the columns
	col1 := columns[0]
	col2 := columns[1]

	if len(col1) != len(col2) {
		panic("Lengths of two parsed columns are different")
	}

	// Cast the columns into integers for sorting.
	dataCol1 := make([]int, len(col1))
	dataCol2 := make([]int, len(col2))

	for i := range col1 {
		strToInt, err := strconv.Atoi(col1[i])
		if err != nil {
			log.Fatal(err)
		}
		dataCol1[i] = strToInt
	}

	for i := range col2 {
		strToInt, err := strconv.Atoi(col2[i])
		if err != nil {
			log.Fatal(err)
		}
		dataCol2[i] = strToInt
	}

	return dataCol1, dataCol2
}
