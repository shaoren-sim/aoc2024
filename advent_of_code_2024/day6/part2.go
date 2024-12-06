package main

import (
	"fmt"
)

func GetPath(maze [][]int, startState [3]int) [][3]int {
	// Traverse the maze as per the instruction.
	// Store the path
	pathStates := make([][3]int, 0)
	// Store the initial state
	pathStates = append(pathStates, startState)

	// Traverse the maze
	// While loop until hit an out-of-index error
	// Indicating that the guard has left the valid area.
	x := startState[0]
	y := startState[1]
	direction := startState[2]

	// Store next state.
	for {
		nextState, mazeEnd := MazeStep(maze, [3]int{x, y, direction})

		if mazeEnd {
			break
		}

		nextX := nextState[0]
		nextY := nextState[1]
		nextDirection := nextState[2]

		// Store next state
		x = nextX
		y = nextY
		direction = nextDirection

		pathStates = append(pathStates, [3]int{x, y, direction})
	}

	return pathStates
}

func NoObstaclesAhead(maze [][]int, startState [3]int) bool {
	// Basic heuristic to get rid of easy cases.
	// "Laser" ahead of the starting point to check for obstacles.
	// If there are no obstacles, return true
	rows := len(maze)
	cols := len(maze[0])

	x := startState[0]
	y := startState[1]
	direction := startState[2]

	if direction == 0 {
		// Moving up
		for y > 0 {
			y = y - 1
			row := maze[y]
			if row[x] == 1 {
				return false
			}
		}
	} else if direction == 1 {
		// Moving right
		row := maze[y]
		for x < cols-1 {
			x = x + 1
			if row[x] == 1 {
				return false
			}
		}
	} else if direction == 2 {
		// Moving down
		for y < rows-1 {
			y = y + 1
			row := maze[y]
			if row[x] == 1 {
				return false
			}
		}
	} else if direction == 3 {
		// Moving left
		row := maze[y]
		for x > 0 {
			x = x - 1
			if row[x] == 1 {
				return false
			}
		}
	} else {
		panic("Direction is invalid")
	}
	return true
}

func MazeStep(maze [][]int, state [3]int) ([3]int, bool) {
	rows := len(maze)
	cols := len(maze[0])

	// Do a step in the maze
	x := state[0]
	y := state[1]
	direction := state[2]

	nextX := x
	nextY := y
	nextDirection := direction

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
		panic("Direction is invalid")
	}

	// Check to see if the next position is out of the maze.
	if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
		return [3]int{x, y, direction}, true
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
		return MazeStep(maze, [3]int{nextX, nextY, nextDirection})
	}

	// Store next state
	x = nextX
	y = nextY
	direction = nextDirection

	return [3]int{x, y, direction}, false
}

func StateInList(currentState [3]int, fullStatesList [][3]int) bool {
	for _, state := range fullStatesList {
		if currentState == state {
			return true
		}
	}
	return false
}

func ExhaustiveObstructionSearch(
	maze [][]int,
	pathStates [][3]int,
) [][2]int {
	rows := len(maze)
	cols := len(maze[0])

	// Exhaustively search all possible obstruction locations.
	// Store the obstruction locations.
	validObstructionLocs := make([][2]int, 0)

	// Heuristic 1: We only obstruct the path (1 obstruction limit)
	for i, startState := range pathStates[:len(pathStates)-1] {
		fmt.Println("State", i+1, "of", len(pathStates)-1)
		// Set this as the start point.
		// This skips the steps to get to this point
		x := startState[0]
		y := startState[1]
		direction := startState[2]

		// Place the obstacle
		obstacleLoc := pathStates[i+1]
		obstacleX := obstacleLoc[0]
		obstacleY := obstacleLoc[1]
		// Make a copy of the maze to not affect the source.
		obstacledMaze := make([][]int, rows)
		for i, mazeRow := range maze {
			obstacledMaze[i] = make([]int, cols)
			obstacledRow := obstacledMaze[i]
			for j, el := range mazeRow {
				obstacledRow[j] = el
				if i == obstacleY {
					obstacledRow[obstacleX] = 1
				}
			}
		}
		// fmt.Println("Placing obstacle at", obstacleX, obstacleY)
		// Do the maze.
		// Turn because of the new obstruction.
		direction = direction + 1
		if direction > 3 {
			direction -= 4
		}

		// Heuristic 2: "Laser lookahead"
		// If there are no obstacles, obstacle does not work.
		if NoObstaclesAhead(maze, [3]int{x, y, direction}) {
			// fmt.Println("Failed laser lookahead test")
			continue
		}

		// Heuristic 3: If the same path is traversed twice, it is probably a loop.
		// Store the traversed map
		traversedMap := make(map[[3]int]int)

		// Exhaustive maze crawl.
		for {
			// Do one step in the maze
			newState, mazeEnd := MazeStep(obstacledMaze, [3]int{x, y, direction})

			// If the maze completes, end.
			if mazeEnd {
				// fmt.Println("Maze completes")
				break
			}

			// Check if the state is already in the pathStates list
			// Assume a loop if the same state is hit twice.
			// If it is, presume it is a loop.
			_, exists := traversedMap[newState]
			if exists {
				traversedMap[newState] += 1
			} else {
				traversedMap[newState] = 1
			}
			// traversedMap[newState] += 1
			// fmt.Println(newState, "is in the pathStates list. Count=", traversedMap[newState])
			if traversedMap[newState] == 2 {
				validObstructionLocs = append(
					validObstructionLocs,
					[2]int{obstacleX, obstacleY},
				)
				break
			}

			// Update params for next step
			x = newState[0]
			y = newState[1]
			direction = newState[2]
		}
	}
	return validObstructionLocs
}

func areSlicesEqual(a [][2]int, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Use a map to easily check existence
	counts := make(map[[2]int]bool)

	for _, el := range a {
		counts[el] = true
	}
	for _, el := range b {
		// Check if the element exists in the map
		_, exists := counts[el]
		if !exists {
			fmt.Println(el, "does not exist in slice a. Valid elements are", a)
			return false
		}
	}

	return true
}

func testObstructionSolve() {
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
	start := starts[0]

	pathStates := GetPath(maze, start)

	validObstructionLocs := ExhaustiveObstructionSearch(maze, pathStates)
	truths := make([][2]int, 6)
	truths[0] = [2]int{3, 6}
	truths[1] = [2]int{6, 7}
	truths[2] = [2]int{7, 7}
	truths[3] = [2]int{1, 8}
	truths[4] = [2]int{3, 8}
	truths[5] = [2]int{7, 9}

	if !areSlicesEqual(truths, validObstructionLocs) {
		panic("Failed test")
	}
}

func MainPart2() {
	testObstructionSolve()
	// panic("Check")
	// Constants for parsing.
	const startStr = "^"
	const obstacleStr = "#"

	mazeInput := GetInputs()
	// fmt.Println(len(mazeStr), "lines in input.")
	maze, starts := ParseMaze(mazeInput, startStr, obstacleStr)
	start := starts[0] // Only 1 start

	pathStates := GetPath(maze, start)
	fmt.Println(len(pathStates))
	validObstructionLocs := ExhaustiveObstructionSearch(maze, pathStates)
	fmt.Printf("Answer Part 2: %d\n", len(validObstructionLocs))
}
