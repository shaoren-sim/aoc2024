package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
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

func getDownloadedFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	// If the file already exists.
	if !errors.Is(err, os.ErrNotExist) {
		b, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		// Cast []byte to []string
		bString := string(b)

		return bString, nil
	} else {
		return "", err
	}
}

func writeStringToFile(input string, filePath string) error {
	// Open the file with deferred close.
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}

func GetInputLines() []string {
	// Save download to file.
	const downloadedFile string = "input.txt"

	// Try to get a pre-downloaded file.
	input, err := getDownloadedFile(downloadedFile)
	if err != nil {
		// If parsing the downladed file failed,
		// Download input by emulating curl command.
		var inputURL string = "https://adventofcode.com/2024/day/4/input"
		var sessionCookie string = GetCookieFromEnvVar()

		input = DownloadInputWithCookie(inputURL, sessionCookie)

		// After downloading, save a copy of the file.
		writeErr := writeStringToFile(input, downloadedFile)
		if writeErr != nil {
			panic(writeErr)
		}
	}

	// Here, lines are in the format of [][]string
	parsedLines, err := StringToLines(input)
	if err != nil {
		panic("Problem parsing file")
	}

	return parsedLines
}
