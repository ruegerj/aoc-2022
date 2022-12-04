package day04

import "strings"

func Part2(input string) int {
	pairs := strings.Split(input, "\n")

	overlapSectionsCount := 0

	for _, pair := range pairs {
		sectionRanges := strings.Split(pair, ",")

		lowerA, upperA := computeLimits(sectionRanges[0])
		lowerB, upperB := computeLimits(sectionRanges[1])

		if upperA < lowerB || upperB < lowerA {
			continue
		}

		overlapSectionsCount++
	}

	return overlapSectionsCount
}
