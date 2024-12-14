package main

import (
	"fmt"
)

func SolvePart2(lines []string, tokenCosts map[string]int) int {
	buttonConditions, prizePositions := ParseEquations(lines)

	// Shift prize positions by 10000000000000
	for i := range prizePositions {
		prizePosition := prizePositions[i]
		for j, position := range prizePosition {
			newPosition := position + 10000000000000
			prizePosition[j] = newPosition
		}
		prizePositions[i] = prizePosition
	}

	// In Go, maps do not have guaranteed order,
	// To ensure that we can retrack the number of button presses per button,
	// Pre-define the order here.
	buttonOrder := make([]string, 0)
	for buttonName := range buttonConditions[0] {
		buttonOrder = append(buttonOrder, buttonName)
	}

	tokens := 0

	// Solve the equations
	for i, condition := range buttonConditions {
		prizePosition := prizePositions[i]
		buttonPresses := SolveEquation(condition, prizePosition, buttonOrder)
		// fmt.Println("presses for buttons", buttonOrder, buttonPresses)

		// If presses pass all conditions, cast to int and calc the number of tokens.
		tokens += CountTokens(buttonPresses, tokenCosts)
	}

	return tokens
}

func MainPart2() {
	const intCheckThreshold float64 = 0.00001
	tokenCosts := map[string]int{"A": 3, "B": 1}

	lines := GetInputs()
	tokens := SolvePart2(lines, tokenCosts)

	fmt.Printf("Answer Part 2: %d\n", int(tokens))
}
