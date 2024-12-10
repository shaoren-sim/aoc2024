package main

import (
	"fmt"
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
			starts = append(starts, [2]int{yInd, xInd})

			lineSlice[xInd] = intForm
		}
		array[yInd] = lineSlice
	}
	return array, starts
}

func SearchNeighbourhood(maze [][]int, pos [2]int, num int) (int, [][2]int) {
	// Get the upper bounds.
	yDim := len(maze)
	xDim := len(maze[0])

	// Search the neighbourhood for the next number.
	next := num + 1

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
		yMax = yDim
	}
	if xMin < 0 {
		xMin = 0
	}
	if xMax >= xDim {
		xMax = xDim
	}

	nextPositions := make([][2]int, 0)
	neighbourhood := maze[yMin:yMax]
	for yMod, neighbourhoodLine := range neighbourhood {
		for xMod, neighbourhoodNum := range neighbourhoodLine[xMin:xMax] {
			if neighbourhoodNum == next {
				nextPositions = append(nextPositions, [2]int{yMin + yMod, xMin + xMod})
			}
		}
	}

	return next, nextPositions
}

func ScoreStart(maze [][]int, start [2]int, startNum int, endNum int) int {
	score := 0

	// Always start at the start number.
	num := startNum
	positions := make([][2]int, 1)
	positions[0] = start
	// Crawl maze until endNum is reached.
	for num != endNum {
		for _, position := range positions {
			num, positions = SearchNeighbourhood(maze, position, num)
			if num == endNum {
				score += 1
			}
		}
	}

	return score
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
	testSolve("testeasy.txt", 1)

	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	fmt.Printf("Answer Part 1: %d\n", 0)
}
