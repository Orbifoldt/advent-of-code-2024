package day13

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func SolvePart1(useRealInput bool) (int, error) {
	machines, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	totalCosts := int64(0)
	for _, machine := range machines {
		totalCosts += cheapestWinPt2(machine, 0)
	}

	return int(totalCosts), nil
}

func SolvePart2(useRealInput bool) (int64, error) {
	machines, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	prizeIncrement := int64(10000000000000)

	totalCosts := int64(0)
	for _, machine := range machines {
		totalCosts += cheapestWinPt2(machine, prizeIncrement)
	}

	return totalCosts, nil
}

type machine struct{ a, b, prize util.Vec }

// func cheapestWinPt1(machine machine) int {
// 	// Quick and dirty solution: we do loop, but do it somewhat smartly
// 	minCost := math.MaxInt
// 	for i := range 100 {
// 		j := (machine.prize.X - i*machine.a.X) / machine.b.X
// 		if machine.a.Times(i).Plus(machine.b.Times(j)) == machine.prize {
// 			cost := i*3 + j*1
// 			if cost < minCost {
// 				minCost = cost
// 			}
// 		}
// 	}
// 	if minCost == math.MaxInt {
// 		minCost = 0
// 	}
// 	return minCost
// }

func cheapestWinPt2(machine machine, prizeIncrement int64) int64 {
	// After analyzing part 1, the solution is unique if it exists, which I suppose it must be (inverses are unique)
	// We're basically just solving A*v = p, where v=(i, j) is how many times we press A and B resp.,
	// and where p is the prize vector, and A a 2x2 matrix whose columns are the two button vectors.
	// In other words, v = (A^-1)*p
	a, b, c, d := int64(machine.a.X), int64(machine.b.X), int64(machine.a.Y), int64(machine.b.Y)
	px, py := int64(machine.prize.X)+prizeIncrement, int64(machine.prize.Y)+prizeIncrement

	determinant := a*d - b*c
	if determinant == 0 {
		return 0
	}

	// Integer rounded solution (since we use integer division)
	i := (d*px - b*py) / determinant
	j := (-c*px + a*py) / determinant

	// Verify if it's valid
	if i < 0 || j < 0 {
		return 0
	}
	if i*a+j*b != px || i*c+j*d != py {
		return 0
	}

	return int64(3)*i + j
}

func parseInput(useRealInput bool) ([]machine, error) {
	data, err := util.ReadInputMulti(13, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) <= 1 {
		return nil, fmt.Errorf("expected multiple sections of input")
	}

	if err != nil {
		return nil, err
	}

	var machines []machine
	for _, block := range data {
		buttonA := extractToVector(buttonRegex, block[0])
		buttonB := extractToVector(buttonRegex, block[1])
		prize := extractToVector(prizeRegex, block[2])
		machines = append(machines, machine{buttonA, buttonB, prize})
	}

	return machines, nil
}

var buttonRegex = regexp.MustCompile(`Button \w: X\+(?P<x>\d{1,3}), Y\+(?P<y>\d{1,3})`)
var prizeRegex = regexp.MustCompile(`Prize: X=(?P<x>\d{1,6}), Y=(?P<y>\d{1,6})`)

func extractToVector(regex *regexp.Regexp, str string) util.Vec {
	matches := regex.FindStringSubmatch(str)
	x, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return util.Vec{X: x, Y: y}
}
