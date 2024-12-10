package day10

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	grid, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	score, _ := scoreTrailheads(grid)
	return score, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	grid, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	_, score := scoreTrailheads(grid)
	return score, nil
}

type trail struct{ start, end util.Vec }

// Score all trailheads using both scoring methods (part 1 and 2)
func scoreTrailheads(grid [][]int) (int, int) {
	// List all trailheads
	var trailheads []util.Vec
	for y, row := range grid {
		for x, h := range row {

			if h == 0 {
				trailheads = append(trailheads, util.Vec{X: x, Y: y})
			}
		}
	}

	totalScore := 0
	totalScoreUnique := 0
	for _, trailhead := range trailheads {
		trails := findAllTrails(grid, trailhead)

		// Count number of ends that can be reached
		score := 0
		var trailEnds []util.Vec
		for _, trail := range trails {
			if !slices.Contains(trailEnds, trail.end) {
				trailEnds = append(trailEnds, trail.end)
				score += 1
			}
		}
		totalScore += score

		// Count number of total trails
		totalScoreUnique += len(trails)
	}
	return totalScore, totalScoreUnique
}

func findAllTrails(grid [][]int, trailhead util.Vec) (trails []*trail) {
	directions := util.ClockwiseDirections()
	bfsQueue := []util.Vec{trailhead}
	for {
		if len(bfsQueue) == 0 {
			break
		}

		// pop from queue
		coord := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		height := get(grid, coord)

		if height == 9 {
			trails = append(trails, &trail{trailhead, coord})
		} else if height >= 0 { // allow -1 for the non-filled examples
			for _, dir := range directions {
				new := coord.PlusDir(dir)

				newHeight := get(grid, new)
				if newHeight == height+1 {
					bfsQueue = append(bfsQueue, new)
				}
			}
		}
	}
	return trails
}

func get(grid [][]int, coordinate util.Vec) int {
	if coordinate.X < 0 || coordinate.X >= len(grid[0]) || coordinate.Y < 0 || coordinate.Y >= len(grid) {
		return -1
	}
	return grid[coordinate.Y][coordinate.X]
}

func parseInput(useRealInput bool) ([][]int, error) {
	data, err := util.ReadInputMulti(10, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	return parseToGrid(data[0])
}

// to test the other examples
func parseStringInput(input string) ([][]int, error) {
	data := make([]string, 0)
	for _, row := range strings.Split(input, "\n") {
		row = strings.TrimSpace(row)
		if len(row) > 0 {
			data = append(data, row)
		}
	}

	return parseToGrid(data)
}

func parseToGrid(data []string) ([][]int, error) {
	grid := make([][]int, len(data))
	for y, row := range data {
		grid[y] = make([]int, len(row))
		for x, c := range row {
			var h int
			var err error
			if c == rune('.') {
				h = -1
			} else {
				h, err = strconv.Atoi(string(c))
				if err != nil {
					return nil, err
				}

			}
			grid[y][x] = h
		}
	}

	return grid, nil
}
