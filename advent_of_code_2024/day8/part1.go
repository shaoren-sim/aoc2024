package main

import (
	"fmt"
	"slices"
)

func ParsePositions(lines []string, blankChar string) map[string][][2]int {
	// Parse the lines.
	// Returns:
	// map[string][][2]int: For each character, a list of positions.

	positionsMap := make(map[string][][2]int)

	for y, line := range lines {
		for x, charRune := range line {
			char := string(charRune)
			if char != blankChar {
				_, exists := positionsMap[char]
				if !exists {
					// If this is a new character, create the list.
					// Store the first position.
					positions := make([][2]int, 1)
					positions[0] = [2]int{x, y}
					// Add the list to the map
					positionsMap[char] = positions
				} else {
					// If this is a previously seen character, append.
					positions := positionsMap[char]
					positions = append(positions, [2]int{x, y})
					positionsMap[char] = positions
				}
			}
		}
	}

	return positionsMap
}

func GetVector(a [2]int, b [2]int) [2]int {
	newX := b[0] - a[0]
	newY := b[1] - a[1]
	return [2]int{newX, newY}
}

func GetPossibleAntinodes(positions [][2]int) [][2]int {
	// Given the positions, find all the possible antinode locations.
	// This includes nodes that will need to be filtered later on, i.e. out of range.
	antinodes := make([][2]int, 0)

	for i, positionA := range positions[:len(positions)-1] {
		// Check all combinations
		for j := 1; j < len(positions); j++ {
			positionB := positions[j]
			// For the same element, skip
			if i == j {
				continue
			}

			vectorA := GetVector(positionA, positionB)
			antinodeA := [2]int{positionB[0] + vectorA[0], positionB[1] + vectorA[1]}
			antinodes = append(
				antinodes,
				antinodeA,
			)
			// fmt.Println(positionA, positionB, vectorA, antinodeA)

			// Get the negative vector
			vectorB := [2]int{-vectorA[0], -vectorA[1]}
			antinodeB := [2]int{positionA[0] + vectorB[0], positionA[1] + vectorB[1]}
			antinodes = append(antinodes, antinodeB)
			// fmt.Println(positionA, positionB, vectorB, antinodeB)
		}
	}

	return antinodes
}

func FilterAntinodes(antinodes [][2]int, dimX int, dimY int) [][2]int {
	filteredAntinodes := make([][2]int, 0)
	for _, antinode := range antinodes {
		// Remove antinodes that are beyond the boundary.
		if antinode[0] < 0 || antinode[0] >= dimX || antinode[1] < 0 || antinode[1] >= dimY {
			continue
		}

		// Skip antinodes that are already in the list.
		if slices.Contains(filteredAntinodes, antinode) {
			continue
		}

		filteredAntinodes = append(filteredAntinodes, antinode)
	}
	return filteredAntinodes
}

func SolveForAntinodes(lines []string, blankChar string) [][2]int {
	// Track height and width
	dimX := len(lines[0])
	dimY := len(lines)
	// fmt.Println("Input width (X):", dimX)
	// fmt.Println("Input height (Y):", dimY)

	// Get the positions of each unique character.
	positionsMap := ParsePositions(lines, blankChar)
	// for k, v := range positionsMap {
	// 	fmt.Println(k, len(v))
	// }

	possibleAntinodes := make([][2]int, 0)
	for _, positions := range positionsMap {
		possibleAntinodes = append(possibleAntinodes, GetPossibleAntinodes(positions)...)
	}
	// fmt.Println("Number of antinodes before filtering:", len(possibleAntinodes))

	// Filter out antinodes that break the rules.
	antinodes := FilterAntinodes(possibleAntinodes, dimX, dimY)

	return antinodes
}

func testSolve(filepath string, expected int) {
	const blankChar string = "."

	input, err := getDownloadedFile(filepath)
	if err != nil {
		panic(err)
	}

	// Here, lines are in the format of [][]string
	lines, err := GetLines(input)
	if err != nil {
		panic("Problem parsing file")
	}

	antinodes := SolveForAntinodes(lines, blankChar)

	if len(antinodes) != expected {
		fmt.Println("Got", len(antinodes), ", expected", expected)
		panic("Failed test")
	}
}

func MainPart1() {
	testSolve("testeasy.txt", 4)
	testSolve("test.txt", 14)

	const blankChar string = "."

	lines := GetInputs()

	antinodes := SolveForAntinodes(lines, blankChar)
	fmt.Printf("Answer Part 1: %d\n", len(antinodes))
}
