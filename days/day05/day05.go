package day05

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	_, correct, _, err := getCorrectAndIncorrectPages(useRealInput)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, update := range correct {
		sum += update[len(update)/2]
	} 

	return sum, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	rules, _, incorrect, err := getCorrectAndIncorrectPages(useRealInput)
	if err != nil {
		return 0, err
	}

	// Sort each update using bubble sort
	// Assuming there is a unique solution the rules define a well-ordering, hence this should work
	sum := 0
	for _, update := range incorrect {
		out: for {
			for _, rule := range rules {
				a := rule[0]
				b := rule[1]
				idxA := slices.Index(update, a)
				idxB := slices.Index(update, b)

				// If a and b appear in incorrect order, swap
				if idxA != -1 && idxB != -1 && idxA > idxB {
					update[idxA] = b
					update[idxB] = a
					continue out
				}
			}

			// Now, update is sorted
			sum += update[len(update)/2]
			break out
		}
	}
	
	return sum, nil
}

func getCorrectAndIncorrectPages(useRealInput bool) (rules [][2]int, correct [][]int, incorrect [][]int, err error) {
	rules, updates, err := parseInput(useRealInput)
	if err != nil {
		return nil, nil, nil, err
	}

	for _, rule := range rules {
		for idx, update := range updates {
			if update != nil {
				a := slices.Index(update, rule[0])
				b := slices.Index(update, rule[1])
				if a != -1 && b != -1 && a > b {
					incorrect = append(incorrect, update)
					updates[idx] = nil
				}
			}
		}
	}

	for _, update := range updates {
		if update != nil {
			correct = append(correct, update)
		}
	}

	return rules, correct, incorrect, nil
}

func parseInput(useRealInput bool) ([][2]int, [][]int, error) {
	data, err := util.ReadInputMulti(5, useRealInput)
	if err != nil {
		return nil, nil, err
	}
	if len(data) != 2 {
		return nil, nil, fmt.Errorf("expected input data to contain 2 sections, got %d", len(data))
	}

	rulesStrings := data[0]
	updatesStrings := data[1]

	rules := make([][2]int, len(rulesStrings))
	for idx, rule := range rulesStrings {
		a, b, _ := strings.Cut(rule, "|")
		ai, err := strconv.Atoi(a)
		if err != nil {
			return nil, nil, err
		}
		bi, err := strconv.Atoi(b)
		if err != nil {
			return nil, nil, err
		}
		rules[idx] = [2]int{ai, bi}
	}

	updates := make([][]int, len(updatesStrings))
	for idx, updateString := range updatesStrings {
		update := make([]int, 0)
		for _, updateEntryString := range strings.Split(updateString, ",") {
			updateEntry, err := strconv.Atoi(updateEntryString)
			if err != nil {
				return nil, nil, err
			}
			update = append(update, updateEntry)
		}
		updates[idx] = update
	}
	return rules, updates, nil
}

