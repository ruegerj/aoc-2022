package day06

import (
	"container/list"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) *util.Solution {
	chars := strings.Split(input, "")

	const uniquePackageCount = 14
	var startOfPacketIndex int
	charHistory := list.New()

	for i, char := range chars {
		startOfPacketIndex = i

		charHistory.PushBack(char)

		if charHistory.Len() < uniquePackageCount {
			continue
		}

		lastRelevantPackages := getTopNValues(charHistory, uniquePackageCount)

		if set.New(lastRelevantPackages...).Len() == uniquePackageCount {
			break
		}

		oldest := charHistory.Front()
		charHistory.Remove(oldest)
	}

	return util.NewSolution(2, startOfPacketIndex+1)
}
