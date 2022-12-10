package day09

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	commands := strings.Split(input, "\n")

	head := NewPoint(0, 0)
	tail := NewPoint(0, 0)

	visitedCoordinates := [][]int{tail.Coordinates()}

	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		steps := util.MustParseInt(parts[1])

		for i := 0; i < steps; i++ {
			head.MoveByDirection(direction)
			tail.Follow(head)

			tailCoordinates := tail.Coordinates()

			alreadyVisited := util.Includes(visitedCoordinates, func(coords []int, _ int) bool {
				return coords[0] == tailCoordinates[0] && coords[1] == tailCoordinates[1]
			})

			if alreadyVisited {
				continue
			}

			visitedCoordinates = append(visitedCoordinates, tailCoordinates)
		}
	}

	return util.NewSolution(1, len(visitedCoordinates))
}
