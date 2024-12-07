package day07


import (
	"testing"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	expected := int64(3749)
	if solution != expected {
		t.Fatalf(`Expected %d, but was "%d"`, expected, solution)
	}
}


func TestShouldCorrectlyDeterminePar21OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	expected := int64(11387)
	if solution != expected {
		t.Fatalf(`Expected %d, but was "%d"`, expected, solution)
	}
}
