package main

import "fmt"

func GetPossibleAntinodesPart2(positions [][2]int, dimX int, dimY int) [][2]int {
	// Given the positions, find all the possible antinode locations.
	// This includes nodes that will need to be filtered later on, i.e. out of range.
	antinodes := make([][2]int, 0)

	for i, positionA := range positions {
		// Check all combinations
		for j, positionB := range positions {
			// For the same element, add the node to the list.
			// Since the loop can enter the j-loop, this means there is more than 1 antenna.
			if i == j {
				antinodes = append(
					antinodes,
					positionA,
				)
				continue
			}

			vectorA := GetVector(positionA, positionB)
			// Add new antinodes until out of range.
			aMult := 1
			for {
				// Compute position
				newX := positionB[0] + vectorA[0]*aMult
				newY := positionB[1] + vectorA[1]*aMult

				// Break if out of range.
				if newX < 0 || newX >= dimX || newY < 0 || newY >= dimY {
					break
				}
				antinodes = append(
					antinodes,
					[2]int{newX, newY},
				)
				aMult += 1
			}
			// fmt.Println(positionA, positionB, vectorA, antinodeA)

			// Get the negative vector
			vectorB := [2]int{-vectorA[0], -vectorA[1]}
			aMult = 1
			for {
				// Compute position
				newX := positionA[0] + vectorB[0]*aMult
				newY := positionA[1] + vectorB[1]*aMult

				// Break if out of range.
				if newX < 0 || newX >= dimX || newY < 0 || newY >= dimY {
					break
				}
				antinodes = append(
					antinodes,
					[2]int{newX, newY},
				)
				aMult += 1
			}

			antinodeB := [2]int{positionA[0] + vectorB[0], positionA[1] + vectorB[1]}
			antinodes = append(antinodes, antinodeB)
			// fmt.Println(positionA, positionB, vectorB, antinodeB)
		}
	}

	return antinodes
}

func SolveForAntinodesPart2(lines []string, blankChar string) [][2]int {
	// Track height and width
	dimX := len(lines[0])
	dimY := len(lines)
	// fmt.Println("Input width (X):", dimX)
	// fmt.Println("Input height (Y):", dimY)

	// Get the positions of each unique character.
	positionsMap := ParsePositions(lines, blankChar)

	possibleAntinodes := make([][2]int, 0)
	for _, positions := range positionsMap {
		possibleAntinodes = append(
			possibleAntinodes,
			GetPossibleAntinodesPart2(positions, dimX, dimY)...)
	}
	// fmt.Println("Number of antinodes before filtering:", len(possibleAntinodes))

	// Filter out antinodes that break the rules.
	antinodes := FilterAntinodes(possibleAntinodes, dimX, dimY)

	return antinodes
}

func testSolve2(filepath string, expected int) {
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

	// Filter out antinodes that break the rules.
	antinodes := SolveForAntinodesPart2(lines, blankChar)

	if len(antinodes) != expected {
		fmt.Println("Got", len(antinodes), ", expected", expected)
		panic("Failed test")
	}
}

func MainPart2() {
	testSolve2("test.txt", 34)

	const blankChar string = "."

	lines := GetInputs()

	antinodes := SolveForAntinodesPart2(lines, blankChar)

	fmt.Printf("Answer Part 2: %d\n", len(antinodes))
}
