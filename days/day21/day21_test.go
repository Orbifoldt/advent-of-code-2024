package day21

import (
	"advent-of-code-2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 126384, solution)
}

func TestShouldCorrectlyDetermineOptimalMoves(t *testing.T) {
	right, left, up, down := int(util.RIGHT), int(util.LEFT), int(util.UP), int(util.DOWN)

	// numerical pad
	assert.Equal(t, []int{right, down, down, down, A}, optimalMoveCached(7, 0, true))
	assert.Equal(t, []int{up, up, up, left, A}, optimalMoveCached(0, 7, true))
	assert.Equal(t, []int{up, right, right, A}, optimalMoveCached(4, 9, true))
	assert.Equal(t, 0, len(optimalMoveCached(5, 5, true)))
	assert.Equal(t, []int{up, A}, optimalMoveCached(0, 2, true))

	// directional pad
	assert.Equal(t, []int{right, right, up, A}, optimalMoveCached(left, A, false))
	assert.Equal(t, []int{right, up, A}, optimalMoveCached(left, up, false))
	assert.Equal(t, []int{down, left, left, A}, optimalMoveCached(A, left, false))
	assert.Equal(t, 0, len(optimalMoveCached(down, down, false)))
}

func TestShouldCorrectlyDetermineTotalMovesExample1(t *testing.T) {
	// Robot 0: v<A<AA>>^AvAA^<A>A v<<A>>^AvA^A v<A>^A<Av<A>>^AvA^A  v<A<A>>^AvA^<A>A
	// correct: <vA<AA>>^AvAA<^A>A <v<A>>^AvA^A <vA>^A<v<A>^A>AAvA^A <v<A>A>^AAAvA<^A>A
	// 	  	   |        0         |     2      |         9          |         A        |
	// Robot 1: v<<A>>^A <A>A vA^<AA>A v<AAA>^A
	// correct: v<<A>>^A <A>A vA<^AA>A <vAAA>^A
	// 		   |    0   |  2 |    9   |    A    |
	// Robot 2: <A ^A >^^A vvvA
	// correct: <A ^A >^^A vvvA
	// Robot 3: 029A
	// correct: 029A

	right, left, up, down := int(util.RIGHT), int(util.LEFT), int(util.UP), int(util.DOWN)

	// Moves for pressing 0:
	// <vA<AA>>^A
	assert.Equal(t, 10, movesLength([]int{down, left, left, A}, 1, false))
	// vAA<^A>A
	assert.Equal(t, 8, movesLength([]int{right, right, up, A}, 1, false))
	// <vA<AA>>^A vAA<^A>A
	assert.Equal(t, 18, movesLength([]int{left, A}, 2, false))
	assert.Equal(t, 18, movesLength([]int{0}, 3, true))

	// Moves for pressing 0, 2, 9, A:
	// LEVEL 1
	// 0: <vA<AA>>^A<vA<AA>>^A vAA<^A>A
	assert.Equal(t, 18, movesLength([]int{down, left, left, A, right, right, up, A}, 1, false))
	// 2: v<<A>>^AvA^A
	assert.Equal(t, 12, movesLength([]int{left, A, right, A}, 1, false))
	// 9:
	assert.Equal(t, 20, movesLength([]int{down, A, up, left, A, A, right, A}, 1, false))
	// A:
	assert.Equal(t, 18, movesLength([]int{down, left, A, A, A, right, up, A}, 1, false))

	// LEVEL 2
	// 0: v<<A>>^A
	assert.Equal(t, 18, movesLength([]int{left, A}, 2, false))
	// 2: <A>A
	assert.Equal(t, 12, movesLength([]int{up, A}, 2, false))
	// 9: vA^<AA>A
	assert.Equal(t, 20, movesLength([]int{right, up, up, A}, 2, false))
	// A: v<AAA>^A
	assert.Equal(t, 18, movesLength([]int{down, down, down, A}, 2, false))

	// LEVEL 3
	// 0: <A
	assert.Equal(t, 18, movesLength([]int{0}, 3, true))
	// 2: ^A
	assert.Equal(t, 30, movesLength([]int{0, 2}, 3, true))
	// 9: >^^A
	assert.Equal(t, 50, movesLength([]int{0, 2, 9}, 3, true))
	// A: vvvA
	assert.Equal(t, 68, movesLength([]int{0, 2, 9, A}, 3, true))
}

func TestShouldCorrectlyDetermineTotalMovesExample2to5(t *testing.T) {
	assert.Equal(t, 60, movesLength([]int{9, 8, 0, A}, 3, true))
	assert.Equal(t, 68, movesLength([]int{1, 7, 9, A}, 3, true))
	assert.Equal(t, 64, movesLength([]int{4, 5, 6, A}, 3, true))
	assert.Equal(t, 64, movesLength([]int{3, 7, 9, A}, 3, true))
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 111, solution)
}
