package day06

import (
	"container/list"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	chars := strings.Split(input, "")

	const uniquePackageCount = 4
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

	return util.NewSolution(1, startOfPacketIndex+1)
}

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

func getTopNValues(list *list.List, n int) []any {
	item := list.Front()

	values := make([]any, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			item = item.Next()
		}

		values[i] = item.Value
	}

	return values
}
