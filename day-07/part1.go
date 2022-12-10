package day07

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
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

	return util.NewSolution(1, totalSize)
}
