package day12

import (
	"sort"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) *util.Solution {
	grid := parseGrid(input)
	start := findNodes(grid, "S")[0]
	end := findNodes(grid, "E")[0]
	startingNodes := findNodes(grid, "a")

	startingNodes = append(startingNodes, start)

	steps := []int{}

	for _, node := range startingNodes {
		stepCount := bfs(grid, node, end)

		if stepCount == notFound {
			continue
		}

		steps = append(steps, stepCount)
	}

	sort.Ints(steps)

	return util.NewSolution(2, steps[0])
}
