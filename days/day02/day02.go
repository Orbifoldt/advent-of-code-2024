package day02

import (
	"advent-of-code-2024/util"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	reports, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	totalValid := 0
	for _, report := range reports {
		if isValid(report) {
			totalValid++
		}

	}

	return totalValid, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	reports, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	totalValid := 0
	for _, report := range reports {
		if isValid(report) {
			totalValid++
			continue
		}

		for i := range len(report){
			ithRemoved := copyWithoutEntry(report, i)
			if isValid(ithRemoved) {
				totalValid++
				break
			}
		}
	}

	return totalValid, nil
}

func copyWithoutEntry(slice []int, indexToRemove int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice[:indexToRemove]...)
	newSlice = append(newSlice, slice[indexToRemove+1:]...)
	return newSlice
}

func isValid(report []int) bool {
	var isIncreasing bool
	for i := range len(report) - 1 {
		a := report[i]
		b := report[i+1]

		if a == b {
			// fmt.Printf("Two subsequent values are equal at index %d\n", i)
			return false
		}

		if i == 0 {
			isIncreasing = b > a
		} else {
			if (isIncreasing && b < a) || (!isIncreasing && b > a) {
				// fmt.Printf("Not strictly monotonic at index %d\n", i)
				return false
			}
		}

		dif := b - a
		maxDiff := 3
		if dif < -maxDiff || dif > maxDiff {
			// fmt.Printf("Dif larger than %d at index %d\n", maxDiff, i)
			return false
		}

	}
	return true
}

func parseInput(useRealInput bool) ([][]int, error) {
	data, err := util.ReadInput(2, useRealInput)
	if err != nil {
		return nil, err
	}

	var reports [][]int
	for _, line := range data {
		reportEntries := strings.Split(line, " ")
		var report []int
		for _, entry := range reportEntries {
			entry = strings.TrimSpace(entry)
			if len(entry) > 0 {
				level, err := strconv.Atoi(entry)
				if err != nil {
					return nil, err
				}
				report = append(report, level)
			}
		}

		reports = append(reports, report)
	}
	return reports, nil
}
