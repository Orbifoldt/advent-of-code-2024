package day04

import (
	"advent-of-code-2024/util"
	"slices"
)

func SolvePart1(useRealInput bool) (int, error) {
	board, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	width, height := len(board[0]), len(board)

	allowedDirections := util.ClockwiseDiagDirections()

	xmasCount := 0
	for y, line := range board {
		for x, c := range line {
			if c == 'X' {
				start := util.Vec{X: x, Y: y}
				for _, dir := range allowedDirections {
					newPos := start.PlusDirDiag(dir)
					second := get(board, width, height, newPos)
					if second == 'M' {
						newPos.MoveDirDiag(dir)
						if get(board, width, height, newPos) == 'A' {
							newPos.MoveDirDiag(dir)
							if get(board, width, height, newPos) == 'S' {
								// fmt.Printf("Found XMAS starting at (%d, %d) going %v\n", x, y, dir)
								xmasCount++
							}
						}
					}
				}
			}

		}
	}

	return xmasCount, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	board, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	width, height := len(board[0]), len(board)

	allowedDirections := []util.DiagDirection{util.NE, util.SE, util.SW, util.NW}

	xmasCount := 0
	for y, line := range board {
		if y == 0 || y == height-1 {
			continue
		}

		for x, c := range line {
			// X-MASes are a 3x3 square with the center being an A
			if x == 0 || x == width-1 {
				continue
			}

			if c == 'A' {
				start := util.Vec{X: x, Y: y}
				
				// The corners should contain S, S, M and M, or any cyclic permutation thereof
				var chars string
				for _, dir := range allowedDirections {
					chars += string(get(board, width, height, start.PlusDirDiag(dir)))
				}
				permutations := []string{"MMSS", "MSSM", "SSMM", "SMMS"}
				if slices.Contains(permutations, chars) {
					xmasCount++
				}
			}

		}
	}

	return xmasCount, nil
}

func get(board [][]rune, width, height int, coordinate util.Vec) rune {
	if coordinate.X < 0 || coordinate.X >= width || coordinate.Y < 0 || coordinate.Y >= height {
		return 0
	}
	return board[coordinate.Y][coordinate.X]
}

func parseInput(useRealInput bool) ([][]rune, error) {
	data, err := util.ReadInput(4, useRealInput)
	if err != nil {
		return nil, err
	}

	output := make([][]rune, len(data))
	for y, line := range data {
		output[y] = []rune(line)
	}

	return output, nil
}
