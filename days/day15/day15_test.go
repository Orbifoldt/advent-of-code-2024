package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 10092, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(true)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 9021, solution)
}
