package main

import (
	"advent-of-code-2024/days/day01"
	"fmt"
)

func main() {
	fmt.Println("Day 01:")
	sol1, err := day01.SolvePart1(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt1: Sum of distances is: %d\n", sol1)
	sol2, err := day01.SolvePart2(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pt2: Similarity score: %d\n", sol2)
}