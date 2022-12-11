package day11

import (
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	const roundCount = 20

	monkeyDefinitions := strings.Split(input, "\n\n")

	monkeys := make([]*Monkey, len(monkeyDefinitions))

	for _, definition := range monkeyDefinitions {
		nr, monkey := NewMonkeyFromDefinition(definition)

		monkeys[nr] = monkey
	}

	for round := 0; round < roundCount; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				worryLevel := monkey.InspectItem(item)
				worryLevel /= 3

				targetNr := monkey.GetTarget(worryLevel)

				monkeys[targetNr].CatchItem(worryLevel)
			}

			monkey.Items = []int{}
		}
	}

	inspectCounts := make([]int, len(monkeys))

	for i, monkey := range monkeys {
		inspectCounts[i] = monkey.InspectCount
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))

	monkeyBusinessScore := inspectCounts[0] * inspectCounts[1]

	return util.NewSolution(1, monkeyBusinessScore)
}
