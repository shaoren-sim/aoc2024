package main

import "fmt"

func testSolvePart2(inputFile string, expected int) {
	// Parse file into array.
	input, err := getDownloadedFile(inputFile)
	// Here, lines are in the format of [][]string
	lines, err := GetLines(input)
	if err != nil {
		panic("Problem parsing file")
	}
	fmt.Println(lines)

	score := 0

	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func MainPart2() {
	testSolvePart2("test.txt", 81)
	lines := GetInputs()
	fmt.Println(len(lines), "lines in input.")

	fmt.Printf("Answer Part 2: %d\n", 0)
}
