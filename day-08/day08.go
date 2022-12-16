package day08

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
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

	return util.NewSolution(1, visibleTrees)
}

func Part2(input string) *util.Solution {
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

	return util.NewSolution(2, util.Max(scenicScores))
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

func parseTreeGrid(input string) (map[int][]int, map[int][]int) {
	lines := strings.Split(input, "\n")
	rows := make(map[int][]int, len(lines))
	cols := make(map[int][]int, len(lines))

	for i, row := range lines {
		for j, num := range strings.Split(row, "") {
			tree := util.MustParseInt(num)

			rows[i] = append(rows[i], tree)
			cols[j] = append(cols[j], tree)
		}
	}

	return rows, cols
}
