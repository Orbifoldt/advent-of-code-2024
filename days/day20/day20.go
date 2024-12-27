package day20

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
)

func SolvePart1(useRealInput bool) (int, error) {
	maze, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	path, distMap := findPath(maze)

	var minToSave int
	if useRealInput {
		minToSave = 100
	} else {
		minToSave = 10
	}

	count := countCheats(path, distMap, minToSave, 2, 2)

	return count, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	maze, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	path, distMap := findPath(maze)

	var minToSave int
	if useRealInput {
		minToSave = 100
	} else {
		minToSave = 50
	}

	count := countCheats(path, distMap, minToSave, 1, 20)
	return count, nil
}

func countCheats(path []util.Vec, distMap map[util.Vec]int, minDistToSave int, minCheatDist, maxCheatDist int) int {
	cheatedRuns := make(map[[2]util.Vec]int)
	distSavedCounter := make(map[int]int)
	counter := 0

	for i, v := range path {
		dist := distMap[v]

		// Try to jump ahead
		for _, cheatPos := range path[i:] {
			cheatDistance := util.Abs(cheatPos.X-v.X) + util.Abs(cheatPos.Y-v.Y)
			if cheatDistance < minCheatDist || cheatDistance > maxCheatDist {
				continue
			}

			savedDist := distMap[cheatPos] - dist - cheatDistance
			if savedDist > 0 {
				cheatedRuns[[2]util.Vec{v, cheatPos}] = savedDist
				distSavedCounter[savedDist] = distSavedCounter[savedDist] + 1
				if savedDist >= minDistToSave {
					counter++
				}
			}
		}
	}

	// for dist, count := range distSavedCounter {
	// 	if dist >= minDistToSave {
	// 		fmt.Printf("There are %2d path(s) that save distance: %d\n", count, dist)
	// 	}
	// }
	return counter
}

// Very ugly, but it works...
func findPath(maze *Maze) ([]util.Vec, map[util.Vec]int) {
	dirs := util.ClockwiseDirections()

	floodFillDists := make(map[util.Vec]int)
	floodFillDists[maze.start] = 0
	floodFillPredecessor := make(map[util.Vec]util.Vec)
	floodFillPredecessor[maze.start] = maze.start
	toCheck := []util.Vec{maze.start}

	for len(toCheck) > 0 {
		current := toCheck[0]
		toCheck = toCheck[1:]
		currentDist := floodFillDists[current]

		if current == maze.end {
			break
		}

		for _, dir := range dirs {
			next := current.PlusDir(dir)
			_, visited := floodFillDists[next]
			if !visited && !maze.walls[next] {
				floodFillDists[next] = currentDist + 1
				floodFillPredecessor[next] = current
				toCheck = append(toCheck, next)
			}
		}
	}

	path := []util.Vec{maze.end}
	pathDist := make(map[util.Vec]int)
	current := maze.end
	pathDist[current] = floodFillDists[current]
	for {
		predecessor := floodFillPredecessor[current]
		path = append(path, predecessor)
		pathDist[predecessor] = floodFillDists[predecessor]
		current = predecessor
		if current == maze.start {
			break
		}
	}
	slices.Reverse(path)

	return path, pathDist
}

type Maze struct {
	start util.Vec
	end   util.Vec
	walls map[util.Vec]bool
	size  int
}

func parseInput(useRealInput bool) (*Maze, error) {
	data, err := util.ReadInputMulti(20, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single section of input")
	}

	var start util.Vec
	var end util.Vec
	walls := make(map[util.Vec]bool)
	for y, row := range data[0] {
		for x, r := range row {
			coord := util.Vec{X: x, Y: y}
			if r == 'S' {
				start = coord
			} else if r == 'E' {
				end = coord
			}

			walls[coord] = (r == '#')
		}
	}

	return &Maze{start: start, end: end, walls: walls, size: len(data[0])}, nil
}
