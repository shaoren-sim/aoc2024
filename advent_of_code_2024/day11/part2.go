package main

import "fmt"

func blink(vals []int, blinks int) ([]int, int, int) {
	remainingVals := MutateInput(vals)
	return remainingVals, blinks - 1, len(remainingVals)
}

func solveInfiniteLoop(input []int, blinks int) int {
	count := 0

	for _, val := range input {
		fmt.Println("Solving for value", val)
		// For each value, make it a slice to match the expected input.
		trackedVals := []int{val}
		remainingBlinks := blinks
		countAtStep := 0
		for {
			trackedVals, remainingBlinks, countAtStep = blink(trackedVals, remainingBlinks)
			fmt.Printf("%d blinks remaining.\n", remainingBlinks)
			if remainingBlinks == 0 {
				count += countAtStep
				break
			}
		}
	}

	return count
}

func testSolvePart1(input []int, blinks int, expected int) {
	count := solveInfiniteLoop(input, blinks)

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

	input := PrepareInput()

	// This didn't work, memory use blows up.
	// for i := range 75 {
	// 	fmt.Println("Doing step", i)
	// 	input = MutateInput(input)
	// }

	count := solveInfiniteLoop(input, 75)

	fmt.Printf("Answer Part 2: %d\n", count)
}
