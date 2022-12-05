package day05

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

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

		stackNr, err := strconv.Atoi(nr)

		if err != nil {
			fmt.Println("Failed to parse number")
			os.Exit(1)
		}

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
			parsedNr, err := strconv.Atoi(match)

			if err != nil {
				fmt.Println("Failed to parse number")
			}

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
