package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsePositionAndVelocity(line string) ([2]int, [2]int) {
	stringParts := strings.Split(line, "=")

	positionParts := strings.Split(strings.Split(stringParts[1], " v")[0], ",")
	position := [2]int{}
	for i, positionPart := range positionParts {
		positionInt, err := strconv.Atoi(positionPart)
		if err != nil {
			panic(fmt.Errorf("Error parsing position %s into int.", positionPart))
		}
		position[i] = positionInt
	}

	velocity := [2]int{}
	velocityParts := strings.Split(stringParts[2], ",")
	for i, velocityPart := range velocityParts {
		velocityInt, err := strconv.Atoi(velocityPart)
		if err != nil {
			panic(fmt.Errorf("Error parsing velocity %s into int.", velocityPart))
		}
		velocity[i] = velocityInt
	}
	return position, velocity
}

func Move(position [2]int, velocity [2]int, bounds [2]int) [2]int {
	newX := position[0] + velocity[0]
	newY := position[1] + velocity[1]

	// Handle teleport case.
	if newX < 0 {
		newX += bounds[0]
	} else if newX >= bounds[0] {
		newX -= bounds[0]
	}

	if newY < 0 {
		newY += bounds[1]
	} else if newY >= bounds[1] {
		newY -= bounds[1]
	}

	return [2]int{newX, newY}
}

func GetQuadrant(position [2]int, bounds [2]int) int {
	// Returns the quadrant of the position.
	// 0 = upper-left
	// 1 = upper-right
	// 2 = lower-right
	// 3 = lower-left
	// -1 = on the horizontal/vertical midpoint.

	sideX := position[0] - bounds[0]/2
	sideY := position[1] - bounds[1]/2

	if sideX == 0 || sideY == 0 {
		// Case where the position is on the midpoint.
		return -1
	}
	if sideX < 0 && sideY > 0 {
		return 0
	} else if sideX < 0 && sideY < 0 {
		return 3
	} else if sideX > 0 && sideY > 0 {
		return 1
	} else if sideX > 0 && sideY < 0 {
		return 2
	} else {
		panic(fmt.Errorf("Invalid side found for (%d, %d)", position[0], position[1]))
	}
}

func GetQuadrantAfterSteps(line string, timesMoved int, bounds [2]int) int {
	position, velocity := ParsePositionAndVelocity(line)
	// fmt.Println("Initial position:", position, "velocity", velocity)
	for range timesMoved {
		position = Move(position, velocity, bounds)
	}
	return GetQuadrant(position, bounds)
}

func GetQuadrantCounts(lines []string, timesMoved int, bounds [2]int) [4]int {
	quadrants := [4]int{}
	for _, line := range lines {
		quadrant := GetQuadrantAfterSteps(line, timesMoved, bounds)
		if quadrant != -1 {
			// If the quadrant is valid, add to the bucket.
			quadrants[quadrant] += 1
		}
	}

	return quadrants
}

func testSolve(inputFile string, expected int) {
	const timesMoved int = 100
	bounds := [2]int{11, 7}

	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}

	quadrants := GetQuadrantCounts(lines, timesMoved, bounds)

	// Calculate safety rating as the product.
	totalProduct := 1
	for _, count := range quadrants {
		totalProduct *= count
	}

	if totalProduct != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, totalProduct))
	}
}

func MainPart1() {
	const timesMoved int = 100
	bounds := [2]int{101, 103}

	testSolve("test.txt", 12)
	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	quadrants := GetQuadrantCounts(lines, timesMoved, bounds)

	// Calculate safety rating as the product.
	totalProduct := 1
	for _, count := range quadrants {
		totalProduct *= count
	}

	fmt.Printf("Answer Part 1: %d\n", totalProduct)
}
