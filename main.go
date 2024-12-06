package main

import (
	"advent-of-code-2024/days/day01"
	"advent-of-code-2024/days/day02"
	"advent-of-code-2024/days/day03"
	"advent-of-code-2024/days/day05"
	"advent-of-code-2024/days/day06"
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

	// TODO: day 4

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
	start := time.Now()
	sol, err = day06.SolvePart2(true)
	fmt.Printf("Took %s\n", time.Since(start))
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Number of obstruction positions that cause loops: %d\n", sol)

}
