package day08

import (
	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) int {
	rows, cols := parseTreeGrid(input)
	scenicScores := make([]int, 0)

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows); j++ {
			tree := rows[i][j]

			viewingDistanceLeft := calcViewingDistance(tree, util.Reverse(rows[i][:j]))
			viewingDistanceAbove := calcViewingDistance(tree, util.Reverse(cols[j][:i]))
			viewingDistanceRight := calcViewingDistance(tree, rows[i][j+1:])
			viewingDistanceBelow := calcViewingDistance(tree, cols[j][i+1:])

			scenicScore := viewingDistanceLeft * viewingDistanceAbove * viewingDistanceRight * viewingDistanceBelow

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
		distance = i + 1

		if height <= tree {
			break
		}
	}

	return distance
}
