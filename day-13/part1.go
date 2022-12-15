package day13

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	pairs := strings.Split(input, "\n\n")
	correctOrderSum := 0

	for pairIndex, pair := range pairs {
		packets := strings.Split(pair, "\n")

		left, _ := parsePacket(packets[0], 0)
		right, _ := parsePacket(packets[1], 0)

		order := left.compareTo(right)

		if order == ORDERED {
			correctOrderSum += pairIndex + 1
		}
	}

	return util.NewSolution(1, correctOrderSum)
}
