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

	regionsSlice := solve(garden)
	regions := make(map[int][]util.Vec, 0)
	numSides := make(map[int]int, 0)
	price := 0
	for _, region := range regionsSlice {
		regions[get(garden, region.coordinates[0])] = region.coordinates
		numSides[get(garden, region.coordinates[0])] = region.cornerCount
		price += len(region.coordinates) * region.cornerCount
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 4, len(regionA))
	assert.Equal(t, 4, numSides[int(rune('A'))])

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 4, len(regionB))
	assert.Equal(t, 4, numSides[int(rune('B'))])

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 4, len(regionC))
	assert.Equal(t, 8, numSides[int(rune('C'))])

	regionD := regions[int(rune('D'))]
	assert.Equal(t, 1, len(regionD))
	assert.Equal(t, 4, numSides[int(rune('D'))])

	regionE := regions[int(rune('E'))]
	assert.Equal(t, 3, len(regionE))
	assert.Equal(t, 4, numSides[int(rune('E'))])

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

	regionsSlice := solve(garden)
	regions := make(map[int][]util.Vec, 0)
	numSides := make(map[int]int, 0)
	price := 0
	for _, region := range regionsSlice {
		regions[get(garden, region.coordinates[0])] = region.coordinates
		numSides[get(garden, region.coordinates[0])] = region.cornerCount
		price += len(region.coordinates) * region.cornerCount
	}

	regionE := regions[int(rune('E'))]
	assert.Equal(t, 17, len(regionE))
	assert.Equal(t, 12, numSides[int(rune('E'))])

	regionX := regions[int(rune('X'))]
	assert.Equal(t, 4, len(regionX))
	assert.Equal(t, 4, numSides[int(rune('X'))])

	regionY := regions[int(rune('Y'))]
	assert.Equal(t, 4, len(regionY))
	assert.Equal(t, 4, numSides[int(rune('Y'))])

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

	regionsSlice := solve(garden)
	regions := make(map[int][]util.Vec, 0)
	numSides := make(map[int]int, 0)
	price := 0
	for _, region := range regionsSlice {
		regions[get(garden, region.coordinates[0])] = region.coordinates
		numSides[get(garden, region.coordinates[0])] = region.cornerCount
		price += len(region.coordinates) * region.cornerCount
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 28, len(regionA))
	assert.Equal(t, 12, numSides[int(rune('A'))])

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 4, len(regionB))
	assert.Equal(t, 4, numSides[int(rune('B'))])

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 4, len(regionC))
	assert.Equal(t, 4, numSides[int(rune('C'))])

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

	regionsSlice := solve(garden)
	regions := make(map[int][]util.Vec, 0)
	numSides := make(map[int]int, 0)
	price := 0
	for _, region := range regionsSlice {
		regions[get(garden, region.coordinates[0])] = region.coordinates
		numSides[get(garden, region.coordinates[0])] = region.cornerCount
		price += len(region.coordinates) * region.cornerCount
	}

	regionA := regions[int(rune('A'))]
	assert.Equal(t, 7, len(regionA))
	assert.Equal(t, 10, numSides[int(rune('A'))])

	regionB := regions[int(rune('B'))]
	assert.Equal(t, 1, len(regionB))
	assert.Equal(t, 4, numSides[int(rune('B'))])

	regionC := regions[int(rune('C'))]
	assert.Equal(t, 1, len(regionC))
	assert.Equal(t, 4, numSides[int(rune('C'))])
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
