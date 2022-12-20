package day15

import (
	"regexp"
	"sort"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	const targetY = 2000000

	sensors, nearestBeacons := parseInput(input)
	potentialBeacons := set.New()

	for _, sensor := range sensors {
		beacon := nearestBeacons[sensor]

		scanRange := sensor.ManhattanDistance(beacon)
		positionsCoveredEachSide := scanRange - util.Abs(sensor.Y-targetY)

		if positionsCoveredEachSide < 0 {
			continue
		}

		potentialBeacons.Insert(sensor.X)

		coveredRangeStart := sensor.X - positionsCoveredEachSide
		coveredRangeEnd := sensor.X + positionsCoveredEachSide

		for i := coveredRangeStart; i <= coveredRangeEnd; i++ {
			potentialBeacons.Insert(i)
		}
	}

	for _, beacon := range nearestBeacons {
		if beacon.Y != targetY {
			continue
		}

		potentialBeacons.Remove(beacon.X)
	}

	return util.NewSolution(1, potentialBeacons.Len())
}

func Part2(input string) *util.Solution {
	const lowerBound = 0
	// const upperBound = 20
	const upperBound = 4000000
	const tuningFrequencyMultiplier = 4000000

	sensors, nearestBeacons := parseInput(input)

	for currentY := lowerBound; currentY < upperBound; currentY++ {
		sort.Slice(sensors, func(i, j int) bool {
			sensorA := sensors[i]
			sensorB := sensors[j]

			scanRangeA := sensorA.ManhattanDistance(nearestBeacons[sensorA])
			scanRangeB := sensorB.ManhattanDistance(nearestBeacons[sensorB])
			uncoveredRangeA := scanRangeA - util.Abs(sensorA.Y-currentY)
			uncoveredRangeB := scanRangeB - util.Abs(sensorB.Y-currentY)

			return sensorA.X-uncoveredRangeA < sensorB.X-uncoveredRangeB
		})

		var coveredVerticalRange int
		for _, sensor := range sensors {
			scanRange := sensor.ManhattanDistance(nearestBeacons[sensor])
			positionsCoveredEachSide := scanRange - util.Abs(sensor.Y-currentY)

			if positionsCoveredEachSide < 0 {
				continue
			}

			firstCoveredPosAtTarget := sensor.X - positionsCoveredEachSide

			if coveredVerticalRange < firstCoveredPosAtTarget-1 {
				tuningFrequency := (firstCoveredPosAtTarget-1)*tuningFrequencyMultiplier + currentY

				return util.NewSolution(2, tuningFrequency)
			}

			if coveredVerticalRange < sensor.X+positionsCoveredEachSide {
				coveredVerticalRange = sensor.X + positionsCoveredEachSide
			}
		}
	}

	return util.NewSolution(2, -1)
}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}

// Calculates the Manhattan Distance between the pointA and pointB
func (pointA *Point) ManhattanDistance(pointB *Point) int {
	absDiffX := util.Abs(pointA.X - pointB.X)
	absDiffY := util.Abs(pointA.Y - pointB.Y)

	return absDiffX + absDiffY
}

func parseInput(input string) ([]*Point, map[*Point]*Point) {
	sensorBeaconPairs := strings.Split(input, "\n")
	coordinateMatcher := regexp.MustCompile(`^.+x=(?P<sensorX>[-\d]+).+y=(?P<sensorY>[-\d]+).+x=(?P<beaconX>[-\d]+).+y=(?P<beaconY>[-\d]+)$`)

	sensors := make([]*Point, 0)
	nearestBeacons := make(map[*Point]*Point, 0)

	for _, pair := range sensorBeaconPairs {
		coordinateMap := util.MatchNamedSubgroups(coordinateMatcher, pair)

		sensorX := util.MustParseInt(coordinateMap["sensorX"])
		sensorY := util.MustParseInt(coordinateMap["sensorY"])
		beaconX := util.MustParseInt(coordinateMap["beaconX"])
		beaconY := util.MustParseInt(coordinateMap["beaconY"])

		sensor := NewPoint(sensorX, sensorY)

		nearestBeacons[sensor] = NewPoint(beaconX, beaconY)
		sensors = append(sensors, sensor)
	}

	return sensors, nearestBeacons
}
