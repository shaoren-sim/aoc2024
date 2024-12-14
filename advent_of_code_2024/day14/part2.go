// Unfortunately the "lazy" parsing from part1 cannot be used here.
// Rewrite parsing to load all positions into memory for drawing purposes.
package main

import (
	"fmt"
)

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
	const maxMovements int = 10
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

		fmt.Println("Step", step)
		draw(positions, bounds)
		fmt.Println("===========")
	}
}
