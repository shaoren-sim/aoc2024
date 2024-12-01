package main

import "fmt"

func getCountsMap(col1 []int, col2 []int) map[int]int {
	// Using a map to store the counts.
	counts := make(map[int]int)

	for _, val := range col1 {
		// Check if value already exists in the map
		_, ok := counts[val]

		if !ok {
			counts[val] = 0

			for _, occ := range col2 {
				if occ == val {
					counts[val] += 1
				}
			}
		}
	}

	return counts
}

func calculateSimilarityScore(countsMap map[int]int) int {
	var similarityScore int

	for key, value := range countsMap {
		if value == 0 {
			continue
		}
		similarityScore += key * value
	}

	return similarityScore
}

func MainPart2() {
	dataCol1, dataCol2 := GetNumericalColumns()

	countsMap := getCountsMap(dataCol1, dataCol2)
	similarityScore := calculateSimilarityScore(countsMap)

	fmt.Printf("Answer Part 2: %d\n", similarityScore)
}
