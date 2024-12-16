package day16

import (
	"advent-of-code-2024/util"
	"fmt"
	"math"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func SolvePart1(useRealInput bool) (int, error) {
	n, _, err := solve(useRealInput)
	return n, err
}

func SolvePart2(useRealInput bool) (int, error) {
	_, n, err := solve(useRealInput)
	return n, err
}

func solve(useRealInput bool) (int, int, error) {
	maze, start, end, err := parseInput(useRealInput)
	if err != nil {
		return 0, 0, err
	}
	startLoc := location{*start, util.RIGHT}
	allDirs := util.ClockwiseDirections()

	// Dijkstra with priority queue
	// We use negative distance since it's a max-priority queue (instead of min-prio)
	pq := prque.New()
	pq.Push(startLoc, 0)
	dists := make(map[location]float32)
	dists[startLoc] = 0
	paths := make(map[location][][]location, 0)
	paths[startLoc] = [][]location{{startLoc}}

	for !pq.Empty() {
		// Pop closest by location
		tmpLoc, minDist := pq.Pop()
		minLoc, ok := tmpLoc.(location)
		if !ok {
			panic("ohno")
		}

		// For each possible next location, see if we could reach that with less distance
		for _, next := range nextLocations(maze, minLoc, minDist) {
			previousDist, visited := dists[next.loc]
			if !visited {
				previousDist = -math.MaxFloat32
			}

			if next.dist > previousDist {
				// If we gotten there in less distance (inverted) then update it as the best
				dists[next.loc] = next.dist
				pq.Push(next.loc, next.dist)

				// replace all old paths since now we've gotten quicker
				paths[next.loc] = extendPaths(paths, minLoc, next.loc)
			} else if next.dist == previousDist {
				// add new paths that got there in the same distance
				paths[next.loc] = append(paths[next.loc], extendPaths(paths, minLoc, next.loc)...)
			}
		}
	}

	// Min distance is minimum distance of all directions how we could have ended at the end
	minDist := float32(-math.MaxFloat32)
	for _, dir := range allDirs {
		dist := dists[location{*end, util.Direction(dir)}]
		if dist > minDist {
			minDist = dist
		}
	}

	tiles := make(map[util.Vec]struct{})
	for _, dir := range allDirs {
		endLoc := location{*end, util.Direction(dir)}
		dist := dists[endLoc]
		if dist == minDist {
			for _, p := range paths[endLoc] {
				for _, loc := range p {
					tiles[loc.p] = struct{}{}
				}
			}
		}
	}

	return -int(minDist), len(tiles), nil
}

func get(maze [][]bool, p util.Vec) bool {
	return maze[p.Y][p.X]
}

type location struct {
	p   util.Vec
	dir util.Direction
}

type locationDist struct {
	loc  location
	dist float32
}

func nextLocations(maze [][]bool, minLoc location, minDist float32) []locationDist {
	newLocations := make([]locationDist, 0)

	// Move a step if possible
	nextPosition := minLoc.p.PlusDir(minLoc.dir)
	nextLoc := location{nextPosition, minLoc.dir}
	if get(maze, nextPosition) {
		newDist := minDist - 1
		newLocations = append(newLocations, locationDist{nextLoc, newDist})
	}

	// Rotate clockwise or counter clockwise
	newLocations = append(newLocations,
		locationDist{location{minLoc.p, minLoc.dir.RotateClockwise()}, minDist - 1000},
		locationDist{location{minLoc.p, minLoc.dir.RotateCounterClockwise()}, minDist - 1000},
	)
	return newLocations
}

func extendPaths(paths map[location][][]location, curLoc, nextLoc location) (newPaths [][]location) {
	for _, oldPath := range paths[curLoc] {
		extendedPath := make([]location, len(oldPath)+1)
		copy(extendedPath, oldPath)
		extendedPath[len(oldPath)] = nextLoc
		newPaths = append(newPaths, extendedPath)
	}
	return
}

func parseInput(useRealInput bool) ([][]bool, *util.Vec, *util.Vec, error) {
	data, err := util.ReadInputMulti(16, useRealInput)
	if err != nil {
		return nil, nil, nil, err
	}
	if len(data) != 1 {
		return nil, nil, nil, fmt.Errorf("expected single line of input")
	}

	maze := make([][]bool, len(data[0]))
	var start util.Vec
	var end util.Vec
	for y, row := range data[0] {
		maze[y] = make([]bool, len(row))
		for x, r := range row {
			maze[y][x] = (r != rune('#'))
			if r == rune('S') {
				start = util.Vec{X: x, Y: y}
			} else if r == rune('E') {
				end = util.Vec{X: x, Y: y}
			}
		}
	}
	return maze, &start, &end, nil
}
