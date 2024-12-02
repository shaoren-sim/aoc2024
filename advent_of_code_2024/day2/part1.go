package main

import (
	"fmt"
)

func getDifferences(line []int) []int {
	// Stopping condition
	totalNumbers := len(line)

	// Initialize a difference array
	differences := make([]int, totalNumbers-1)

	// Getting the differences in values.
	for i, num := range line {
		if i+1 >= totalNumbers {
			break
		}
		nextNum := line[i+1]

		difference := num - nextNum
		differences[i] = difference
	}

	return differences
}

func checkZeros(differences []int) bool {
	for _, num := range differences {
		if num == 0 {
			return false
		}
	}
	return true
}

func checkMonotonic(differences []int) bool {
	// Initialize the first direction.
	// Since we already discard lines with zero earlier, should be safe.
	// decreasing: true;
	// increasing: false;
	direction := differences[0] < 0

	for i := 1; i < len(differences); i++ {
		if direction {
			// Handle case where the direction is positive.
			if differences[i] > 0 {
				return false
			}
		} else {
			// Altcase where initial direction is negative
			if differences[i] < 0 {
				return false
			}
		}

	}

	return true
}

func checkInRange(differences []int, minval int, maxval int) bool {
	for _, num := range differences {
		// Absolute value.
		if num < 0 {
			num = -num
		}

		if num < minval {
			return false
		}
		if num > maxval {
			return false
		}
	}
	return true
}

func isSafe(differences []int) bool {
	// Condition 1: Any stagnation (diff = 0)
	if !checkZeros(differences) {
		// fmt.Println(differences, "fails the zeros check")
		return false
	}

	// Condition 2: Always increasing or decreasing.
	if !checkMonotonic(differences) {
		// fmt.Println(differences, "fails the monotonic check")
		return false
	}

	// Condition 3: Differences are at least 1, at most 3.
	if !checkInRange(differences, 1, 3) {
		// fmt.Println(differences, "fails the range check between", 1, "and", 3)
		return false
	}

	return true
}

func MainPart1() {
	// https://adventofcode.com/2024/day/2#_
	parsedLines := GetNumericInput()

	var safeCount int

	for _, line := range parsedLines {
		// fmt.Println("Original", line)
		// Get differences.
		differences := getDifferences(line)

		// Do all of the checks
		if isSafe(differences) {
			// fmt.Println(differences, "passes all tests")
			safeCount += 1
		}
	}

	fmt.Printf("Answer Part 1: %d\n", safeCount)
}
