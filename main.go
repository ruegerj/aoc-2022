package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	day01 "github.com/ruegerj/aoc-2022/day-01"
	day02 "github.com/ruegerj/aoc-2022/day-02"
	"github.com/ruegerj/aoc-2022/util"
	"golang.org/x/exp/slices"
)

type dailyChallenge func(string) int

func main() {
	fmt.Println("     ___       ______     ______     ___     ___    ___    ___   ")
	fmt.Println("    /   \\     /  __  \\   /      |   |__ \\   / _ \\  |__ \\  |__ \\  ")
	fmt.Println("   /  ^  \\   |  |  |  | |  ,----'      ) | | | | |    ) |    ) | ")
	fmt.Println("  /  /_\\  \\  |  |  |  | |  |          / /  | | | |   / /    / /  ")
	fmt.Println(" /  _____  \\ |  `--'  | |  `----.    / /_  | |_| |  / /_   / /_  ")
	fmt.Println("/__/     \\__\\ \\______/   \\______|   |____|  \\___/  |____| |____| ")
	fmt.Println("-------------------------------")
	fmt.Println("🎄 Happy Coding & festive season")

	daysLeft := util.DaysUntilXMas()
	timeLeftMsg := "🎅 XMas has already passed"

	if daysLeft > 0 {
		timeLeftMsg = fmt.Sprintf("⏱ ~%d days left until XMas", daysLeft)
	}

	fmt.Println(timeLeftMsg + "\n")

	args := os.Args[1:]
	printHelp := slices.Contains(args, "help") || slices.Contains(args, "h")

	if printHelp {
		fmt.Println("usage: go run . <day-nr>")
		return
	}

	dayNr, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("❌ Invalid day number...")
		return
	}

	dayRegistry := map[int]func(){
		1: func() { runDay(1, day01.Part1, day01.Part2) },
		2: func() { runDay(2, day02.Part1, day02.Part2) },
	}

	requestedDay := dayRegistry[dayNr]

	if requestedDay == nil {
		fmt.Println("🛠 Not implemented")
		return
	}

	requestedDay()
}

func runDay(nr int, part1 dailyChallenge, part2 dailyChallenge) {
	input := util.LoadDailyInput(nr)
	normalizedNr := util.PadNumber(nr)

	fmt.Printf("⭐️ Day %s\n", normalizedNr)

	start1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(input), time.Since(start1))

	start2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(input), time.Since(start2))
}
