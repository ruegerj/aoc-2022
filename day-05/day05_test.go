package day05

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var dailyInput string

func TestMain(m *testing.M) {
	input, err := os.ReadFile(path.Join("..", "data", "05.txt"))

	if err != nil {
		fmt.Println("Failed to load input file", err)
		os.Exit(1)
	}

	dailyInput = string(input)

	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	expected := "WSFTMRHPP"
	result := Part1(dailyInput)

	if result != expected {
		t.Errorf("Expected %s, produced %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := "GSLCMFBRP"
	result := Part2(dailyInput)

	if result != expected {
		t.Errorf("Expected %s, produced %s", expected, result)
	}
}
