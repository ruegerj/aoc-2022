package day03

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
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

	return util.NewSolution(1, priorityScore)
}
