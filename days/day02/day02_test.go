package day02

import (
	"testing"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if solution != 2 {
		t.Fatalf(`Expected 2, but was "%d"`, solution)
	}
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if solution != 4 {
		t.Fatalf(`Expected 4, but was "%d"`, solution)
	}
}
