package day01

import (
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
