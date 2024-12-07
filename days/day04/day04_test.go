package day04

import (
	"advent-of-code-2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 18, solution)
}

func TestShouldCorrectlyGetEntriesFromInput(t *testing.T) {
	board, err := parseInput(false)
	assert.NoError(t, err)
	width, height := len(board[0]), len(board)

	c := get(board, width, height, util.Vec{X: 6, Y: 0})
	assert.Equal(t, 'M', c)
}


func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 9, solution)
}

