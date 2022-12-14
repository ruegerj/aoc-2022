package day10

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	const noopCmd = "noop"
	const addXCmd = "addx"

	instructions := strings.Split(input, "\n")

	instructionPointer := 0
	register := 1
	cycle := 1
	cmdCycle := 0
	signalStrengths := make([]int, 0)

	for true {
		if instructionPointer >= len(instructions) {
			break
		}

		instruction := instructions[instructionPointer]
		parts := strings.Split(instruction, " ")
		cmd := parts[0]

		if cmd == noopCmd {
			signalStrengths = append(signalStrengths, cycle*register)
			cycle++
			instructionPointer++
			continue
		}

		if cmdCycle < 2 {
			signalStrengths = append(signalStrengths, cycle*register)
			cycle++
			cmdCycle++
			continue
		}

		cmdCycle = 0
		register += util.MustParseInt(parts[1])
		instructionPointer++
	}

	signalStrengthSum := 0

	for i := 19; i < len(signalStrengths); i += 40 {
		signalStrengthSum += signalStrengths[i]
	}

	return util.NewSolution(1, signalStrengthSum)
}

func Part2(input string) *util.Solution {
	const noopCmd = "noop"
	const addXCmd = "addx"
	const screenWidth = 40

	instructions := strings.Split(input, "\n")

	instructionPointer := 0
	register := 1
	cycle := 1
	cmdCycle := 0
	screen := [][]string{make([]string, screenWidth)}
	renderPosition := 0
	renderRow := 0
	signalStrengths := make([]int, 0)

	for true {
		if instructionPointer >= len(instructions) {
			break
		}

		if renderPosition >= screenWidth {
			screen = append(screen, make([]string, screenWidth))
			renderPosition = 0
			renderRow++
		}

		instruction := instructions[instructionPointer]
		parts := strings.Split(instruction, " ")
		cmd := parts[0]

		if cmd == noopCmd {
			signalStrengths = append(signalStrengths, cycle*register)
			printPixel(screen, renderRow, renderPosition, register)
			renderPosition++
			cycle++
			instructionPointer++
			continue
		}

		if cmdCycle < 2 {
			signalStrengths = append(signalStrengths, cycle*register)
			printPixel(screen, renderRow, renderPosition, register)
			renderPosition++
			cycle++
			cmdCycle++
			continue
		}

		cmdCycle = 0
		register += util.MustParseInt(parts[1])
		instructionPointer++
	}

	return util.NewSolution(2, screen)
}

func printPixel(screen [][]string, row int, pos int, register int) {
	char := "."

	if register == pos || register == pos-1 || register == pos+1 {
		char = "#"
	}

	screen[row][pos] = char
}
