package day09

import (
	"container/list"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) int {
	const ropeLength = 10
	commands := strings.Split(input, "\n")

	rope, head, tail := createRope(ropeLength)

	visitedCoordinates := [][]int{tail.Coordinates()}

	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		steps := util.MustParseInt(parts[1])

		for i := 0; i < steps; i++ {
			head.MoveByDirection(direction)

			for knot := rope.Front(); knot != nil; knot = knot.Next() {
				prevKnot := knot.Prev()

				if prevKnot == nil {
					continue
				}

				knot.Value.(*Point).Follow(prevKnot.Value.(*Point))
			}

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

	return len(visitedCoordinates)
}

func createRope(length int) (*list.List, *Point, *Point) {
	head := NewPoint(0, 0)
	tail := NewPoint(0, 0)

	rope := list.New()
	rope.PushFront(head)

	for i := 0; i < length-2; i++ {
		rope.PushBack(NewPoint(0, 0))
	}

	rope.PushBack(tail)

	return rope, head, tail
}
