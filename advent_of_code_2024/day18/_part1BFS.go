package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseBytePositions(lines []string) [][2]int {
	// Parse the whole list into a single slice of (x, y)
	positions := make([][2]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Errorf("Error parsing line '%s', %s into int.", line, parts[0]))
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Errorf("Error parsing line '%s', %s into int.", line, parts[1]))
		}

		positions[i] = [2]int{x, y}
	}
	return positions
}

func printMaze(positions [][2]int, path [][2]int, xMax int, yMax int) {
	// Helper function to print the maze.
	for y := range yMax + 1 {
		for x := range xMax + 1 {
			pos := [2]int{x, y}
			if slices.Contains(positions, pos) {
				fmt.Printf("#")
			} else if slices.Contains(path, pos) {
				fmt.Printf("O")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

type PathData struct {
	Position [2]int
	Path     [][2]int
	Cost     int
}

func getNextSteps(current PathData) []PathData {
	nextSteps := make([]PathData, 4)

	// Extract the current data.
	x := current.Position[0]
	y := current.Position[1]
	// For the path and cost, append accordingly.
	path := append(current.Path, current.Position)
	cost := current.Cost + 1

	// Step 1: Up
	nextSteps[0] = PathData{
		Position: [2]int{x, y - 1},
		Path:     path,
		Cost:     cost,
	}
	// Step 2: Right
	nextSteps[1] = PathData{
		Position: [2]int{x + 1, y},
		Path:     path,
		Cost:     cost,
	}
	// Step 3: Down
	nextSteps[2] = PathData{
		Position: [2]int{x, y + 1},
		Path:     path,
		Cost:     cost,
	}
	// Step 4: Left
	nextSteps[3] = PathData{
		Position: [2]int{x - 1, y},
		Path:     path,
		Cost:     cost,
	}
	return nextSteps
}

func getAllPaths(blockPositions [][2]int, start [2]int, end [2]int, xMax int, yMax int) []PathData {
	// Find the shortest path.
	// Uses BFS as per the code from day16.

	// Store all the valid paths that reach the end.
	validPaths := make([]PathData, 0)

	// Init path queue.
	queue := []PathData{{Position: start, Path: [][2]int{}, Cost: 0}}

	// Pruning logic.
	costCache := make(map[[2]int]int)

	for len(queue) > 0 {
		// Pop from front of queue
		current := queue[0]
		queue = queue[1:]

		// Extract existing values.
		position := current.Position
		currentCost := current.Cost
		x := position[0]
		y := position[1]

		// Break conditions.
		// Condition 1: If out of bounds
		if x < 0 || y < 0 || x > xMax || y > yMax {
			// fmt.Println("Out of bounds", position)
			continue
		}
		// Condition 2: If is wall
		if slices.Contains(blockPositions, position) {
			// fmt.Println("Is wall", position)
			continue
		}
		// Condition 3: If already traversed.
		if slices.Contains(current.Path, position) {
			// fmt.Println("Already traversed", position)
			continue
		}
		// Condition 4: If at end.
		if position == end {
			validPaths = append(validPaths, current)
			// fmt.Println("Reached end", position)
			continue
		}

		// Check the cost cache.
		existingCost, exists := costCache[position]
		if exists {
			if existingCost < currentCost {
				// fmt.Println("Previously have better cost", position)
				continue
			}
		} else {
			costCache[position] = currentCost
		}
		// Append next possible steps to the queue.
		nextSteps := getNextSteps(current)
		queue = append(queue, nextSteps...)
	}
	return validPaths
}

func getShortestPath(paths []PathData) [][2]int {
	shortestPath := paths[0].Path
	for _, pathObj := range paths {
		if len(pathObj.Path) < len(shortestPath) {
			shortestPath = pathObj.Path
		}
	}
	return shortestPath
}

func testSolve(inputFile string, expected int) {
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
	score := SolvePart1(blockPositions, 0, 12, start, end, edgeLength, edgeLength)
	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func SolvePart1(allBlockPositions [][2]int, blockStart int, blockEnd int, start [2]int, end [2]int, xMax int, yMax int) int {
	blockPositions := allBlockPositions[blockStart:blockEnd]
	printMaze(blockPositions, [][2]int{}, xMax, yMax)
	paths := getAllPaths(blockPositions, start, end, xMax, yMax)

	shortestPath := getShortestPath(paths)
	return len(shortestPath)
}

func MainPart1() {
	testSolve("test.txt", 22)
	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	edgeLength := 70
	var start [2]int = [2]int{0, 0}
	var end [2]int = [2]int{edgeLength, edgeLength}
	blockPositions := parseBytePositions(lines)
	score := SolvePart1(blockPositions, 0, 1024, start, end, edgeLength, edgeLength)
	fmt.Printf("Answer Part 1: %d\n", score)
}
