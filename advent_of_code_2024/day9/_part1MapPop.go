// A scrapped attempt at doing a map-based approach, much slower.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetFillsAndBlanks(input string) ([]int, []int) {
	inputLength := len(input)
	lengthMod := 0
	if inputLength%2 != 0 {
		lengthMod += 1
	}

	// Represent the input as 2 arrays
	filled := make([]int, inputLength/2+lengthMod)
	blanks := make([]int, inputLength/2)
	for i, char := range input {
		intForm, err := strconv.Atoi(string(char))
		if err != nil {
			panic(fmt.Errorf("Error casting %s to integer form.", string(char)))
		}
		if i%2 == 0 {
			filled[i/2] = intForm
		} else {
			blanks[i/2] = intForm
		}
	}
	return filled, blanks
}

func DebugStringRepresentation(filled []int, blanks []int) string {
	// Basic debugging function to check the representation.
	// Likely should not be used for performance concerns.
	if len(filled) > 10 {
		panic("More than 10 fill values. Since printing only supports 0-9, cannot get string representation.")
	}

	var sb strings.Builder

	fillId := 0
	for i, fillCount := range filled {
		// Add the fill values.
		for range fillCount {
			sb.WriteString(fmt.Sprintf("%d", fillId))
		}
		fillId += 1

		// Add the blanks
		// Break if out of bounds.
		if i >= len(blanks) {
			break
		}
		blankCount := blanks[i]
		for range blankCount {
			sb.WriteString(".")
		}
	}

	return sb.String()
}

func GetIDCount(filled []int) int {
	total := 0
	for _, count := range filled {
		total += count
	}
	return total
}

func GetStack(values []int) []int {
	stack := make([]int, 0)

	for i, count := range values {
		for range count {
			stack = append(stack, i)
		}
	}
	return stack
}

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

func (stackMap StackMap) PopStackMapRight() (int, bool) {
	// Return a popped value and whether the stack is depleted.
	for id := len(stackMap) - 1; id >= 0; id-- {
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

func GetChecksum(filled []int, blanks []int) int {
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
				val, depleted := stackMap.PopStackMapLeft()
				if depleted {
					return sum
				}
				sum += currentInd * val
				currentInd += 1
				idsUsedLeft += 1
			}
		} else {
			// Odd case: use the blanks, and pull the filled.
			values := blanks[i/2]
			for range values {
				// Break the loop when all IDs have been depleted.
				val, depleted := stackMap.PopStackMapRight()
				if depleted {
					return sum
				}
				sum += currentInd * val
				currentInd += 1
				idsUsedRight += 1
			}
		}
	}
	return sum
}

func testSolve(input string, expected int) {
	// Parse the input to get block representation.
	filled, blanks := GetFillsAndBlanks(input)
	// fmt.Println(filled, blanks)
	// stringRep := DebugStringRepresentation(filled, blanks)
	// fmt.Println(stringRep)

	checksum := GetChecksum(filled, blanks)

	if checksum != expected {
		panic(fmt.Errorf("Expected checksum = %d, got %d", expected, checksum))
	}
}

func MainPart1() {
	testSolve("12345", 60)
	testSolve("2333133121414131402", 1928)

	const blankChar string = "."

	input := GetInputs()

	if len(input) != 1 {
		panic("Input has more than 1 line.")
	}

	filled, blanks := GetFillsAndBlanks(input[0])
	checksum := GetChecksum(filled, blanks)

	fmt.Printf("Answer Part 1: %d\n", checksum)
}
