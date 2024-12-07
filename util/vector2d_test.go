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