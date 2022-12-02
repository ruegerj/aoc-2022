package main

import (
	"fmt"

	"github.com/ruegerj/aoc-2022/util"
)

const Rock = "rock"
const Paper = "paper"
const Scissors = "scissors"

var operandScores = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func main() {
	input := util.LoadDailyInput(2)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
