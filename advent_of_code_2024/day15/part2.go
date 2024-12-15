// Unfortunately the "lazy" parsing from part1 cannot be used here.
// Rewrite parsing to load all positions into memory for drawing purposes.
package main

import (
	"fmt"
)

func MainPart2() {
	testSolve("test.txt", 12)
	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	score := 0

	fmt.Printf("Answer Part 2: %d\n", score)
}
