package day08

import (
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) int {
	treeGrid := parseTreeGrid(input)
	visibleTrees := len(treeGrid)*4 - 4

	for i, row := range treeGrid {
		if i == 0 || i == len(row)-1 {
			continue
		}

		for j, tree := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}

			treesLeft := treeGrid[i][:j]
			treesRight := treeGrid[i][j+1:]
			treesAbove := make([]int, 0)
			treesBelow := make([]int, 0)

			for k, r := range treeGrid {
				if k == i {
					continue
				}

				if k < i {
					treesAbove = append(treesAbove, r[j])
					continue
				}

				treesBelow = append(treesBelow, r[j])
			}

			highestLeft := util.Max(treesLeft)
			highestRight := util.Max(treesRight)
			highestAbove := util.Max(treesAbove)
			highestBelow := util.Max(treesBelow)

			if tree <= highestLeft && tree <= highestRight && tree <= highestAbove && tree <= highestBelow {
				continue
			}

			visibleTrees++
		}
	}

	return visibleTrees
}
