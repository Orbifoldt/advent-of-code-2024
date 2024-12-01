package day01

import (
	"testing"
)

func TestShouldCorrectlyDetermineTotalDistanceOnExampleInput(t *testing.T) {
	distance, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if distance != 11 {
		t.Fatalf(`Expected total distance to be 11, but was "%d"`, distance)
	}
}


func TestShouldCorrectlyDetermineSimilarityScoreOnExampleInput(t *testing.T) {
	similarity, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	if similarity != 31 {
		t.Fatalf(`Expected similarity score to be 31, but was "%d"`, similarity)
	}
}
