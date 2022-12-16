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

const Rock = "rock"
const Paper = "paper"
const Scissors = "scissors"

var operandScores = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var loosingCombination = map[string]string{
	Paper:    Rock,
	Scissors: Paper,
	Rock:     Scissors,
}

var winningCombination = map[string]string{
	Scissors: Rock,
	Rock:     Paper,
	Paper:    Scissors,
}
