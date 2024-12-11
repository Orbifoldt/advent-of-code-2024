package day11

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int64, error) {
	stones, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	countCache := make(map[state]int64, 0)
	total := int64(0)
	for _, stone := range stones {
		total += countStones(countCache, stone, 25)
	}

	return total, nil
}

type state struct {
	stone      int64
	iterations int
}

func countStones(countCache map[state]int64, startStone int64, numIterations int) int64 {
	if numIterations == 0 {
		return 1
	}
	cachedCount := countCache[state{startStone, numIterations}]
	if cachedCount == 0 {
		var count int64
		if startStone == 0 {
			count = countStones(countCache, 1, numIterations-1)
		} else if str := fmt.Sprint(startStone); len(str)%2 == 0 {
			a, b := str[:len(str)/2], str[len(str)/2:]
			x1, err := strconv.ParseInt(a, 10, 64)
			if err != nil {
				panic(err)
			}
			x2, err := strconv.ParseInt(b, 10, 64)
			if err != nil {
				panic(err)
			}
			count = countStones(countCache, x1, numIterations-1) + countStones(countCache, x2, numIterations-1)
		} else {
			count = countStones(countCache, startStone*2024, numIterations-1)
		}
		countCache[state{startStone, numIterations}] = count
		return count
	} else {
		return cachedCount
	}
}

func SolvePart2(useRealInput bool) (int64, error) {
	stones, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	countCache := make(map[state]int64, 0)
	total := int64(0)
	for _, stone := range stones {
		total += countStones(countCache, stone, 75)
	}

	return total, nil
}

func parseInput(useRealInput bool) ([]int64, error) {
	data, err := util.ReadInputMulti(11, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	stones := make([]int64, 0)
	for _, str := range strings.Split(data[0][0], " ") {
		x, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		stones = append(stones, x)
	}

	return stones, nil
}
