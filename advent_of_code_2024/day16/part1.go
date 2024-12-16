package main

import (
	"fmt"
	"slices"
)

func ParseMaze(lines []string) ([][]bool, [2]int, [2]int) {
	// Parse the maze.
	// Returns:
	// [][]bool: wall = true, path = false.
	// [3]int: Start location {y, x}
	// [2]int: End location {y, x}
	// Directions are in CSS order:
	// 0: Up
	// 1: Right
	// 2: Down
	// 3: Left.
	maze := make([][]bool, len(lines))
	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for y, line := range lines {
		row := make([]bool, len(line))
		for x, charRune := range line {
			char := string(charRune)
			if char == "#" {
				// Wall case.
				row[x] = true
			} else if char == "S" {
				// Start position.
				start = [2]int{y, x}
			} else if char == "E" {
				// End position
				end = [2]int{y, x}
			} else if char == "." {
				// Path with no obstructions.
				row[x] = false
			} else {
				panic(fmt.Errorf("Invalid character %s", char))
			}
		}
		maze[y] = row
	}

	return maze, start, end
}

func getValueAtPosition(maze [][]bool, coords [2]int) bool {
	// Helper function to get value at coordinates.
	y := coords[0]
	x := coords[1]

	row := maze[y]
	return row[x]
}

func translateInDirection(position [2]int, direction int) [2]int {
	// Helper function to move in the specified direction.
	if direction == 0 {
		// Up case.
		return [2]int{position[0] - 1, position[1]}
	} else if direction == 1 {
		// Right case.
		return [2]int{position[0], position[1] + 1}
	} else if direction == 2 {
		// Down case
		return [2]int{position[0] + 1, position[1]}
	} else if direction == 3 {
		// Left case
		return [2]int{position[0], position[1] - 1}
	} else {
		panic(fmt.Errorf("Invalid direction %d, must be 0-3", direction))
	}
}

func getPossibleSteps(maze [][]bool, currentPath [][2]int, position [2]int) ([][2]int, []int) {
	// Counts the number of next steps that the crawler can take.
	possibleSteps := make([][2]int, 0)
	possibleDirections := make([]int, 0)
	// For all possible steps, check each direction.
	// Only allow non-walls and non-previously-traversed paths.
	// Up case.
	pos := [2]int{position[0] - 1, position[1]}
	if !slices.Contains(currentPath, pos) && !getValueAtPosition(maze, pos) {
		possibleSteps = append(possibleSteps, pos)
		possibleDirections = append(possibleDirections, 0)
	}
	// Right case
	pos = [2]int{position[0], position[1] + 1}
	if !slices.Contains(currentPath, pos) && !getValueAtPosition(maze, pos) {
		possibleSteps = append(possibleSteps, pos)
		possibleDirections = append(possibleDirections, 1)
	}
	// Down case
	pos = [2]int{position[0] + 1, position[1]}
	if !slices.Contains(currentPath, pos) && !getValueAtPosition(maze, pos) {
		possibleSteps = append(possibleSteps, pos)
		possibleDirections = append(possibleDirections, 2)
	}
	// Left case
	pos = [2]int{position[0], position[1] - 1}
	if !slices.Contains(currentPath, pos) && !getValueAtPosition(maze, pos) {
		possibleSteps = append(possibleSteps, pos)
		possibleDirections = append(possibleDirections, 3)
	}

	return possibleSteps, possibleDirections
}

func deepCopyPaths(s [][2]int) [][2]int {
	return append(s[:0:0], s...)
}

func deepCopyDirections(s []int) []int {
	return append(s[:0:0], s...)
}

func crawl(
	maze [][]bool,
	end [2]int,
	validPaths [][][2]int,
	validDirectionLists [][]int,
	currentPath [][2]int,
	currentDirectionList []int,
	position [2]int,
	direction int,
) ([][][2]int, [][]int) {
	fmt.Println(position, direction, len(currentPath), len(validPaths), "valid paths found")
	// Break conditions:
	// Condition 1: Reached a wall.
	// if getValueAtPosition(maze, position) {
	// 	fmt.Println("Wall hit")
	// 	// Delete the path in memory.
	// 	currentPath = make([][2]int, 0)
	// 	currentDirectionList = make([]int, 0)
	// 	return validPaths, validDirectionLists
	// }

	currentPath = append(currentPath, position)
	currentDirectionList = append(currentDirectionList, direction)

	// Condition 2: Reached the end.
	if position == end {
		// Since this is a valid path, append.
		validPaths = append(validPaths, currentPath)
		validDirectionLists = append(validDirectionLists, currentDirectionList)
		fmt.Println("Found path reaching end")
		return validPaths, validDirectionLists
	}

	// Get the number of possible next paths.
	possibleSteps, possibleDirections := getPossibleSteps(maze, currentPath, position)
	if len(possibleSteps) != len(possibleDirections) {
		panic("Number of possible steps does not match number of directions.")
	}
	for i := range possibleSteps {
		validPaths, validDirectionLists = crawl(maze, end, validPaths, validDirectionLists, deepCopyPaths(currentPath), deepCopyDirections(currentDirectionList), possibleSteps[i], possibleDirections[i])
		// validPaths, validDirectionLists = crawl(maze, end, validPaths, validDirectionLists, currentPath, currentDirectionList, possibleSteps[i], possibleDirections[i])
	}
	return validPaths, validDirectionLists
}
func GetValidPaths(maze [][]bool, start [2]int, end [2]int) ([][][2]int, [][]int) {
	// Initialize the valid path store.
	validPaths := make([][][2]int, 0)
	validDirectionLists := make([][]int, 0)

	// Initialize the starting state.
	path := make([][2]int, 0)
	directionList := make([]int, 0)

	// // Initialize the blank path storage.
	// // Used to handle branching.
	// pathStorage := make([][][2]int, 0)
	// directionListStorage := make([][]int, 0)
	// validPaths, validDirectionLists = crawl(maze, end, validPaths, validDirectionLists, pathStorage, directionListStorage, path, directionList, start, 1)

	validPaths, validDirectionLists = crawl(maze, end, validPaths, validDirectionLists, path, directionList, start, 1)
	return validPaths, validDirectionLists
}

func calculateScore(path [][2]int, directionList []int) int {
	// Following the rules:
	// For every step taken, add 1.
	// Subtract 1 because the first position is not a step taken.
	score := len(path) - 1

	// For every change in direction 90 degrees, 1000 points.
	for i := range len(directionList) - 1 {
		a := directionList[i]
		b := directionList[i+1]

		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		score += diff * 1000
	}
	return score
}

func Solve(maze [][]bool, start [2]int, end [2]int) int {

	validPaths, validDirections := GetValidPaths(maze, start, end)
	// Check to ensure number of paths match number of dirtect
	if len(validPaths) != len(validDirections) {
		panic("Number of paths does not equal number of directions.")
	}

	// Calculate the lowest score.
	score := calculateScore(validPaths[0], validDirections[0])
	for i := range validPaths[1:] {
		newScore := calculateScore(validPaths[i], validDirections[i])
		if newScore < score {
			score = newScore
		}
	}
	return score
}

func testSolve(inputFile string, expected int) {
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
	// fmt.Println("Start", start, "End", end)
	// for _, line := range maze {
	// 	fmt.Println(line)
	// }

	score := Solve(maze, start, end)
	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func MainPart1() {
	testSolve("test.txt", 7036)
	testSolve("test2.txt", 11048)
	// panic("Check tests")
	lines := GetInputs()
	fmt.Println(len(lines), "lines in input.")

	maze, start, end := ParseMaze(lines)
	score := Solve(maze, start, end)
	fmt.Printf("Answer Part 1: %d\n", score)
}
