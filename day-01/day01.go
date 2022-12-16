package day01

import (
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	if !strings.HasSuffix(input, "\n") {
		input += "\n"
	}

	food := strings.Split(input, "\n")

	maxCalories := 0
	caloriesSum := 0

	for _, calories := range food {
		if calories == "" {
			if caloriesSum > maxCalories {
				maxCalories = caloriesSum
			}

			caloriesSum = 0
			continue
		}

		caloriesSum += util.MustParseInt(calories)
	}

	return util.NewSolution(1, maxCalories)
}

func Part2(input string) *util.Solution {
	if !strings.HasSuffix(input, "\n") {
		input += "\n"
	}

	food := strings.Split(input, "\n")

	caloriesSum := make([]int, 0)
	caloriesOfElf := 0

	for _, calories := range food {
		if calories == "" {
			caloriesSum = append(caloriesSum, caloriesOfElf)
			caloriesOfElf = 0
			continue
		}

		caloriesOfElf += util.MustParseInt(calories)
	}

	elfCount := len(caloriesSum)
	sort.Ints(caloriesSum)

	topThree := caloriesSum[elfCount-3 : elfCount]
	topThreeSum := 0

	for _, calories := range topThree {
		topThreeSum += calories
	}

	return util.NewSolution(2, topThreeSum)
}
