package day08

import (
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) int {
	rows, cols := parseTreeGrid(input)
	visibleTrees := len(rows)*4 - 4

	for i := 0; i < len(rows); i++ {
		if i == 0 || i == len(rows)-1 {
			continue
		}

		for j := 0; j < len(rows); j++ {
			if j == 0 || j == len(rows)-1 {
				continue
			}

			tree := rows[i][j]

			highestLeft := util.Max(rows[i][:j])
			highestAbove := util.Max(cols[j][:i])
			highestRight := util.Max(rows[i][j+1:])
			highestBelow := util.Max(cols[j][i+1:])

			if tree <= highestLeft && tree <= highestRight && tree <= highestAbove && tree <= highestBelow {
				continue
			}

			visibleTrees++
		}
	}

	return visibleTrees
}
