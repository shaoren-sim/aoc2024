package main

import (
	"fmt"
	"slices"
)

func InitBoolMatrix(yDim int, xDim int) [][]bool {
	matrix := make([][]bool, yDim)
	for y := range yDim {
		line := make([]bool, xDim)
		matrix[y] = line
	}
	return matrix
}

func PrepareInput(lines []string) (map[string][][2]int, map[string][][]bool) {
	yDim := len(lines)
	xDim := len(lines[0])

	// Return value initialization
	// // 1: The string array.
	// array := make([][]string, yDim)
	// 1. The positions of every unique character.
	charPosMaps := make(map[string][][2]int)
	// 2: The map of every unique character.
	charMaps := make(map[string][][]bool)

	for y, line := range lines {
		for x, charRune := range line {
			char := string(charRune)

			// For each character, get a boolean map of the positions.
			if charMap, exists := charMaps[char]; exists {
				charMapLine := charMap[y]
				charMapLine[x] = true
			} else {
				// If does not exist, init the matrix.
				charMap := InitBoolMatrix(yDim, xDim)
				charMapLine := charMap[y]
				charMapLine[x] = true
				charMaps[char] = charMap
			}

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
	return charPosMaps, charMaps
}

func GetArea(charMap [][]bool) int {
	// Flatten the 2D slice for easier looping.
	flatten := slices.Concat(charMap...)

	// Area is just the total number of true values.
	area := 0

	for _, el := range flatten {
		if el == true {
			area += 1
		}
	}

	return area
}

func GetValAtCoord(charMap [][]bool, y int, x int) bool {
	// Port from day10 path finding.
	// Difference: Here we allow out-of-bounds values to count perimeters.
	// Helper function to get a matrix value by coordinate.
	// Mainly because I'm lazy to do a check for boundaries.
	// Returns:
	// int: The value at said coordinate, -1 if invalid.

	// Get the upper bounds.
	yDim := len(charMap)
	xDim := len(charMap[0])

	// If out-of-bounds, return 1 (i.e. assume False)
	if y < 0 || y >= yDim || x < 0 || x >= xDim {
		return false
	}

	row := charMap[y]

	return row[x]
}

func CountAdjacentFalses(charMap [][]bool, y int, x int) int {
	// Counts the number of adjacent false values.
	adjacentFalses := 0

	// Case 1: Up
	val := GetValAtCoord(charMap, y-1, x)
	if val == false {
		adjacentFalses += 1
	}
	// Case 2: Down
	val = GetValAtCoord(charMap, y+1, x)
	if val == false {
		adjacentFalses += 1
	}
	// Case 3: Left
	val = GetValAtCoord(charMap, y, x-1)
	if val == false {
		adjacentFalses += 1
	}
	// Case 4: Right
	val = GetValAtCoord(charMap, y, x+1)
	if val == false {
		adjacentFalses += 1
	}
	return adjacentFalses
}

func GetPerimeter(charMap [][]bool) int {
	// Naive way of calculating the perimeter.
	// Likely inefficient.
	// Any side that neighbours another element adds 1 to perimeter.
	// Using modified + sign search code from day10

	perimeter := 0 // The count of Adjacent false values.
	for y, line := range charMap {
		for x, isChar := range line {
			// Only consider true cases.
			if isChar {
				adjacentFalses := CountAdjacentFalses(charMap, y, x)
				perimeter += adjacentFalses
			}
		}
	}
	return perimeter
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
