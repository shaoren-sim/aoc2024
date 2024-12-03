package main

import (
	"fmt"
	"strings"
)

func heuristicValidRule(input string) (string, string, bool) {
	// Here, assume that we are starting after a "don't()"
	_, doCut, doFound := strings.Cut(input, "do()")
	if !doFound {
		return "", "", true
	}
	validPart, nextPart, stop := strings.Cut(doCut, "don't()")
	return validPart, nextPart, !stop
}

func ExtractValidParts(input string) []string {
	// Heuristic-based filtering.
	var validParts []string

	// Heuristic 1: Sequence starts enabled
	// This means anything before the first "don't()" is valid
	start, nextPart, _ := strings.Cut(input, "don't()")
	validParts = append(validParts, start)

	// Heuristic 2: Anything between a "do()" and "don't()" will be valid.
	for {
		validPart, newNextPart, stopCondition := heuristicValidRule(nextPart)
		validParts = append(validParts, validPart)
		nextPart = newNextPart
		if stopCondition {
			break
		}
	}
	return validParts
}

func ParseCorruptedMemoryWithDos(input string) int {
	validParts := ExtractValidParts(input)
	validString := strings.Join(validParts[:], ",")

	numStrs := ExtractNumStrings(validString)
	sumOfProducts := 0
	for _, numStr := range numStrs {
		numInts := NumStringToInts(numStr)
		sumOfProducts += MultNumInts(numInts)
	}

	return sumOfProducts
}

func TestParseCorruptedMemoryWithDos() {
	inputTestString := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	if ParseCorruptedMemoryWithDos(inputTestString) != 48 {
		panic("Failed test case.")
	}

	inputTestString = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()mul(4,2)don't()do()mul(3,mul(3,4))"
	if ParseCorruptedMemoryWithDos(inputTestString) != 60 {
		panic("Failed test case.")
	}

	inputTestString = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()mul(4,2)don't()do()mul(3,mul(3,4))xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()mul(4,2)don't()do()mul(3,mul(3,4))"
	if ParseCorruptedMemoryWithDos(inputTestString) != 120 {
		panic("Failed test case.")
	}

	inputTestString = "xmul(2,4)&mul[3,7]!^don't()don't()_mul(5,5)+mul(32,64](mul(11,8)undo()do()do()do()?don't()do()?mul(8,5))don't()mul(4,2)don't()do()mul(3,mul(3,4))xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()mul(4,2)don't()do()mul(3,mul(3,4))"
	if ParseCorruptedMemoryWithDos(inputTestString) != 120 {
		panic("Failed test case.")
	}

}
func MainPart2() {
	TestParseCorruptedMemoryWithDos()

	// https://adventofcode.com/2024/day/2#_
	parsedLines := GetInputLines()
	// For simplicity, merge all lines to a single string.
	inputString := strings.Join(parsedLines[:], "")

	sumOfProducts := ParseCorruptedMemoryWithDos(inputString)

	fmt.Printf("Answer Part 2: %d\n", sumOfProducts)
}
