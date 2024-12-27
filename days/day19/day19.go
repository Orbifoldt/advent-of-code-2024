package day19

import (
	"advent-of-code-2024/util"
	"fmt"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	availablePatterns, designs, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	possibleCount := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if isPossible(design, availablePatterns, cache) {
			possibleCount++
		}
	}

	return possibleCount, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	availablePatterns, designs, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	possibleCount := 0
	cache := make(map[string]int)
	for _, design := range designs {
		possibleCount += countPossibilities(design, availablePatterns, cache)
	}

	return possibleCount, nil
}

func isPossible(design string, availablePatterns []string, cache map[string]bool) bool {
	if poss, wasChecked := cache[design]; wasChecked {
		return poss
	}

	if len(design) == 0 {
		return true
	}

	for _, pattern := range availablePatterns {
		tail, hasPrefix := strings.CutPrefix(design, pattern)

		if hasPrefix {
			poss := isPossible(tail, availablePatterns, cache)
			if poss {
				cache[design] = true
				return true
			}
		}
	}
	cache[design] = false
	return false
}

func countPossibilities(design string, availablePatterns []string, cache map[string]int) int {
	if poss, wasChecked := cache[design]; wasChecked {
		return poss
	}

	if len(design) == 0 {
		return 1
	}

	differentWays := 0
	for _, pattern := range availablePatterns {
		tail, hasPrefix := strings.CutPrefix(design, pattern)

		if hasPrefix {
			differentWays += countPossibilities(tail, availablePatterns, cache)
		}
	}
	cache[design] = differentWays
	return differentWays
}

func parseInput(useRealInput bool) ([]string, []string, error) {
	data, err := util.ReadInputMulti(19, useRealInput)
	if err != nil {
		return nil, nil, err
	}
	if len(data) != 2 {
		return nil, nil, fmt.Errorf("expected two sections of input")
	}

	availablePatterns := strings.Split(data[0][0], ", ")

	designs := data[1]

	return availablePatterns, designs, nil
}
