package main

import (
	"fmt"
	"slices"
)

var (
	Symbols     []string       = []string{"#", "O", "@"}
	movementMap map[string]int = map[string]int{
		"^": 0,
		">": 1,
		"v": 2,
		"<": 3,
	}
)

type PositionMap map[string][][2]int

func (positionMap PositionMap) GetRobotPosition() [2]int {
	// Gets the current location of the robot
	slice := positionMap["@"]
	// Here, this position is always of length 1
	// (unless more robots are introduced.)
	return slice[0]
}

func (positionMap PositionMap) SetRobotPosition(newPos [2]int) {
	slice := [][2]int{newPos}
	positionMap["@"] = slice
}

func (positionMap PositionMap) moveBoxes(
	positionToConsider [2]int,
	direction int,
) bool {
	// Function to recursively move the boxes.
	boxPositions := positionMap["O"]

	// Break condition: If there is no box at that position.
	// Search for the existence of this slice, and it's position.
	ind, exists := searchForPosition(boxPositions, positionToConsider)
	if !exists {
		// No box exists, just move.
		return true
	} else {
		// Further recursive moving and checks.
		// This prevents two boxes from having the same coords.
		nextBoxPosition := translateCoordsByDirection(positionToConsider, direction)
		// Check to ensure that the new position of the box is not a wall.
		if slices.Contains(positionMap["#"], nextBoxPosition) {
			return false
		}

		canMove := positionMap.moveBoxes(nextBoxPosition, direction)
		if !canMove {
			return false
		}

		// Move the found box.
		// Replace the coords in-place.
		boxPositions[ind] = nextBoxPosition
	}

	return true
}

func (positionMap PositionMap) CalculateScorePart1() int {
	boxPosition := positionMap["O"]

	score := 0
	for _, pos := range boxPosition {
		y := pos[0]
		x := pos[1]

		score += 100*y + x
	}
	return score
}

func (positionMap PositionMap) drawState(xMax int, yMax int) {
	// Debug code to draw the state.
	// Init the canvas
	canvas := make([][]string, yMax)
	for i := range canvas {
		slice := make([]string, xMax)
		for j := range slice {
			slice[j] = "."
		}
		canvas[i] = slice
	}

	// Draw objects.
	for _, symbol := range Symbols {
		positions := positionMap[symbol]
		for _, pos := range positions {
			y := pos[0]
			x := pos[1]
			canvasLine := canvas[y]
			canvasLine[x] = symbol
			canvas[y] = canvasLine
		}
	}

	// Draw to terminal.
	for _, line := range canvas {
		fmt.Println(line)
	}
}

func PrepareInputs(lines []string) (PositionMap, []int) {
	// Parse the input
	// Returns:
	// positionMap: map[string][][2]int: All positions of walls etc.
	// movementPath: [][]int: Movement path.
	positionMap := make(map[string][][2]int)

	// Init the positionMap.
	for _, symbol := range Symbols {
		initSlice := make([][2]int, 0)
		positionMap[symbol] = initSlice
	}

	// For positions, we parse all lines that start with "#".
	startFrom := 0
	for y, line := range lines {
		// Check to ensure that the line has characters.
		if len(line) == 0 {
			// Nothing to parse.
			// Assume this is the break for the next parse step.
			startFrom = y + 1
			break
		}
		// Check to ensure that the line starts with #.
		if string(line[0]) != "#" {
			break
		}

		for x, charRune := range line {
			char := string(charRune)
			// Add to the slice in the map.
			slice := positionMap[char]
			slice = append(slice, [2]int{y, x})
			positionMap[char] = slice
		}
	}

	// Next, parse the movements.
	movementPath := make([]int, 0)
	for _, line := range lines[startFrom:] {
		fmt.Println(line)
		// Check to ensure that the line has characters.
		if len(line) == 0 {
			// Nothing to parse.
			break
		}
		for _, charRune := range line {
			char := string(charRune)
			fmt.Println(charRune, movementMap[char])
			movementPath = append(movementPath, movementMap[char])
		}
	}

	return positionMap, movementPath
}

func translateCoordsByDirection(coords [2]int, direction int) [2]int {
	// Move the robot by position.
	next := [2]int{coords[0], coords[1]}
	if direction == 0 {
		// Up
		next[0] -= 1
	} else if direction == 1 {
		// Right
		next[1] += 1
	} else if direction == 2 {
		// Down
		next[0] += 1
	} else if direction == 3 {
		// Left
		next[1] -= 1
	} else {
		panic(fmt.Errorf("Invalid direction %d, should be between 0-3", direction))
	}

	return next
}

func searchForPosition(slice [][2]int, position [2]int) (int, bool) {
	for ind, el := range slice {
		if el == position {
			return ind, true
		}
	}
	return -1, false
}

func moveRobot(positionMap PositionMap, direction int) {
	// Uses the positionMap to move the robot, performing the checks.
	currentPosition := positionMap.GetRobotPosition()

	nextPosition := translateCoordsByDirection(currentPosition, direction)

	// Check 1, if the position is a wall "#"
	// Don't move
	if slices.Contains(positionMap["#"], nextPosition) {
		// fmt.Println("Wall found")
		nextPosition = currentPosition
	}

	// Check 2: if the position is a box "O"
	// Move the box and anything else in that direction.
	if slices.Contains(positionMap["O"], nextPosition) {
		canMove := positionMap.moveBoxes(nextPosition, direction)
		if !canMove {
			return
		}
	}

	positionMap.SetRobotPosition(nextPosition)
}

func Solve(lines []string) int {
	positionMap, path := PrepareInputs(lines)
	for _, direction := range path {
		moveRobot(positionMap, direction)
	}

	score := positionMap.CalculateScorePart1()
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

	positionMap, path := PrepareInputs(lines)

	// Debug drawing of state.
	yMax := 0
	xMax := 0
	for _, pos := range positionMap["#"] {
		y := pos[0]
		x := pos[1]
		if x > xMax {
			xMax = x + 1
		}
		if y > yMax {
			yMax = y + 1
		}
	}

	// positionMap.drawState(xMax, yMax)

	for _, direction := range path {
		moveRobot(positionMap, direction)
		// positionMap.drawState(xMax, yMax)
	}

	score := positionMap.CalculateScorePart1()

	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func MainPart1() {
	testSolve("testsmall.txt", 2028)
	testSolve("testlarge.txt", 10092)
	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	score := Solve(lines)

	fmt.Printf("Answer Part 1: %d\n", score)
}
