package day16

import (
	"advent-of-code-2024/util"
	"container/heap"
	"fmt"
	"math"
)

func SolvePart1(useRealInput bool) (int, error) {
	n, _, err := solve(useRealInput)
	return n, err
}

func SolvePart2(useRealInput bool) (int, error) {
	_, n, err := solve(useRealInput)
	return n, err
}

type PriorityQueue []*locationDist // See https://pkg.go.dev/container/heap#example-package-PriorityQueue

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*locationDist)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func solve(useRealInput bool) (int, int, error) {
	maze, start, end, err := parseInput(useRealInput)
	if err != nil {
		return 0, 0, err
	}
	startLoc := location{*start, util.RIGHT}

	// Dijkstra with priority queue
	// We use negative distance since it's a max-priority queue (instead of min-prio)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &locationDist{startLoc, 0, -1})
	dists := make(map[location]int)
	dists[startLoc] = 0
	paths := make(map[location][][]location, 0)
	paths[startLoc] = [][]location{{startLoc}}

	for pq.Len() > 0 {
		// Pop closest by location
		tmp, ok := heap.Pop(&pq).(*locationDist)
		minLoc, minDist := tmp.loc, tmp.dist
		if !ok {
			panic("ohno")
		}

		// For each possible next location, see if we could reach that with less distance
		for _, next := range nextLocations(maze, minLoc, minDist) {
			previousDist, visited := dists[next.loc]
			if !visited {
				previousDist = math.MaxInt
			}

			if next.dist < previousDist {
				// If we gotten there in less distance (inverted) then update it as the best
				dists[next.loc] = next.dist
				heap.Push(&pq, next)

				// replace all old paths since now we've gotten quicker
				paths[next.loc] = extendPaths(paths, minLoc, next.loc)
			} else if next.dist == previousDist {
				// add new paths that got there in the same distance
				paths[next.loc] = append(paths[next.loc], extendPaths(paths, minLoc, next.loc)...)
			}
		}
	}

	// Min distance is minimum distance of all directions how we could have ended at the end
	minDist := math.MaxInt
	allDirs := util.ClockwiseDirections()
	for _, dir := range allDirs {
		dist := dists[location{*end, util.Direction(dir)}]
		if dist < minDist {
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

	// for y, row := range maze {
	// 	for x, b := range row {
	// 		_, visited := tiles[util.Vec{X: x, Y: y}]
	// 		if visited {
	// 			fmt.Print("O")
	// 		} else if b {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Print("#")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return minDist, len(tiles), nil
}

func get(maze [][]bool, p util.Vec) bool {
	return maze[p.Y][p.X]
}

type location struct {
	p   util.Vec
	dir util.Direction
}

type locationDist struct {
	loc   location
	dist  int
	index int
}

func nextLocations(maze [][]bool, minLoc location, minDist int) []*locationDist {
	newLocations := make([]*locationDist, 0)

	// Move a step if possible
	nextPosition := minLoc.p.PlusDir(minLoc.dir)
	nextLoc := location{nextPosition, minLoc.dir}
	if get(maze, nextPosition) {
		newDist := minDist + 1
		newLocations = append(newLocations, &locationDist{nextLoc, newDist, -1})
	}

	// Rotate clockwise or counter clockwise
	newLocations = append(newLocations,
		&locationDist{location{minLoc.p, minLoc.dir.RotateClockwise()}, minDist + 1000, -1},
		&locationDist{location{minLoc.p, minLoc.dir.RotateCounterClockwise()}, minDist + 1000, -1},
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
