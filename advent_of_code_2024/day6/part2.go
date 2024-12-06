package main

import (
	"fmt"
)

func IntersectionHeuristic(directionsMap [4][][]bool) [][2]int {
	// Rough heuristic to get the intersections that are invalid

	obstructionLocs := make([][2]int, 0)
	// Test for validity using the direction.
	for direction, directionMap := range directionsMap {
		// Get the corresponding invalid direction for comparison.
		invalidDirection := direction + 3
		if invalidDirection > 3 {
			invalidDirection -= 3
		}
		invalidDirectionMap := directionsMap[invalidDirection]

		// Compare the two directions for intersects.
		for i, directionMapRow := range directionMap {
			invalidDirectionMapRow := invalidDirectionMap[i]
			for j, el := range directionMapRow {
				invalidDirectionMapEl := invalidDirectionMapRow[j]
				if el && invalidDirectionMapEl {
					xPos := j
					yPos := i
					if invalidDirection == 0 {
						yPos -= 1
					} else if invalidDirection == 1 {
						xPos += 1
					} else if invalidDirection == 2 {
						yPos += 1
					} else if invalidDirection == 3 {
						xPos -= 1
					} else {
						panic("Invalid direction")
					}
					obstructionLocs = append(obstructionLocs, [2]int{xPos, yPos})
				}
			}
		}
	}
	return obstructionLocs
}

func GetDirectionsMaps(maze [][]int, startState [3]int) [4][][]bool {
	// Traverse the maze as per the instruction.
	// Get a map of each direction
	rows := len(maze)
	cols := len(maze[0])

	var directionsMap [4][][]bool
	for dirInd := range directionsMap {
		directionsMap[dirInd] = make([][]bool, rows)
		directionMap := directionsMap[dirInd]
		for rowInd := range rows {
			directionMap[rowInd] = make([]bool, cols)
		}
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
		directionMap := directionsMap[direction]
		directionMapRow := directionMap[y]
		directionMapRow[x] = true

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

	return directionsMap
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
	// traversedMap, directionMap, timesTraversedMap, turningPointStates := GetPathInformation(maze, start)
	directionsMap := GetDirectionsMaps(maze, start)

	for dirInd := range directionsMap {
		fmt.Println("Map for direction", dirInd)
		directionMap := directionsMap[dirInd]
		for _, row := range directionMap {
			fmt.Println(row)
		}
	}

	// Do the intersection heuristic.
	obstructionLocs := IntersectionHeuristic(directionsMap)
	fmt.Println(obstructionLocs)
}

func MainPart2() {
	testObstructionSolve()
	// conditions, results := GetInputs()

	fmt.Printf("Answer Part 2: %d\n", 0)
}
