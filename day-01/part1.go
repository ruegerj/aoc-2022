package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
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

		caloryCount, err := strconv.ParseInt(calories, 10, 64)

		if err != nil {
			fmt.Println("Failed to parse calories", err)
			continue
		}

		caloriesSum += int(caloryCount)
	}

	return maxCalories
}
