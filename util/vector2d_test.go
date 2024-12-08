package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyAddVector(t *testing.T) {
	v := Vec{3, -5}
	w := Vec{-2, -3}
	v.Add(w)
	assert.Equal(t, Vec{1, -8}, v)
	assert.Equal(t, Vec{-2, -3}, w, "w should remain unchanged")
}

func TestShouldCorrectlyPlusVectors(t *testing.T) {
	v := Vec{3, -5}
	w := Vec{-2, -3}
	u := v.Plus(w)
	assert.Equal(t, Vec{1, -8}, u)
	assert.Equal(t, Vec{3, -5}, v, "v should remain unchanged")
	assert.Equal(t, Vec{-2, -3}, w, "w should remain unchanged")
}

func TestShouldCorrectlyMinusVectors(t *testing.T) {
	v := Vec{3, -5}
	w := Vec{-2, -3}
	u := v.Minus(w)
	assert.Equal(t, Vec{5, -2}, u)
	assert.Equal(t, Vec{3, -5}, v, "v should remain unchanged")
	assert.Equal(t, Vec{-2, -3}, w, "w should remain unchanged")
}

func TestShouldCorrectlyTimesAVector(t *testing.T) {
	v := Vec{1, -3}
	u := v.Times(-5)
	assert.Equal(t, Vec{-5, 15}, u)
	assert.Equal(t, Vec{1, -3}, v, "v should remain unchanged")
}

func TestShouldCorrectlyDivideAVector(t *testing.T) {
	v := Vec{-9, 15}
	u := v.Divide(-3)
	assert.Equal(t, Vec{3, -5}, u)
	assert.Equal(t, Vec{-9, 15}, v, "v should remain unchanged")
}

func TestShouldDeterminIfAVectorIsInBounds(t *testing.T) {
	width, height := 5, 7
	assert.Equal(t, true, Vec{3, 0}.IsInBounds(width, height))
	assert.Equal(t, true, Vec{0, 5}.IsInBounds(width, height))
	assert.Equal(t, true, Vec{1, 6}.IsInBounds(width, height))
	assert.Equal(t, true, Vec{0, 0}.IsInBounds(width, height))
	assert.Equal(t, true, Vec{4, 6}.IsInBounds(width, height))
}

func TestShouldDeterminIfAVectorIsNotInBounds(t *testing.T) {
	width, height := 5, 7
	assert.Equal(t, false, Vec{3, -1}.IsInBounds(width, height))
	assert.Equal(t, false, Vec{3, 7}.IsInBounds(width, height))
	assert.Equal(t, false, Vec{3, 8}.IsInBounds(width, height))
	assert.Equal(t, false, Vec{-1, 5}.IsInBounds(width, height))
	assert.Equal(t, false, Vec{5, 5}.IsInBounds(width, height))
	assert.Equal(t, false, Vec{6, 5}.IsInBounds(width, height))
}

func TestShouldCorrectlyUpdateVectorAccordingToDirection(t *testing.T) {
	v := Vec{-33, 17}
	v.MoveDir(DOWN)
	assert.Equal(t, Vec{-33, 18}, v)
}

func TestShouldCorrectlyPlusDirectionToVector(t *testing.T) {
	v := Vec{-33, 17}
	w := v.PlusDir(RIGHT)
	assert.Equal(t, Vec{-32, 17}, w)
	assert.Equal(t, Vec{-33, 17}, v, "v should remain unchanged")
}

func TestShouldCorrectlyUpdateVectorAccordingToDiagonalDirection(t *testing.T) {
	v := Vec{-33, 17}
	v.MoveDirDiag(SW)
	assert.Equal(t, Vec{-34, 18}, v)
}

func TestShouldCorrectlyPlusDiagonalDirectionToVector(t *testing.T) {
	v := Vec{-33, 17}
	w := v.PlusDirDiag(NW)
	assert.Equal(t, Vec{-34, 16}, w)
	assert.Equal(t, Vec{-33, 17}, v, "v should remain unchanged")
}
