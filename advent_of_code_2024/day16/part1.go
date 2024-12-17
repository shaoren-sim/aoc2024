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

type PathData struct {
	Path          [][2]int // The sequence of positions in the path
	DirectionList []int    // The sequence of directions taken
	Position      [2]int   // Current position in the path
	Direction     int
}

func _GetValidPaths(maze [][]bool, start [2]int, end [2]int) ([][][2]int, [][]int) {
	// Result variables to store valid paths and their direction lists.
	var validPaths [][][2]int
	var validDirectionLists [][]int

	// Queue to store paths and their associated data for BFS.

	queue := []PathData{
		{
			Path:          [][2]int{start},
			DirectionList: []int{1},
			Position:      start,
			Direction:     1,
		},
	}

	costCache := make(map[[2]int]int)

	for len(queue) > 0 {
		// Dequeue the first element
		current := queue[0]
		queue = queue[1:]

		// Extend the current path
		// current.Path = append(current.Path, current.Position)

		// Check if we've reached the end
		if current.Position == end {
			// existingCost, exists := costCache[end]
			// currentCost := calculateScore(current.Path, current.DirectionList)
			// if exists {
			// 	if existingCost > currentCost {
			// 		costCache[end] = currentCost
			// 	}
			// } else {
			// 	costCache[end] = currentCost
			// }
			validPaths = append(validPaths, current.Path)
			validDirectionLists = append(validDirectionLists, current.DirectionList)
			continue
		}

		// Get possible steps and directions from the current position
		possibleSteps, possibleDirections := getPossibleSteps(maze, current.Path, current.Position)
		if len(possibleSteps) != len(possibleDirections) {
			panic("Number of possible steps does not match number of directions.")
		}

		// Evaluate branching and costs
		if len(possibleSteps) > 1 {
			currentCost := calculateScore(current.Path, current.DirectionList)
			existingCost, exists := costCache[current.Position]
			if exists && currentCost > existingCost {
				// fmt.Println("Pruning at", current.Position, current.Direction)
				continue
			}
			// Update cost cache
			costCache[current.Position] = currentCost
		}

		// Add possible steps to the queue
		for i := range possibleSteps {
			newPathData := PathData{
				Path: append([][2]int{}, current.Path...), // Deep copy current path
				DirectionList: append(
					[]int{},
					current.DirectionList...), // Deep copy direction list
				Position:  possibleSteps[i],
				Direction: possibleDirections[i],
			}
			newPathData.DirectionList = append(newPathData.DirectionList, possibleDirections[i])
			newPathData.Path = append(newPathData.Path, possibleSteps[i])
			queue = append(queue, newPathData)
		}
	}
	// fmt.Println(costCache)
	// for y, row := range maze {
	// 	for x, b := range row {
	// 		if b {
	// 			fmt.Printf("#")
	// 		} else {
	// 			_, exists := costCache[[2]int{y, x}]
	// 			if exists {
	// 				fmt.Printf("X")
	// 			} else {
	// 				if [2]int{y, x} == start {
	// 					fmt.Printf("S")
	// 				} else if [2]int{y, x} == end {
	// 					fmt.Printf("E")
	// 				} else {
	// 					fmt.Printf(".")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	return validPaths, validDirectionLists
}

func findExistingPath(validPaths [][][2]int, validDirectionLists [][]int, position [2]int, direction int) (int, int, bool) {
	for objInd := range validPaths {
		validPath := validPaths[objInd]
		validDirectionList := validDirectionLists[objInd]
		for sliceInd := range validPath {
			pos := validPath[sliceInd]
			dir := validDirectionList[sliceInd]

			if pos == position && dir == direction {
				return objInd, sliceInd, true
			}
		}
	}

	return -1, -1, false
}

func GetValidPaths(maze [][]bool, start [2]int, end [2]int) ([][][2]int, [][]int) {
	// Result variables to store valid paths and their direction lists.
	var validPaths [][][2]int
	var validDirectionLists [][]int

	// Queue to store paths and their associated data for BFS.

	queue := []PathData{
		{
			Path:          [][2]int{start},
			DirectionList: []int{1},
			Position:      start,
			Direction:     1,
		},
	}

	costCache := make(map[[3]int]int)

	for len(queue) > 0 {
		// Dequeue the first element
		current := queue[0]
		queue = queue[1:]

		// Extend the current path
		// current.Path = append(current.Path, current.Position)

		// Check if we've reached the end
		if current.Position == end {
			// existingCost, exists := costCache[end]
			// currentCost := calculateScore(current.Path, current.DirectionList)
			// if exists {
			// 	if existingCost > currentCost {
			// 		costCache[end] = currentCost
			// 	}
			// } else {
			// 	costCache[end] = currentCost
			// }
			validPaths = append(validPaths, current.Path)
			validDirectionLists = append(validDirectionLists, current.DirectionList)
			continue
		}

		state := [3]int{current.Position[0], current.Position[1], current.Direction}
		currentCost := calculateScore(current.Path, current.DirectionList)
		existingCost, exists := costCache[state]
		if exists {
			if currentCost > existingCost {
				// fmt.Println("Pruning at", current.Position, current.Direction)
				continue
			} else if currentCost == existingCost {
				// fmt.Println("Cost match found at", state)
				// // If any valid paths exists, append.
				// // fmt.Println(validPaths)
				// objInd, sliceInd, found := findExistingPath(validPaths, validDirectionLists, current.Position, current.Direction)
				// if found {
				// 	fmt.Println(objInd, sliceInd)
				// } else {
				// 	// Put it at the back of the queue to reevaluate at the end.
				// 	queue = append(queue, current)
				// }
				// continue
			} else if currentCost < existingCost {
				// Update the cache
				costCache[state] = currentCost
			}
		}
		// Update cost cache
		costCache[state] = currentCost

		// Get possible steps and directions from the current position
		possibleSteps, possibleDirections := getPossibleSteps(maze, current.Path, current.Position)
		if len(possibleSteps) != len(possibleDirections) {
			panic("Number of possible steps does not match number of directions.")
		}

		// Evaluate branching and costs
		// if len(possibleSteps) > 1 {
		// 	state := [3]int{current.Position[0], current.Position[1], current.Direction}
		// 	currentCost := calculateScore(current.Path, current.DirectionList)
		// 	existingCost, exists := costCache[state]
		// 	if exists {
		// 		if currentCost > existingCost {
		// 			// fmt.Println("Pruning at", current.Position, current.Direction)
		// 			continue
		// 		} else if currentCost == existingCost {
		// 			// fmt.Println("Cost match found at", state)
		// 			// // If any valid paths exists, append.
		// 			// fmt.Println(validPaths)
		// 			// objInd, sliceInd, found := findExistingPath(validPaths, validDirectionLists, current.Position, current.Direction)
		// 			// if found {
		// 			// 	fmt.Println(objInd, sliceInd)
		// 			// } else {
		// 			// 	// Put it at the back of the queue to reevaluate at the end.
		// 			// 	queue = append(queue, current)
		// 			// }
		// 			continue
		// 		} else if currentCost < existingCost {
		// 			// Update the cache
		// 			costCache[state] = currentCost
		// 		}
		// 	}
		// 	// Update cost cache
		// 	costCache[state] = currentCost
		// }

		// Add possible steps to the queue
		for i := range possibleSteps {
			newPathData := PathData{
				Path: append([][2]int{}, current.Path...), // Deep copy current path
				DirectionList: append(
					[]int{},
					current.DirectionList...), // Deep copy direction list
				Position:  possibleSteps[i],
				Direction: possibleDirections[i],
			}
			newPathData.DirectionList = append(newPathData.DirectionList, possibleDirections[i])
			newPathData.Path = append(newPathData.Path, possibleSteps[i])
			queue = append(queue, newPathData)
		}
	}
	// fmt.Println(costCache)
	// for y, row := range maze {
	// 	for x, b := range row {
	// 		if b {
	// 			fmt.Printf("#")
	// 		} else {
	// 			_, exists := costCache[[2]int{y, x}]
	// 			if exists {
	// 				fmt.Printf("X")
	// 			} else {
	// 				if [2]int{y, x} == start {
	// 					fmt.Printf("S")
	// 				} else if [2]int{y, x} == end {
	// 					fmt.Printf("E")
	// 				} else {
	// 					fmt.Printf(".")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	return validPaths, validDirectionLists
}
func calculateScore(path [][2]int, directionList []int) int {
	// Following the rules:
	// For every step taken, add 1.
	// Subtract 1 because the first position is not a step taken.
	score := len(path) - 1

	// For every change in direction 90 degrees, 1000 points.
	for i := 0; i < len(directionList)-1; i++ {
		a := directionList[i]
		b := directionList[i+1]

		diff := (b - a + 4) % 4
		if 4-diff < diff {
			diff = 4 - diff
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
	for i := range validPaths {
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
