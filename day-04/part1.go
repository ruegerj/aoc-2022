package day04

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	pairs := strings.Split(input, "\n")

	containedSectionsCount := 0

	for _, pair := range pairs {
		sectionRanges := strings.Split(pair, ",")

		lowerA, upperA := computeLimits(sectionRanges[0])
		lowerB, upperB := computeLimits(sectionRanges[1])

		if (lowerA > lowerB && upperA > upperB) || (lowerB > lowerA && upperB > upperA) {
			continue
		}

		containedSectionsCount++
	}

	return util.NewSolution(1, containedSectionsCount)
}
