package main

import (
	"fmt"
	"strconv"
)

func DoOperatorWithConcat(current int, number int, op string) int {
	if op == "+" {
		return current + number
	} else if op == "*" {
		return current * number
	} else if op == "||" {
		// Do concatenation of number
		concatStr := fmt.Sprintf("%d%d", current, number)
		concat, err := strconv.Atoi(concatStr)
		if err != nil {
			fmt.Println("Error when concatenating", current, "and", number)
			fmt.Println("Got string", concatStr, ".")
			panic("Error when concatenating numbers.")
		}
		return concat
	} else {
		panic("Invalid operator.")
	}
}

func TestOperatorsWithConcat(answer int, numbers []int, ops []string) ([]string, bool) {
	// Tests the operator configurations
	// Returns:
	// []string: the combination of operators, nil if not possible.
	// bool: true if the answer can be formed, false otherwise

	// Generate all possible arrangements.
	arrangements := SampleOpArrangements(ops, len(numbers)-1)

	// Exhaustive search through all arrangements.
	for _, arrangement := range arrangements {
		// Evaluate left to right,
		// i.e. the candidate answer is starts with the first element.
		candidate := numbers[0]
		for i, op := range arrangement {
			candidate = DoOperatorWithConcat(candidate, numbers[i+1], op)
		}
		if candidate == answer {
			return arrangement, true
		}
	}

	return nil, false
}

func testSolveWithConcat() {
	const separator string = ":"

	input, err := getDownloadedFile("test.txt")
	if err != nil {
		panic(err)
	}

	// Here, lines are in the format of [][]string
	lines, err := GetLines(input)
	if err != nil {
		panic("Problem parsing file")
	}
	answers, numbers := ParseLines(lines, separator)

	sumOfAnswers := 0

	for i := range lines {
		_, success := TestOperatorsWithConcat(answers[i], numbers[i], []string{"+", "*", "||"})
		if success {
			sumOfAnswers += answers[i]
		}
	}

	if sumOfAnswers != 11387 {
		fmt.Println("Got sum of answers", sumOfAnswers, ", expected", 3749)
		panic("Failed test")
	}
}

func MainPart2() {
	testSolveWithConcat()

	// Parse input
	const separator string = ":"

	lines := GetInputs()
	answers, numbers := ParseLines(lines, separator)

	// Compute the sum of valid answers
	sumOfAnswers := 0

	for i := range lines {
		_, success := TestOperatorsWithConcat(answers[i], numbers[i], []string{"+", "*", "||"})
		if success {
			// fmt.Println(numbers[i], "is valid with", arrangement, ", producing answer", answers[i])
			sumOfAnswers += answers[i]
		}
	}

	fmt.Printf("Answer Part 2: %d\n", sumOfAnswers)
}
