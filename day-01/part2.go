package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

		caloryCount, err := strconv.ParseInt(calories, 10, 64)

		if err != nil {
			fmt.Println("Failed to parse calories", err)
			continue
		}

		caloriesOfElf += int(caloryCount)
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
