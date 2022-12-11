package day11

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

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

	match := monkeyDataMatcher.FindStringSubmatch(definition)

	definitionArgs := make(map[string]string)
	for i, name := range monkeyDataMatcher.SubexpNames() {
		if i > 0 && i <= len(match) {
			definitionArgs[name] = match[i]
		}
	}

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
