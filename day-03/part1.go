package day03

import (
	"strings"
)

func Part1(input string) int {
	rucksacks := strings.Split(input, "\n")
	priorityScore := 0

	for _, rucksack := range rucksacks {
		itemsPerCompartment := len(rucksack) / 2

		compA := rucksack[:itemsPerCompartment]
		compB := rucksack[itemsPerCompartment:]

		matches := getMatches(compA, compB)

		commonItem := matches[0]

		itemPriority := calcItemPriority(commonItem)

		priorityScore += itemPriority
	}

	return priorityScore
}
