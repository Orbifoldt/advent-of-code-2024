package main

import (
	"advent-of-code-2024/days/day01"
	"advent-of-code-2024/days/day02"
	"advent-of-code-2024/days/day03"
	"advent-of-code-2024/days/day04"
	"advent-of-code-2024/days/day05"
	"advent-of-code-2024/days/day06"
	"advent-of-code-2024/days/day07"
	"advent-of-code-2024/days/day08"
	"advent-of-code-2024/days/day09"
	"advent-of-code-2024/days/day10"
	"advent-of-code-2024/days/day11"
	"advent-of-code-2024/days/day12"
	"advent-of-code-2024/days/day13"
	"advent-of-code-2024/days/day14"
	"advent-of-code-2024/days/day15"
	"advent-of-code-2024/days/day16"
	"advent-of-code-2024/days/day17"
	"advent-of-code-2024/days/day18"
	"advent-of-code-2024/days/day19"
	"advent-of-code-2024/days/day20"
	"advent-of-code-2024/days/day21"
	"advent-of-code-2024/days/day22"
	"advent-of-code-2024/days/day23"
	"advent-of-code-2024/days/day24"
	"advent-of-code-2024/days/day25"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Day 01:")
	sol, err := day01.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of distances is: %d\n", sol)
	sol, err = day01.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Similarity score: %d\n", sol)

	fmt.Println("\n\nDay 02:")
	sol, err = day02.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of valid reports is: %d\n", sol)
	sol, err = day02.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: After applying problem dampener, number of valid reports is: %d\n", sol)

	fmt.Println("\n\nDay 03:")
	sol, err = day03.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of multiplications: %d\n", sol)
	sol, err = day03.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Sum of only enabled multiplications: %d\n", sol)

	fmt.Println("\n\nDay 04:")
	sol, err = day04.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of XMASes: %d\n", sol)
	sol, err = day04.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of X-MASes: %d\n", sol)

	fmt.Println("\n\nDay 05:")
	sol, err = day05.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of middle pages of correct updates: %d\n", sol)
	sol, err = day05.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Sum of middle pages of sorted incorrect updates: %d\n", sol)

	fmt.Println("\n\nDay 06:")
	sol, err = day06.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of visited squares: %d\n", sol)
	start06_2 := time.Now()
	sol, err = day06.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start06_2))
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of obstruction positions that cause loops: %d\n", sol)

	fmt.Println("\n\nDay 07:")
	sol64, err := day07.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of valid test values: %d\n", sol64)
	start7_2 := time.Now()
	sol64, err = day07.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start7_2))
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Sum of valid test values with concatenation: %d\n", sol64)

	fmt.Println("\n\nDay 08:")
	sol, err = day08.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of unique antinode locations: %d\n", sol)
	sol, err = day08.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of collinear antinode locations: %d\n", sol)

	fmt.Println("\n\nDay 09:")
	sol64, err = day09.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Checksum after moving file blocks: %d\n", sol64)
	sol64, err = day09.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Checksum after moving whole files: %d\n", sol64)

	fmt.Println("\n\nDay 10:")
	sol, err = day10.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Score of all trailheads: %d\n", sol)
	sol, err = day10.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Score of all trailheads counting unique trails: %d\n", sol)

	fmt.Println("\n\nDay 11:")
	sol64, err = day11.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of stones after 25 times blinking: %d\n", sol64)
	sol64, err = day11.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of stones after 75 times blinking: %d\n", sol64)

	fmt.Println("\n\nDay 12:")
	sol, err = day12.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Total price of the garden: %d\n", sol)
	sol, err = day12.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total price after discount: %d\n", sol)

	fmt.Println("\n\nDay 13:")
	sol, err = day13.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Costs to win all prizes: %d\n", sol)
	sol64, err = day13.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total price after discount: %d\n", sol64)

	fmt.Println("\n\nDay 14:")
	sol, err = day14.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Safety factor: %d\n", sol)
	sol, err = day14.SolvePart2(true, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total price after discount: %d\n", sol)

	fmt.Println("\n\nDay 15:")
	sol, err = day15.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of GPS: %d\n", sol)
	sol, err = day15.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: After widening, sum of GPS: %d\n", sol)

	fmt.Println("\n\nDay 16:")
	sol, err = day16.SolvePart1(true)  // optimize
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Min distance: %d\n", sol)
	sol, err = day16.SolvePart2(true)  // optimize
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of tiles covered by at least one path: %d\n", sol)

	fmt.Println("\n\nDay 17:")
	solStr, err := day17.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: program output: %s\n", solStr)
	start17_2 := time.Now()
	sol64, err = day17.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start17_2))
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: value of A register to self-replicate the program: %d\n", sol64)

	fmt.Println("\n\nDay 18:")
	sol, err = day18.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Min distance: %d\n", sol)
	solStr, err = day18.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Coordinate of byte preventing exit: '%s'\n", solStr)

	fmt.Println("\n\nDay 19:")
	sol, err = day19.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Total number of possible designs: %d\n", sol)
	sol, err = day19.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of ways all designs can be made: %d\n", sol)

	fmt.Println("\n\nDay 20:")
	sol, err = day20.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Total number of cheats saving 100 or more picoseconds: %d\n", sol)
	sol, err = day20.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total number of cheats of up to 20 ps that save >=100ps: %d\n", sol)

	fmt.Println("\n\nDay 21:")
	sol, err = day21.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Total score: %d\n", sol)
	sol, err = day21.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total score with 26 robots: %d\n", sol)

	fmt.Println("\n\nDay 22:")
	sol64, err = day22.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum 2000th secret numbers: %d\n", sol64)
	start22_2 := time.Now()
	sol64, err = day22.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start22_2))
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Total bananas at optimal price: %d\n", sol64)

	fmt.Println("\n\nDay 23:")
	sol, err = day23.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of K3 subgraphs containing vertex starting with t: %d\n", sol)
	solStr, err = day23.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Largest complete subgraph consists of nodes: '%s'\n", solStr)

	fmt.Println("\n\nDay 24:")
	sol, err = day24.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Output of wires starting with z: %d\n", sol)
	fmt.Println("pt2: solution was found visually/manually")

	fmt.Println("\n\nDay 25:")
	sol, err = day25.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Number of key and lock pairs that fit: %d\n", sol)

	fmt.Printf("\n\nIn total all took %s\n", time.Since(start))
}
