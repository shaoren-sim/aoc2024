package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseLines(lines []string, sep string) ([]int, [][]int) {
	// Parse lines from input.
	// Returns:
	// []int: Answer
	// [][]int: Slice of numbers.

	answers := make([]int, len(lines))
	numbers := make([][]int, len(lines))

	for i, line := range lines {
		// Split by the ":" separator.
		lineParts := strings.Split(line, sep)
		if len(lineParts) != 2 {
			fmt.Println("Problem parsing line", i, ":", line)
			fmt.Println("Expected single ':' as a splitter.")
			panic("Error in line parsing.")
		}

		// Parse the LHS, i.e. the answer.
		answer, err := strconv.Atoi(lineParts[0])
		if err != nil {
			fmt.Println("Error parsing LHS (Answer) for line", line)
			panic(err)
		}
		answers[i] = answer

		// Parse the RHS, i.e. the numbers.
		numbersString := strings.TrimSpace(lineParts[1])
		numbersStringParts := strings.Split(numbersString, " ")
		numberParts := make([]int, len(numbersStringParts))

		for j, numberStringPart := range numbersStringParts {
			numberPart, err := strconv.Atoi(numberStringPart)
			if err != nil {
				fmt.Println("Error parsing RHS (numbers) for line", line)
				panic(err)
			}
			numberParts[j] = numberPart
		}
		numbers[i] = numberParts

	}
	return answers, numbers
}

func DoOperator(current int, number int, op string) int {
	if op == "+" {
		return current + number
	} else if op == "*" {
		return current * number
	} else if op == "-" {
		return current - number
	} else if op == "/" {
		return current / number
	} else {
		panic("Invalid operator.")
	}
}

func PowerOf(a int, b int) int {
	// Wrapper to get the power of for integers
	return int(math.Pow(float64(a), float64(b)))
}

func SampleOpArrangements(ops []string, num int) [][]string {
	// Generates all possible arrangements of the operators
	// Port of https://docs.python.org/3/library/itertools.html#itertools.product
	choices := len(ops)
	total := PowerOf(choices, num)
	arrangements := make([][]string, total)

	for i := range arrangements {
		arrangement := make([]string, num)

		// Generate the arrangement for this index `i`
		for j := 0; j < num; j++ {
			// For each position `j`, we map the current index to an operator
			arrangement[j] = ops[(i/(PowerOf(choices, num-j-1)))%choices]
		}

		arrangements[i] = arrangement
	}

	return arrangements
}

func TestOperators(answer int, numbers []int, ops []string) ([]string, bool) {
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
			candidate = DoOperator(candidate, numbers[i+1], op)
		}
		if candidate == answer {
			return arrangement, true
		}
	}

	return nil, false
}

func testSolve() {
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
		_, success := TestOperators(answers[i], numbers[i], []string{"+", "*"})
		if success {
			sumOfAnswers += answers[i]
		}
	}

	if sumOfAnswers != 3749 {
		fmt.Println("Got sum of answers", sumOfAnswers, ", expected", 3749)
		panic("Failed test")
	}
}

func MainPart1() {
	testSolve()

	// Parse input
	const separator string = ":"

	lines := GetInputs()
	answers, numbers := ParseLines(lines, separator)

	// Compute the sum of valid answers
	sumOfAnswers := 0

	for i := range lines {
		_, success := TestOperators(answers[i], numbers[i], []string{"+", "*"})
		if success {
			// fmt.Println(numbers[i], "is valid with", arrangement, ", producing answer", answers[i])
			sumOfAnswers += answers[i]
		}
	}

	fmt.Printf("Answer Part 1: %d\n", sumOfAnswers)
}
