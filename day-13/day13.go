package day13

import (
	"sort"
	"strconv"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	pairs := strings.Split(input, "\n\n")
	correctOrderSum := 0

	for pairIndex, pair := range pairs {
		packets := strings.Split(pair, "\n")

		left, _ := parsePacket(packets[0], 0)
		right, _ := parsePacket(packets[1], 0)

		order := left.compareTo(right)

		if order == ORDERED {
			correctOrderSum += pairIndex + 1
		}
	}

	return util.NewSolution(1, correctOrderSum)
}

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

const (
	ORDERED   int = -1
	EQUAL     int = 0
	UNORDERED int = 1
)

type Packet struct {
	Number     int
	SubPackets Packets
}

func NewNumberPacket(number int) *Packet {
	return &Packet{Number: number}
}

func NewCollectionPackage(collection []*Packet) *Packet {
	if collection == nil {
		collection = make([]*Packet, 0)
	}

	return &Packet{SubPackets: collection}
}

func NewDividerPackage(number int) *Packet {
	return NewCollectionPackage([]*Packet{NewCollectionPackage([]*Packet{NewNumberPacket(number)})})
}

func (packet *Packet) IsNumber() bool {
	if packet.SubPackets != nil {
		return false
	}

	return true
}

func (packet *Packet) IsCollection() bool {
	if packet.SubPackets == nil {
		return false
	}

	return true
}

func (packet *Packet) IsDivider() bool {
	if !packet.HasPacketAt(0) || !packet.SubPackets[0].HasPacketAt(0) {
		return false
	}

	keyPart := packet.SubPackets[0].SubPackets[0].Number

	return keyPart == 2 || keyPart == 6
}

func (packet *Packet) HasPacketAt(index int) bool {
	if !packet.IsCollection() {
		return false
	}

	return len(packet.SubPackets)-1 >= index
}

func (packetA *Packet) compareTo(packetB *Packet) int {
	if packetA.IsNumber() && packetB.IsNumber() {
		diff := packetA.Number - packetB.Number

		if diff < 0 {
			return ORDERED
		}

		if diff > 0 {
			return UNORDERED
		}

		return EQUAL
	}

	if packetA.IsCollection() && packetB.IsNumber() {
		wrappedB := NewCollectionPackage([]*Packet{packetB})

		return packetA.compareTo(wrappedB)
	}

	if packetA.IsNumber() && packetB.IsCollection() {
		wrappedA := NewCollectionPackage([]*Packet{packetA})

		return wrappedA.compareTo(packetB)
	}

	for i := 0; i < len(packetA.SubPackets); i++ {
		if !packetA.HasPacketAt(i) {
			return ORDERED
		}

		if !packetB.HasPacketAt(i) {
			return UNORDERED
		}

		a := packetA.SubPackets[i]
		b := packetB.SubPackets[i]

		order := a.compareTo(b)

		if order == EQUAL {
			continue
		}

		return order
	}

	if len(packetA.SubPackets) > len(packetB.SubPackets) {
		return UNORDERED
	}

	if len(packetA.SubPackets) < len(packetB.SubPackets) {
		return ORDERED
	}

	return EQUAL
}

// Implement sort interface
type Packets []*Packet

func (packets Packets) Len() int {
	return len(packets)
}

func (packets Packets) Swap(i, j int) {
	packets[i], packets[j] = packets[j], packets[i]
}

func (packets Packets) Less(i, j int) bool {
	return packets[i].compareTo(packets[j]) == ORDERED
}

func parsePacket(packetStream string, position int) (*Packet, int) {
	root := NewCollectionPackage(nil)

	var current int
	start := position
	bits := strings.Split(packetStream, "")

	for current = position; current < len(bits); current++ {
		bit := bits[current]

		if bit == "," {
			continue
		}

		if bit == "[" {
			sub, processed := parsePacket(packetStream, current+1)
			current += processed
			root.SubPackets = append(root.SubPackets, sub)
			continue
		}

		if bit == "]" {
			break
		}

		num := util.MustParseInt(bit)
		nextNum, err := strconv.Atoi(bits[current+1])

		if err == nil {
			num = num*10 + nextNum
			current++
		}

		root.SubPackets = append(root.SubPackets, NewNumberPacket(num))
	}

	current += 1

	return root, current - start
}
