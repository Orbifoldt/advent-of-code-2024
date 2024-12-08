package util

import "fmt"

type Vec struct{ X, Y int }

func (v Vec) String() string {
	return fmt.Sprintf("Vec(X: %d, Y: %d)", v.X, v.Y)
}

// Update receiver by adding argument vector to it
func (v *Vec) Add(w Vec) {
	v.X += w.X
	v.Y += w.Y
}

// Create new vector that is sum of the receiver and argument
func (v Vec) Plus(w Vec) Vec {
	return Vec{v.X + w.X, v.Y + w.Y}
}

// Create new vector that is difference between the receiver and argument
func (v Vec) Minus(w Vec) Vec {
	return Vec{v.X - w.X, v.Y - w.Y}
}

// Create new vector by mulitplying both arguments of the receiver by c
func (v Vec) Times(c int) Vec {
	return Vec{v.X * c, v.Y * c}
}

// Create new vector by dividing both arguments of the receiver by c
func (v Vec) Divide(c int) Vec {
	return Vec{v.X / c, v.Y / c}
}

// Test if a point is within the square at (0, 0) in the positive quadrant 
func (v Vec) IsInBounds(width, height int) bool {
	return 0 <= v.X && v.X < width && 0 <= v.Y && v.Y < height
}

//go:generate stringer -type=Direction
type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) ToVec() Vec {
	switch d {
	case UP:
		return Vec{0, -1}
	case RIGHT:
		return Vec{1, 0}
	case DOWN:
		return Vec{0, 1}
	case LEFT:
		return Vec{-1, 0}
	}
	panic("Invalid direction")
}

// Update receiver by adding the corresponding DiagDirection vector
func (v *Vec) MoveDir(d Direction) {
	v.Add(d.ToVec())
}

// Create new vector that is result of summing receiver and the DiagDirectionr vector
func (v Vec) PlusDir(d Direction) Vec {
	return v.Plus(d.ToVec())
}

//go:generate stringer -type=DiagDirection
type DiagDirection int

const (
	N DiagDirection = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func ClockwiseDiagDirections() [8]DiagDirection {
	return [8]DiagDirection{N, NE, E, SE, S, SW, W, NW}
}

func (d DiagDirection) ToVec() Vec {
	switch d {
	case N:
		return Vec{0, -1}
	case NE:
		return Vec{1, -1}
	case E:
		return Vec{1, 0}
	case SE:
		return Vec{1, 1}
	case S:
		return Vec{0, 1}
	case SW:
		return Vec{-1, 1}
	case W:
		return Vec{-1, 0}
	case NW:
		return Vec{-1, -1}
	}
	panic("Invalid direction")
}

// Update receiver by adding the corresponding DiagDirection vector
func (v *Vec) MoveDirDiag(d DiagDirection) {
	v.Add(d.ToVec())
}

// Create new vector that is result of summing receiver and the DiagDirectionr vector
func (v Vec) PlusDirDiag(d DiagDirection) Vec {
	return v.Plus(d.ToVec())
}
