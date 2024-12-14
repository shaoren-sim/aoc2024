package main

import (
	"fmt"
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

func Solve2x2Eqn(lhs [][]int, rhs []int) ([]int, bool) {
	// Primary School Algorithm.
	// Initial implementation uses divides and floats, but that fails part2.
	// Reimplement only using ints.
	// Int validity is returned as a boolean: false if divide has remaineder.
	if len(lhs) != 2 && len(lhs[0]) != 2 {
		panic("Solve2x2Eqn only works for 2x2 left-hand-side inputs.")
	}
	if len(rhs) != 2 {
		panic("Solve2x2Eqn only works for 1x2 right-hand-side inputs (x, y).")
	}

	// Set the first value to be the same scale as the 2nd.
	// Cloning to not affect the underlying equation.
	eqnClone := append(lhs[:0:0], lhs...)
	lineq1 := eqnClone[0]
	lineq2 := eqnClone[1]

	newLineq1 := append(lineq1[:0:0], lineq1...)
	for i, val := range newLineq1 {
		newLineq1[i] = val * lineq2[0]
	}

	newLineq2 := append(lineq2[:0:0], lineq2...)
	for i, val := range newLineq2 {
		newLineq2[i] = val * lineq1[0]
	}

	// Early break condition to test if the int is valid
	newRhs1 := rhs[0]*lineq2[0] - rhs[1]*lineq1[0]
	if newRhs1%(newLineq1[1]-newLineq2[1]) != 0 {
		return nil, false
	}

	// Solve for value at index 1.
	ans1 := (rhs[0]*lineq2[0] - rhs[1]*lineq1[0]) / (newLineq1[1] - newLineq2[1])

	// Yet another early break condition, check if the index 0 value is a valid int.
	newRhs2 := rhs[0] - lineq1[1]*ans1
	if newRhs2%lineq1[0] != 0 {
		return nil, false
	}

	// Solve for value at index 0.
	ans0 := (rhs[0] - lineq1[1]*ans1) / lineq1[0]

	return []int{ans0, ans1}, true
}

func SolveEquation(buttonConditions map[string][]int, prizePosition []int, buttonOrder []string) map[string]int {
	// Store the solved answer of number of presses.
	buttonPresses := make(map[string]int)

	// Build the equation.
	lhs := make([][]int, len(buttonConditions))
	for i := range len(buttonConditions) {
		eqn := make([]int, 0)
		for _, buttonName := range buttonOrder {
			eqn = append(eqn, buttonConditions[buttonName][i])
		}
		lhs[i] = eqn
	}

	// Solve the equation.
	ans, isValidInt := Solve2x2Eqn(lhs, prizePosition)

	// Early return if solving gets an invalid int.
	if !isValidInt {
		return nil
	}

	// Fill up the output dict.
	for i, buttonName := range buttonOrder {
		buttonPresses[buttonName] = ans[i]
	}
	return buttonPresses
}

func belowMaxThreshold(buttonPresses map[string]int, threshold int) bool {
	for _, presses := range buttonPresses {
		if presses > threshold {
			// fmt.Println("Button presses exceed threshold.")
			return false
		}
	}
	return true
}

func CountTokens(buttonPresses map[string]int, tokenCosts map[string]int) int {
	tokens := 0
	for buttonName, presses := range buttonPresses {
		tokens += presses * tokenCosts[buttonName]
	}
	return tokens
}

func Solve(lines []string, maxButtonPresses int, tokenCosts map[string]int) int {
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

		// Early break condition if solving obtained invalid int.
		if buttonPresses == nil {
			continue
		}

		// Filter to see if the button presses are valid.
		// Condition 1: no button presses beyond 100.
		if !belowMaxThreshold(buttonPresses, maxButtonPresses) {
			continue
		}

		// If presses pass all conditions, cast to int and calc the number of tokens.
		tokens += CountTokens(buttonPresses, tokenCosts)
	}

	return tokens
}

func testSolve(inputFile string, expected int) {
	const maxButtonPresses int = 100
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

	tokens := Solve(lines, maxButtonPresses, tokenCosts)
	fmt.Println(tokens)

	if tokens != expected {
		panic(fmt.Errorf("Expected %d, got %d", expected, tokens))
	}
}

func MainPart1() {
	const maxButtonPresses int = 100
	tokenCosts := map[string]int{"A": 3, "B": 1}

	testSolve("test.txt", 480)
	lines := GetInputs()
	tokens := Solve(lines, maxButtonPresses, tokenCosts)

	fmt.Printf("Answer Part 1: %d\n", tokens)
}
