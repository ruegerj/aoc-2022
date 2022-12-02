package main

import (
	"fmt"

	"github.com/ruegerj/aoc-2022/util"
)

func main() {
	input := util.LoadDailyInput(1)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
