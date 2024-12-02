package main

import (
	"fmt"
)

func removeFromSliceByInd(slice []int, ind int) []int {
	result := make([]int, 0, len(slice)-1)    // Create a new slice with the right capacity
	result = append(result, slice[:ind]...)   // Add elements before the index
	result = append(result, slice[ind+1:]...) // Add elements after the index
	return result
}

func isSafePremutations(line []int) bool {
	for i := 0; i < len(line); i++ {
		permutedLine := removeFromSliceByInd(line, i)
		fmt.Println("Permutation", permutedLine)
		// Get differences.
		differences := getDifferences(permutedLine)

		// Do all of the checks
		if isSafe(differences) {
			return true
		}
	}
	return false
}

func MainPart2() {
	// https://adventofcode.com/2024/day/2#_
	parsedLines := GetNumericInput()

	var safeCount int

	for _, line := range parsedLines {
		if isSafePremutations(line) {
			safeCount += 1
		}
	}

	fmt.Printf("Answer Part 2: %d\n", safeCount)
}
