package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IntPow(a int, b int) int {
	// a^b
	if b == 0 {
		return 1
	}

	if b == 1 {
		return a
	}

	result := a
	for i := 2; i <= b; i++ {
		result *= b
	}
	return result
}

func castTo3Bit(a int) [3]bool {
	if a == 0 {
		return [3]bool{false, false, false}
	} else if a == 1 {
		return [3]bool{false, false, true}
	} else if a == 2 {
		return [3]bool{false, true, false}
	} else if a == 3 {
		return [3]bool{false, true, true}
	} else if a == 4 {
		return [3]bool{true, false, false}
	} else if a == 5 {
		return [3]bool{true, false, true}
	} else if a == 6 {
		return [3]bool{true, true, false}
	} else if a == 7 {
		return [3]bool{true, true, true}
	} else {
		panic(fmt.Errorf("Invalid 3-bit number %d", a))
	}
}

func castFrom3Bit(a [3]bool) int {
	if a == [3]bool{false, false, false} {
		return 0
	} else if a == [3]bool{false, false, true} {
		return 1
	} else if a == [3]bool{false, true, false} {
		return 2
	} else if a == [3]bool{false, true, true} {
		return 3
	} else if a == [3]bool{true, false, false} {
		return 4
	} else if a == [3]bool{true, false, true} {
		return 5
	} else if a == [3]bool{true, true, false} {
		return 6
	} else if a == [3]bool{true, true, true} {
		return 7
	} else {
		panic(fmt.Errorf("Invalid 3-bit number %v", a))
	}
}

func xor3bit(a [3]bool, b [3]bool) [3]bool {
	out := [3]bool{false, false, false}
	for i := range a {
		out[i] = (a[i] || b[i]) && !(a[i] && b[i])
	}
	return out
}

type Program struct {
	A       int
	B       int
	C       int
	Pointer int
	Output  []int
}

func (program Program) operandRules(operand int) int {
	if operand == 0 || operand == 1 || operand == 2 || operand == 3 {
		return operand
	} else if operand == 4 {
		return program.A
	} else if operand == 5 {
		return program.B
	} else if operand == 6 {
		return program.C
	} else {
		panic(fmt.Errorf("Invalid operand %d", operand))
	}
}

func (program *Program) regDiv(operand int) int {
	// Divide
	denominator := IntPow(2, operand)
	currentVal := program.A
	return currentVal / denominator
}

func (program *Program) op0advA(operand int) {
	operand = program.operandRules(operand)
	program.A = program.regDiv(operand)

	program.Pointer += 2
}

func (program *Program) op1bxl(operand int) {
	// operand = program.operandRules(operand)
	// op3bit := castTo3Bit(operand)
	// b3bit := castTo3Bit(program.B)

	// program.B = castFrom3Bit(xor3bit(b3bit, op3bit))
	program.B = program.B ^ operand
	program.Pointer += 2
}

func (program *Program) op2bst(operand int) {
	operand = program.operandRules(operand)
	program.B = operand % 8
	program.Pointer += 2
}

func (program *Program) op3jnz(operand int) {
	// operand = program.operandRules(operand)
	if program.A == 0 {
		program.Pointer += 2
	} else {
		program.Pointer = operand
	}
}

func (program *Program) op4bxc(operand int) {
	// operand = program.operandRules(operand)
	// b3bit := castTo3Bit(program.B)
	// c3bit := castTo3Bit(program.C)
	//
	// program.B = castFrom3Bit(xor3bit(b3bit, c3bit))
	program.B = program.B ^ program.C
	program.Pointer += 2
}

func (program *Program) op5out(operand int) {
	operand = program.operandRules(operand)
	program.Output = append(program.Output, operand%8)
	program.Pointer += 2
}

func (program *Program) op6bdv(operand int) {
	operand = program.operandRules(operand)
	program.B = program.regDiv(operand)
	program.Pointer += 2
}

func (program *Program) op7cdv(operand int) {
	operand = program.operandRules(operand)
	program.C = program.regDiv(operand)
	program.Pointer += 2
}

func (program *Program) executeInstruction(instructionVal int, operand int) {
	if instructionVal == 0 {
		program.op0advA(operand)
	} else if instructionVal == 1 {
		program.op1bxl(operand)
	} else if instructionVal == 2 {
		program.op2bst(operand)
	} else if instructionVal == 3 {
		program.op3jnz(operand)
	} else if instructionVal == 4 {
		program.op4bxc(operand)
	} else if instructionVal == 5 {
		program.op5out(operand)
	} else if instructionVal == 6 {
		program.op6bdv(operand)
	} else if instructionVal == 7 {
		program.op7cdv(operand)
	} else {
		panic(fmt.Errorf("Invalid instruction %d", instructionVal))
	}
}

func (program *Program) executeInstructions(instructions []int) {
	for program.Pointer < len(instructions) {
		instructionOp := instructions[program.Pointer]
		operand := instructions[program.Pointer+1]
		fmt.Println("Pointer:", program.Pointer, "Opcode:", instructionOp, "Operand:", operand)
		program.executeInstruction(instructionOp, operand)
		fmt.Println(program)
	}
}

func (program *Program) stringifyOut() string {
	var sb strings.Builder

	for _, val := range program.Output {
		fmt.Fprintf(&sb, "%d,", val)
	}

	outputString := sb.String()
	outputString = strings.TrimSuffix(outputString, ",")

	return outputString
}

func parseInput(lines []string) (Program, []int) {
	program := Program{}
	instructions := make([]int, 0)

	for _, line := range lines {
		if strings.Contains(line, "Register A:") {
			parts := strings.Split(line, "Register A: ")
			intForm, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Errorf("Error casting %s to integer", parts[1]))
			}
			program.A = intForm
		} else if strings.Contains(line, "Register B:") {
			parts := strings.Split(line, "Register B: ")
			intForm, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Errorf("Error casting %s to integer", parts[1]))
			}
			program.B = intForm
		} else if strings.Contains(line, "Register C:") {
			parts := strings.Split(line, "Register C: ")
			intForm, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Errorf("Error casting %s to integer", parts[1]))
			}
			program.C = intForm
		} else if strings.Contains(line, "Program:") {
			parts := strings.Split(line, "Program: ")
			intStrings := strings.Split(parts[1], ",")
			for _, intString := range intStrings {
				intForm, err := strconv.Atoi(intString)
				if err != nil {
					panic(fmt.Errorf("Error casting %s to integer", intString))
				}
				instructions = append(instructions, intForm)
			}
		}
	}

	return program, instructions
}

func testSolve(inputFile string, expected string) {
	// Parse the input file into a 2D array.
	rawInput, err := getDownloadedFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}
	lines, err := GetLines(rawInput)
	if err != nil {
		panic(fmt.Errorf("Error parsing file %s.", inputFile))
	}

	program, instructions := parseInput(lines)
	fmt.Println(program)
	fmt.Println(instructions)

	program.executeInstructions(instructions)
	output := program.stringifyOut()

	if output != expected {
		panic(fmt.Errorf("Expected %s got %s", expected, output))
	}
}

func runBasicTests() {
	program := Program{
		C: 9,
	}
	program.executeInstructions([]int{2, 6})
	if program.B != 1 {
		panic("Failed test 1.")
	}

	program = Program{
		A: 10,
	}
	program.executeInstructions([]int{5, 0, 5, 1, 5, 4})
	if program.stringifyOut() != "0,1,2" {
		panic("Failed test 2.")
	}

	program = Program{
		A: 2024,
	}
	program.executeInstructions([]int{0, 1, 5, 4, 3, 0})
	if program.stringifyOut() != "4,2,5,6,7,7,7,7,3,1,0" {
		panic("Failed test 3.")
	}
	if program.A != 0 {
		panic("Failed test 3.")
	}

	program = Program{
		B: 29,
	}
	program.executeInstructions([]int{1, 7})
	if program.B != 26 {
		panic("Failed test 4.")
	}

	program = Program{
		B: 2024,
		C: 43690,
	}
	program.executeInstructions([]int{4, 0})
	if program.B != 44354 {
		panic("Failed test 5.")
	}
}

func MainPart1() {
	runBasicTests()

	testSolve("test.txt", "4,6,3,5,6,3,5,2,1,0")
	panic("Check tests")
	lines := GetInputs()
	fmt.Println(len(lines), "in input.")

	program, instructions := parseInput(lines)
	fmt.Println(program)
	fmt.Println(instructions)

	program.executeInstructions(instructions)
	output := program.stringifyOut()

	fmt.Printf("Answer Part 1: %s\n", output)
}
