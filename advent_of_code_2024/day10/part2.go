package main

import "fmt"

func SearchNeighbourhoodPart2(
	maze [][]int,
	positions [][2]int,
	num int,
	endNum int,
) (int, [][2]int, bool) {
	// Recursive function to find ends.
	// Returns:
	// num: The current number to search for.
	// nextPositions: The next positions to start searching from.
	// ended: Signal to end infinite loop. true if end.

	// Break condition for successful recursion.
	// If the number is equal to the end number,
	if num == endNum {
		fmt.Println("End condition reached, end found.")
		return -1, positions, true
	}
	// Alternate break condition, if no positions provided.
	if len(positions) == 0 {
		return -1, make([][2]int, 0), true
	}

	// Search in a + shape for the next number.
	next := num + 1
	nextPositions := make([][2]int, 0)

	// Search in a + shape
	for _, pos := range positions {
		y := pos[0]
		x := pos[1]
		// Case 1: Up
		val, valid := GetValAtCoord(maze, y-1, x)
		if valid && val == next {
			position := [2]int{y - 1, x}
			nextPositions = append(nextPositions, position)
		}
		// Case 2: Down
		val, valid = GetValAtCoord(maze, y+1, x)
		if valid && val == next {
			position := [2]int{y + 1, x}
			nextPositions = append(nextPositions, position)
		}
		// Case 3: Left
		val, valid = GetValAtCoord(maze, y, x-1)
		if valid && val == next {
			position := [2]int{y, x - 1}
			nextPositions = append(nextPositions, position)
		}
		// Case 4: Right
		val, valid = GetValAtCoord(maze, y, x+1)
		if valid && val == next {
			position := [2]int{y, x + 1}
			nextPositions = append(nextPositions, position)
		}
	}

	return next, nextPositions, false
}

func ScoreStartPart2(maze [][]int, start [2]int, startNum int, endNum int) int {
	// Always start at the start number.
	num := startNum
	positions := [][2]int{start}

	// Infinite loop function to emulate "recursive" functions.
	ended := false
	for {
		num, positions, ended = SearchNeighbourhoodPart2(
			maze,
			positions,
			num,
			endNum,
		)
		if num == endNum {
			break
		}
		if ended {
			break
		}
	}
	return len(positions)
}

func testSolvePart2(inputFile string, expected int) {
	const startNum int = 0
	const endNum int = 9

	// Parse file into array.
	input, err := getDownloadedFile(inputFile)
	// Here, lines are in the format of [][]string
	lines, err := GetLines(input)
	if err != nil {
		panic("Problem parsing file")
	}

	// Cast the file into an array.
	maze, starts := LinesToArr(lines, startNum)

	// For each start, calculate the score.
	totalScore := 0
	for _, start := range starts {
		totalScore += ScoreStartPart2(maze, start, startNum, endNum)
	}

	if totalScore != expected {
		panic(fmt.Errorf("Expected score = %d, got %d", expected, totalScore))
	}
}

func MainPart2() {
	const startNum int = 0
	const endNum int = 9
	testSolvePart2("test.txt", 81)
	lines := GetInputs()

	// Cast the file into an array.
	maze, starts := LinesToArr(lines, startNum)

	// For each start, calculate the score.
	totalScore := 0
	for _, start := range starts {
		totalScore += ScoreStartPart2(maze, start, startNum, endNum)
	}
	fmt.Printf("Answer Part 2: %d\n", totalScore)
}
