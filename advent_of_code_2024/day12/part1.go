package main

import (
	"fmt"
	"slices"
)

func PrepareInput(lines []string) map[string][][2]int {
	// Return value initialization
	charPosMaps := make(map[string][][2]int)

	for y, line := range lines {
		for x, charRune := range line {
			char := string(charRune)

			// Also, store a list of positions for each character.
			if charPosMap, exists := charPosMaps[char]; exists {
				charPosMap = append(charPosMap, [2]int{y, x})
				charPosMaps[char] = charPosMap
			} else {
				// If does not exist, init a slice.
				slice := [][2]int{{y, x}}
				charPosMaps[char] = slice
			}
		}
	}
	return charPosMaps
}

func GetPerimeterFromPosition(positions [][2]int) int {
	// Perimeter can be found by searching adjacency in a + shape.
	// Since we already assign values to components through a similar process,
	// We can already assume that anything outside of our position list is false.
	// Thus, calculate sides that are not adjacent to the same component.
	perimeter := 0
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

func crawl(path [][2]int, positions [][2]int, position [2]int) [][2]int {
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

	return path
}

func GetConnectedComponents(positions [][2]int) [][][2]int {
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

func Solve(lines []string) int {
	charPosMaps := PrepareInput(lines)

	// Store the price.
	totalPrice := 0

	// Calculate the area and perimeter.
	for _, charPositions := range charPosMaps {
		// fmt.Println("==============")
		// fmt.Println(char)

		// For all positions, group into components.
		components := GetConnectedComponents(charPositions)
		for _, positions := range components {
			// Area is just the number of filled spots.
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

	totalPrice := Solve(lines)

	if totalPrice != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, totalPrice))
	}
}

func MainPart1() {
	testSolve("testeasy.txt", 140)
	testSolve("test.txt", 1930)
	// panic("Break")
	lines := GetInputs()
	totalPrice := Solve(lines)

	fmt.Printf("Answer Part 1: %d\n", totalPrice)
}
