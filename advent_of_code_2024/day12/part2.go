package main

import (
	"fmt"
	"slices"
)

func CountInSlice(s []bool) int {
	count := 0
	for _, el := range s {
		if el {
			count += 1
		}
	}
	return count
}

func getTDirection(adjacency []bool) int {
	// Returns a direction corresponsing to the stem of the T,
	// In other words, the direction not part of the straight line.
	// 0: up
	// 1: down
	// 2: left
	// 3: right
	if !adjacency[0] {
		return 1
	} else if !adjacency[1] {
		return 0
	} else if !adjacency[2] {
		return 3
	} else if !adjacency[3] {
		return 2
	} else {
		panic(fmt.Errorf("Invalid adjacency array."))
	}
}

func getLDirection(adjacency []bool) int {
	// Returns the diagonal index to check for filled in L-cases.
	// Diagonal indices
	// 0: Upper-left
	// 1: Lower-left
	// 2: Upper-right.
	// 3: Lower-right

	// Adjacency indices.
	// 0: up
	// 1: down
	// 2: left
	// 3: right
	if adjacency[0] && adjacency[2] {
		return 0
	} else if adjacency[0] && adjacency[3] {
		return 2
	} else if adjacency[1] && adjacency[2] {
		return 1
	} else if adjacency[1] && adjacency[3] {
		return 3
	} else {
		panic(fmt.Errorf("Invalid adjacency array."))
	}
}

func getTFilledDiagonals(diagonals []bool, direction int) int {
	// Uses the direction from getTDirection.
	filledCount := 0
	if direction == 0 {
		// T points up, need upper-left and upper-right.
		if diagonals[0] {
			filledCount += 1
		}
		if diagonals[2] {
			filledCount += 1
		}
	} else if direction == 1 {
		// T points down, need lower-left and lower-right.
		if diagonals[1] {
			filledCount += 1
		}
		if diagonals[3] {
			filledCount += 1
		}
	} else if direction == 2 {
		// T points left, need upper-left and lower-left.
		if diagonals[0] {
			filledCount += 1
		}
		if diagonals[1] {
			filledCount += 1
		}
	} else if direction == 3 {
		// T points right, need upper-right and lower-right.
		if diagonals[2] {
			filledCount += 1
		}
		if diagonals[3] {
			filledCount += 1
		}
	}

	return filledCount
}

func CountCorners(positions [][2]int, position [2]int) int {
	// This only considers the + shape and the diagonals.
	// This stores the adjacency in the following:
	// 0: up
	// 1: down
	// 2: left
	// 3: right
	adjacency := make([]bool, 4)

	y := position[0]
	x := position[1]

	// Check to see if adjacent values are in our list of positions.
	// Case 0: Up
	adjacency[0] = slices.Contains(positions, [2]int{y - 1, x})
	// Case 1: Down
	adjacency[1] = slices.Contains(positions, [2]int{y + 1, x})
	// Case 2: Left
	adjacency[2] = slices.Contains(positions, [2]int{y, x - 1})
	// Case 3: Right
	adjacency[3] = slices.Contains(positions, [2]int{y, x + 1})

	// Calculate diagonals.
	diagonals := make([]bool, 4)

	// Case 0: Upper-left
	diagonals[0] = slices.Contains(positions, [2]int{y - 1, x - 1})
	// Case 1: Lower-left
	diagonals[1] = slices.Contains(positions, [2]int{y + 1, x - 1})
	// Case 2: Upper-right.
	diagonals[2] = slices.Contains(positions, [2]int{y - 1, x + 1})
	// Case 3: Lower-right
	diagonals[3] = slices.Contains(positions, [2]int{y + 1, x + 1})

	// Count-based rules
	adjacencyCount := CountInSlice(adjacency)

	if adjacencyCount == 0 {
		// Rule 0: None adjacent, 4 corners.
		// In the case where there is no adjacency,
		// This is an individual pixel. 4 corners.
		// ...
		// .X.
		// ...
		return 4
	} else if adjacencyCount == 1 {
		// Rule 2: If only 1 adjacent, 2 corners.
		// ...
		// oX.
		// ...
		return 2
	} else if adjacencyCount == 4 {
		// Rule 1, if surrounded, count the number of filled diagonals.
		// .o.
		// oXo
		// .o.
		return 4 - CountInSlice(diagonals)
	} else if adjacencyCount == 3 {
		// ...
		// oXo
		// .o.
		// Depends, a pure T-shape returns 2 corners.
		// But a filled T-shape might have 1 or 0 corners.
		// Also, we need to check this before the straight line case
		// This is because a T is a superset of a straight line.
		// Start by getting direction of the T-shape
		direction := getTDirection(adjacency)
		filledDiagonals := getTFilledDiagonals(diagonals, direction)
		if filledDiagonals == 2 {
			return 0
		} else if filledDiagonals == 1 {
			return 1
		} else if filledDiagonals == 0 {
			return 2
		}
	} else if adjacencyCount == 2 {
		// When 2 adjacent, depends on situation.
		if (adjacency[0] && adjacency[1]) || (adjacency[2] && adjacency[3]) {
			// Rule 3a: If straight line, depends.
			// Straight lines count be an |-shape (no corners)
			// Or it could be a C-shape (2 corners)
			// Or it could be an I-shape (4 corners)
			// Note we already check the T-case in the earlier adjacencyCount==3 case
			// .o.
			// .X.
			// .o.

			// Strongly considering whether this is needed.
			// This should already be handled through other conditions.
			// The problem is that we should already count these through the case below with the 90-degree angle.
			// Should be able to just count the number of filled diagonals.
			// return CountInSlice(diagonals)
			return 0
		} else {
			// Rule 3b: If 90-degree angle, depends.
			// This is because an L shape might either be 2 corners or single corner.
			// Dependent on the diagonal.
			// Note: The upper left and bottom right both count.
			// .o.
			// oX.
			// ...
			diagonalToCheck := getLDirection(adjacency)
			// If the diagonal is filled, 1 corner.
			// 2 otherwise
			if diagonals[diagonalToCheck] {
				return 1
			} else {
				return 2
			}
		}
	}

	panic(fmt.Errorf("Got adjacency of %d, should be between 0 and 4.", adjacencyCount))
}

func CountSidesByPosition(positions [][2]int) int {
	// Sides can be calculated by a simple heuristic.
	// The number of sides == number of corners.

	sides := 0

	// Loop through all the positions, and count the number of corners.
	for _, position := range positions {
		// Use heuristic: num sides = num corners.
		sides += CountCorners(positions, position)
	}
	return sides
}

func SolvePart2(lines []string) int {
	charPosMaps := PrepareInput(lines)

	// Store the price.
	totalPrice := 0

	// Calculate the area and side count.
	for _, charPositions := range charPosMaps {
		// fmt.Println("==============")
		// fmt.Println(char)

		// For all positions, group into components.
		components := GetConnectedComponents(charPositions)
		for _, positions := range components {
			area := len(positions)
			sides := CountSidesByPosition(positions)
			price := area * sides
			totalPrice += price
			// fmt.Println("Component", i+1, "of", len(components), "for character", char)
			// fmt.Println("Area:", area)
			// fmt.Println("Sides:", sides)
			// fmt.Println("Price:", price)
		}
	}

	return totalPrice
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

	totalPrice := SolvePart2(lines)

	if totalPrice != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, totalPrice))
	}
}

func MainPart2() {
	testSolvePart2("testeasy.txt", 80)
	testSolvePart2("test.txt", 1206)
	lines := GetInputs()
	totalPrice := SolvePart2(lines)

	fmt.Printf("Answer Part 2: %d\n", totalPrice)
}
