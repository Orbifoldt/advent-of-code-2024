package util

type Vec struct{ X, Y int }

// Update receiver by adding argument vector to it
func (v *Vec) Add(w Vec) {
	v.X += w.X
	v.Y += w.Y
}

// Create new vector that is sum of the receiver and argument
func (v Vec) Plus(w Vec) Vec {
	return Vec{v.X + w.X, v.Y + w.Y}
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
