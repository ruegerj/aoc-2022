package day10

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

var dailyInput string

func TestMain(m *testing.M) {
	input, err := os.ReadFile(path.Join("..", "data", "10.txt"))

	if err != nil {
		fmt.Println("Failed to load input file", err)
		os.Exit(1)
	}

	dailyInput = string(input)

	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	expected := 13440
	solution := Part1(dailyInput)

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}

func TestPart2(t *testing.T) {
	expected := []string{
		"# # # . . # # # . . # # # # . . # # . . # # # . . . # # . . # # # # . . # # . .",
		"# . . # . # . . # . . . . # . # . . # . # . . # . # . . # . . . . # . # . . # .",
		"# . . # . # # # . . . . # . . # . . . . # . . # . # . . # . . . # . . # . . # .",
		"# # # . . # . . # . . # . . . # . # # . # # # . . # # # # . . # . . . # # # # .",
		"# . . . . # . . # . # . . . . # . . # . # . # . . # . . # . # . . . . # . . # .",
		"# . . . . # # # . . # # # # . . # # # . # . . # . # . . # . # # # # . # . . # .",
	}

	solution := Part2(dailyInput)

	for i, row := range solution.Result.([][]string) {
		printedRow := strings.Join(row, " ")

		if printedRow != expected[i] {
			t.Errorf("Expected %s\n produced %s", expected, printedRow)
		}
	}
}
