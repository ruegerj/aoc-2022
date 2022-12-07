package day07

import (
	"strings"
)

func Part1(input string) int {
	const dirSizeThreshold = 100000

	lines := strings.Split(input, "\n")

	dirMap := parseInput(lines)

	var totalSize int = 0

	for _, dir := range dirMap {
		size := dir.EffectiveSize()

		if size > dirSizeThreshold {
			continue
		}

		totalSize += size
	}

	return totalSize
}
