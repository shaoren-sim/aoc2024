package main

import "fmt"

func CountSidesByPosition(positions [][2]int) int {
	// Perimeter can be found by searching adjacency in a + shape.
	// Since we already assign values to components through a similar process,
	// We can already assume that anything outside of our position list is false.
	// Thus, calculate sides that are not adjacent to the same component.
	sides := 0
	for _, position := range positions {
		y := position[0]
		x := position[1]

		// Check to see if adjacent values are in our list of positions.
		// Case 1: Up
		if !slices.Contains(positions, [2]int{y - 1, x}) {
			perimeter += 1
		}
		// Case 2: Down
		if !slices.Contains(positions, [2]int{y + 1, x}) {
			perimeter += 1
		}
		// Case 3: Left
		if !slices.Contains(positions, [2]int{y, x - 1}) {
			perimeter += 1
		}
		// Case 4: Right
		if !slices.Contains(positions, [2]int{y, x + 1}) {
			perimeter += 1
		}
	}
	return perimeter
}
func SolvePart2(lines []string) int {
	charPosMaps, _ := PrepareInput(lines)

	// Store the price.
	totalPrice := 0

	// Calculate the area and perimeter.
	for _, charPositions := range charPosMaps {
		// fmt.Println("==============")
		// fmt.Println(char)

		// For all positions, group into components.
		components := GetConnectedComponents(charPositions)
		for _, positions := range components {
			area := len(positions)
			perimeter := GetPerimeterFromPosition(positions)
			price := area * perimeter
			totalPrice += price
			// fmt.Println("Component", i+1, "of", len(components), "for character", char)
			// fmt.Println("Area:", area)
			// fmt.Println("Perimeter:", perimeter)
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
	testSolve("testeasy.txt", 436)
	testSolve("test.txt", 1206)
	// panic("Break")
	lines := GetInputs()
	totalPrice := Solve(lines)

	fmt.Printf("Answer Part 2: %d\n", totalPrice)
}
