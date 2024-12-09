package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, int64(1928), solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, int64(2858), solution)
}
