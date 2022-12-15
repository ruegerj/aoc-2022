package day13

import (
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) *util.Solution {
	var packets Packets = []*Packet{
		NewDividerPackage(2),
		NewDividerPackage(6),
	}

	for _, pair := range strings.Split(input, "\n\n") {
		for _, rawPacket := range strings.Split(pair, "\n") {
			packet, _ := parsePacket(rawPacket, 0)

			packets = append(packets, packet)
		}
	}

	sort.Sort(packets)

	dividerIndexes := []int{}

	for i, packet := range packets {
		if !packet.IsDivider() {
			continue
		}

		dividerIndexes = append(dividerIndexes, i+1)
	}

	decoderKey := dividerIndexes[0] * dividerIndexes[1]

	return util.NewSolution(2, decoderKey)
}
