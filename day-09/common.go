package day09

import "github.com/ruegerj/aoc-2022/util"

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}

func (point *Point) MoveX(count int) {
	point.X += count
}

func (point *Point) MoveY(count int) {
	point.Y += count
}

func (point *Point) MoveByDirection(direction string) {
	if direction == "U" {
		point.MoveY(1)
	}

	if direction == "R" {
		point.MoveX(1)
	}

	if direction == "D" {
		point.MoveY(-1)
	}

	if direction == "L" {
		point.MoveX(-1)
	}
}

func (pointA *Point) Overlaps(pointB *Point) bool {
	return pointA.X == pointB.X && pointA.Y == pointB.Y
}

func (point *Point) Coordinates() []int {
	return []int{point.X, point.Y}
}

func (pointA *Point) Follow(pointB *Point) {
	if pointA.Overlaps(pointB) {
		return
	}

	rowDiff := util.Abs(pointB.X - pointA.X)
	colDiff := util.Abs(pointB.Y - pointA.Y)

	if colDiff == 0 && rowDiff != 0 {
		steps := rowDiff - 1

		if steps == 0 {
			return
		}

		if pointB.X < pointA.X {
			steps *= -1
		}

		pointA.MoveX(steps)
		return
	}

	if colDiff != 0 && rowDiff == 0 {
		steps := colDiff - 1

		if steps == 0 {
			return
		}

		if pointB.Y < pointA.Y {
			steps *= -1
		}

		pointA.MoveY(steps)
		return
	}

	if colDiff == 1 && rowDiff == 1 {
		return
	}

	stepsX := util.Abs(pointB.X - pointA.X)
	stepsY := util.Abs(pointB.Y - pointA.Y)

	if stepsX > stepsY {
		stepsX--
	} else if stepsY > stepsX {
		stepsY--
	} else {
		stepsX--
		stepsY--
	}

	if pointB.X < pointA.X {
		stepsX *= -1
	}

	if pointB.Y < pointA.Y {
		stepsY *= -1
	}

	pointA.MoveX(stepsX)
	pointA.MoveY(stepsY)
}
