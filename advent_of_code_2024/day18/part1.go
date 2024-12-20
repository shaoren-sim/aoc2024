package main

import (
	"container/heap"
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

// PriorityQueue implements a priority queue for PathData.
type PriorityQueue []PathData

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest cost, so we use Less based on Cost.
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(PathData)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func getShortestPath(blockPositions [][2]int, start [2]int, end [2]int, xMax int, yMax int) *PathData {
	// Priority queue for exploration.
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize with the start position.
	heap.Push(pq, PathData{Position: start, Path: [][2]int{}, Cost: 0})

	// Cost cache to avoid revisiting with higher cost.
	costCache := make(map[[2]int]int)

	for pq.Len() > 0 {
		// Pop the lowest-cost path.
		current := heap.Pop(pq).(PathData)
		position := current.Position
		currentCost := current.Cost
		x, y := position[0], position[1]

		// Skip if out of bounds.
		if x < 0 || y < 0 || x > xMax || y > yMax {
			continue
		}

		// Skip if it's a wall.
		if slices.Contains(blockPositions, position) {
			continue
		}

		// Check if we reached the end.
		if position == end {
			return &current
		}

		// Skip if a better cost already exists in the cache.
		if cachedCost, exists := costCache[position]; exists && cachedCost <= currentCost {
			continue
		}

		// Update the cost cache.
		costCache[position] = currentCost

		// Explore next steps.
		for _, nextStep := range getNextSteps(current) {
			heap.Push(pq, nextStep)
		}
	}

	// Return nil if no path is found.
	return nil
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
	// printMaze(blockPositions, [][2]int{}, xMax, yMax)
	pathObj := getShortestPath(blockPositions, start, end, xMax, yMax)

	return len(pathObj.Path)
}

func MainPart1() {
	testSolve("test.txt", 22)
	lines := GetInputs()
	// fmt.Println(len(lines), "in input.")

	edgeLength := 70
	var start [2]int = [2]int{0, 0}
	var end [2]int = [2]int{edgeLength, edgeLength}
	blockPositions := parseBytePositions(lines)
	score := SolvePart1(blockPositions, 0, 1024, start, end, edgeLength, edgeLength)
	fmt.Printf("Answer Part 1: %d\n", score)
}
