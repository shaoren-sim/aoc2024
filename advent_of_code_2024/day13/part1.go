package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetButtonModifiers(line string) []int {
	modifiers := make([]int, 0)

	modiferParts := strings.Split(line, "+")
	// Part 1:
	modiferString1 := strings.Split(modiferParts[1], ",")[0]
	mod1, err := strconv.Atoi(modiferString1)
	if err != nil {
		panic(fmt.Errorf("Error parsing %s", modiferString1))
	}
	modifiers = append(modifiers, mod1)

	// Part 2:
	mod2, err := strconv.Atoi(modiferParts[2])
	if err != nil {
		panic(fmt.Errorf("Error parsing %s", modiferParts[2]))
	}
	modifiers = append(modifiers, mod2)
	return modifiers
}

func GetPrizePosition(line string) []int {
	position := make([]int, 2)
	positionParts := strings.Split(line, "=")
	// Part 1:
	positionString1 := strings.Split(positionParts[1], ",")[0]
	pos1, err := strconv.Atoi(positionString1)
	if err != nil {
		panic(fmt.Errorf("Error parsing %s", positionString1))
	}
	position[0] = pos1

	// Part 2:
	pos2, err := strconv.Atoi(positionParts[2])
	if err != nil {
		panic(fmt.Errorf("Error parsing %s", positionParts[2]))
	}
	position[1] = pos2
	return position
}

func ParseEquations(lines []string) ([]map[string][]int, [][]int) {
	buttonConditions := make([]map[string][]int, 0)
	prizePositions := make([][]int, 0)

	// Disgusting parsing code.
	buttonCondition := make(map[string][]int)

	for _, line := range lines {
		// Skip empty lines
		if line == "\n" {
			continue
		}

		// If the line describes the buttons, append to the map.
		if strings.Contains(line, "Button ") {
			// Parse out the button name.
			buttonName := string(line[7])

			// Parse out the button conditions.
			modifiers := GetButtonModifiers(line)

			// Add the button to the map.
			buttonCondition[buttonName] = modifiers
		}

		// If the line describes the prize, assign the position.
		if strings.Contains(line, "Prize: ") {
			// Append to the full map and list.
			buttonConditions = append(buttonConditions, buttonCondition)
			prizePositions = append(prizePositions, GetPrizePosition(line))

			// Then, reset the map
			buttonCondition = make(map[string][]int)
		}
	}

	// Before return, sanity check.
	if len(buttonConditions) != len(prizePositions) {
		panic("Length of conditions (LHS) does not match prize positions (RHS).")
	}

	return buttonConditions, prizePositions
}

func Solve2x2Eqn(lhs [][]float64, rhs []float64) []float64 {
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
	multBy := lineq2[0] / lineq1[0]

	newLineQ1 := append(lineq1[:0:0], lineq1...)
	for i, val := range newLineQ1 {
		newLineQ1[i] = val * multBy
	}
	// Solve for value at index 1.
	ans1 := (rhs[0]*multBy - rhs[1]) / (newLineQ1[1] - lineq2[1])
	// Solve for value at index 0.
	ans0 := (rhs[0] - lineq1[1]*ans1) / lineq1[0]

	return []float64{ans0, ans1}
}

func SolveEquation(buttonConditions map[string][]int, prizePosition []int, buttonOrder []string) map[string]float64 {
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
	ans := Solve2x2Eqn(lhs, prizePositionFloat)
	fmt.Println(ans)

	// Fill up the output dict.
	for i, buttonName := range buttonOrder {
		buttonPresses[buttonName] = ans[i]
	}
	return buttonPresses
}

func belowMaxThreshold(buttonPresses map[string]float64, threshold float64) bool {
	for _, presses := range buttonPresses {
		if presses > threshold {
			fmt.Println("Button presses exceed threshold.")
			return false
		}
	}
	return true
}

func closeToInt(buttonPresses map[string]float64, threshold float64) bool {
	for _, presses := range buttonPresses {
		roundInt := math.Round(presses)
		if math.Abs(presses-roundInt) > threshold {
			fmt.Println("Likely not an integer.")
			return false
		}
	}
	return true
}

func CountTokens(buttonPresses map[string]float64, tokenCosts map[string]int) int {
	tokens := 0
	for buttonName, presses := range buttonPresses {
		tokens += int(math.Round(presses)) * tokenCosts[buttonName]
	}
	return tokens
}

func Solve(lines []string, maxButtonPresses float64, intCheckThreshold float64, tokenCosts map[string]int) int {
	buttonConditions, prizePositions := ParseEquations(lines)

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
		// Condition 1: no button presses beyond 100.
		if !belowMaxThreshold(buttonPresses, maxButtonPresses) {
			continue
		}
		// Condition 2: No such thing as floating point integer presses.
		if !closeToInt(buttonPresses, intCheckThreshold) {
			continue
		}

		// If presses pass all conditions, cast to int and calc the number of tokens.
		tokens += CountTokens(buttonPresses, tokenCosts)
	}

	return tokens
}

func testSolve(inputFile string, expected int) {
	const maxButtonPresses float64 = 100
	const intCheckThreshold float64 = 0.00001
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

	tokens := Solve(lines, maxButtonPresses, intCheckThreshold, tokenCosts)

	if tokens != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, tokens))
	}
}

func MainPart1() {
	const maxButtonPresses float64 = 100
	const intCheckThreshold float64 = 0.00001
	tokenCosts := map[string]int{"A": 3, "B": 1}

	testSolve("test.txt", 480)
	lines := GetInputs()
	tokens := Solve(lines, maxButtonPresses, intCheckThreshold, tokenCosts)

	fmt.Printf("Answer Part 1: %d\n", tokens)
}
