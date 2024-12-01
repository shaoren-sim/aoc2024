package main

import (
	"fmt"
	"sort"
)

func sortAscending(arr []int) []int {
	// Sorting the dataCols in ascending order
	// https://stackoverflow.com/a/40932847
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr
}

func MainPart1() {
	// https://adventofcode.com/2024/day/1#_
	dataCol1, dataCol2 := GetNumericalColumns()

	dataCol1 = sortAscending(dataCol1)
	dataCol2 = sortAscending(dataCol2)

	// fmt.Println(dataCol1)
	// fmt.Println(dataCol2)

	// Done with parsing, do the math.
	// Using a loop to calculate the total distance between the two lists.
	sumOfDistances := 0
	for i := 0; i < len(dataCol1); i++ {
		distance := dataCol1[i] - dataCol2[i]

		// Golang surprisingly does not have Abs for ints.
		if distance < 0 {
			distance = -distance
		}

		sumOfDistances += distance
	}

	fmt.Printf("Answer Part 1: %d\n", sumOfDistances)
}
