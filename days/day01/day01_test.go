package day01

import (
	"testing"
)

func TestAdd(t *testing.T){
	a := 1
	b := 2
	sum := Add(a, b)
	if sum != 3 {
		t.Fatalf(`Expected "%d + %d" to equal 3, but actual was "%d"`, a, b, sum)
	}
}