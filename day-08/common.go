package day08

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func parseTreeGrid(input string) [][]int {
	lines := strings.Split(input, "\n")
	treeGrid := make([][]int, len(lines))

	for i, row := range lines {
		treeGrid[i] = make([]int, len(row))

		for j, tree := range strings.Split(row, "") {
			treeGrid[i][j] = util.MustParseInt(tree)
		}
	}

	return treeGrid
}
