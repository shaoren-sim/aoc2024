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

func GetChecksumPart2(filled []int, blanks []int) int {
	// Calculate the checksum by summing index*value
	// Includes the blank-filling logic.
	// Store a count of the number of ids used
	idsUsedLeft := 0
	idsUsedRight := 0

	// Get a Map of the filledValues to pop from.
	stackMap := GetStackMap(filled)
	blankMap := GetStackMap(blanks)

	// Track the sum
	sum := 0
	currentInd := 0
	for i := 0; i < len(filled)+len(blanks); i++ {
		if i%2 == 0 {
			fmt.Println(i, "stack", stackMap)
			// Even case: use the filled values.
			values := filled[i/2]
			for range values {
				// Break the loop when all IDs have been depleted.
				val, _ := stackMap.PopStackMapLeft()
				sum += currentInd * val
				currentInd += 1
				idsUsedLeft += 1
			}
		} else {
			// Odd case: use the blanks, and pull the filled.
			fmt.Println(i, "blank", blankMap)
			elements := 0
			blankSize := blanks[i/2]
			for blankSize > elements {
				fmt.Println(i, "stack", stackMap)
				fmt.Println(i, "blank", blankMap)
				valSlice, depleted := stackMap.PopStackMapRightWithSize(blankSize - elements)
				if depleted {
					break
				}
				for _, val := range valSlice {
					sum += currentInd * val
					currentInd += 1
					idsUsedRight += 1
					elements += 1
					blankMap[i/2] -= 1
				}
			}
		}
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

	fmt.Printf("Answer Part 1: %d\n", checksum)
}
