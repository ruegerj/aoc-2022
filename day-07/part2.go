package day07

import (
	"sort"
	"strings"
)

func Part2(input string) int {
	const requiredSpace = 30000000
	const diskSpace = 70000000

	lines := strings.Split(input, "\n")

	dirMap := parseInput(lines)
	sizeList := make(directorySizeList, 0)

	totalSize := dirMap["/"].EffectiveSize()
	spaceToReclaim := requiredSpace - (diskSpace - totalSize)

	for path, dir := range dirMap {
		size := dir.EffectiveSize()

		if size < spaceToReclaim {
			continue
		}

		sizeList = append(sizeList, directorySize{path, size})
	}

	sort.Sort(sizeList)

	return sizeList[0].size
}

// Implement sort interface for Directory array
type directorySize struct {
	path string
	size int
}

type directorySizeList []directorySize

func (list directorySizeList) Len() int {
	return len(list)
}

func (list directorySizeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list directorySizeList) Less(i, j int) bool {
	return list[i].size < list[j].size
}
