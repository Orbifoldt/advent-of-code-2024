package day03

import (
	"advent-of-code-2024/util"
	"regexp"
	"strconv"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	data, err := util.ReadInput(3, useRealInput)
	if err != nil {
		return 0, err
	}

	regex, err := regexp.Compile(`mul\((?P<a>\d{1,3}),(?P<b>\d{1,3})\)`)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range data {
		for _, match := range regex.FindAllStringSubmatch(line, -1) {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}
			sum += a * b
		}
	}

	return sum, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	data, err := util.ReadInput(3, useRealInput)
	if err != nil {
		return 0, err
	}

	var sb strings.Builder
	for _, line := range data {
		sb.WriteString(line)
	}
	concatenatedInstructions := sb.String()

	return sumEnabledMultiplications(concatenatedInstructions)

}

func sumEnabledMultiplications(concatenatedInstructions string) (int, error) {
	regex, err := regexp.Compile(`mul\((?P<a>\d{1,3}),(?P<b>\d{1,3})\)`)
	if err != nil {
		return 0, err
	}

	// Idea: first split at all do's. Then split those sections at all dont's and
	// discard everything after a don't
	sum := 0
	split := strings.SplitN(concatenatedInstructions, "do()", -1)
	for _, section := range split {
		doSection, _, _ := strings.Cut(section, "don't()")

		for _, match := range regex.FindAllStringSubmatch(doSection, -1) {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}
			sum += a * b
		}
	}

	return sum, nil
}
