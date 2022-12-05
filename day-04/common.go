package day04

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func computeLimits(sectionRange string) (int, int) {
	limits := strings.Split(sectionRange, "-")

	lower := util.MustParseInt(limits[0])
	upper := util.MustParseInt(limits[1])

	return lower, upper
}
