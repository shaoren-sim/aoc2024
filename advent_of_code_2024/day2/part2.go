package main

import (
	"fmt"
)

// func removeFromSliceByInd(slice []int, ind int) []int {
// 	// Because we are operating on differences,
// 	// We need to add the two values together.
// 	return append(slice[:ind], slice[ind+1:]...)
// }

func removeFromSliceByInd(slice []int, ind int) []int {
	// Because we are operating on differences,
	// We need to add the two values together.
	if ind == len(slice)-1 {
		return slice[:ind]
	}
	if ind == 0 {
		return slice[1:]
	}
	newSlice := append(slice[:ind-1], slice[ind-1]+slice[ind])
	return append(newSlice, slice[ind+1:]...)
}

func checkZerosWithTolerance(differences []int, tolerance int, toleranceLimit int) ([]int, int, bool) {
	// If tolerance is already triggered, break
	if tolerance > toleranceLimit {
		return differences, tolerance, false
	}

	loopMax := len(differences)
	for i := 0; i < loopMax; i++ {
		num := differences[i]
		if num == 0 {
			if tolerance < toleranceLimit {
				tolerance += 1
				fmt.Println(differences, differences[i], "at index", i, "fails condition.")
				differences = removeFromSliceByInd(differences, i)
				fmt.Println("Post removal", differences)
				loopMax -= 1
			} else {
				return differences, tolerance, false
			}
		}
	}
	return differences, tolerance, true
}

func checkMonotonicWithTolerace(differences []int, tolerance int, toleranceLimit int) ([]int, int, bool) {
	// If tolerance is already triggered, break
	if tolerance > toleranceLimit {
		return differences, tolerance, false
	}

	// Initialize the first direction.
	// Since we already discard lines with zero earlier, should be safe.
	// decreasing: true;
	// increasing: false;
	direction := differences[0] < 0
	loopMax := len(differences)
	for i := 1; i < loopMax; i++ {
		if direction {
			// Handle case where the direction is positive.
			if differences[i] > 0 {
				if tolerance < toleranceLimit {
					fmt.Println(differences, differences[i], "at index", i, "fails condition.")
					tolerance += 1
					differences = removeFromSliceByInd(differences, i)
					loopMax -= 1
					fmt.Println("Post removal", differences)
				} else {
					return differences, tolerance, false
				}
			}
		} else {
			// Altcase where initial direction is negative
			if differences[i] < 0 {
				if tolerance < toleranceLimit {
					fmt.Println(differences, differences[i], "at index", i, "fails condition.")
					tolerance += 1
					differences = removeFromSliceByInd(differences, i)
					loopMax -= 1
					fmt.Println("Post removal", differences)
				} else {
					return differences, tolerance, false
				}
			}
		}

	}

	return differences, tolerance, true
}

func checkInRangeWithTolerance(differences []int, tolerance int, minval int, maxval int, toleranceLimit int) ([]int, int, bool) {
	// If tolerance is already triggered, break.
	if tolerance > toleranceLimit {
		return differences, tolerance, false
	}
	loopMax := len(differences)
	for i := 0; i < loopMax; i++ {
		num := differences[i]
		// Absolute value.
		if num < 0 {
			num = -num
		}

		if num < minval {
			if tolerance < toleranceLimit {
				tolerance += 1
				fmt.Println(differences, differences[i], "at index", i, "fails condition.")
				differences = removeFromSliceByInd(differences, i)
				fmt.Println("Post removal", differences)
				loopMax -= 1
			} else {
				return differences, tolerance, false
			}
		}
		if num > maxval {
			if tolerance < toleranceLimit {
				tolerance += 1
				fmt.Println(differences, differences[i], "at index", i, "fails condition.")
				differences = removeFromSliceByInd(differences, i)
				fmt.Println("Post removal", differences)
				loopMax -= 1
			} else {
				return differences, tolerance, false
			}
		}
	}
	return differences, tolerance, true
}

func isSafeWithTolerance(differences []int, toleranceLimit int) bool {
	// initialize tolerance
	tolerance := 0

	// Condition 1: Any stagnation (diff = 0)
	differences, tolerance, pass := checkZerosWithTolerance(differences, tolerance, toleranceLimit)
	if !pass {
		fmt.Println(differences, "fails the zeros check")
		return false
	}

	// Condition 2: Always increasing or decreasing.
	differences, tolerance, pass = checkMonotonicWithTolerace(differences, tolerance, toleranceLimit)
	fmt.Println(differences)
	if !pass {
		fmt.Println(differences, "fails the monotonic check")
		return false
	}

	// Condition 3: Differences are at least 1, at most 3.
	differences, tolerance, pass = checkInRangeWithTolerance(differences, tolerance, 1, 3, toleranceLimit)
	if !pass {
		fmt.Println(differences, "fails the range check between", 1, "and", 3)
		return false
	}

	return true
}

func MainPart2() {
	const toleranceLimit int = 1
	// testLine := []int{33, 37, 38, 37, 38, 41}
	// differences := getDifferences(testLine)
	//
	// fmt.Println("Original", testLine)
	// fmt.Println("Differences", differences)
	// isSafeWithTolerance(differences, 1)
	//
	// panic("Break")

	// https://adventofcode.com/2024/day/2#_
	parsedLines := GetNumericInput()

	var safeCount int

	for _, line := range parsedLines {
		// fmt.Println("Original", line)
		// Get differences.
		differences := getDifferences(line)
		fmt.Println("-------------")
		fmt.Println("Original", line)
		fmt.Println("Differences", differences)

		// Do all of the checks
		if isSafeWithTolerance(differences, toleranceLimit) {
			// fmt.Println(differences, "passes all tests")
			safeCount += 1
		}
	}

	fmt.Printf("Answer Part 2: %d\n", safeCount)
}
