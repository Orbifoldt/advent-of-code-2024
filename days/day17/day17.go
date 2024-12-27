package day17

import (
	"advent-of-code-2024/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (string, error) {
	program, err := parseInput(useRealInput)
	if err != nil {
		return "", err
	}

	var out []int64
	if useRealInput {
		out = program.runOptimized()
	} else {
		out = program.run()
	}

	outputString := fmt.Sprintf("%d", out[0])
	for i := 1; i < len(out); i++ {
		outputString = fmt.Sprintf("%s,%d", outputString, out[i])
	}
	return outputString, nil
}

func SolvePart2(useRealInput bool) (int64, error) {
	program, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	// Some debugging/analysis shows:
	// - register A only ever decreases (we divide it), with minimum value of 0
	// - when register A then the program ends (we don't jump anymore)
	// - Only the last instruction is a JNZ instruction, and it always jumps to 0
	// - Only the second to last instruction is an OUT instruction
	// => If and only if A reaches 0 the program ends
	// => The program consists of cycles of fixed length
	// => Last action of a cycle is outputting a fixed combo opperand (register A in sample input, register B in my input)

	// Next, we try to find out exactly what a single cycle outputs
	// My input:
	// - 2,4	BST combo(4)	B=A%8				B = last 3 bits of A
	// - 1,1	BXL lit(1)  	B=B^1
	// - 7,5	CDV combo(5)	C=A/(2**B)
	// - 4,4	BXC ignored 	B=B^C
	// - 1,4	BXL lit(4)  	B=B^4
	// - 0,3	ADV combo(3)	A=A/(2**3)			A = bit shift A 3 to right
	// - 5,5	OUT combo(5)	output <- B%8
	// - 3,0	JNZ lit(0)  	Jump to 0 if A!=0
	// Notice how every cycle B and C get new values, so the output doesn't depend on B and C
	// Also, every cycle A get's bitshifted by exactly 3, it doesn't change in any other way
	// Furthermore, the output depends on the last 3 to 7 bits of A (C is set to A bitshifted max of 7)
	// => if the first output matches we no longer need to change the last 3 bits of A intial value
	// => we have 16 values in the program, so we need to consider A's up to 2**(3*16 + 4) or less (4 for the effect on C)
	// => Total values we need to check is way lower, only 16 times 2**7

	// For completenes, outputed value is:
	// output <- (((((A % 8) ^ 1) ^ (A / (2 ** (A % 8) ^ 1)))) ^ 4) % 8)
	// This doesn't clarify that much, so we don't use it

	a := int64(0)
	currentIndex := 0
	for a < math.MaxInt64 {
		// fmt.Printf("Running program with A=%d", a)
		p := programState{a, 0, 0, 0, program.instructions}
		output := p.runOptimized()

		numMatches := 0
		for i := 0; i < min(len(output), len(program.instructions)); i++ {
			if output[i] == int64(program.instructions[i]) {
				numMatches++
			} else {
				break
			}
		}

		if numMatches == len(program.instructions) {
			break
		}

		if numMatches > 2*currentIndex+4 {
			// fmt.Printf("\n\nMATCH!!!\nA=%d - curIdx=%d,matches=%d - ", a, currentIndex, numMatches)
			// for _, x := range output {
			// 	fmt.Printf("%d,", x)
			// }
			// fmt.Println()
			currentIndex++
		}
		a += util.Pow64(2, int64(3*currentIndex*2))
	}

	return a, nil
}

func (p *programState) runOptimized() (out []int64) {
	a := p.a
	for a != 0 {
		b := a % 8
		b = b ^ 1
		c := (a / util.Pow64(2, b))
		b = b ^ c
		b = b ^ 4
		a = a / util.Pow64(2, 3)
		outValue := (b % 8)
		// fmt.Printf("\n\nA=%d\nB=%d\nC=%d\nOutput+=%d\n", a, b, c, outValue)
		out = append(out, outValue)
	}
	return
}

func (p *programState) run() (out []int64) {
	for {
		terminated := p.execute(&out)

		if terminated {
			break
		}

	}
	return
}

func (p *programState) execute(output *[]int64) (terminated bool) {
	if p.instructionPointer >= len(p.instructions) {
		return true
	}

	instruction := Instruction(p.instructions[p.instructionPointer])
	operand := p.instructions[p.instructionPointer+1]
	shouldIncrementPointer := true

	switch instruction {
	case ADV:
		p.a = p.a / util.Pow64(2, p.getComboOperand(operand))
	case BXL:
		p.b = p.b ^ int64(operand)
	case BST:
		p.b = (p.getComboOperand(operand) % 8)
	case JNZ:
		if p.a != 0 {
			p.instructionPointer = int(operand)
			shouldIncrementPointer = false
		}
	case BXC:
		p.b = p.b ^ p.c
	case OUT:
		outValue := p.getComboOperand(operand) % 8
		*output = append(*output, outValue)
		// fmt.Printf("\n\nA=%d\nB=%d\nC=%d\nOutput+=%d\n", p.a, p.b, p.c, outValue)
		// outputed = true
	case BDV:
		p.b = p.a / util.Pow64(2, p.getComboOperand(operand))
	case CDV:
		p.c = p.a / util.Pow64(2, p.getComboOperand(operand))
	}

	if shouldIncrementPointer {
		p.instructionPointer += 2
	}
	return false
}

func (p programState) getComboOperand(operand int8) int64 {
	switch {
	case 0 <= operand && operand <= 3:
		return int64(operand)
	case operand == 4:
		return p.a
	case operand == 5:
		return p.b
	case operand == 6:
		return p.c
	default:
		panic("invalid combo operand")
	}
}

type Instruction int

const (
	ADV Instruction = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type programState struct {
	a                  int64
	b                  int64
	c                  int64
	instructionPointer int
	instructions       []int8
}

func parseInput(useRealInput bool) (*programState, error) {
	data, err := util.ReadInputMulti(17, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 sections of input")
	}

	split := strings.Split(data[0][0], ": ")
	a, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return nil, err
	}
	split = strings.Split(data[0][1], ": ")
	b, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return nil, err
	}
	split = strings.Split(data[0][2], ": ")
	c, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return nil, err
	}

	split = strings.Split(data[1][0], ": ")
	instructions := make([]int8, 0)
	for _, opcodeStr := range strings.Split(split[1], ",") {
		opcode, err := strconv.ParseInt(opcodeStr, 10, 8)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, int8(opcode))
	}

	return &programState{a: a, b: b, c: c, instructionPointer: 0, instructions: instructions}, nil
}
