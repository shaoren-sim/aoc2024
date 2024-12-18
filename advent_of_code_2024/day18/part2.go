package main

import (
	"fmt"
)

func testSolvePart2(inputFile string, expected [2]int) {
	edgeLength := 6

	var start [2]int = [2]int{0, 0}
	var end [2]int = [2]int{edgeLength, edgeLength}
	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}

	blockPositions := parseBytePositions(lines)
	// // As per the question, only consider 12 bytes.
	// blockPositions = blockPositions[:12]
	// // printMaze(blockPositions, [][2]int{}, edgeLength, edgeLength)
	//
	// paths := getAllPaths(blockPositions, start, end, edgeLength, edgeLength)
	// // for _, pathObj := range paths {
	// // 	fmt.Println("Cost:", pathObj.Cost)
	// // 	printMaze(blockPositions, pathObj.Path, edgeLength, edgeLength)
	// // }
	//
	// shortestPath := getShortestPath(paths)
	// score := len(shortestPath)
	lastByte := SolvePart2(blockPositions, start, end, edgeLength, edgeLength, 0)
	if lastByte != expected {
		panic(fmt.Errorf("Expected %v, got %v", expected, lastByte))
	}
}

func SolvePart2(allBlockPositions [][2]int, start [2]int, end [2]int, xMax int, yMax int, startFromBlock int) [2]int {
	// The most basic optimiztions applied.
	// Start from backwards, and try to solve the maze until a path can be found.
	// This saves time because the less spaces there are, the faster the path finding algorithm.
	for lastBlock := len(allBlockPositions); lastBlock > startFromBlock; lastBlock-- {
		blockPositions := allBlockPositions[:lastBlock]
		// printMaze(blockPositions, [][2]int{}, xMax, yMax)
		pathObj := getShortestPath(blockPositions, end, start, xMax, yMax)
		if pathObj != nil {
			return allBlockPositions[lastBlock]
		}
	}
	panic("Did not find a blocked path.")
}

func MainPart2() {
	testSolvePart2("test.txt", [2]int{6, 1})
	lines := GetInputs()

	edgeLength := 70
	var start [2]int = [2]int{0, 0}
	var end [2]int = [2]int{edgeLength, edgeLength}
	blockPositions := parseBytePositions(lines)
	// From part1, we know we can search from 1024 bytes onwards
	lastByte := SolvePart2(blockPositions, start, end, edgeLength, edgeLength, 1025)
	fmt.Printf("Answer Part 2: %v\n", lastByte)
}
