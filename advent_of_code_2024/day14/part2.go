// Unfortunately the "lazy" parsing from part1 cannot be used here.
// Rewrite parsing to load all positions into memory for drawing purposes.
package main

import (
	"fmt"
	"slices"
)

// Reuse code from day12 to find connected components.
// Assumes that the Christmas Tree is connected.
func crawl(path [][2]int, positions [][2]int, position [2]int) [][2]int {
	// Modified from day12, includes diagonals as well.
	if !slices.Contains(path, position) {
		path = append(path, position)
	}

	y := position[0]
	x := position[1]

	// Case 1: Up
	if !slices.Contains(path, [2]int{y - 1, x}) && slices.Contains(positions, [2]int{y - 1, x}) {
		path = crawl(path, positions, [2]int{y - 1, x})
	}
	// Case 2: Down
	if !slices.Contains(path, [2]int{y + 1, x}) && slices.Contains(positions, [2]int{y + 1, x}) {
		path = crawl(path, positions, [2]int{y + 1, x})
	}
	// Case 3: Left
	if !slices.Contains(path, [2]int{y, x - 1}) && slices.Contains(positions, [2]int{y, x - 1}) {
		path = crawl(path, positions, [2]int{y, x - 1})
	}
	// Case 4: Right
	if !slices.Contains(path, [2]int{y, x + 1}) && slices.Contains(positions, [2]int{y, x + 1}) {
		path = crawl(path, positions, [2]int{y, x + 1})
	}

	// Added diagonal cases.
	if !slices.Contains(path, [2]int{y - 1, x - 1}) && slices.Contains(positions, [2]int{y - 1, x - 1}) {
		path = crawl(path, positions, [2]int{y - 1, x - 1})
	}
	if !slices.Contains(path, [2]int{y + 1, x + 1}) && slices.Contains(positions, [2]int{y + 1, x + 1}) {
		path = crawl(path, positions, [2]int{y + 1, x + 1})
	}
	if !slices.Contains(path, [2]int{y + 1, x - 1}) && slices.Contains(positions, [2]int{y + 1, x - 1}) {
		path = crawl(path, positions, [2]int{y + 1, x - 1})
	}
	if !slices.Contains(path, [2]int{y - 1, x + 1}) && slices.Contains(positions, [2]int{y - 1, x + 1}) {
		path = crawl(path, positions, [2]int{y - 1, x + 1})
	}
	return path
}

func GetConnectedComponents(positions [][2]int) [][][2]int {
	// Code from day12 to find connected components.
	components := make([][][2]int, 0)

	// Track the used positions.
	usedPosition := make([][2]int, 0)

	for _, position := range positions {
		if slices.Contains(usedPosition, position) {
			continue
		}
		proposedComponent := [][2]int{position}
		proposedComponent = crawl(proposedComponent, positions, position)

		components = append(components, proposedComponent)
		usedPosition = append(usedPosition, proposedComponent...)
	}

	return components
}

func GetLengthLargestComponent(components [][][2]int) int {
	maxLength := 0
	for _, component := range components {
		if len(component) > maxLength {
			maxLength = len(component)
		}
	}

	return maxLength
}

func moveAll(positions [][2]int, velocities [][2]int, bounds [2]int) [][2]int {
	newPositions := make([][2]int, len(positions))
	for i := range positions {
		position := positions[i]
		velocity := velocities[i]
		position = Move(position, velocity, bounds)

		newPositions[i] = position
	}
	return newPositions
}

func draw(positions [][2]int, bounds [2]int) {
	// Initializing the blank "canvas"
	matrix := make([][]int, bounds[1])
	for i := range bounds[1] {
		row := make([]int, bounds[0])
		matrix[i] = row
	}

	// Add a "robot" by position.
	for _, position := range positions {
		// Increment the value
		row := matrix[position[1]]
		row[position[0]] += 1
	}

	// Drawing.
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func MainPart2() {
	// Define the search space.
	// maxMovements: The number of times to search.
	// lengthThreshold: Longest set of connected robots.
	const maxMovements int = 10000
	const lengthThreshold int = 20
	bounds := [2]int{101, 103}

	// Parse out all the initial positions and velocities.
	lines := GetInputs()
	positions := make([][2]int, len(lines))
	velocities := make([][2]int, len(lines))

	for i, line := range lines {
		position, velocity := ParsePositionAndVelocity(line)
		positions[i] = position
		velocities[i] = velocity
	}

	// Draw the initial arrangement.
	// Likely unnecessary since the answer is obviously not 0 steps.
	// draw(positions, bounds)
	for step := range maxMovements {
		// Move all robots 1 step.
		positions = moveAll(positions, velocities, bounds)
		length := GetLengthLargestComponent(GetConnectedComponents(positions))
		if length > lengthThreshold {
			fmt.Println(step+1, length)
			draw(positions, bounds)
		}
	}
}
