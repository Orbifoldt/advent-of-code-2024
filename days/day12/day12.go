package day12

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

func SolvePart1(useRealInput bool) (int, error) {
	garden, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	price := 0
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				numFences, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				price += len(region) * numFences
			}
		}
	}

	return price, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	garden, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	price := 0
	width, height := len(garden[0]), len(garden)
	processedCoordinates := make(map[util.Vec]bool, 0)
	for y := range height {
		for x := range width {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				_, region := findRegionAndCountFences(garden, coord, processedCoordinates)
				numSides := calculateNumSides(region)
				price += numSides * len(region)
			}
		}
	}

	return price, nil
}

type edge struct{ a, b util.Vec }

func (e edge) String() string {
	return fmt.Sprintf("Edge(a=%s, b=%s)", e.a, e.b)
}

func findRegionAndCountFences(
	garden [][]int,
	plantCoordinate util.Vec,
	processedCoordinates map[util.Vec]bool,
) (numFences int, region []util.Vec) {
	directions := util.ClockwiseDirections()
	targetPlant := get(garden, plantCoordinate)

	toCheckQueue := []util.Vec{plantCoordinate}
	for len(toCheckQueue) > 0 {
		// Pop plant coordinate from queue
		currentCoord := toCheckQueue[0]
		toCheckQueue = toCheckQueue[1:]
		processedCoordinates[currentCoord] = true
		region = append(region, currentCoord)

		// Check all neighbors of the same type
		numNeighbors := 0
		for _, dir := range directions {
			nextCoord := currentCoord.PlusDir(dir)
			nextPlant := get(garden, nextCoord)
			if nextPlant == targetPlant {
				numNeighbors += 1

				// If wasn't checked yet, add to queue
				if !processedCoordinates[nextCoord] && !slices.Contains(toCheckQueue, nextCoord) {
					toCheckQueue = append(toCheckQueue, nextCoord)
				}
			}
		}
		numFences += (4 - numNeighbors)
	}
	return
}

func get(grid [][]int, coordinate util.Vec) int {
	if coordinate.X < 0 || coordinate.X >= len(grid[0]) || coordinate.Y < 0 || coordinate.Y >= len(grid) {
		return -1
	}
	return grid[coordinate.Y][coordinate.X]
}

func calculateNumSides(region []util.Vec) int {
	// All edges of a square with its top left corner at (0, 0)
	possibleEdges := map[util.Direction]edge{
		util.UP:    {util.Vec{X: 0, Y: 0}, util.Vec{X: 1, Y: 0}},
		util.RIGHT: {util.Vec{X: 1, Y: 0}, util.Vec{X: 1, Y: 1}},
		util.DOWN:  {util.Vec{X: 1, Y: 1}, util.Vec{X: 0, Y: 1}},
		util.LEFT:  {util.Vec{X: 0, Y: 1}, util.Vec{X: 0, Y: 0}},
	}

	// List all edges of all squares that form the region
	var edges []edge
	for _, coordinate := range region {
		for _, edge := range possibleEdges {
			edge = edge.Offset(coordinate)
			edges = append(edges, edge)
		}
	}

	// Filter to only keep external edges
	var externalEdges []edge
	var externalEdgesReadOnly []edge
	for i, edge := range edges {
		anySame := false
		for j, otherEdge := range edges {
			// Any two distinct edges that are the same are not external edges
			if i != j && edge.IsSameAs(otherEdge) {
				anySame = true
				break
			}
		}
		if !anySame {
			externalEdges = append(externalEdges, edge)
			externalEdgesReadOnly = append(externalEdgesReadOnly, edge)
		}
	}

	// Create array of all sides
	var sides [][]edge
	for len(externalEdges) > 0 {
		// Pop an edge that is not part of any side yet
		extEdge := externalEdges[0]
		externalEdges = externalEdges[1:]

		// Find all other edges that belong to same side as this edge
		side := []edge{extEdge}
		for i := 0; i < len(externalEdges); i++ {
			otherEdge := externalEdges[i]

			// Check if an edge can belong to that side
			for _, sideEdge := range side {
				// The edge belongs to the side if there is an otherEdge to which it connects and is not perpendicular
				if otherEdge.ConnectAndNotPerpendicular(sideEdge) {
					// Two possibilities: normal straight line, or a cross
					if arePartOfCross(externalEdgesReadOnly, otherEdge, sideEdge) {
						continue
					}

					// otherEdge is truly part of this side
					side = append(side, otherEdge)
					externalEdges = slices.Delete(externalEdges, i, i+1)
					i = -1
					break
				}
			}
		}
		sides = append(sides, side)
	}
	return len(sides)
}

func (e edge) IsSameAs(other edge) bool {
	return (e.a == other.a && e.b == other.b) || (e.a == other.b && e.b == other.a)
}

func (e edge) Offset(offset util.Vec) edge {
	return edge{a: e.a.Plus(offset), b: e.b.Plus(offset)}
}

func (e edge) ConnectAndNotPerpendicular(other edge) bool {
	perpendicular := (e.a.Minus(e.b)).IsPerpendicularTo(other.a.Minus(other.b))
	return !perpendicular && e.Connects(other)
}

func (e edge) Connects(other edge) bool {
	return (e.a == other.a || e.a == other.b || e.b == other.a || e.b == other.b)
}

func arePartOfCross(allEdges []edge, edgeA, edgeB edge) bool {
	// We have a cross if there is at least one edge that also connects to the other two but is perpendicular
	// Example of a cross:
	// AAX
	// AYA
	// AAA
	for _, thirdEdge := range allEdges {
		if !thirdEdge.IsSameAs(edgeA) &&
			!thirdEdge.IsSameAs(edgeB) &&
			thirdEdge.Connects(edgeA) &&
			thirdEdge.Connects(edgeB) {
			return true
		}
	}
	return false
}

func parseInput(useRealInput bool) ([][]int, error) {
	data, err := util.ReadInputMulti(12, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	garden := make([][]int, len(data[0]))
	for y, line := range data[0] {
		garden[y] = make([]int, len(line))
		for x, r := range line {
			garden[y][x] = int(r)
		}
	}

	return garden, nil
}
