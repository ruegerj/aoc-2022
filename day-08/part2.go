package day08

import (
	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) int {
	treeGrid := parseTreeGrid(input)
	scenicScores := make([]int, 0)

	for i, row := range treeGrid {
		for j, tree := range row {
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

			viewScoreLeft := calcViewingDistance(tree, util.Reverse(treesLeft))
			viewScoreRight := calcViewingDistance(tree, treesRight)
			viewScoreAbove := calcViewingDistance(tree, util.Reverse(treesAbove))
			viewScoreBelow := calcViewingDistance(tree, treesBelow)

			scenicScore := viewScoreLeft * viewScoreAbove * viewScoreRight * viewScoreBelow

			if scenicScore <= 0 {
				continue
			}

			scenicScores = append(scenicScores, scenicScore)
		}
	}

	return util.Max(scenicScores)
}

func calcViewingDistance(height int, trees []int) int {
	if len(trees) == 0 {
		return 0
	}

	distance := 0

	for i, tree := range trees {
		distance = i

		if height <= tree {
			break
		}
	}

	return distance + 1
}
