package day18

import (
	"advent-of-code-2024/util"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	corruptions, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	size, maxTime := getParameters(useRealInput)

	return solve(corruptions, size, maxTime), nil
}

func SolvePart2(useRealInput bool) (string, error) {
	corruptions, err := parseInput(useRealInput)
	if err != nil {
		return "none", err
	}

	size, minTime := getParameters(useRealInput)

	// We keep on increasing time and seeing what would be minimum distance to exit
	// Once our path is blocked this min distance will be 0
	var corruptionIndex int
	for i := range len(corruptions) - minTime {
		maxTime := i + minTime

		if solve(corruptions, size, maxTime) == 0 {
			corruptionIndex = maxTime - 1
			break
		}
	}

	blockingByte := corruptions[corruptionIndex]
	return fmt.Sprintf("%d,%d", blockingByte.X, blockingByte.Y), nil
}

func getParameters(useRealInput bool) (size int, maxTime int) {
	if useRealInput {
		size = 71 // (0,0) to (70, 70)
		maxTime = 1024
	} else {
		size = 7
		maxTime = 12
	}
	return
}

func solve(corruptions []util.Vec, size int, time int) int {
	board := make([][]bool, size)
	for y := range size {
		board[y] = make([]bool, size)
	}
	for i := range time {
		v := corruptions[i]
		board[v.Y][v.X] = true
	}

	start := util.Vec{X: 0, Y: 0}
	end := util.Vec{X: size - 1, Y: size - 1}
	dirs := util.ClockwiseDirections()

	// Dijkstra init
	pq := make(util.PriorityQueueVec, 0)
	heap.Init(&pq)
	heap.Push(&pq, util.NewEntry(start, 0))
	minDists := make(map[util.Vec]int, 0)
	minDists[start] = 0

	// Dijkstra loop
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*util.Entry)

		for _, dir := range dirs {
			nextPos := current.V.PlusDir(dir)
			nextDist := current.Dist + 1
			if canMoveTo(board, nextPos, size) && nextDist < getDist(minDists, nextPos) {
				heap.Push(&pq, util.NewEntry(nextPos, nextDist))
				minDists[nextPos] = nextDist
			}
		}
	}
	return minDists[end]
}

func canMoveTo(board [][]bool, target util.Vec, size int) bool {
	isInBounds := target.X >= 0 && target.X < size && target.Y >= 0 && target.Y < size
	return isInBounds && !board[target.Y][target.X]
}

func getDist(minDists map[util.Vec]int, pos util.Vec) int {
	minDist, visited := minDists[pos]
	if visited {
		return minDist
	} else {
		return math.MaxInt
	}
}

func parseInput(useRealInput bool) ([]util.Vec, error) {
	data, err := util.ReadInputMulti(18, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	output := make([]util.Vec, len(data[0]))
	for t, entry := range data[0] {
		splitResult := strings.Split(entry, ",")
		x, err := strconv.Atoi(splitResult[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(splitResult[1])
		if err != nil {
			return nil, err
		}
		output[t] = util.Vec{X: x, Y: y}
	}

	return output, nil
}
