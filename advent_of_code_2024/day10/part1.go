package main

import (
	"fmt"
	"slices"
	"strconv"
)

func LinesToArr(lines []string, startNum int) ([][]int, [][2]int) {
	array := make([][]int, len(lines))

	// Track the occurences of starting points.
	starts := make([][2]int, 0)

	for yInd, line := range lines {
		lineSlice := make([]int, len(line))
		for xInd, char := range line {
			intForm, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			// If the integer matches our startNum, append
			if intForm == startNum {
				starts = append(starts, [2]int{yInd, xInd})
			}

			lineSlice[xInd] = intForm
		}
		array[yInd] = lineSlice
	}
	return array, starts
}

func SearchNeighbourhoodWithDiagonals(
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

	// Get the upper bounds.
	yDim := len(maze)
	xDim := len(maze[0])

	// Search the neighbourhood for the next number.
	next := num + 1
	nextPositions := make([][2]int, 0)
	for _, pos := range positions {
		// Define range.
		yMin := pos[0] - 1
		yMax := pos[0] + 1
		xMin := pos[1] - 1
		xMax := pos[1] + 1

		// Truncate neighbourhood at boundaries
		if yMin < 0 {
			yMin = 0
		}
		if yMax >= yDim {
			yMax = yDim - 1
		}
		if xMin < 0 {
			xMin = 0
		}
		if xMax >= xDim {
			xMax = xDim - 1
		}

		neighbourhood := maze[yMin : yMax+1]

		// Debug statment to visualize the neighbourhood.
		fmt.Println("===========")
		fmt.Printf("Searching for %d\n", next)
		fmt.Println("-----------")
		for _, line := range neighbourhood {
			truncLine := line[xMin : xMax+1]
			fmt.Println(truncLine)
		}

		for yMod, neighbourhoodLine := range neighbourhood {
			for xMod, neighbourhoodNum := range neighbourhoodLine[xMin : xMax+1] {
				if neighbourhoodNum == next {
					position := [2]int{yMin + yMod, xMin + xMod}
					fmt.Println("Found", next, "at", position)
					if !slices.Contains(nextPositions, position) {
						nextPositions = append(nextPositions, [2]int{yMin + yMod, xMin + xMod})
					}
				}
			}
		}
	}

	return next, nextPositions, false
}

func GetValAtCoord(array [][]int, y int, x int) (int, bool) {
	// Helper function to get a matrix value by coordinate.
	// Mainly because I'm lazy to do a check for boundaries.
	// Returns:
	// int: The value at said coordinate, -1 if invalid.
	// bool: Whether the value is valid.

	// Get the upper bounds.
	yDim := len(array)
	xDim := len(array[0])

	// Check for invalid values.
	if y < 0 || y >= yDim || x < 0 || x >= xDim {
		return -1, false
	}

	row := array[y]

	return row[x], true
}

func SearchNeighbourhood(
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
			if !slices.Contains(nextPositions, position) {
				nextPositions = append(nextPositions, position)
			}
		}
		// Case 2: Down
		val, valid = GetValAtCoord(maze, y+1, x)
		if valid && val == next {
			position := [2]int{y + 1, x}
			if !slices.Contains(nextPositions, position) {
				nextPositions = append(nextPositions, position)
			}
		}
		// Case 3: Left
		val, valid = GetValAtCoord(maze, y, x-1)
		if valid && val == next {
			position := [2]int{y, x - 1}
			if !slices.Contains(nextPositions, position) {
				nextPositions = append(nextPositions, position)
			}
		}
		// Case 4: Right
		val, valid = GetValAtCoord(maze, y, x+1)
		if valid && val == next {
			position := [2]int{y, x + 1}
			if !slices.Contains(nextPositions, position) {
				nextPositions = append(nextPositions, position)
			}
		}
	}

	return next, nextPositions, false
}

func ScoreStart(maze [][]int, start [2]int, startNum int, endNum int) int {
	// Always start at the start number.
	num := startNum
	positions := [][2]int{start}

	// Infinite loop function to emulate "recursive" functions.
	ended := false
	for {
		num, positions, ended = SearchNeighbourhood(
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

func testSolve(inputFile string, expected int) {
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
		totalScore += ScoreStart(maze, start, startNum, endNum)
	}

	if totalScore != expected {
		panic(fmt.Errorf("Expected score = %d, got %d", expected, totalScore))
	}
}

func MainPart1() {
	const startNum int = 0
	const endNum int = 9
	testSolve("testeasy.txt", 1)
	testSolve("test.txt", 36)
	lines := GetInputs()
	fmt.Println(len(lines), "lines in input.")

	// Cast the file into an array.
	maze, starts := LinesToArr(lines, startNum)

	// For each start, calculate the score.
	totalScore := 0
	for _, start := range starts {
		totalScore += ScoreStart(maze, start, startNum, endNum)
	}
	fmt.Printf("Answer Part 1: %d\n", totalScore)
}
