package main

import (
	"fmt"
)

func ParseMaze(lines []string, startStr string, obstacleStr string) ([][]int, [][3]int) {
	// Parse the maze, returns:
	// 1: Maze in the form of []int, where 0=clear path, 1=obstacle
	// 2: Starting positions in the form of [][2]int, [[x, y, direction]]
	// Direction in state is 0=up, 1=right, 2=down, 3=left
	rows := len(lines)
	cols := len(lines[0])

	// Initialize the maze array.
	maze := make([][]int, rows)

	// Parse the lines.
	// Store the number of starts found.
	starts := make([][3]int, 0)
	for i, row := range lines {
		maze[i] = make([]int, cols)

		for j, char := range row {
			// Since []int slices default to 0,
			// blank spaces do not need to be handled.
			if string(char) == startStr {
				// Store the starting locations.
				loc := [3]int{j, i, 0}
				starts = append(starts, loc)
			} else if string(char) == obstacleStr {
				// Add a 1 to the obstacle for parsing.
				col := maze[i]
				col[j] = 1
			}
		}
	}

	return maze, starts
}

func GetTraversedPath(maze [][]int, startState [3]int) [][]bool {
	// Traverse the maze as per the instruction.
	// Get an output array of the same shape as maze
	// false = not traversed
	// true = traversed.
	rows := len(maze)
	cols := len(maze[0])

	traversedMap := make([][]bool, rows)
	for rowInd := range maze {
		traversedMap[rowInd] = make([]bool, cols)
	}

	// Traverse the maze
	// While loop until hit an out-of-index error
	// Indicating that the guard has left the valid area.
	x := startState[0]
	y := startState[1]
	direction := startState[2]

	// Store next state.
	// These are used for lookahead purposes.
	nextX := x
	nextY := y
	nextDirection := direction
	for {
		// Mark the traversed paths by setting value to True.
		traversedRow := traversedMap[y]
		traversedRow[x] = true

		// Move a step
		if direction == 0 {
			nextY = y - 1
		} else if direction == 1 {
			nextX = x + 1
		} else if direction == 2 {
			nextY = y + 1
		} else if direction == 3 {
			nextX = x - 1
		} else {
			fmt.Println("Invalid direction=", direction)
			panic("Direction is invalid")
		}

		// Check to see if the next position is out of the maze.
		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			break
		}

		// Check if the next step would hit an obstacle.
		nextCol := maze[nextY]
		if nextCol[nextX] == 1 {
			// If an obstacle is found, reset position.
			nextX = x
			nextY = y
			// Turn according to rules.
			nextDirection = direction + 1
			// If the direction is beyond 3 = left, reset to 0=up
			if nextDirection > 3 {
				// fmt.Println("Direction is", nextDirection)
				nextDirection -= 4
				// fmt.Println("After fix:", nextDirection)
			}
		}

		// Debug code to print the maze state
		// fmt.Println("=======")
		// for _, traverse := range traversedMap {
		// 	fmt.Println(traverse)
		// }

		// Store next state
		x = nextX
		y = nextY
		direction = nextDirection
	}

	return traversedMap
}

func SumPaths(traversedMaps [][][]bool) [][]bool {
	summedTraversal := traversedMaps[0]

	for i := 1; i < len(traversedMaps); i++ {
		traversal := traversedMaps[i]
		for j, row := range traversal {
			summedRow := summedTraversal[j]
			for k, el := range row {
				summedRow[k] = el
			}
		}
	}
	return summedTraversal
}

func CountTraversal(traversal [][]bool) int {
	count := 0

	for _, row := range traversal {
		for _, el := range row {
			if el == true {
				count += 1
			}
		}
	}
	return count
}

func GetTraversal(maze [][]int, starts [][3]int) [][]bool {
	// NOTE: Probably a premature optimization
	// But presumably there might be more than 1 start, so this handles that.
	traversedMaps := make([][][]bool, len(starts))
	for startInd, start := range starts {
		traversedMap := GetTraversedPath(maze, start)
		traversedMaps[startInd] = traversedMap
	}

	// Sum the paths.
	return SumPaths(traversedMaps)
}

func testSolve() {
	const startStr = "^"
	const obstacleStr = "#"

	input, err := getDownloadedFile("test.txt")
	if err != nil {
		panic(err)
	}

	// Here, lines are in the format of [][]string
	lines, err := GetLines(input)
	if err != nil {
		panic("Problem parsing file")
	}

	maze, starts := ParseMaze(lines, startStr, obstacleStr)
	traversal := GetTraversal(maze, starts)

	// Count the traversed paths.
	traversalCount := CountTraversal(traversal)
	if traversalCount != 41 {
		fmt.Println("Counted", traversalCount, "traversals, expected 41.")
		panic("Failed test")
	}
}

func MainPart1() {
	testSolve()
	// Constants for parsing.
	const startStr = "^"
	const obstacleStr = "#"

	mazeInput := GetInputs()
	// fmt.Println(len(mazeStr), "lines in input.")
	maze, starts := ParseMaze(mazeInput, startStr, obstacleStr)
	traversal := GetTraversal(maze, starts)
	traversalCount := CountTraversal(traversal)

	fmt.Printf("Answer Part 1: %d\n", traversalCount)
}
