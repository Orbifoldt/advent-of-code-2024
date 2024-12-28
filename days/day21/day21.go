package day21

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

func SolvePart1(useRealInput bool) (int, error) {
	codes, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	totalScore := 0
	for _, code := range codes {
		totalScore += score(code[:], 3)
	}

	return totalScore, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	codes, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	totalScore := 0
	for _, code := range codes {
		totalScore += score(code[:], 26)
	}

	return totalScore, nil
}

var cache = make(map[string]int)

func movesLength(path []int, robot int, topLevel bool) int {
	// Caching results with key of format "level_pathString", e.g. "13_v^A<<AA>A"
	key := fmt.Sprintf("%d_", robot)
	for _, r := range path {
		key += string(toChar(r, false))
	}
	if !topLevel {
		if cachedVal, cached := cache[key]; cached {
			return cachedVal
		}
	}

	// for range 4 - robot {
	// 	fmt.Print(" ")
	// }
	// fmt.Printf("Robot %2d moving %s\n", robot, toString(path, topLevel))

	if robot == 0 {
		return len(path)
	}

	totalLength := 0
	from := A
	for _, to := range path {
		if from == to {
			// Just press the same button again by pressing A
			totalLength += 1

			// for range 4 - robot + 1 {
			// 	fmt.Print(" ")
			// }
			// fmt.Printf("Robot %2d moving A\n", robot-1)
		} else {
			// Else, see what a robot one level down needs to do
			movesOneLevelDown := optimalMoveCached(from, to, topLevel)
			length := movesLength(movesOneLevelDown, robot-1, false)

			totalLength += length
		}
		from = to
	}

	// pathString := toString(path, topLevel)
	// check := getSequenceLength(pathString, robot)
	// if totalLength != check {
	// 	panic(fmt.Errorf("oh no: %d isn't equal to check value %d for %s", totalLength, check, pathString))
	// }

	if !topLevel {
		cache[key] = totalLength
	}

	return totalLength
}

func score(code []int, robot int) int {
	numericPart := 0
	for _, c := range code {
		if c == A {
			break
		}
		numericPart = numericPart*10 + c
	}
	// fmt.Printf("%dA:\n", numericPart)

	numMoves := movesLength(code, robot, true)
	return numericPart * numMoves
}

// var numKeypadPositions map[util.Vec]int
// var numKeypadPositionsReverse map[int]util.Vec

// var dirKeypadPositions map[util.Vec]int
// var dirKeypadPositionsReverse map[int]util.Vec

// func optimalMove(from, to int, numerical bool) []int {
// 	if numKeypadPositions == nil {
// 		numKeypadPositions, numKeypadPositionsReverse = numericKeypadPositions()
// 	}
// 	if dirKeypadPositions == nil {
// 		dirKeypadPositions, dirKeypadPositionsReverse = directionalKeypadPositions()
// 	}

// 	var fromVec util.Vec
// 	var toVec util.Vec
// 	if numerical {
// 		fromVec = numKeypadPositionsReverse[from]
// 		toVec = numKeypadPositionsReverse[to]
// 	} else {
// 		fromVec = dirKeypadPositionsReverse[from]
// 		toVec = dirKeypadPositionsReverse[to]
// 	}

// 	divVec := toVec.Minus(fromVec)

// 	// We should alway move the maximal distance in one direction to minimize moves a level up
// 	// Because of missing corners the order does matter, so to not get stuck we use: RIGHT, DOWN, UP, LEFT
// 	moves := make([]int, 0)
// 	if !numerical {
// 		if divVec.X > 0 {
// 			for range divVec.X {
// 				moves = append(moves, int(util.RIGHT))
// 			}
// 		}
// 		if divVec.Y > 0 {
// 			for range divVec.Y {
// 				moves = append(moves, int(util.DOWN))
// 			}
// 		}
// 		if divVec.Y < 0 {
// 			for range -divVec.Y {
// 				moves = append(moves, int(util.UP))
// 			}
// 		}
// 		if divVec.X < 0 {
// 			for range -divVec.X {
// 				moves = append(moves, int(util.LEFT))
// 			}
// 		}
// 	} else {
// 		if divVec.Y < 0 {
// 			for range -divVec.Y {
// 				moves = append(moves, int(util.UP))
// 			}
// 		}
// 		if divVec.X < 0 {
// 			for range -divVec.X {
// 				moves = append(moves, int(util.LEFT))
// 			}
// 		}
// 		if divVec.X > 0 {
// 			for range divVec.X {
// 				moves = append(moves, int(util.RIGHT))
// 			}
// 		}
// 		if divVec.Y > 0 {
// 			for range divVec.Y {
// 				moves = append(moves, int(util.DOWN))
// 			}
// 		}
// 	}

// 	moves = append(moves, A)
// 	return moves
// }

type move struct {
	from, to  int
	numerical bool
}

// var optimalMoveCache = make(map[move][]int)

func optimalMoveCached(from, to int, numerical bool) []int {
	key := move{from, to, numerical}
	return paths2[key]
	// if from == to {
	// 	return []int{}
	// }
	// if cacheResult, cached := optimalMoveCache[key]; cached {
	// 	return cacheResult
	// }
	// calculatedResult := optimalMove(from, to, numerical)
	// optimalMoveCache[key] = calculatedResult
	// return calculatedResult

	// a, b := toChar(from, numerical), toChar(to, numerical)
	// pathString := paths[buttonPair{a, b}]
	// // fmt.Printf(">> a=%s, b=%s    %s\n", string(a), string(b), pathString)

	// pathIntSlice := make([]int, 0)
	// for _, r := range pathString {
	// 	pathIntSlice = append(pathIntSlice, toInt(r, false))
	// }
	// return pathIntSlice
}

const A int = 10

// func numericKeypadPositions() (map[util.Vec]int, map[int]util.Vec) {
// 	// +---+---+---+
// 	// | 7 | 8 | 9 |
// 	// +---+---+---+
// 	// | 4 | 5 | 6 |
// 	// +---+---+---+
// 	// | 1 | 2 | 3 |
// 	// +---+---+---+
// 	//     | 0 | A |
// 	//     +---+---+
// 	positions := make(map[util.Vec]int)
// 	for y := range 3 {
// 		for x := range 3 {
// 			positions[util.Vec{X: x, Y: y}] = x + 1 + 3*(2-y)
// 		}
// 	}
// 	positions[util.Vec{X: 1, Y: 3}] = 0
// 	positions[util.Vec{X: 2, Y: 3}] = A
// 	reversePositions := make(map[int]util.Vec)
// 	for k, v := range positions {
// 		reversePositions[v] = k
// 	}
// 	return positions, reversePositions
// }

// func directionalKeypadPositions() (map[util.Vec]int, map[int]util.Vec) {
// 	//     +---+---+
// 	//     | ^ | A |
// 	// +---+---+---+
// 	// | < | v | > |
// 	// +---+---+---+
// 	positions := make(map[util.Vec]int)
// 	positions[util.Vec{X: 1, Y: 0}] = int(util.UP)
// 	positions[util.Vec{X: 2, Y: 0}] = A
// 	positions[util.Vec{X: 0, Y: 1}] = int(util.LEFT)
// 	positions[util.Vec{X: 1, Y: 1}] = int(util.DOWN)
// 	positions[util.Vec{X: 2, Y: 1}] = int(util.RIGHT)
// 	reversePositions := make(map[int]util.Vec)
// 	for k, v := range positions {
// 		reversePositions[v] = k
// 	}
// 	return positions, reversePositions
// }

func parseInput(useRealInput bool) ([][4]int, error) {
	data, err := util.ReadInputMulti(21, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	out := make([][4]int, 0)
	for _, row := range data[0] {
		rowOut := [4]int{}
		for i, r := range row {
			if r == 'A' {
				rowOut[i] = A
			} else {
				integer, err := strconv.Atoi(string(r))
				if err != nil {
					return nil, err
				}
				rowOut[i] = integer
			}
		}
		out = append(out, rowOut)
	}

	return out, nil
}

// func toString(path []int, isNumerical bool) string {
// 	runes := make([]rune, 0)
// 	var sb strings.Builder
// 	for _, p := range path {
// 		sb.WriteRune(toChar(p, isNumerical))
// 		runes = append(runes, toChar(p, isNumerical))
// 	}
// 	return string(runes)
// }

func toChar(m int, isNumerical bool) rune {
	var c rune
	if !isNumerical {
		switch m {
		case 0:
			c = '^'
		case 1:
			c = '>'
		case 2:
			c = 'v'
		case 3:
			c = '<'
		case A:
			c = 'A'
		default:
			c = []rune(fmt.Sprintf("%d", m))[0]
		}
	} else {
		if m == A {
			c = 'A'
		} else {
			c = []rune(fmt.Sprintf("%d", m))[0]
		}
	}
	return c
}

// func toInt(c rune, isNumerical bool) int {
// 	if isNumerical {
// 		if c == 'A' {
// 			return A
// 		} else {
// 			x, _ := strconv.Atoi(string(c))
// 			return x
// 		}
// 	} else {
// 		switch c {
// 		case '^':
// 			return 0
// 		case '>':
// 			return 1
// 		case 'v':
// 			return 2
// 		case '<':
// 			return 3
// 		case 'A':
// 			return A
// 		}
// 	}
// 	return -1
// }

// Was stuck, adjusted this from https://topaz.github.io/paste/#XQAAAQD6EQAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDinguAj8sFyD4ftJ8OW5ekoqnVIaEYm1TRzRozXdrSWph6uPxJTZjFA79rnO0DuC0UuIh7ERk/Duk3psptqVA73J8Z4I2hW2SL6gJB3Q1/1XR38DhDEO8md9rWyVWBo4CJDydIWxuyMJcWVQ1ufgwX0ZmJoE5ZQfJjzbh+DZJ+rn4Sbosya9WNQ6/9qJVmsYOKTaGQSnTXeXuKSxXcsEM8aOYrwwokHmz1qm/XBD0xY3AAVJTkzrYdTl4OQox1mjh84ro0qEU4/rUv+M6fqUGeLv9d6MjvSvdq9zb3kVlUg0EIXhQNEfOjoUDdnZjwo6W1fT4LIOQRKQDTXkLNCovRuWLphjKGCHFmcJyrSRLcBrpa3UPBz2cIWDABX80asng9GqkqlBPZ00Xwo7HcF2g/+CbHp5krjy5MghjW+IFMXZ6KzO3bqOGmmXmXmAJbcPpaASmOyJuZAIqZPlGWT8cfySbmkncDrAdGAl5s9SaZp+ouygkV74I+YpoBLf0T+bbJ4jHatgdyEeX+hb35+eN08OB0E0S06WUzSAHJfIJV6QqkQAhlu3aJTiam9LsuRQa8l65jvSVum1zympU0thS5R73tSq43ExG3n7ZoPEt54+GoAqWEb8DYTiiSaLSL0E+aDYNsxeBhJNU8aw3+TPNHS5AOzEHuKfY1fAfcvsx8tcpeEB1bSGtPhmjbDOJtnzhdS76jrV68K4hjgzpqJ/YfVeoIhVBPJpnc9AuNMyMECjKIN9gOhKqzooSr6teDDrCi6EFCvwAWoA9BQtKq+m4nH0QPUznsSCp4WVVzzeuuUQU9VKbtJClZlWGdFOp+iaoaZg3jec+rsSrPbW/Egz9Wz6G+BT+ehdYjxUjkSj27qxJK0O+0prbEkq1E4rjNgZIq0l1n76ERO7TPmYhDOzOP/URcnhJJ27nB6eeeABAlr7ovC+1nFP51FNvugYf8NvGQ1T5Tk/X1yGvomNHwHvvJ6PWDQn4FgDPQj1f6AEqnD1aStgkAszG7CSoF1shmf6ODrEcuELJ+UcVUrD2QFMiIqpi5xaez/noF5INiALPoYedsYG9cQ1ogzqBK80652Q2JK7YxdAX9N2ANeYF9tdZ6kY4Gh4t9sIa16T1ujIPde/SFaEFy/uxVkjwMsxUMdLCyIKJJ9YtPI9Ahfmt+3HziHQHM3IQ8sqTlPnyFPp2JLazKBkPlOuaZ1Q+xFRmTf5urms6GyI4GZ6nsjRCoGqJ391vaD8LhgSUlbTVPLgPolti1H4WVKv/DMee2bSRnpmlbff9/BzPP/qbip1F1iw2Za9WZ7ycf4zMn2/zK0VW8OtNAfRTStX3wqWNStrdPSPppYtRSnsHTkGYND3q5lLI0JyEJNu11IOiM2XQMVzF6QgTOuiCx8qKdFG+pUtWL2onkj2iNOvURnsU0jTqWH2rCVdG6T9JeXEjZno7bABLfBmyMNBB+v9ajn+Nv0pW7Ag9cHPgeogkZVvx6zaJTpPQQ/cD2juyVFpEsdBmXEqaIVrKOCIUYnZUwEmEF3zeDqdJZqlYwrQcQzD4xX4SpDxWRXDTR24Rg7zj/vuA/eO1eb4IA/qLxhltMAb7VLqZOU0imaKmDZHOjTP14WoRXD/uy6Y7p9BThBCmg+fX3QXFu5HsCgN82aaAbKJ6laKX2CbA6//SMmnE
// and that worked...
var paths2 = map[move][]int{
	// (from, to, numerical) -> moves[]
	{0, 1, false}:  {2, 1, 10},
	{0, 1, true}:   {0, 3, 10},
	{0, 2, false}:  {2, 10},
	{0, 2, true}:   {0, 10},
	{0, 3, false}:  {2, 3, 10},
	{0, 3, true}:   {0, 1, 10},
	{0, 4, true}:   {0, 3, 0, 10},
	{0, 5, true}:   {0, 0, 10},
	{0, 6, true}:   {0, 0, 1, 10},
	{0, 7, true}:   {0, 0, 0, 3, 10},
	{0, 8, true}:   {0, 0, 0, 10},
	{0, 9, true}:   {0, 0, 0, 1, 10},
	{0, 10, false}: {1, 10},
	{0, 10, true}:  {1, 10},
	{1, 0, false}:  {3, 0, 10},
	{1, 0, true}:   {1, 2, 10},
	{1, 2, false}:  {3, 10},
	{1, 2, true}:   {1, 10},
	{1, 3, false}:  {3, 3, 10},
	{1, 3, true}:   {1, 1, 10},
	{1, 4, true}:   {0, 10},
	{1, 5, true}:   {0, 1, 10},
	{1, 6, true}:   {0, 1, 1, 10},
	{1, 7, true}:   {0, 0, 10},
	{1, 8, true}:   {0, 0, 1, 10},
	{1, 9, true}:   {0, 0, 1, 1, 10},
	{1, 10, false}: {0, 10},
	{1, 10, true}:  {1, 1, 2, 10},
	{2, 0, false}:  {0, 10},
	{2, 0, true}:   {2, 10},
	{2, 1, false}:  {1, 10},
	{2, 1, true}:   {3, 10},
	{2, 3, false}:  {3, 10},
	{2, 3, true}:   {1, 10},
	{2, 4, true}:   {3, 0, 10},
	{2, 5, true}:   {0, 10},
	{2, 6, true}:   {0, 1, 10},
	{2, 7, true}:   {3, 0, 0, 10},
	{2, 8, true}:   {0, 0, 10},
	{2, 9, true}:   {0, 0, 1, 10},
	{2, 10, false}: {0, 1, 10},
	{2, 10, true}:  {2, 1, 10},
	{3, 0, false}:  {1, 0, 10},
	{3, 0, true}:   {3, 2, 10},
	{3, 1, false}:  {1, 1, 10},
	{3, 1, true}:   {3, 3, 10},
	{3, 2, false}:  {1, 10},
	{3, 2, true}:   {3, 10},
	{3, 4, true}:   {3, 3, 0, 10},
	{3, 5, true}:   {3, 0, 10},
	{3, 6, true}:   {0, 10},
	{3, 7, true}:   {3, 3, 0, 0, 10},
	{3, 8, true}:   {3, 0, 0, 10},
	{3, 9, true}:   {0, 0, 10},
	{3, 10, false}: {1, 1, 0, 10},
	{3, 10, true}:  {2, 10},
	{4, 0, true}:   {1, 2, 2, 10},
	{4, 1, true}:   {2, 10},
	{4, 2, true}:   {2, 1, 10},
	{4, 3, true}:   {2, 1, 1, 10},
	{4, 5, true}:   {1, 10},
	{4, 6, true}:   {1, 1, 10},
	{4, 7, true}:   {0, 10},
	{4, 8, true}:   {0, 1, 10},
	{4, 9, true}:   {0, 1, 1, 10},
	{4, 10, true}:  {1, 1, 2, 2, 10},
	{5, 0, true}:   {2, 2, 10},
	{5, 1, true}:   {3, 2, 10},
	{5, 2, true}:   {2, 10},
	{5, 3, true}:   {2, 1, 10},
	{5, 4, true}:   {3, 10},
	{5, 6, true}:   {1, 10},
	{5, 7, true}:   {3, 0, 10},
	{5, 8, true}:   {0, 10},
	{5, 9, true}:   {0, 1, 10},
	{5, 10, true}:  {2, 2, 1, 10},
	{6, 0, true}:   {3, 2, 2, 10},
	{6, 1, true}:   {3, 3, 2, 10},
	{6, 2, true}:   {3, 2, 10},
	{6, 3, true}:   {2, 10},
	{6, 4, true}:   {3, 3, 10},
	{6, 5, true}:   {3, 10},
	{6, 7, true}:   {3, 3, 0, 10},
	{6, 8, true}:   {3, 0, 10},
	{6, 9, true}:   {0, 10},
	{6, 10, true}:  {2, 2, 10},
	{7, 0, true}:   {1, 2, 2, 2, 10},
	{7, 1, true}:   {2, 2, 10},
	{7, 2, true}:   {2, 2, 1, 10},
	{7, 3, true}:   {2, 2, 1, 1, 10},
	{7, 4, true}:   {2, 10},
	{7, 5, true}:   {2, 1, 10},
	{7, 6, true}:   {2, 1, 1, 10},
	{7, 8, true}:   {1, 10},
	{7, 9, true}:   {1, 1, 10},
	{7, 10, true}:  {1, 1, 2, 2, 2, 10},
	{8, 0, true}:   {2, 2, 2, 10},
	{8, 1, true}:   {3, 2, 2, 10},
	{8, 2, true}:   {2, 2, 10},
	{8, 3, true}:   {2, 2, 1, 10},
	{8, 4, true}:   {3, 2, 10},
	{8, 5, true}:   {2, 10},
	{8, 6, true}:   {2, 1, 10},
	{8, 7, true}:   {3, 10},
	{8, 9, true}:   {1, 10},
	{8, 10, true}:  {2, 2, 2, 1, 10},
	{9, 0, true}:   {3, 2, 2, 2, 10},
	{9, 1, true}:   {3, 3, 2, 2, 10},
	{9, 2, true}:   {3, 2, 2, 10},
	{9, 3, true}:   {2, 2, 10},
	{9, 4, true}:   {3, 3, 2, 10},
	{9, 5, true}:   {3, 2, 10},
	{9, 6, true}:   {2, 10},
	{9, 7, true}:   {3, 3, 10},
	{9, 8, true}:   {3, 10},
	{9, 10, true}:  {2, 2, 2, 10},
	{10, 0, false}: {3, 10},
	{10, 0, true}:  {3, 10},
	{10, 1, false}: {2, 10},
	{10, 1, true}:  {0, 3, 3, 10},
	{10, 2, false}: {3, 2, 10},
	{10, 2, true}:  {3, 0, 10},
	{10, 3, false}: {2, 3, 3, 10},
	{10, 3, true}:  {0, 10},
	{10, 4, true}:  {0, 0, 3, 3, 10},
	{10, 5, true}:  {3, 0, 0, 10},
	{10, 6, true}:  {0, 0, 10},
	{10, 7, true}:  {0, 0, 0, 3, 3, 10},
	{10, 8, true}:  {3, 0, 0, 0, 10},
	{10, 9, true}:  {0, 0, 0, 10},
}
