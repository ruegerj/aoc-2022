package day03

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
	"golang.org/x/exp/slices"
)

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
