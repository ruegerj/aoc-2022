package day05

import (
	"strings"
)

func Part1(input string) string {
	parts := strings.Split(input, "\n\n")

	stacks := parseStacks(parts[0])
	instructions := parseInstructions(parts[1])

	for _, instruction := range instructions {
		count := instruction[0]
		from := stacks[instruction[1]-1]
		to := stacks[instruction[2]-1]

		for i := 0; i < count; i++ {
			crate := from.Pop()
			to.Push(crate)
		}
	}

	return getTopLevelCrates(stacks)
}
