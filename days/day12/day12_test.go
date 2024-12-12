package day12

import (
	"advent-of-code-2024/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 1930, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 1206, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput_sample1(t *testing.T) {
	garden := parse(`
	AAAA
	BBCD
	BBCC
	EEEC
	`)
	assert.Equal(t, 4, len(garden))
	assert.Equal(t, 4, len(garden[0]))

	price := 0
	regions := make(map[int][]util.Vec, 0)
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				_, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				regions[get(garden, coord)] = region
				price += len(region) * calculateNumSides(region)
			}
		}
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 4, len(regionA))
	assert.Equal(t, 4, calculateNumSides(regionA))

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 4, len(regionB))
	assert.Equal(t, 4, calculateNumSides(regionB))

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 4, len(regionC))
	assert.Equal(t, 8, calculateNumSides(regionC))

	regionD := regions[int(rune('D'))]
	assert.Equal(t, 1, len(regionD))
	assert.Equal(t, 4, calculateNumSides(regionD))

	regionE := regions[int(rune('E'))]
	assert.Equal(t, 3, len(regionE))
	assert.Equal(t, 4, calculateNumSides(regionE))

	assert.Equal(t, 80, price)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput_sample2(t *testing.T) {
	garden := parse(`
	EEEEE
	EXXXX
	EEEEE
	EYYYY
	EEEEE
	`)
	assert.Equal(t, 5, len(garden))
	assert.Equal(t, 5, len(garden[0]))

	price := 0
	regions := make(map[int][]util.Vec, 0)
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				_, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				regions[get(garden, coord)] = region
				price += len(region) * calculateNumSides(region)
			}
		}
	}

	regionE := regions[int(rune('E'))]
	assert.Equal(t, 17, len(regionE))
	assert.Equal(t, 12, calculateNumSides(regionE))

	regionX := regions[int(rune('X'))]
	assert.Equal(t, 4, len(regionX))
	assert.Equal(t, 4, calculateNumSides(regionX))

	regionY := regions[int(rune('Y'))]
	assert.Equal(t, 4, len(regionY))
	assert.Equal(t, 4, calculateNumSides(regionY))

	assert.Equal(t, 236, price)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput_sample3(t *testing.T) {
	garden := parse(`
	AAAAAA
	AAABBA
	AAABBA
	ACCAAA
	ACCAAA
	AAAAAA
	`)
	assert.Equal(t, 6, len(garden))
	assert.Equal(t, 6, len(garden[0]))

	price := 0
	regions := make(map[int][]util.Vec, 0)
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				_, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				regions[get(garden, coord)] = region
				price += len(region) * calculateNumSides(region)
			}
		}
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 28, len(regionA))
	assert.Equal(t, 12, calculateNumSides(regionA))

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 4, len(regionB))
	assert.Equal(t, 4, calculateNumSides(regionB))

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 4, len(regionC))
	assert.Equal(t, 4, calculateNumSides(regionC))

	assert.Equal(t, 368, price)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput_sample4(t *testing.T) {
	garden := parse(`
	AAB
	ACA
	AAA
	`)
	assert.Equal(t, 3, len(garden))
	assert.Equal(t, 3, len(garden[0]))

	regions := make(map[int][]util.Vec, 0)
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				_, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				regions[get(garden, coord)] = region
			}
		}
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 7, len(regionA))
	assert.Equal(t, 10, calculateNumSides(regionA))

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 1, len(regionB))
	assert.Equal(t, 4, calculateNumSides(regionB))

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 1, len(regionC))
	assert.Equal(t, 4, calculateNumSides(regionC))
}

func parse(str string) [][]int {
	data := strings.Split(str, "\n")
	garden := make([][]int, 0)
	for _, line := range data {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		gardenRow := make([]int, len(line))
		for x, r := range line {
			gardenRow[x] = int(r)
		}
		garden = append(garden, gardenRow)
	}
	return garden
}
