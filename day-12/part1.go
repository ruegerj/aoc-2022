package day12

import (
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	grid := parseGrid(input)
	start := findNodes(grid, "S")[0]
	end := findNodes(grid, "E")[0]

	steps := bfs(grid, start, end)

	return util.NewSolution(1, steps)
}
