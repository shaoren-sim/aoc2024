package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ParseRule(condition string, conditionSplitter string) []int {
	conditionStrings := strings.Split(condition, conditionSplitter)

	// var ruleArr []int
	ruleArr := make([]int, len(conditionStrings))

	for j, conditionString := range conditionStrings {
		conditionInt, err := strconv.Atoi(conditionString)
		if err != nil {
			panic(err)
		}
		ruleArr[j] = conditionInt
	}
	return ruleArr
}

func ParseRules(conditions []string, conditionSplitter string) [][]int {
	rules := make([][]int, len(conditions))

	for i, condition := range conditions {
		rules[i] = ParseRule(condition, conditionSplitter)
	}
	return rules
}

func ParseResult(result string, resultSplitter string) []int {
	// Split the result into parts
	resultParts := strings.Split(result, resultSplitter)

	parsedResult := make([]int, len(resultParts))

	for i, resultPart := range resultParts {
		// Cast to integer.
		resultPartInt, err := strconv.Atoi(resultPart)
		if err != nil {
			panic(err)
		}
		parsedResult[i] = resultPartInt
	}

	return parsedResult
}

func RulesToMap(rules [][]int, reverse bool) map[int][]int {
	adjacencyMap := make(map[int][]int)

	for _, rule := range rules {
		// Reverse if requested
		if reverse {
			slices.Reverse(rule)
		}
		root := rule[0]

		// if the root node already does not exist in the adjacencyList.
		// Create the []int slice
		if _, ok := adjacencyMap[root]; !ok {
			adjacencyMap[root] = make([]int, 0)
		}
		for j := 1; j < len(rule); j++ {
			adjacencyMap[root] = append(adjacencyMap[root], rule[j])
		}
	}

	return adjacencyMap
}

func ResultIsValid(resultParts []int, ruleMap map[int][]int) bool {
	// Checking forward rules.
	for i, page := range resultParts[:len(resultParts)-1] {
		// fmt.Println("Checking page", page)
		// If the part is in the rule adjacency list, do comparisons.
		if conditions, ok := ruleMap[page]; ok {
			// fmt.Println("Page", page, "is in rule tree.")
			// fmt.Println("Conditions for page", page, ":", conditions)
			// Loop through the remaining pages.
			// Ensure that any pages that have a rule are after.
			for _, nextPage := range resultParts[i+1:] {
				if !slices.Contains(conditions, nextPage) {
					return false
				}
			}
		}
	}

	return true
}

func GetMiddleValue(result []int) int {
	if len(result)%2 != 1 {
		panic("To get the middle value, input must have an odd number length.")
	}

	midpoint := len(result) / 2

	return result[midpoint]
}

func MainPart1() {
	// Constants for parsing.
	ruleSplitter := "|"
	resultSplitter := ","

	conditions, results := GetInputs()
	fmt.Println(len(conditions), "condition lines from input.")
	fmt.Println(len(results), "result lines from input")

	// Parse the rules into [2]int arrays
	rules := ParseRules(conditions, ruleSplitter)
	// Convert the rules into a "tree" for easier execution.
	forwardRuleMap := RulesToMap(rules, false)
	reverseRuleMap := RulesToMap(rules, true)

	sum := 0

	for _, result := range results {
		resultParts := ParseResult(result, resultSplitter)
		// Forward check
		if !ResultIsValid(resultParts, forwardRuleMap) {
			// fmt.Println(result, "breaks the forward rules.")
			continue
		}
		// fmt.Println(result, "passes all forward rules.")

		// Reverse check
		slices.Reverse(resultParts)
		if !ResultIsValid(resultParts, reverseRuleMap) {
			// fmt.Println(result, "breaks the reverse rules.")
			continue
		}
		// fmt.Println(result, "passes all reverse rules.")

		sum += GetMiddleValue(resultParts)
	}

	fmt.Printf("Answer Part 1: %d\n", sum)
}
