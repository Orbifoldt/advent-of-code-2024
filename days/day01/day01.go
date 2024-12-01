package day01

import (
	"advent-of-code-2024/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	firstList, secondList, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDistance := 0
	for idx, a := range firstList {
		b := secondList[idx]
		if b >= a {
			totalDistance += b - a
		} else {
			totalDistance += a - b
		}
	}

	return totalDistance, nil
}


func SolvePart2(useRealInput bool) (int, error) {
	left, right, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	rightFrequency := make(map[int]int)
	for _, entry := range right {
		rightFrequency[entry] += 1  // If entry doesn't exist, it returns 0, so this works
	}

	similarity := 0
	for _, entry := range left {
		similarity += entry * rightFrequency[entry]
	}

	return similarity, nil
}


func parseInput(useRealInput bool) ([]int, []int, error) {
	data, err := util.ReadInput(1, useRealInput)
	if err != nil {
		return nil, nil, err
	}

	var firstList []int
	var secondList []int
	for _, line := range data {
		a, b, _ := strings.Cut(line, " ")

		ai, err := strconv.Atoi(strings.TrimSpace(a))
		if err != nil {
			return nil, nil, err
		}
		bi, err := strconv.Atoi(strings.TrimSpace(b))
		if err != nil {
			return nil, nil, err
		}

		firstList = append(firstList, ai)
		secondList = append(secondList, bi)
	}
	if len(firstList) != len(secondList) {
		return nil, nil, fmt.Errorf("the two lists have unequal lengths: %d and %d respectively", len(firstList), len(secondList))
	}
	return firstList, secondList, nil

}