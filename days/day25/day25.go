package day25

import (
	"advent-of-code-2024/util"
	"fmt"
)

func SolvePart1(useRealInput bool) (int, error) {
	locks, keys, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				count++
			}
		}
	}

	return count, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	_, _, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func fit(lock LockKey, key LockKey) bool {
	for i, l := range lock {
		k := key[i]
		if k+l > 7 {
			return false
		}
	}
	return true
}

type LockKey [5]int

func parseInput(useRealInput bool) ([]LockKey, []LockKey, error) {
	data, err := util.ReadInputMulti(25, useRealInput)
	if err != nil {
		return nil, nil, err
	}
	if len(data) <= 1 {
		return nil, nil, fmt.Errorf("expected more than 1 section of input")
	}

	locks := make([]LockKey, 0)
	keys := make([]LockKey, 0)

	for _, entry := range data {
		isLock := true
		for _, k := range entry[0] {
			if k == '.' {
				isLock = false
				break
			}
		}

		configuration := LockKey{}
		for _, row := range entry {
			for i, k := range row {
				if k == '#' {
					configuration[i] = configuration[i] + 1
				}
			}
		}

		if isLock {
			locks = append(locks, configuration)
		} else {
			keys = append(keys, configuration)
		}
	}

	return locks, keys, nil
}
