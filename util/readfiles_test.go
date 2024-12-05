package util

import (
	"testing"
	"strings"
)

func TestShouldCorrectlyReadSampleInputFile(t *testing.T) {
	data, err := ReadInput(1, false)
	if err != nil {
		t.Fatalf("Expected to not return an error, but got: %v", err)
	}
	if data == nil {
		t.Fatal("Expected ReadInput to not return a nil value")
	}

	length := len(data)
	if  length != 6 {
		t.Fatalf("Expected to read 6 non empty lines, but got '%d' instead", length)
	}

	if !strings.HasPrefix(data[0], "3   4") {
		t.Fatal("Expected read input for day 1 to start with '3   4'")
	}
}

func TestShouldCorrectlyReadRealInputFile(t *testing.T) {
	data, err := ReadInput(1, true)
	if err != nil {
		t.Fatalf("Expected to not return an error, but got: %v", err)
	}
	if data == nil {
		t.Fatal("Expected ReadInput to not return a nil value")
	}

	length := len(data)
	if  length != 1000 {
		t.Fatalf("Expected to read 1000 non empty lines, but got '%d' instead", length)
	}

	if !strings.HasPrefix(data[0], "15131   78158") {
		t.Fatal("Expected read input for day 1 to start with '15131   78158'")
	}
}

func TestShouldCorrectlyReadSampleInputFileWithMultipleSections(t *testing.T) {
	data, err := ReadInputMulti(5, false)
	if err != nil {
		t.Fatalf("Expected to not return an error, but got: %v", err)
	}
	if data == nil {
		t.Fatal("Expected ReadInput to not return a nil value")
	}

	length := len(data)
	if  length != 2 {
		t.Fatalf("Expected to read 2 sections, but got '%d' instead", length)
	}

	if !strings.HasPrefix(data[0][0], "47|53") {
		t.Fatal("Expected first section of sample day 5 to start with '47|53'")
	}

	if !strings.HasPrefix(data[1][0], "75,47,61,53,29") {
		t.Fatal("Expected second section of sample day 5 to start with '75,47,61,53,29'")
		
	}
}
