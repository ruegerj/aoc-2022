package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while fetching current dir", err)
		return
	}

	rawInput, err := os.ReadFile(path.Join(cwd, "day-01", "01.txt"))
	input := string(rawInput)

	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
