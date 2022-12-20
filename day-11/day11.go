package day11

import (
	"regexp"
	"sort"
	"strconv"
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

func Part2(input string) *util.Solution {
	const roundCount = 10000
	limit := 1

	monkeyDefinitions := strings.Split(input, "\n\n")

	monkeys := make([]*Monkey, len(monkeyDefinitions))

	for _, definition := range monkeyDefinitions {
		nr, monkey := NewMonkeyFromDefinition(definition)
		limit *= monkey.ThrowTest.DivisibleBy

		monkeys[nr] = monkey
	}

	for round := 0; round < roundCount; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				worryLevel := monkey.InspectItem(item) % limit

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

	return util.NewSolution(2, monkeyBusinessScore)
}

type Monkey struct {
	Items                []int
	InspectCount         int
	WorryLevelCalculator *WorryLevelCalculator
	ThrowTest            *ThrowTest
}

type WorryLevelCalculator struct {
	LeftOperand  string
	Operator     string
	RightOperand string
}

type ThrowTest struct {
	DivisibleBy  int
	PassedTarget int
	FailedTarget int
}

func NewMonkeyFromDefinition(definition string) (int, *Monkey) {
	monkeyDataMatcher := regexp.MustCompile(`^Monkey (?P<nr>\d+):\n +Starting items: (?P<items>[\d, ]+)\n +Operation: new = (?P<leftOperand>[old\d]+) (?P<operator>[+*]{1}) (?P<rightOperand>[old\d]+)\n +Test: divisible by (?P<divisibleBy>\d+)\n +If true: throw to monkey (?P<passedTarget>\d+)\n +If false: throw to monkey (?P<failedTarget>\d+)$`)

	definitionArgs := util.MatchNamedSubgroups(monkeyDataMatcher, definition)

	monkey := Monkey{
		InspectCount: 0,
		Items:        make([]int, 0),
		WorryLevelCalculator: &WorryLevelCalculator{
			LeftOperand:  definitionArgs["leftOperand"],
			Operator:     definitionArgs["operator"],
			RightOperand: definitionArgs["rightOperand"],
		},
		ThrowTest: &ThrowTest{
			DivisibleBy:  util.MustParseInt(definitionArgs["divisibleBy"]),
			PassedTarget: util.MustParseInt(definitionArgs["passedTarget"]),
			FailedTarget: util.MustParseInt(definitionArgs["failedTarget"]),
		},
	}

	items := strings.Split(definitionArgs["items"], ", ")

	for _, item := range items {
		monkey.Items = append(monkey.Items, util.MustParseInt(item))
	}

	return util.MustParseInt(definitionArgs["nr"]), &monkey
}

func (monkey *Monkey) CatchItem(item int) {
	monkey.Items = append(monkey.Items, item)
}

func (monkey *Monkey) InspectItem(item int) int {
	monkey.InspectCount++

	left, leftErr := strconv.Atoi(monkey.WorryLevelCalculator.LeftOperand)

	if leftErr != nil {
		left = item
	}

	right, rightErr := strconv.Atoi(monkey.WorryLevelCalculator.RightOperand)

	if rightErr != nil {
		right = item
	}

	if monkey.WorryLevelCalculator.Operator == "*" {
		return left * right
	}

	return left + right
}

func (monkey *Monkey) GetTarget(item int) int {
	if item%int(monkey.ThrowTest.DivisibleBy) != 0 {
		return monkey.ThrowTest.FailedTarget
	}

	return monkey.ThrowTest.PassedTarget
}
