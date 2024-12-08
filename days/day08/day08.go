package day08

import (
	"advent-of-code-2024/util"
	"slices"
)

func SolvePart1(useRealInput bool) (int, error) {
	antennas, width, height, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	antinodes := make([]util.Vec, 0)
	for _, locations := range antennas {
		possibleAntinodes := make([]util.Vec, 0)
		for _, v := range locations {
			for _, w := range locations {
				if v != w {
					// If v and w are antenna's of same kind, then (v-w) is the vector pointing from w to v
					// Then, the location a = v + (v - w) is the mirror of w in v, and this satisfies:
					//    ||a - v|| = ||v - w||    and    ||a - w|| = 2*||v - w||
					antinode := v.Minus(w).Plus(v)
					possibleAntinodes = append(possibleAntinodes, antinode)
				}
			}
		}

		// An antinode can only ocur if there is no antenna of that same type at its location
		// They can neither ocur outside of the map. Also, we only want unique locations
		for _, a := range possibleAntinodes {
			if !slices.Contains(locations, a) && 
			a.IsInBounds(width, height) &&
			!slices.Contains(antinodes, a) {
				antinodes = append(antinodes, a)
			}
		}

	}

	return len(antinodes), nil
}


func SolvePart2(useRealInput bool) (int, error) {
	antennas, width, height, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	// A quick solution would be to just loop through all coordinates and all antenna pairs,
	// and then check which coordinates are colinear with at least one pair of points
	// We, however, choose to use an exact solution...

	antinodes := make([]util.Vec, 0)
	for _, locations := range antennas {
		possibleAntinodes := make([]util.Vec, 0)
		for _, v := range locations {
			for _, w := range locations {
				if v != w {
					// If v and w are antenna's of same kind, then d = (w - v) is the vector pointing from v to w
					// If we find the GCD of d's coordinates we can scale d down so all grid locations on the line through
					// v and w can be reached by integer multiples of this scaled down vector
					d := v.Minus(w)
					gcd := util.Abs(util.Gcd(d.X, d.Y))
					scaledD := d.Divide(gcd)

					// Now find all coordinates within the grid, these are of the form  v + m*scaledD   for integers m
					minMultiple := - max(util.Abs(v.X / scaledD.X), util.Abs(v.Y / scaledD.Y)) 
					toEdge := util.Vec{X: width, Y: height}.Minus(v)
					maxMultiple := min(util.Abs(toEdge.X / scaledD.X), util.Abs(toEdge.Y / scaledD.Y)) 

					for m := minMultiple; m <= maxMultiple; m++ {
						location := v.Plus(scaledD.Times(m))
						possibleAntinodes = append(possibleAntinodes, location)
					}
				}
			}
		}

		// An antinode now cannot ocur outside of the map. Also, we only want unique locations
		for _, a := range possibleAntinodes {
			if a.IsInBounds(width, height) && !slices.Contains(antinodes, a) {
				antinodes = append(antinodes, a)
			}
		}

	}

	return len(antinodes), nil
}




func parseInput(useRealInput bool) (antennas map[rune][]util.Vec, width, height int, err error) {
	data, err := util.ReadInput(8, useRealInput)
	if err != nil {
		return nil, 0, 0, err
	}
	height, width = len(data), len(data[0])

	antennas = make(map[rune][]util.Vec, 0)
	for y, line := range data {
		for x, c := range line {
			if c != rune('.') {
				antennas[c] = append(antennas[c], util.Vec{X: x, Y: y})
			}
		}
	}

	return antennas, width, height, nil
}
