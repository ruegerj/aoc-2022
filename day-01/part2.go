package day01

import (
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) int {
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

	return topThreeSum
}
