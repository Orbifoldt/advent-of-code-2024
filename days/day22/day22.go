package day22

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

func SolvePart1(useRealInput bool) (int64, error) {
	secretVals, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	sum := int64(0)
	for _, secret := range secretVals {
		x := evolve(secret, 2000, &[]int8{})
		sum += x
	}

	return sum, nil
}

func SolvePart2(useRealInput bool) (int64, error) {
	secretVals, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}
	N := 2000

	allPrices := make([][]int8, len(secretVals))
	for i, secret := range secretVals {
		p := make([]int8, 0, N+1)
		evolve(secret, N, &p)
		allPrices[i] = p
	}

	// Calculate all divs
	allDiffs := make([][]int8, len(secretVals))
	for i, prices := range allPrices {
		allDiffs[i] = make([]int8, N) // N prices changes => N+1 prices => N divs
		prev := int8(-1)
		for j, p := range prices {
			if prev != int8(-1) {
				div := p - prev
				allDiffs[i][j-1] = div
			}
			prev = p
		}
	}

	// For each difference-sequence pattern, determine the total price
	cache := make(map[int]int)
	for i := range allDiffs {
		// since a pattern can occur more than once, keep track of the ones we already checked
		alreadyChecked := make(map[int]struct{})
		for t := range N - 6 {
			// since the diffs are 4 int8's, we can make it a single int and use that as cache key
			key := toInt(allDiffs[i][t : t+4])

			if _, checked := alreadyChecked[key]; !checked {
				pattern := [4]int8{allDiffs[i][t], allDiffs[i][t+1], allDiffs[i][t+2], allDiffs[i][t+3]}
				cache[key] += int(totalBananas(allPrices[i], allDiffs[i], pattern))
				alreadyChecked[key] = struct{}{}
			}
		}
	}

	maxTotal := 0
	for _, totalBananas := range cache {
		if totalBananas > maxTotal {
			maxTotal = totalBananas
		}
	}

	return int64(maxTotal), nil
}

func mix(given, secret int64) int64 {
	return given ^ secret
}

func prune(secret int64) int64 {
	return secret % int64(16777216)
}

func evolve(secret int64, n int, prices *[]int8) int64 {
	*prices = append(*prices, int8(secret%10))
	for range n {
		secret = prune(mix(secret*64, secret))
		secret = prune(mix(secret/32, secret))
		secret = prune(mix(secret*2048, secret))
		*prices = append(*prices, int8(secret%10))
	}
	return secret
}

// Find price after first occurrence of pattern price-differences, or 0 if pattern doesn't occur
func totalBananas(prices, divs []int8, pattern [4]int8) int8 {
	for i := 0; i < len(divs)-5; i++ {
		if divs[i] == pattern[0] && divs[i+1] == pattern[1] && divs[i+2] == pattern[2] && divs[i+3] == pattern[3] {
			return prices[i+4] // we sell 1 moment after div i+3
		}
	}
	return 0
}

func toInt(pattern []int8) int {
	key := 0
	for _, p := range pattern {
		key = (key << 8) + int(p)
	}
	return key
}

func parseInput(useRealInput bool) ([]int64, error) {
	data, err := util.ReadInputMulti(22, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	nums := make([]int64, 0)
	for _, line := range data[0] {
		a, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, a)
	}

	return nums, nil
}
