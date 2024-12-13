package main

import (
	"fmt"
)

func Solve2x2EqnPart2(lhs [][]float64, rhs []float64) []float64 {
	// answer := make([]float64, 2)
	if len(lhs) != 2 && len(lhs[0]) != 2 {
		panic("Solve2x2Eqn only works for 2x2 left-hand-side inputs.")
	}
	if len(rhs) != 2 {
		panic("Solve2x2Eqn only works for 1x2 right-hand-side inputs (x, y).")
	}

	// Primary School Algorithm.
	// Set the first value to be the same scale as the 2nd.
	// Cloning to not affect the underlying equation.
	eqnClone := append(lhs[:0:0], lhs...)
	lineq1 := eqnClone[0]
	lineq2 := eqnClone[1]

	// Try deriving a new line
	// Assuming
	// AX+BY = I
	// CX+DY = J
	// Remove the large value of 10000000000 as N.
	// AX + BY = I + N		-- (1)
	// CX + DY = J + N		-- (2)
	// This means we can do (3) = (1) - (2) to get a tractable solution.
	// The problem is that even with 3, we will still need to use either (1) or (2)
	lineq3 := append(lineq1[:0:0], lineq1...)
	for i, val := range lineq3 {
		lineq3[i] = val - lineq2[i]
	}
	rhsq3 := rhs[0] - rhs[1]
	multBy := lineq2[0] / lineq3[0]

	newLineQ3 := append(lineq3[:0:0], lineq3...)
	for i, val := range newLineQ3 {
		newLineQ3[i] = val * multBy
	}
	// Solve for value at index 1.
	ans1 := (rhsq3*multBy - rhs[1]) / (newLineQ3[1] - lineq2[1])
	// Solve for value at index 0.
	ans0 := (rhs[0] - newLineQ3[1]*ans1) / newLineQ3[0]

	return []float64{ans0, ans1}
}

func SolveEquationPart2(buttonConditions map[string][]int, prizePosition []int, buttonOrder []string) map[string]float64 {
	// Store the solved answer of number of presses.
	buttonPresses := make(map[string]float64)

	// Build the equation.
	lhs := make([][]float64, len(buttonConditions))
	for i := range len(buttonConditions) {
		eqn := make([]float64, 0)
		for _, buttonName := range buttonOrder {
			eqn = append(eqn, float64(buttonConditions[buttonName][i]))
		}
		// for _, conditions := range buttonConditions {
		// 	eqn = append(eqn, conditions[i])
		// }
		lhs[i] = eqn
	}

	// Cast prizePosition to float as well.
	prizePositionFloat := make([]float64, len(prizePosition))
	for i, pos := range prizePosition {
		prizePositionFloat[i] = float64(pos)
	}

	// Solve the equation.
	fmt.Println(lhs)
	fmt.Println(buttonOrder)
	fmt.Println(prizePosition)
	ans := Solve2x2EqnPart2(lhs, prizePositionFloat)
	fmt.Println(ans)

	// Fill up the output dict.
	for i, buttonName := range buttonOrder {
		buttonPresses[buttonName] = ans[i]
	}
	return buttonPresses
}
func SolvePart2(lines []string, intCheckThreshold float64, tokenCosts map[string]int) int {
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
		fmt.Println("presses for buttons", buttonOrder, buttonPresses)
		// Filter to see if the button presses are valid.
		// Condition 2: No such thing as floating point integer presses.
		if !closeToInt(buttonPresses, intCheckThreshold) {
			continue
		}
		panic("BReak")

		// If presses pass all conditions, cast to int and calc the number of tokens.
		tokens += CountTokens(buttonPresses, tokenCosts)
	}

	return tokens
}

func CountTokensPart2(buttonPresses map[string]int, tokenCosts map[string]int) int {
	tokens := 0
	for buttonName, presses := range buttonPresses {
		tokens += presses * tokenCosts[buttonName]
	}
	return tokens
}

func testIntIdea(buttonConditions map[string][]int, rhs []int, buttonPresses map[string]float64, buttonOrder []string) (bool, map[string]int) {
	const searchIntRange int = 10000
	// Store the solved answer of number of presses.
	// buttonPressesInt := make(map[string]int)

	// Build the equation.
	lhs := make([][]int, len(buttonConditions))
	for i := range len(buttonConditions) {
		eqn := make([]int, 0)
		for _, buttonName := range buttonOrder {
			eqn = append(eqn, buttonConditions[buttonName][i])
		}
		// for _, conditions := range buttonConditions {
		// 	eqn = append(eqn, conditions[i])
		// }
		lhs[i] = eqn
	}

	// Try deriving a new line
	// Assuming
	// AX+BY = I
	// CX+DY = J
	// Remove the large value of 10000000000 as N.
	// AX + BY = I + N		-- (1)
	// CX + DY = J + N		-- (2)
	// This means we can do (3) = (1) - (2) to get a tractable solution.
	// Derive eqn (3)
	lineq1 := lhs[0]
	lineq2 := lhs[1]

	// Try deriving a new line
	// Assuming
	// AX+BY = I
	// CX+DY = J
	// Remove the large value of 10000000000 as N.
	// AX + BY = I + N		-- (1)
	// CX + DY = J + N		-- (2)
	// This means we can do (3) = (1) - (2) to get a tractable solution.
	// The problem is that even with 3, we will still need to use either (1) or (2)
	lineq3 := append(lineq1[:0:0], lineq1...)
	for i, val := range lineq3 {
		lineq3[i] = val - lineq2[i]
	}
	rhsq3 := rhs[0] - rhs[1]
	fmt.Println(lineq3, rhsq3)

	// Loop through values and test for integer precision.
	// We use the tractable solution to test using ints.
	// fmt.Println(buttonPresses)
	// for buttonName, presses := range buttonPresses {
	// 	fmt.Println(buttonName, int(presses))
	// }
	buttonPressesInt := make(map[string]int)

	for _, buttonName := range buttonOrder {
		floatVal := buttonPresses[buttonName]
		intVal := int(floatVal)
		buttonPressesInt[buttonName] = intVal
	}

	fmt.Println("Casted to ints:", buttonPressesInt)
	rhsCheck := lineq3[0]*buttonPressesInt[buttonOrder[0]] + lineq3[1]*buttonPressesInt[buttonOrder[1]]
	if rhsCheck == rhsq3 {
		return true, buttonPressesInt
	}
	return false, buttonPressesInt
}
func testSolvePart2(inputFile string) {
	const maxButtonPresses float64 = 100
	tokenCosts := map[string]int{"A": 3, "B": 1}

	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}

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
		fmt.Println("presses for buttons", buttonOrder, buttonPresses)
		works, buttonPressesInt := testIntIdea(condition, prizePosition, buttonPresses, buttonOrder)

		// If presses pass all conditions, cast to int and calc the number of tokens.
		if works {
			fmt.Println("Works for", i)
			tokens += CountTokensPart2(buttonPressesInt, tokenCosts)
		}
	}
}

func MainPart2() {
	testSolvePart2("test.txt")

	const intCheckThreshold float64 = 0.00001
	tokenCosts := map[string]int{"A": 3, "B": 1}

	lines := GetInputs()
	tokens := SolvePart2(lines, intCheckThreshold, tokenCosts)

	fmt.Printf("Answer Part 2: %d\n", int(tokens))
}
