package day14

import (
	"advent-of-code-2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func SolvePart1(useRealInput bool) (int, error) {
	guards, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	var width, height int
	if useRealInput {
		width, height = 101, 103
	} else {
		width, height = 11, 7
	}

	for _, g := range guards {
		g.move(100, width, height)
	}

	return countQuadrants(guards, width, height), nil
}

func SolvePart2(useRealInput bool, drawOutput bool) (int, error) {

	guards, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	var width, height int
	if useRealInput {
		width, height = 101, 103
	} else {
		width, height = 11, 7
	}

	for i := range 10000 {
		for _, g := range guards {
			g.move(1, width, height)
		}
		if detectChristmasTree(guards, width, height, i, drawOutput) {
			return i+1, nil
		}
	}

	return -1, nil
}

func detectChristmasTree(guards []*guard, width, height, iteration int, draw bool) bool {
	positions := make(map[util.Vec]int)
	for _, g := range guards {
		positions[g.p] += 1
	}

	containsDiagonalLine := false

	for p, _ := range positions {
		p1 := p.PlusDirDiag(util.NE)
		p2 := p1.PlusDirDiag(util.NE)
		p3 := p2.PlusDirDiag(util.NE)
		p4 := p3.PlusDirDiag(util.NE)
		p5 := p4.PlusDirDiag(util.NE)
		p6 := p5.PlusDirDiag(util.NE)
		if positions[p1] > 0 && positions[p2] > 0 && positions[p3] > 0 && positions[p4] > 0 && positions[p5] > 0 && positions[p6] > 0 {
			containsDiagonalLine = true
			break
		}
	}

	if !containsDiagonalLine {
		return false
	}

	if draw {
		fmt.Printf("\n\n============= ITERATION %d ===============\n", iteration)
		for y := range height {
			for x := range width {
				count := positions[util.Vec{X: x, Y: y}]
				if count > 0 {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}
	return true
}

type guard struct{ p, v util.Vec }

func (g *guard) move(n, width, height int) {
	g.p.Add(g.v.Times(n))
	g.p.X = util.Mod(g.p.X, width)
	g.p.Y = util.Mod(g.p.Y, height)
}

func countQuadrants(guards []*guard, width, height int) int {
	nw, ne, sw, se := 0, 0, 0, 0

	for _, g := range guards {
		x, y := g.p.X, g.p.Y
		switch {
		case x < width/2 && y < height/2:
			nw++
		case x < width/2 && y > height/2:
			sw++
		case x > width/2 && y < height/2:
			ne++
		case x > width/2 && y > height/2:
			se++
		default: 
			// noop, Exactly in middle
		}
	}
	return nw * ne * sw * se
}

func parseInput(useRealInput bool) ([]*guard, error) {
	data, err := util.ReadInputMulti(14, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	var guardRegex = regexp.MustCompile(`p=(?P<px>\d{1,3}),(?P<py>\d{1,3}) v=(?P<vx>-?\d{1,3}),(?P<vy>-?\d{1,3})`)
	var guards []*guard
	for _, line := range data[0] {
		matches := guardRegex.FindStringSubmatch(line)
		p := util.Vec{X: mustConvertToInt(matches[1]), Y: mustConvertToInt(matches[2])}
		v := util.Vec{X: mustConvertToInt(matches[3]), Y: mustConvertToInt(matches[4])}
		guards = append(guards, &guard{p, v})
	}

	return guards, nil
}

func mustConvertToInt(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}
