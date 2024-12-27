package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}

	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", solution)
}

func TestShouldCorrectlyDeterminePart1OnRealInput(t *testing.T) {
	solution, err := SolvePart1(true)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}

	assert.Equal(t, "7,4,2,0,5,0,5,3,7", solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(true)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, int64(202991746427434), solution)
}
