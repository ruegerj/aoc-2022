package day03

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
	"golang.org/x/exp/slices"
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

func Part2(input string) *util.Solution {
	rucksacks := strings.Split(input, "\n")
	priorityScore := 0

	groups := chunkSlice(rucksacks, 3)

	for _, group := range groups {
		rucksackA := group[0]
		rucksackB := group[1]
		rucksackC := group[2]

		matchesAB := getMatches(rucksackA, rucksackB)
		badge := ""

		for _, itemC := range strings.Split(rucksackC, "") {
			if !slices.Contains(matchesAB, itemC) {
				continue
			}

			badge = itemC
			break
		}

		itemPriority := calcItemPriority(badge)
		priorityScore += itemPriority
	}

	return util.NewSolution(2, priorityScore)
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func getMatches(a string, b string) []string {
	matches := make([]string, 0)

	for _, item := range strings.Split(a, "") {
		if !strings.Contains(b, item) {
			continue
		}

		matches = append(matches, item)
	}

	return matches
}

func calcItemPriority(item string) int {
	itemPriority := 0

	lowerCaseAlphabet := "abcdefghijklmnopqrstuvwxyz"

	if strings.Index(lowerCaseAlphabet, item) == -1 {
		item = strings.ToLower(item)
		itemPriority += 26
	}

	return itemPriority + strings.Index(lowerCaseAlphabet, item) + 1
}
