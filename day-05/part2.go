package day05

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) *util.Solution {
	parts := strings.Split(input, "\n\n")

	stacks := parseStacks(parts[0])
	instructions := parseInstructions(parts[1])

	for _, instruction := range instructions {
		count := instruction[0]
		from := stacks[instruction[1]-1]
		to := stacks[instruction[2]-1]

		batch := make([]interface{}, count)

		for i := count - 1; i >= 0; i-- {
			batch[i] = from.Pop()
		}

		for i := 0; i < count; i++ {
			to.Push(batch[i])
		}
	}

	return util.NewSolution(2, getTopLevelCrates(stacks))
}
