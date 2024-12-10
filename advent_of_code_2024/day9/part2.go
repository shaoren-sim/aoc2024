package main

import "fmt"

type StackMap map[int]int

func GetStackMap(values []int) StackMap {
	stack := make(map[int]int, len(values))

	for i, count := range values {
		stack[i] = count
	}

	return stack
}

func FillInTheBlanks(filled []int, blanks []int) []int {
	// Initialize the arranged array.
	// Defaults to all 0s, which is what we want.
	values := make([]int, GetIDCount(filled)+GetIDCount(blanks))

	// Get a Map of the filledValues to pop from.
	stackMap := GetStackMap(filled)
	blankMap := GetStackMap(blanks)

	// Initializing a store for the blanks.
	blankStack := make([][]int, len(blanks))
	// Also track whether this blank is fully filled.
	blankFilledCount := make([]int, len(blanks))

	for blankInd, blankSize := range blanks {
		// Defaults to 0, which is what we want.
		blankStack[blankInd] = make([]int, blankSize)
	}

	// Step 1: Following instructions, fill blanks from right to left.
	for fillInd := len(filled) - 1; fillInd >= 0; fillInd-- {
		fillCount := filled[fillInd]

		// Following instructions, fill blanks from left to right.
		// But only to the left of the file.
		for blankInd, blankSize := range blanks[:fillInd] {
			// If the fill count is too large, skip this blank
			if fillCount > blankSize {
				continue
			}
			// Otherwise, fill in the blank.
			blankSlice := blankStack[blankInd]
			for range fillCount {
				indToFill := blankFilledCount[blankInd]

				blankSlice[indToFill] = fillInd
				stackMap[fillInd] -= 1
				blankMap[fillInd-1] += 1
				// blankMap[blankInd] -= 1
				blankFilledCount[blankInd] += 1
			}

			// Decrement the blank count here.
			blanks[blankInd] -= fillCount
			break
		}
	}

	// Step 2: For each blank, expand the slice to match the new blank sizes.
	for blankInd, blankSize := range blankMap {
		blankSlice := blankStack[blankInd]
		for len(blankSlice) < blankSize {
			blankSlice = append(blankSlice, 0)
		}
		blankStack[blankInd] = blankSlice
	}

	// Step 3: Merge into the rearranged sequence.
	currentInd := 0
	for ind := range blanks {
		// Fill values always first.
		fillCount := stackMap[ind]
		for range fillCount {
			values[currentInd] = ind
			currentInd += 1
		}

		// Next, fill the blanks.
		blankSize := blankMap[ind]
		if len(blankStack[ind]) != blankSize {
			panic("Size mismatch between blankStack and blankSize")
		}
		for _, el := range blankStack[ind] {
			values[currentInd] = el
			currentInd += 1
		}
	}
	// fmt.Println("Values =", values)

	return values
}

func GetChecksumPart2(filled []int, blanks []int) int {
	// Calculate the checksum by summing index*value
	// Start by filling in the blanks.
	rearranged := FillInTheBlanks(filled, blanks)

	sum := 0
	for i, val := range rearranged {
		sum += val * i
	}
	return sum
}

func testSolvePart2(input string, expected int) {
	// Parse the input to get block representation.
	filled, blanks := GetFillsAndBlanks(input)

	checksum := GetChecksumPart2(filled, blanks)

	if checksum != expected {
		panic(fmt.Errorf("Expected checksum = %d, got %d", expected, checksum))
	}
}

func MainPart2() {
	testSolvePart2("2333133121414131402", 2858)

	const blankChar string = "."

	input := GetInputs()

	if len(input) != 1 {
		panic("Input has more than 1 line.")
	}

	filled, blanks := GetFillsAndBlanks(input[0])
	checksum := GetChecksumPart2(filled, blanks)

	fmt.Printf("Answer Part 2: %d\n", checksum)
}
