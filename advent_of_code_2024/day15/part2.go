// Unfortunately the "lazy" parsing from part1 cannot be used here.
// Rewrite parsing to load all positions into memory for drawing purposes.
package main

import (
	"fmt"
	"slices"
)

var SymbolsP2 []string = []string{"#", "[", "]", "@"}

type PositionMapP2 map[string][][2]int

func (positionMap PositionMapP2) GetRobotPosition() [2]int {
	// Gets the current location of the robot
	slice := positionMap["@"]
	// Here, this position is always of length 1
	// (unless more robots are introduced.)
	return slice[0]
}

func (positionMap PositionMapP2) SetRobotPosition(newPos [2]int) {
	slice := [][2]int{newPos}
	positionMap["@"] = slice
}

func (positionMap PositionMapP2) moveBoxesLR(
	positionToConsider [2]int,
	direction int,
) bool {
	// Function to recursively move the boxes.
	// Left/Right cases.
	// Break condition: If there is no box at that position.
	// Search for the existence of the box (either side)
	indLeft, isLeft := searchForPosition(positionMap["["], positionToConsider)
	indRight, isRight := searchForPosition(positionMap["]"], positionToConsider)

	if !isLeft && !isRight {
		// No box exists, just move.
		return true
	} else {
		// fmt.Println("Entering recursive call with", positionToConsider, isLeft, isRight)
		// Further recursive moving and checks.
		// This prevents two boxes from having the same coords.
		coordsLeft := positionToConsider
		coordsRight := positionToConsider
		if isLeft {
			coordsRight = [2]int{positionToConsider[0], positionToConsider[1] + 1}
		} else {
			coordsLeft = [2]int{positionToConsider[0], positionToConsider[1] - 1}
		}

		nextBoxCoordsLeft := translateCoordsByDirection(coordsLeft, direction)
		nextBoxCoordsRight := translateCoordsByDirection(coordsRight, direction)

		nextBoxCoords := nextBoxCoordsLeft
		if isRight {
			nextBoxCoords = nextBoxCoordsRight
		}

		// Check to ensure that the new position of the box is not a wall.
		if slices.Contains(positionMap["#"], nextBoxCoordsLeft) || slices.Contains(positionMap["#"], nextBoxCoordsRight) {
			return false
		}
		// fmt.Println("Making recursive call for", nextBoxCoords, direction)
		canMove := positionMap.moveBoxesLR(nextBoxCoords, direction)
		if !canMove {
			return false
		}

		// Move the found box.
		// Replace the coords in-place.
		// Not as straightforward, need to consider the 2 sides of the box.
		sliceLeft := positionMap["["]
		sliceRight := positionMap["]"]
		// fmt.Println("Left [", sliceLeft)
		// fmt.Println("Right ]", sliceRight)
		if isLeft {
			// Handle left.
			// fmt.Println("Moving [", sliceLeft[indLeft], "to", nextBoxCoordsLeft)
			sliceLeft[indLeft] = nextBoxCoordsLeft
			positionMap["["] = sliceLeft
		} else {
			// fmt.Println("Moving ]", sliceRight[indRight], "to", nextBoxCoordsRight)
			sliceRight[indRight] = nextBoxCoordsRight
			positionMap["]"] = sliceRight
		}
	}
	return true
}

func (positionMap PositionMapP2) checkCanMove(
	positionToConsider [2]int,
	direction int,
) bool {
	// Up/Down case.
	// Search for the existence of the box (either side)
	_, isLeft := searchForPosition(positionMap["["], positionToConsider)
	_, isRight := searchForPosition(positionMap["]"], positionToConsider)

	if !isLeft && !isRight {
		// No box exists, just move.
		return true
	} else {
		// When considering up/down, both sides need to be handled.
		// Further recursive moving and checks.
		// This prevents two boxes from having the same coords.
		coordsLeft := positionToConsider
		coordsRight := positionToConsider
		if isLeft {
			coordsRight = [2]int{positionToConsider[0], positionToConsider[1] + 1}
		} else {
			coordsLeft = [2]int{positionToConsider[0], positionToConsider[1] - 1}
		}

		nextBoxCoordsLeft := translateCoordsByDirection(coordsLeft, direction)
		nextBoxCoordsRight := translateCoordsByDirection(coordsRight, direction)
		// Check to ensure that the new position of the box is not a wall.
		if slices.Contains(positionMap["#"], nextBoxCoordsLeft) || slices.Contains(positionMap["#"], nextBoxCoordsRight) {
			return false
		}
		// fmt.Println("Making recursive call for", nextBoxCoordsLeft, direction)
		// Here, perform a lookahead to see if all boxes can be moved.

		canMoveL := positionMap.checkCanMove(nextBoxCoordsLeft, direction)
		canMoveR := positionMap.checkCanMove(nextBoxCoordsRight, direction)
		return canMoveL && canMoveR
	}
}

func (positionMap PositionMapP2) moveBoxesUD(
	positionToConsider [2]int,
	direction int,
) bool {
	// Up/Down case.
	// Search for the existence of the box (either side)
	indLeft, isLeft := searchForPosition(positionMap["["], positionToConsider)
	indRight, isRight := searchForPosition(positionMap["]"], positionToConsider)

	if !isLeft && !isRight {
		// No box exists, just move.
		return true
	} else {
		// When considering up/down, both sides need to be handled.
		// Further recursive moving and checks.
		// This prevents two boxes from having the same coords.
		coordsLeft := positionToConsider
		coordsRight := positionToConsider
		if isLeft {
			coordsRight = [2]int{positionToConsider[0], positionToConsider[1] + 1}
		} else {
			coordsLeft = [2]int{positionToConsider[0], positionToConsider[1] - 1}
		}

		nextBoxCoordsLeft := translateCoordsByDirection(coordsLeft, direction)
		nextBoxCoordsRight := translateCoordsByDirection(coordsRight, direction)
		// Check to ensure that the new position of the box is not a wall.
		if slices.Contains(positionMap["#"], nextBoxCoordsLeft) || slices.Contains(positionMap["#"], nextBoxCoordsRight) {
			return false
		}
		// fmt.Println("Making recursive call for", nextBoxCoordsLeft, direction)
		// Here, perform a lookahead to see if all boxes can be moved.
		if !positionMap.checkCanMove(nextBoxCoordsLeft, direction) || !positionMap.checkCanMove(nextBoxCoordsRight, direction) {
			return false
		}

		canMove := positionMap.moveBoxesUD(nextBoxCoordsLeft, direction)
		if !canMove {
			return false
		}
		canMove = positionMap.moveBoxesUD(nextBoxCoordsRight, direction)
		if !canMove {
			return false
		}
		// Move the found box.
		// Replace the coords in-place.
		// Not as straightforward, need to consider the 2 sides of the box.
		sliceLeft := positionMap["["]
		sliceRight := positionMap["]"]
		// fmt.Println("Left [", sliceLeft)
		// fmt.Println("Right ]", sliceRight)
		if isLeft {
			// Handle left.
			// fmt.Println("Moving [", sliceLeft[indLeft], "to", nextBoxCoordsLeft)
			sliceLeft[indLeft] = nextBoxCoordsLeft
			positionMap["["] = sliceLeft
			// Handle right.
			searchPos := [2]int{positionToConsider[0], positionToConsider[1] + 1}
			ind, exists := searchForPosition(sliceRight, searchPos)
			if !exists {
				panic(fmt.Errorf("Found box without RHS [ at pos:(%d,%d)", searchPos[0], searchPos[1]))
			}
			// fmt.Println("Moving ]", sliceRight[ind], "to", nextBoxCoordsRight)
			sliceRight[ind] = nextBoxCoordsRight
			positionMap["]"] = sliceRight
		} else {
			// Handle left.
			searchPos := [2]int{positionToConsider[0], positionToConsider[1] - 1}
			ind, exists := searchForPosition(sliceLeft, searchPos)
			if !exists {
				panic(fmt.Errorf("Found box without LHS [ at pos:(%d,%d)", searchPos[0], searchPos[1]))
			}
			// fmt.Println("Moving [", sliceLeft[ind], "to", nextBoxCoordsLeft)
			sliceLeft[ind] = nextBoxCoordsLeft
			positionMap["["] = sliceLeft
			// Handle right.
			// fmt.Println("Moving ]", sliceRight[indRight], "to", nextBoxCoordsRight)
			sliceRight[indRight] = nextBoxCoordsRight
			positionMap["]"] = sliceRight
		}
	}
	return true
}

func (positionMap PositionMapP2) CalculateScore() int {
	boxPosition := positionMap["["]

	score := 0
	for _, pos := range boxPosition {
		y := pos[0]
		x := pos[1]

		score += 100*y + x
	}
	return score
}

func (positionMap PositionMapP2) drawState(xMax int, yMax int) {
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
	for _, symbol := range SymbolsP2 {
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

func PrepareInputsP2(lines []string) (PositionMapP2, []int) {
	// Parse the input
	// Returns:
	// positionMap: map[string][][2]int: All positions of walls etc.
	// movementPath: [][]int: Movement path.
	positionMap := make(map[string][][2]int)

	// Init the positionMap.
	for _, symbol := range SymbolsP2 {
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

		// For each character, cast using the rules described.
		for x, charRune := range line {
			char := string(charRune)

			widerChars := ".."
			if char == "#" {
				widerChars = "##"
			} else if char == "O" {
				widerChars = "[]"
			} else if char == "@" {
				widerChars = "@."
			}
			// Add to the slice in the map.
			for i, widerCharRune := range widerChars {
				widerChar := string(widerCharRune)
				slice := positionMap[widerChar]
				slice = append(slice, [2]int{y, (x * 2) + i})
				positionMap[widerChar] = slice
			}
		}
	}

	// Next, parse the movements.
	movementPath := make([]int, 0)
	for _, line := range lines[startFrom:] {
		// fmt.Println(line)
		// Check to ensure that the line has characters.
		if len(line) == 0 {
			// Nothing to parse.
			break
		}
		for _, charRune := range line {
			char := string(charRune)
			movementPath = append(movementPath, movementMap[char])
		}
	}

	return positionMap, movementPath
}

func testParse(inputFile string) {
	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	positionMap, path := PrepareInputsP2(lines)
	for key, slice := range positionMap {
		fmt.Println(key, slice)
	}
	fmt.Println(path)

	// Debug drawing of state.
	yMax := 0
	xMax := 0
	for _, pos := range positionMap["#"] {
		y := pos[0]
		x := pos[1]
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	xMax += 1
	yMax += 1

	positionMap.drawState(xMax, yMax)
}

func moveRobotP2(positionMap PositionMapP2, direction int) {
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
	if slices.Contains(positionMap["["], nextPosition) ||
		slices.Contains(positionMap["]"], nextPosition) {
		if direction == 1 || direction == 3 {
			canMove := positionMap.moveBoxesLR(nextPosition, direction)
			if !canMove {
				return
			}
		} else {
			canMove := positionMap.moveBoxesUD(nextPosition, direction)
			if !canMove {
				return
			}
		}
	}

	positionMap.SetRobotPosition(nextPosition)
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

	positionMap, path := PrepareInputsP2(lines)

	// Debug drawing of state.
	yMax := 0
	xMax := 0
	for _, pos := range positionMap["#"] {
		y := pos[0]
		x := pos[1]
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	xMax += 1
	yMax += 1

	// positionMap.drawState(xMax, yMax)

	for _, direction := range path {
		moveRobotP2(positionMap, direction)
		// positionMap.drawState(xMax, yMax)
	}

	score := positionMap.CalculateScore()

	if score != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, score))
	}
}

func MainPart2() {
	// testParse("testsmall2.txt")
	// testParse("testlarge.txt")
	testSolvePart2("testsmall2.txt", 618)
	testSolvePart2("testlarge.txt", 9021)
	// testSolvePart2("testbed.txt", 0)
	// panic("BReak")
	lines := GetInputs()

	positionMap, path := PrepareInputsP2(lines)
	// fmt.Println(path)

	// Debug drawing of state.
	// yMax := 0
	// xMax := 0
	// for _, pos := range positionMap["#"] {
	// 	y := pos[0]
	// 	x := pos[1]
	// 	if x > xMax {
	// 		xMax = x
	// 	}
	// 	if y > yMax {
	// 		yMax = y
	// 	}
	// }
	// xMax += 1
	// yMax += 1

	// positionMap.drawState(xMax, yMax)
	for _, direction := range path {
		moveRobotP2(positionMap, direction)
	}

	score := positionMap.CalculateScore()

	fmt.Printf("Answer Part 2: %d\n", score)
}
