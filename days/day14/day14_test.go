package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 12, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(true, true)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 7492, solution)
}
