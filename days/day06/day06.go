package day06

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
)

func SolvePart1(useRealInput bool) (int, error) {
	board, pos, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	height := len(board)
	width := len(board[0])

	visited := walk(board, width, height, pos[0], pos[1])

	visitedCount := 0
	for y := range height {
		for x := range width {
			if visited[y][x] {
				visitedCount++
			}
		}
	}

	return visitedCount, nil
}

func walk(board [][]bool, width, height, startX, startY int) (visited [][]bool) {
	pos := [2]int{startX, startY}

	visited = make([][]bool, height)
	for y := range height {
		visited[y] = make([]bool, width)
	}
	visited[pos[1]][pos[0]] = true  // start position should be included!!!

	dir := UP
	for {
		newPos := nextCoordinate(pos, dir)
		x, y := newPos[0], newPos[1]

		if x < 0 || x >= width || y < 0 || y >= height {
			// fmt.Printf("Left board at (%d, %d)\n", pos[0], pos[1])
			// printBoard(board, visited, pos)
			return visited
		}

		if board[y][x] {
			// hit a wall/box
			dir = turnRight(dir)
			// fmt.Printf("Hit wall at (%d, %d), turning toward %v \n", x, y, dir)
		} else {
			// Only move if we didn't hit anything
			pos = newPos
			visited[y][x] = true
		}
	}
}

func SolvePart2(useRealInput bool) (int, error) {
	board, originalPos, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	height := len(board)
	width := len(board[0])

	possibleVisited := walk(board, width, height, originalPos[0], originalPos[1])

	var loopCount atomic.Int32
	var wg sync.WaitGroup
	wg.Add(width * height)

	for obsY := range height {
		for obsX := range width {
			go func(obsX int, obsY int) {
				defer wg.Done()

				if obsX == originalPos[0] && obsY == originalPos[1] {
					return  // can't place obstacle where guard starts
				} else if !possibleVisited[obsY][obsX] {
					return // If original path doesn't reach, placing an obstacle there doesn't matter
				}

				checkIfLoops(board, width, height, originalPos[0], originalPos[1], obsX, obsY, &loopCount) 
			}(obsX, obsY)
		}
	}
	
	wg.Wait()
	return int(loopCount.Load()), nil
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (dir Direction) String(d Direction) string {
	switch d {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	}
	panic("Invalid direction received")
}

func turnRight(d Direction) (turned Direction) {
	switch d {
	case UP:
		turned = RIGHT
	case RIGHT:
		turned = DOWN
	case DOWN:
		turned = LEFT
	case LEFT:
		turned = UP
	}
	return
}

func nextCoordinate(current [2]int, direction Direction) [2]int {
	switch direction {
	case UP:
		return [2]int{current[0], current[1] - 1}
	case RIGHT:
		return [2]int{current[0] + 1, current[1]}
	case DOWN:
		return [2]int{current[0], current[1] + 1}
	case LEFT:
		return [2]int{current[0] - 1, current[1]}
	}
	panic("Unreachable, direction should be exhaustive check")
}

func printBoard(board [][]bool, visited [][]bool, exitPosition [2]int) {
	for y, row := range board {
		for x, wall := range row {
			if wall {
				fmt.Print("#")
			} else if visited[y][x] {
				fmt.Print("X")
			} else if exitPosition[0] == x && exitPosition[1] == y {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func checkIfLoops(board [][]bool, width, height, startX, startY, obsX, obsY int, loopCounter *atomic.Int32) {
	pos := [2]int{startX, startY}

	// Keep track of all directions that we went through a particular position
	// If you are on same tile and same direction, then you're in a loop
	visited := make([][][]Direction, height)
	for y := range height {
		visited[y] = make([][]Direction, width)
		for x := range width {
			visited[y][x] = make([]Direction, 0)
		}
	}

	dir := UP
	visited[pos[1]][pos[0]] = append(visited[pos[1]][pos[0]], dir) 

	for {
		newPos := nextCoordinate(pos, dir)
		x, y := newPos[0], newPos[1]

		if x < 0 || x >= width || y < 0 || y >= height {
			// left board
			return 
		}

		if board[y][x] || (x == obsX && y == obsY) {
			// hit a wall/box
			dir = turnRight(dir)
		} else {
			if slices.Contains(visited[y][x], dir) {
				// Loop detected!
				loopCounter.Add(1)
				return 
			} else {
				// Move to next tile
				pos = newPos
				visited[y][x] = append(visited[y][x], dir)
			}
		}
	}
}

func parseInput(useRealInput bool) (board [][]bool, start [2]int, err error) {
	data, err := util.ReadInputMulti(6, useRealInput)
	if err != nil {
		return nil, [2]int{0, 0}, err
	}
	if len(data) != 1 {
		return nil, [2]int{0, 0}, fmt.Errorf("expected to get 1 section in the input, got %d isntead", len(data))
	}

	board = make([][]bool, len(data[0]))
	for y, row := range data[0] {
		board[y] = make([]bool, len(row))
		for x, c := range row {
			if c == '^' {
				start = [2]int{x, y}
			}
			board[y][x] = (c == '#')
		}
	}
	return
}
