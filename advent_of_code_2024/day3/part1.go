package main

import "fmt"
import "strings"
import "strconv"

func ExtractNumStrings(input string) []string {
	parts := strings.Split(input, "mul(")
	// for _, part := range parts {
	// 	fmt.Println(part)
	// }

	var numStrs []string
	for _, part := range parts[1:] {
		numStr := strings.Split(part, ")")
		if len(numStr) > 1 && strings.Contains(numStr[0], ",") {
			numStrs = append(numStrs, numStr[0])
		}
	}

	return numStrs
}

func NumStringToInts(numStr string) []int {
	numParts := strings.Split(numStr, ",")
	numInts := make([]int, 2)
	// fmt.Println(numParts, len(numParts))
	if len(numParts) == 2 {
		for i, numPart := range numParts {
			numInt, err := strconv.Atoi(numPart)
			if err != nil {
				// fmt.Println("String conversion failed for", numPart, "in", numParts)
				return make([]int, 2)
			}
			numInts[i] = numInt
		}
	}

	return numInts
}

func MultNumInts(numInts []int) int {
	product := 1
	for _, numInt := range numInts {
		product *= numInt
	}
	return product
}

func ParseCorruptedMemory(input string) int {
	numStrs := ExtractNumStrings(input)
	// fmt.Println(numStrs)
	sumOfProducts := 0
	for _, numStr := range numStrs {
		numInts := NumStringToInts(numStr)
		// fmt.Println(numInts)
		sumOfProducts += MultNumInts(numInts)
		// fmt.Println(sumOfProducts)
	}

	return sumOfProducts
}

func TestParseCorruptedMemory() {
	if ParseCorruptedMemory("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))mul(mul()mul(1)mul(1,))") != 161 {
		panic("Failed test case.")
	}

}

func MainPart1() {
	TestParseCorruptedMemory()

	// https://adventofcode.com/2024/day/2#_
	parsedLines := GetInputLines()
	fmt.Println(len(parsedLines), "lines from input.")

	sumOfProducts := 0
	for _, parsedLine := range parsedLines {
		sumOfProducts += ParseCorruptedMemory(parsedLine)
	}

	fmt.Printf("Answer Part 1: %d\n", sumOfProducts)
}
