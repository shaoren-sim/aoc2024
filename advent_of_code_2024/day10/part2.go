package main

import "fmt"

func testSolvePart2(input string, expected int) {

	value := 0
	if value != expected {
		panic(fmt.Errorf("Expected checksum = %d, got %d", expected, value))
	}
}

func MainPart2() {
	testSolvePart2("", 2858)

	_ = GetInputs()
	fmt.Printf("Answer Part 2: %d\n", 0)
}
