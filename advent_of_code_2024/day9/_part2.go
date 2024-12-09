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
	values := make([]int, size)
	filledValues := 0
	// Return a popped value and whether the stack is depleted.
	for id := len(stackMap) - 1; id >= 0; id-- {
		count := stackMap[id]
		// If the leftmost value is depleted, go to the next.
		if count == 0 {
			continue
		}

		// As per the conditions, if size is insufficient, skip
		if count > size-filledValues {
			continue
		}

		for range count {
			values[filledValues] = id
			filledValues += 1
			if filledValues > size {
				break
			}
			fmt.Println(values)
		}

		// Set the count to 0
		stackMap[id] = 0
	}
	if filledValues > 0 {
		return values, false
	}

	return nil, true
}

func GetChecksumPart2(filled []int, blanks []int) int {
	// Calculate the checksum by summing index*value
	// Includes the blank-filling logic.
	// Store a count of the number of ids used
	idsUsedLeft := 0
	idsUsedRight := 0

	// Get a Map of the filledValues to pop from.
	stackMap := GetStackMap(filled)
	// fmt.Println(stackMap)

	// Track the sum
	sum := 0
	currentInd := 0
	for i := 0; i < len(filled)+len(blanks); i++ {
		if i%2 == 0 {
			// Even case: use the filled values.
			values := filled[i/2]
			for range values {
				// Break the loop when all IDs have been depleted.
				val, depleted := stackMap.PopStackMapLeftById(i)
				if depleted {
					blanks[i] += 1
				}
				sum += currentInd * val
				currentInd += 1
				idsUsedLeft += 1
			}
		} else {
			// Odd case: use the blanks, and pull the filled.
			values := blanks[i/2]
			fmt.Println(values)
			// Break the loop when all IDs have been depleted.
			valSlice, depleted := stackMap.PopStackMapRightWithSize(values)
			if depleted {
				return sum
			}
			fmt.Println(valSlice)
			for _, val := range valSlice {
				sum += currentInd * val
				currentInd += 1
				idsUsedRight += 1
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
