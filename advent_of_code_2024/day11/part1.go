package main

import (
	"fmt"
	"strconv"
	"strings"
)

func InputToArr(input string) []int {
	numArr := make([]int, 0)
	for _, numStr := range strings.Split(input, " ") {
		numInt, err := strconv.Atoi(numStr)
		numArr = append(numArr, numInt)
		if err != nil {
			panic(fmt.Errorf("Error converting %s to integer.", numStr))
		}
	}
	return numArr
}

func MutateInput(input []int) []int {
	newArrangement := make([]int, 0)
	for i, num := range input {
		numStr := strconv.Itoa(num)
		if num == 0 {
			// Rule 1: If number is 0, replace with 1.
			newArrangement = append(newArrangement, 1)
		} else if len(numStr)%2 == 0 {
			// Rule 2: If number has even number of digits, split.
			part, err := strconv.Atoi(numStr[:len(numStr)/2])
			if err != nil {
				panic(fmt.Errorf("Error converting %s to integer.", numStr[:len(numStr)/2]))
			}
			newArrangement = append(newArrangement, part)
			part, err = strconv.Atoi(numStr[len(numStr)/2:])
			if err != nil {
				panic(fmt.Errorf("Error converting %s to integer.", numStr[len(numStr)/2:]))
			}
			newArrangement = append(newArrangement, part)
		} else {
			// Rule 3: If no other rules apply, multiply the stone by 2024
			newArrangement = append(newArrangement, input[i]*2024)
		}
	}
	return newArrangement
}

func testSolve(input []int, blinks int, expected int) {
	for range blinks {
		input = MutateInput(input)
	}

	if len(input) != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, len(input)))
	}
}

func PrepareInput() []int {
	// Parse input.
	lines := GetInputs()
	fmt.Println(len(lines), "lines in input.")

	// Convert input into numeric array.
	input := InputToArr(lines[0])
	return input
}

func MainPart1() {
	testSolve([]int{0, 1, 10, 99, 999}, 1, 7)
	testSolve([]int{125, 17}, 1, 3)
	testSolve([]int{125, 17}, 2, 4)
	testSolve([]int{125, 17}, 3, 5)
	testSolve([]int{125, 17}, 4, 9)
	testSolve([]int{125, 17}, 5, 13)
	testSolve([]int{125, 17}, 6, 22)
	testSolve([]int{125, 17}, 25, 55312)

	input := PrepareInput()

	for range 25 {
		input = MutateInput(input)
	}

	fmt.Printf("Answer Part 1: %d\n", len(input))
}
