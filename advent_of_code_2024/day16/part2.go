package main

import (
	"fmt"
	"slices"
)

func SolvePart2(maze [][]bool, start [2]int, end [2]int) int {
	// Instead solve as previously, but after obtaining the best score, find tiles.
	validPaths, validDirections := GetValidPaths(maze, start, end)
	// Check to ensure number of paths match number of directions
	if len(validPaths) != len(validDirections) {
		panic("Number of paths does not equal number of directions.")
	}

	// Calculate the lowest score.
	pathScoreMap := make(map[int]int)
	bestScore := calculateScore(validPaths[0], validDirections[0])
	for i := range validPaths {
		newScore := calculateScore(validPaths[i], validDirections[i])
		fmt.Println(len(validPaths[i]), newScore)
		pathScoreMap[i] = newScore
		if newScore < bestScore {
			bestScore = newScore
		}
	}

	tiles := make([][2]int, 0)
	fmt.Println(pathScoreMap)
	for pathInd, score := range pathScoreMap {
		if score == bestScore {
			// Save time on the first append.
			if len(tiles) == 0 {
				tiles = append(tiles, validPaths[pathInd]...)
			} else {
				for _, tile := range validPaths[pathInd] {
					if !slices.Contains(tiles, tile) {
						tiles = append(tiles, tile)
					}
				}
			}
		}
	}
	fmt.Println(tiles)
	fmt.Println(len(tiles))

	for y, row := range maze {
		for x, b := range row {
			if b {
				fmt.Printf("#")
			} else {
				if slices.Contains(tiles, [2]int{y, x}) {
					fmt.Printf("X")
				} else if [2]int{y, x} == start {
					fmt.Printf("S")
				} else if [2]int{y, x} == end {
					fmt.Printf("E")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Printf("\n")
	}
	return len(tiles)
}

func testSolvePart2(inputFile string, expected int) {
	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	maze, start, end := ParseMaze(lines)

	score := SolvePart2(maze, start, end)
	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func MainPart2() {
	testSolvePart2("test.txt", 45)
	testSolvePart2("test2.txt", 64)
	// panic("Check tests")
	lines := GetInputs()
	fmt.Println(len(lines), "lines in input.")

	maze, start, end := ParseMaze(lines)
	score := SolvePart2(maze, start, end)
	fmt.Printf("Answer Part 2: %d\n", score)
}
