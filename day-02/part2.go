package day02

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part2(input string) *util.Solution {
	operandLookup := map[string]string{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	games := strings.Split(input, "\n")
	score := 0

	for _, game := range games {
		operands := strings.Split(game, " ")
		opponent := operandLookup[operands[0]]
		outcome := operands[1]

		if outcome == "Y" {
			score += operandScores[opponent] + 3
			continue
		}

		if outcome == "Z" {
			score += operandScores[winningCombination[opponent]] + 6
			continue
		}

		score += operandScores[loosingCombination[opponent]]
	}

	return util.NewSolution(2, score)
}
