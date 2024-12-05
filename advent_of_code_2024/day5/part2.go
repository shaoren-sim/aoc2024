package main

import (
	"fmt"
	"slices"
)

func GetElementDepthMap(s []int, ruleTree map[int][]int) map[int]int {
	depthMap := make(map[int]int)

	for _, el := range s {
		depthMap[el] = 0
	}

	for _, startEl := range s {
		currentEl := startEl
		for _, el := range s[:] {
			if !slices.Contains(ruleTree[currentEl], el) {
				// fmt.Println("Element", el, "does not belongs after element", currentEl, ".")
			} else {
				// fmt.Println("Element", el, "belongs after element", currentEl, ". Adding to depth.")
				depthMap[el] += 1
			}
		}
	}
	return depthMap
}

func hasDupeValues(m map[int]int) bool {
	// Check if a map has duplicate values.
	// Function from https://stackoverflow.com/a/57237165
	x := make(map[int]struct{})

	for _, v := range m {
		if _, has := x[v]; has {
			return true
		}
		x[v] = struct{}{}
	}

	return false
}

func RearrangeByRules(s []int, ruleTree map[int][]int) []int {
	rearrangedSlice := make([]int, len(s))

	// For each value, identify it's position.
	depthMap := GetElementDepthMap(s, ruleTree)
	// fmt.Println(depthMap)

	// Test to verify that all values in the map are unique.
	for num, depth := range depthMap {
		rearrangedSlice[depth] = num
	}

	return rearrangedSlice
}

func testDepthOrdering() {
	// Constants for parsing.
	ruleSplitter := "|"
	resultSplitter := ","

	truths := make([][]int, 3)
	truths[0] = []int{97, 75, 47, 61, 53}
	truths[1] = []int{61, 29, 13}
	truths[2] = []int{97, 75, 47, 29, 13}

	input, err := getDownloadedFile("test.txt")
	if err != nil {
		panic(err)
	}
	conditions, results, err := GetConditionsResults(input, "")

	// Parse the rules into [2]int arrays
	rules := ParseRules(conditions, ruleSplitter)
	// Convert the rules into a "tree" for easier execution.
	forwardRuleTree := RulesToTree(rules, false)
	reverseRuleTree := RulesToTree(rules, true)

	// sum := 0
	errorInd := 0
	for _, result := range results {
		resultParts := ParseResult(result, resultSplitter)
		reverseResultParts := append(resultParts[:0:0], resultParts...)
		slices.Reverse(reverseResultParts)

		// For part 2, skip anything that passes both rules.
		if ResultIsValid(resultParts, forwardRuleTree) && ResultIsValid(reverseResultParts, reverseRuleTree) {
			// fmt.Println(resultParts, "passes all tests.")
			continue
		}

		rearrangedParts := RearrangeByRules(resultParts, forwardRuleTree)
		for i, part := range rearrangedParts {
			if part != truths[errorInd][i] {
				fmt.Println("Failed test for input", resultParts)
				fmt.Println("Got", rearrangedParts)
				fmt.Println("Expected", truths[errorInd])
				panic("Failed test")
			}
		}
		errorInd += 1
	}

}

func MainPart2() {
	testDepthOrdering()

	// Constants for parsing.
	ruleSplitter := "|"
	resultSplitter := ","

	conditions, results := GetInputs()

	// Parse the rules into [2]int arrays
	rules := ParseRules(conditions, ruleSplitter)
	// Convert the rules into a "tree" for easier execution.
	forwardRuleTree := RulesToTree(rules, false)
	reverseRuleTree := RulesToTree(rules, true)

	sum := 0

	for _, result := range results {
		resultParts := ParseResult(result, resultSplitter)
		reverseResultParts := append(resultParts[:0:0], resultParts...)
		slices.Reverse(reverseResultParts)

		// For part 2, skip anything that passes both rules.
		if ResultIsValid(resultParts, forwardRuleTree) || ResultIsValid(reverseResultParts, reverseRuleTree) {
			continue
		}

		rearrangedParts := RearrangeByRules(resultParts, forwardRuleTree)
		sum += GetMiddleValue(rearrangedParts)
	}

	fmt.Printf("Answer Part 2: %d\n", sum)
}
