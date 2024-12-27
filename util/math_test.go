package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDetermineGcd(t *testing.T) {
	assert.Equal(t, 21, Gcd(1071, 462))
}

func TestShouldCorrectlyDetermineGcdWhenOrderIsFlipped(t *testing.T) {
	assert.Equal(t, 21, Gcd(462, 1071))
}

func TestShouldCorrectlyDetermineGcdRecurisvely(t *testing.T) {
	assert.Equal(t, 21, GcdRecursive(1071, 462))
}

func TestShouldCorrectlyDetermineGcdRecursivelyWhenOrderIsFlipped(t *testing.T) {
	assert.Equal(t, 21, GcdRecursive(462, 1071))
}

func TestShouldReturnPositiveValueForNegativeInput(t *testing.T) {
	assert.Equal(t, 16, Abs(-16))
}

func TestShouldReturnInputValueForPositiveInput(t *testing.T) {
	assert.Equal(t, 33, Abs(33))
}

func TestShouldReturnPowerOfInteger(t *testing.T) {
	assert.Equal(t, 1024, Pow(2, 10))
	assert.Equal(t, 16807, Pow(7, 5))
}
