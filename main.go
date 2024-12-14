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
	"fmt"
	"time"
)

func main() {
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
	start := time.Now()
	sol64, err = day07.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start))
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
}
