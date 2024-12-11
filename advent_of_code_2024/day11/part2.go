package main

import "fmt"

func _blink(val int, blinks int) int {
	// This is still too slow.
	// Need to add a cache to make this feasible.
	count := 0

	// Recursion to avoid OOM from infinite loop.
	if blinks <= 0 {
		// Value no longer needs to be decomposed,
		// Return a count of 1
		// This breaks recursion.
		return 1
	}
	newInput := MutateInput([]int{val})

	for _, val := range newInput {
		count += _blink(val, blinks-1)
	}
	return count
}

// Making a global cache
var cache = make(map[[2]int]int)

func blinkWithCache(val int, blinks int) int {
	if ret, ok := cache[[2]int{val, blinks}]; ok {
		// fmt.Println("Using cache.")
		return ret
	}

	count := 0

	// Recursion to avoid OOM from infinite loop.
	if blinks <= 0 {
		// Value no longer needs to be decomposed,
		// Return a count of 1
		// This breaks recursion.
		return 1
	}
	newInput := MutateInput([]int{val})

	for _, val := range newInput {
		count += blinkWithCache(val, blinks-1)
	}

	cache[[2]int{val, blinks}] = count
	return count
}

func solveRecursion(input []int, blinks int) int {
	count := 0

	for _, val := range input {
		fmt.Println("Solving for value", val)
		// For each value, make it a slice to match the expected input.
		count += blinkWithCache(val, blinks)
	}

	return count
}

func testSolvePart1(input []int, blinks int, expected int) {
	count := solveRecursion(input, blinks)

	if count != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, count))
	}
	fmt.Println("Passed for input", input, "with", blinks, "blinks.")
}

func MainPart2() {
	// Testing the same part1 tests with the new State-based method.
	testSolvePart1([]int{0, 1, 10, 99, 999}, 1, 7)
	testSolvePart1([]int{125, 17}, 1, 3)
	testSolvePart1([]int{125, 17}, 2, 4)
	testSolvePart1([]int{125, 17}, 3, 5)
	testSolvePart1([]int{125, 17}, 4, 9)
	testSolvePart1([]int{125, 17}, 5, 13)
	testSolvePart1([]int{125, 17}, 6, 22)
	testSolvePart1([]int{125, 17}, 25, 55312)
	// panic("Break")
	input := PrepareInput()

	// This didn't work, memory use blows up.
	// for i := range 75 {
	// 	fmt.Println("Doing step", i)
	// 	input = MutateInput(input)
	// }
	count := solveRecursion(input, 75)

	fmt.Printf("Answer Part 2: %d\n", count)
}
