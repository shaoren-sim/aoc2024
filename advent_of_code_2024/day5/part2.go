package main

import (
	"fmt"
	"slices"
)

// func RearrangeByRules(s []int, ruleTree map[int][]int) []int {
// 	rearrangedSlice := make([]int, len(s))
//
// 	// For each value, identify it's position.
// 	for _, page := range s {
//
// 	}
//
// }

func MainPart2() {
	// Constants for parsing.
	ruleSplitter := "|"
	resultSplitter := ","

	conditions, results := GetInputs()
	fmt.Println(len(conditions), "condition lines from input.")
	fmt.Println(len(results), "result lines from input")

	// Parse the rules into [2]int arrays
	rules := ParseRules(conditions, ruleSplitter)
	// Convert the rules into a "tree" for easier execution.
	forwardRuleTree := RulesToTree(rules, false)
	fmt.Println(forwardRuleTree)
	reverseRuleTree := RulesToTree(rules, true)
	fmt.Println(reverseRuleTree)

	sum := 0

	for _, result := range results {
		resultParts := ParseResult(result, resultSplitter)
		reverseResultParts := append(resultParts[:0:0], resultParts...)
		slices.Reverse(reverseResultParts)

		// For part 2, skip anything that passes both rules.
		if ResultIsValid(resultParts, forwardRuleTree) || ResultIsValid(reverseResultParts, reverseRuleTree) {
			continue
		}

		// rearrangedParts := RearrangeByRules(resultParts, forwardRuleTree)

		// // Forward check
		// if !ResultIsValid(resultParts, forwardRuleTree) {
		// 	fmt.Println(result, "breaks the forward rules.")
		// 	continue
		// }
		//
		// fmt.Println(result, "passes all forward rules.")
		// // Reverse check
		// slices.Reverse(resultParts)
		// if !ResultIsValid(resultParts, reverseRuleTree) {
		// 	fmt.Println(result, "breaks the reverse rules.")
		// 	continue
		// }
		// fmt.Println(result, "passes all reverse rules.")

		sum += GetMiddleValue(resultParts)
	}

	fmt.Printf("Answer Part 2: %d\n", sum)
}
