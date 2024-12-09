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

func (stackMap StackMap) PopStackMapLeft() (int, bool) {
	// Return a popped value and whether the stack is depleted.
	for id := 0; id < len(stackMap); id++ {
		count := stackMap[id]
		// If the leftmost value is depleted, go to the next.
		if count == 0 {
			continue
		}

		// Decrement the count by 1
		stackMap[id] -= 1

		return id, false
	}
	return -1, true
}

func (stackMap StackMap) PopStackMapLeftById(id int) (int, bool) {
	count := stackMap[id]
	// If the leftmost value is depleted, go to the next.
	if count == 0 {
		return -1, true
	}

	// Decrement the count by 1
	stackMap[id] -= 1

	return id, false
}

func (stackMap StackMap) PopStackMapRightWithSize(size int) ([]int, bool) {
	// Return a popped value and whether the stack is depleted.
	for id := len(stackMap) - 1; id >= 0; id-- {
		count := stackMap[id]
		// If the leftmost value is depleted, go to the next.
		if count == 0 {
			continue
		}

		// As per the conditions, if size is insufficient, skip
		if count > size {
			continue
		}
		values := make([]int, count)

		for i := range count {
			values[i] = id
		}

		// Set the count to 0
		stackMap[id] = 0

		return values, false
	}

	return nil, true
}

func _GetChecksumPart2(filled []int, blanks []int) int {
	// Calculate the checksum by summing index*value
	// Includes the blank-filling logic.
	// Store a count of the number of ids used
	idsUsedLeft := 0
	idsUsedRight := 0

	// Get a Map of the filledValues to pop from.
	// Also create a 'source' copy to track positions.
	stackMap := GetStackMap(filled)
	sourceMap := GetStackMap(filled)

	// fmt.Println(stackMap)

	// Track the sum
	sum := 0
	currentInd := 0
	for i := 0; i < len(filled)+len(blanks); i++ {
		fmt.Println(i)
		// Start by filling in the blanks
		// We do this because moving the values around can cause new blanks to emerge.
		if i%2 != 0 {
			// Odd case: use the blanks, and pull the filled.
			elements := 0
			blankSize := blanks[i/2]
			fmt.Println("blankSize", blankSize)
			// Break the loop when all IDs have been depleted.
			fmt.Println("blackSize, elements", blankSize, elements)
			for blankSize > elements {
				fmt.Println(stackMap)
				valSlice, depleted := stackMap.PopStackMapRightWithSize(blankSize - elements)
				fmt.Println("valSlice", valSlice)
				if depleted {
					break
				}
				fmt.Println(i, blankSize, valSlice, stackMap)
				for _, val := range valSlice {
					sum += currentInd * val
					currentInd += 1
					idsUsedRight += 1
					elements += 1
				}
			}
		} else {
			// Even case: pre-existing filled values.
			// For this loop, only use these to increment the counters.
			// This allows us to keep track of the blanks.
			values := filled[i/2]
			for range values {
				// Break the loop when all IDs have been depleted.
				_, depleted := sourceMap.PopStackMapLeft()
				fmt.Println(sourceMap)
				if depleted {
					break
				}
				currentInd += 1
				idsUsedLeft += 1
			}
		}

		// Second loop, handle the remaining
	}
	return sum
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

	// Following instructions, fill blanks from right to left.
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

	// For each blank, expand the slice to match the new blank sizes.
	for blankInd, blankSize := range blankMap {
		blankSlice := blankStack[blankInd]
		for len(blankSlice) < blankSize {
			blankSlice = append(blankSlice, 0)
		}
		blankStack[blankInd] = blankSlice
	}

	// Merge the maps into values.
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
