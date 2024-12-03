package main

import (
	"bufio"
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

func GetInputLines() []string {
	// Download input by emulating curl command.
	var inputURL string = "https://adventofcode.com/2024/day/3/input"
	var sessionCookie string = GetCookieFromEnvVar()

	input := DownloadInputWithCookie(inputURL, sessionCookie)

	// Here, lines are in the format of [][]string
	parsedLines, err := StringToLines(input)
	if err != nil {
		panic("Problem downloading file")
	}

	return parsedLines
}
