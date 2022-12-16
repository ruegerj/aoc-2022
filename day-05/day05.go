package day05

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
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

	return util.NewSolution(1, getTopLevelCrates(stacks))
}

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

func parseStacks(input string) []*stack.Stack {
	stackLines := strings.Split(input, "\n")

	for i, j := 0, len(stackLines)-1; i < j; i, j = i+1, j-1 {
		stackLines[i], stackLines[j] = stackLines[j], stackLines[i]
	}

	stacks := make([]*stack.Stack, 0)
	stackNumbers := strings.Split(stackLines[0], " ")

	for _, nr := range stackNumbers {
		if nr == "" {
			continue
		}

		stackNr := util.MustParseInt(nr)

		offset := strings.IndexRune(stackLines[0], rune(nr[0]))

		if len(stacks) < stackNr {
			stacks = append(stacks, stack.New())
		}

		currentStack := stacks[stackNr-1]

		for j := 1; j < len(stackLines); j++ {
			line := stackLines[j]

			if len(line) <= offset {
				continue
			}

			crate := line[offset : offset+1]

			if crate == " " {
				continue
			}

			currentStack.Push(crate)
		}
	}

	return stacks
}

func parseInstructions(input string) [][]int {
	matcherExp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	parsedInstructions := make([][]int, 0)

	for _, instruction := range strings.Split(input, "\n") {
		matches := matcherExp.FindStringSubmatch(instruction)[1:]

		parsedMatches := make([]int, 3)

		for i, match := range matches {
			parsedNr := util.MustParseInt(match)

			parsedMatches[i] = parsedNr
		}

		parsedInstructions = append(parsedInstructions, parsedMatches)
	}

	return parsedInstructions
}

func getTopLevelCrates(stacks []*stack.Stack) string {
	topLevelCrates := ""

	for _, stack := range stacks {
		topLevelCrates += fmt.Sprint(stack.Peek())
	}

	return topLevelCrates
}
