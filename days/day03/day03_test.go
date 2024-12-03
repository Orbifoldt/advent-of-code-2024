package day03

import "testing"

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if solution != 161 {
		t.Fatalf(`Expected 161, but was "%d"`, solution)
	}
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := sumEnabledMultiplications("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if solution != 48 {
		t.Fatalf(`Expected 48, but was "%d"`, solution)
	}
}
