package main

import "fmt"

type State struct {
	Zeros          int
	Ones           int
	Twos           int
	Fours          int
	Twenty         int
	TwentyFour     int
	TwoOhTwoFour   int
	EverythingElse []int
}

func initState() State {
	return State{0, 0, 0, 0, 0, 0, 0, make([]int, 0)}
}

func (state *State) addValue(val int) {
	if val == 0 {
		state.Zeros += 1
	} else if val == 1 {
		state.Ones += 1
	} else if val == 2 {
		state.Twos += 1
	} else if val == 4 {
		state.Fours += 1
	} else if val == 20 {
		state.Twenty += 1
	} else if val == 24 {
		state.TwentyFour += 1
	} else if val == 2024 {
		state.TwoOhTwoFour += 1
	} else {
		state.EverythingElse = append(state.EverythingElse, val)
	}
}

func (state *State) prepareState(input []int) {
	// Initial method call to cast the initial input into a state.
	for _, val := range input {
		state.addValue(val)
	}
}

func (state State) mutate() State {
	newState := initState()

	// Rule 1: 0 -> 1
	newState.Ones += state.Zeros

	// Rule 2: 1 -> 2024
	newState.TwoOhTwoFour += state.Ones

	// Rule 3: -> 2024 -> 20 | 24
	newState.Twenty += state.TwoOhTwoFour
	newState.TwentyFour += state.TwoOhTwoFour

	// Rule 4a: 20 -> 2|0
	newState.Twos += state.Twenty
	newState.Zeros += state.Twenty
	// Rule 4b: 24 -> 2|4
	newState.Twos += state.TwentyFour
	newState.Fours += state.TwentyFour

	// Mutate everything else accordingly.
	everythingElse := MutateInput(state.EverythingElse)

	// But, we also need to sort out any values we want to track.
	for _, val := range everythingElse {
		newState.addValue(val)
	}

	// Apart from the values generated from mutation
	// There are also values formed from 2s and 4s.
	// Rule 5a: 2 -> 4096
	for range state.Twos {
		newState.EverythingElse = append(newState.EverythingElse, 4096)
	}
	// Rule 5b: 4 -> 9192
	for range state.Fours {
		newState.EverythingElse = append(newState.EverythingElse, 9096)
	}
	return newState
}

func (state State) getCount() int {
	count := 0

	count += state.Zeros
	count += state.Ones
	count += state.Twos
	count += state.Fours
	count += state.Twenty
	count += state.TwentyFour
	count += state.TwoOhTwoFour
	count += len(state.EverythingElse)

	return count
}

func CalcByHeuristic(input []int, blinks int) int {
	// A heuristic-based method to get the number of stones at the end.
	// Heuristic 1: The arrangement does not matter.
	// We only care about the number of stones.

	// Heuristic 2: Based on rule 1 and rule 3
	// 0s will always follow the same progression.
	// 0 -> 1 -> 2024 -> 20|24 -> 2|0|2|4 -> 4048|1|4048|9182
	// Here, we see that:
	// 1. 0s always turn into 1s.
	// 2. 1s always turn into 2024s.
	// 3. 2024 becomes 20 then 24
	// 4. 20|24 becomes 2|0|2|4
	// 5. Here, we have another 0.
	// Hence, we use a state that stores the 0, 2, 4, 20, 24, 2024 counts.

	// Start by casting the input into a state.
	state := initState()
	state.prepareState(input)

	return 0
}

func testSolvePart1(input []int, blinks int, expected int) {
	state := initState()
	state.prepareState(input)
	for range blinks {
		state = state.mutate()
	}

	for _, val := range state.EverythingElse {
		if val == 0 || val == 1 || val == 2 || val == 4 || val == 20 || val == 24 || val == 2024 {
			panic("Invalid value found.")
		}
	}
	fmt.Println(state.EverythingElse)
	count := state.getCount()
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
	panic("Break")

	input := PrepareInput()

	// This didn't work, memory use blows up.
	// for i := range 75 {
	// 	fmt.Println("Doing step", i)
	// 	input = MutateInput(input)
	// }

	fmt.Printf("Answer Part 2: %d\n", len(input))
}
