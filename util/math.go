package util

// Determine Greatest Common Divisor of a and b using Euclidean algorithm
// See https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func GcdRecursive(a, b int) int {  // Looks nicer imo, but Go doesn't do Tail-Call Optimization it seems
	if b == 0 {
		return a
	}
	return GcdRecursive(b, a % b)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Mod(a int, n int) int {
	a = a % n
	if a < 0 {
		a += n
	}
	return a
}

