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

	regions := solve(garden)

	price := 0
	for _, region := range regions {
		price += len(region.coordinates) * region.edgeCount
	}

	return price, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	garden, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	regions := solve(garden)

	price := 0
	for _, region := range regions {
		price += len(region.coordinates) * region.cornerCount
	}

	return price, nil
}

type region struct {
	coordinates []util.Vec
	edgeCount   int
	cornerCount int
}

func solve(garden [][]int) (regions []region) {
	diagDirections := util.ClockwiseDiagDirections()

	processedCoordinates := make(map[util.Vec]bool, 0)
	for y, row := range garden {
		for x := range row {
			coord := util.Vec{X: x, Y: y}
			if !processedCoordinates[coord] {
				edgeCount, regionCoords := findRegionAndCountFences(garden, coord, processedCoordinates)

				// Count number of sides. Since this is same as number of corners, we count those instead
				corners := 0
				for _, v0 := range regionCoords {
					p0 := get(garden, v0)

					// Loop through 4 cardinal directions
					for i := range 4 {
						// Look at neighboring plant types:
						dir1, dir2, dir3 := diagDirections[2*i], diagDirections[(2*i+1)%8], diagDirections[(2*i+2)%8]
						v1, v2, v3 := v0.PlusDirDiag(dir1), v0.PlusDirDiag(dir2), v0.PlusDirDiag(dir3)
						p1, p2, p3 := get(garden, v1), get(garden, v2), get(garden, v3)

						// inner corner, if p0=X in bottom left then:
						// XO
						// XX
						if p0 == p1 && p0 == p3 && p0 != p2 {
							corners++
						}

						// outer corner, if p0=X in bottom left then:
						// O?
						// XO
						if p0 != p1 && p0 != p3 {
							corners++
						}
					}
				}
				regions = append(regions, region{regionCoords, edgeCount, corners})
			}
		}
	}

	return regions
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
