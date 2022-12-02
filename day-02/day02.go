package main

import (
	"fmt"
	"os"
	"path"
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
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while fetching current dir", err)
		return
	}

	rawInput, err := os.ReadFile(path.Join(cwd, "day-02", "02.txt"))
	input := string(rawInput)

	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
