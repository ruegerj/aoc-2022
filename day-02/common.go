package day02

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
