package day07

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)


func SolvePart1(useRealInput bool) (int64, error) {
	testValues, operands, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	var sum int64
	for idx, testValue := range testValues {
		operandsLine := operands[idx]
		if canBeSolved(testValue, operandsLine) {
			sum += testValue
		}
	}

	return sum, nil
}

func SolvePart2(useRealInput bool) (int64, error) {
	testValues, operands, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	var sum atomic.Int64
	var wg sync.WaitGroup
	for idx, testValue := range testValues {
		operandsLine := operands[idx]
		wg.Add(1)

		go func() {
			defer wg.Done()
			if canBeSolvedWithConcatenation(testValue, operandsLine) {
				sum.Add(testValue)
			}
		}()
	}
	wg.Wait()

	return sum.Load(), nil
}

func canBeSolved(testValue int64, operands []int64) bool {

	var solveIt func(int64, int64, []int64)bool
	solveIt = func (testValue int64, soFar int64, operands []int64) bool {
		if len(operands) == 0 {
			return testValue == soFar
		}
		
		next, leftOver := operands[0], operands[1:]
		return solveIt(testValue, soFar * next, leftOver) || solveIt(testValue, soFar + next, leftOver)
	}

	return solveIt(testValue, 0, operands)
}


func canBeSolvedWithConcatenation(testValue int64, operands []int64) bool {

	var solveIt func(int64, int64, []int64)bool
	solveIt = func (testValue int64, soFar int64, operands []int64) bool {
		if len(operands) == 0 {
			return testValue == soFar
		}
		
		next, leftOver := operands[0], operands[1:]
		return solveIt(testValue, soFar * next, leftOver) || 
			solveIt(testValue, soFar + next, leftOver) || 
			solveIt(testValue, concatenate(soFar, next), leftOver)
	}

	return solveIt(testValue, 0, operands)
}

func concatenate(a, b int64) int64 {
	c, err := strconv.ParseInt(strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10), 10, 64)
	if err != nil {
		panic(err)
	}
	return c
}

func parseInput(useRealInput bool) (testValues []int64, operands [][]int64, err error) {
	data, err := util.ReadInputMulti(7, useRealInput)
	if err != nil {
		return nil, nil, err
	}
	if len(data) != 1 {
		return nil, nil, fmt.Errorf("expected to get 1 section in the input, got %d isntead", len(data))
	}

	for _, line := range data[0] {
		testValueString, operandsString, _ := strings.Cut(line, ":")

		testValue, err := strconv.ParseInt(testValueString, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		testValues = append(testValues, testValue)

		operandsLine := make([]int64, 0)
		for _, operandString := range strings.Split(strings.TrimSpace(operandsString), " ") {
			operand, err := strconv.ParseInt(operandString, 10, 64)
			if err != nil {
				return nil, nil, err
			}
			operandsLine = append(operandsLine, operand)
		}
		operands = append(operands, operandsLine)
	}
	return
}