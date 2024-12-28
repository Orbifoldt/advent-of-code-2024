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
		x := evolve(secret, 2000)
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

	allSecrets, allPrices := make([][]int64, len(secretVals)), make([][]int8, len(secretVals))

	for i, secret := range secretVals {
		s, p := make([]int64, 0, N+1), make([]int8, 0, N+1)
		evolvePt2(secret, N, &s, &p)
		allSecrets[i] = s
		allPrices[i] = p
	}

	allDivs := make([][]int8, len(secretVals))
	for i, prices := range allPrices {
		allDivs[i] = make([]int8, N) // N prices changes => N+1 prices => N divs
		// secrets := allSecrets[i]
		prev := int8(-1)
		for j, p := range prices {
			if prev != int8(-1) {
				div := p - prev
				allDivs[i][j-1] = div
				// fmt.Printf("%9d: %d (%3d)\n", secrets[j], p, div)
			} else {
				// fmt.Printf("%9d: %d\n", secrets[j], p)

			}
			prev = p
		}
	}

	// caches := make([]map[[4]int8]int, len(allPrices))
	cache := make(map[[4]int8]int)
	for i := range allDivs {
		alreadyChecked := make(map[[4]int8]struct{})
		for t := range N - 6 {
			pattern := [4]int8{allDivs[i][t], allDivs[i][t+1], allDivs[i][t+2], allDivs[i][t+3]}
			if _, checked := alreadyChecked[pattern]; !checked {
				cache[pattern] += int(totalBananas(allPrices[i], allDivs[i], pattern))
				alreadyChecked[pattern] = struct{}{}
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

func evolve(secret int64, n int) int64 {
	for range n {
		secret = prune(mix(secret*64, secret))
		secret = prune(mix(secret/32, secret))
		secret = prune(mix(secret*2048, secret))
	}
	return secret
}

func evolvePt2(secret int64, n int, secrets *[]int64, prices *[]int8) int64 {
	*secrets = append(*secrets, secret)
	*prices = append(*prices, int8(secret%10))

	for range n {
		secret = prune(mix(secret*64, secret))
		secret = prune(mix(secret/32, secret))
		secret = prune(mix(secret*2048, secret))

		*secrets = append(*secrets, secret)
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
