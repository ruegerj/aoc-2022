package day08

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

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
