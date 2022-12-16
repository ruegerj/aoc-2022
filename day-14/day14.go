package day14

import (
	"container/list"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	caveSystem := parseCaveSystem(input)
	height := len(caveSystem)

	var movingSand *Point
	sandUnitCount := 0

	for true {
		if movingSand == nil {
			movingSand = NewPoint(sandSourceX, sandSourceY)
		}

		if movingSand.Y >= height-1 {
			break
		}

		if caveSystem[movingSand.Y+1][movingSand.X] == air {
			movingSand.Y++
			continue
		}

		if caveSystem[movingSand.Y+1][movingSand.X-1] == air {
			movingSand.Y++
			movingSand.X--
			continue
		}

		if caveSystem[movingSand.Y+1][movingSand.X+1] == air {
			movingSand.Y++
			movingSand.X++
			continue
		}

		caveSystem[movingSand.Y][movingSand.X] = sand
		movingSand = nil
		sandUnitCount++
	}

	return util.NewSolution(1, sandUnitCount)
}

func Part2(input string) *util.Solution {
	const floorOffset = 2

	caveSystem := parseCaveSystem(input)
	caveSystem = addFloor(caveSystem, floorOffset)

	var movingSand *Point
	sandUnitCount := 0

	for true {
		if movingSand == nil {
			movingSand = NewPoint(sandSourceX, sandSourceY)
		}

		if caveSystem[sandSourceY][sandSourceX] == sand {
			break
		}

		if caveSystem[movingSand.Y+1][movingSand.X] == air {
			movingSand.Y++
			continue
		}

		if caveSystem[movingSand.Y+1][movingSand.X-1] == air {
			movingSand.Y++
			movingSand.X--
			continue
		}

		if caveSystem[movingSand.Y+1][movingSand.X+1] == air {
			movingSand.Y++
			movingSand.X++
			continue
		}

		caveSystem[movingSand.Y][movingSand.X] = sand
		movingSand = nil
		sandUnitCount++
	}

	return util.NewSolution(2, sandUnitCount)
}

const air rune = 46   // .
const rock rune = 35  // #
const sand rune = 111 // o
const caveWidth = 1000
const sandSourceX = 500
const sandSourceY = 0

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}

func parseCaveSystem(input string) [][]rune {
	rockOutlines := make([]*list.List, 0)
	rockFormations := strings.Split(input, "\n")
	maxY := 0

	for i, rockFormation := range rockFormations {
		rockOutlines = append(rockOutlines, list.New())
		coordinates := strings.Split(rockFormation, " -> ")

		for _, coordinate := range coordinates {
			parts := strings.Split(coordinate, ",")
			x := util.MustParseInt(parts[0])
			y := util.MustParseInt(parts[1])

			if y > maxY {
				maxY = y
			}

			rockOutlines[i].PushBack(NewPoint(x, y))
		}
	}

	height := maxY + 1
	caveSystem := make([][]rune, height)

	for y := 0; y < height; y++ {
		caveSystem[y] = make([]rune, caveWidth)

		for x := 0; x < caveWidth; x++ {
			caveSystem[y][x] = air
		}
	}

	for _, rockOutline := range rockOutlines {
		for current := rockOutline.Front(); current != nil; current = current.Next() {
			currentRock := current.Value.(*Point)
			prev := current.Prev()

			caveSystem[currentRock.Y][currentRock.X] = rock

			if prev == nil {
				continue
			}

			prevRock := prev.Value.(*Point)

			diffX := (prevRock.X - currentRock.X) * -1
			diffY := (prevRock.Y - currentRock.Y) * -1

			for diffX != 0 || diffY != 0 {
				x := prevRock.X + diffX
				y := prevRock.Y + diffY

				caveSystem[y][x] = rock

				if diffX > 0 {
					diffX--
				}

				if diffX < 0 {
					diffX++
				}

				if diffY > 0 {
					diffY--
				}

				if diffY < 0 {
					diffY++
				}
			}
		}
	}

	return caveSystem
}

func addFloor(caveSystem [][]rune, offset int) [][]rune {
	height := len(caveSystem)
	width := len(caveSystem[0])

	caveSystemWithFloor := make([][]rune, height)
	copy(caveSystemWithFloor, caveSystem)

	for i := 0; i < offset; i++ {
		caveSystemWithFloor = append(caveSystemWithFloor, make([]rune, width))

		for j := 0; j < width; j++ {
			if i+1 < offset {
				caveSystemWithFloor[height+i][j] = air
				continue
			}

			caveSystemWithFloor[height+i][j] = rock
		}
	}

	return caveSystemWithFloor
}
