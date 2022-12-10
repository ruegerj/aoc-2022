package day02

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	games := strings.Split(input, "\n")
	score := 0

	operandLookup := map[string]string{
		"A": Rock,
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissors,
		"Z": Scissors,
	}

	for _, game := range games {
		operands := strings.Split(game, " ")
		opponent := operandLookup[operands[0]]
		own := operandLookup[operands[1]]

		score += operandScores[own]

		if opponent == own {
			score += 3
			continue
		}

		if winningCombination[opponent] != own {
			continue
		}

		score += 6
	}

	return util.NewSolution(1, score)
}
