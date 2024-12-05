package day05

import "testing"

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	expected := 143
	if solution != expected {
		t.Fatalf(`Expected %d, but was "%d"`, expected, solution)
	}
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	expected := 123
	if solution != expected {
		t.Fatalf(`Expected %d, but was "%d"`, expected, solution)
	}
}