package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExample1(t *testing.T) {
	grid, err := parseStringInput(`
	...0...
    ...1...
    ...2...
    6543456
    7.....7
    8.....8
    9.....9
	`)
	if err != nil {
		t.Fatalf("Error during parsing input: %v", err)
	}
	solution, _ := scoreTrailheads(grid)
	assert.Equal(t, 2, solution)
}

func TestShouldCorrectlyDeterminePart1OnExample2(t *testing.T) {
	grid, err := parseStringInput(`
	..90..9
    ...1.98
    ...2..7
    6543456
    765.987
    876....
    987....
	`)
	if err != nil {
		t.Fatalf("Error during parsing input: %v", err)
	}
	solution, _ := scoreTrailheads(grid)
	assert.Equal(t, 4, solution)
}

func TestShouldCorrectlyDeterminePart1OnExample3_multipleHeads(t *testing.T) {
	grid, err := parseStringInput(`
	10..9..
    2...8..
    3...7..
    4567654
    ...8..3
    ...9..2
    .....01
	`)
	if err != nil {
		t.Fatalf("Error during parsing input: %v", err)
	}
	solution, _ := scoreTrailheads(grid)
	assert.Equal(t, 3, solution)
}

func TestShouldCorrectlyDeterminePart1OnExample4(t *testing.T) {
	grid, err := parseStringInput(`
	89010123
	78121874
	87430965
	96549874
	45678903
	32019012
	01329801
	10456732
	`)
	if err != nil {
		t.Fatalf("Error during parsing input: %v", err)
	}
	solution, _ := scoreTrailheads(grid)
	assert.Equal(t, 36, solution)
}

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 36, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 81, solution)
}
