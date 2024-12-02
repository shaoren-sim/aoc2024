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

func GetNumericInput() [][]int {
	// Download input by emulating curl command.
	var inputURL string = "https://adventofcode.com/2024/day/2/input"
	var sessionCookie string = GetCookieFromEnvVar()

	input := DownloadInputWithCookie(inputURL, sessionCookie)

	// Here, lines are in the format of [][]string
	parsedLines := ParseCSVLikeFile(input, " ")

	numericLines := make([][]int, len(parsedLines))
	for i, line := range parsedLines {
		singleLine := make([]int, len(line))
		for j, numStr := range line {
			numInt, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			singleLine[j] = numInt
		}
		numericLines[i] = singleLine
	}

	return numericLines
}
