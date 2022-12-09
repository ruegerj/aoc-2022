package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	day01 "github.com/ruegerj/aoc-2022/day-01"
	day02 "github.com/ruegerj/aoc-2022/day-02"
	day03 "github.com/ruegerj/aoc-2022/day-03"
	day04 "github.com/ruegerj/aoc-2022/day-04"
	day05 "github.com/ruegerj/aoc-2022/day-05"
	day06 "github.com/ruegerj/aoc-2022/day-06"
	day07 "github.com/ruegerj/aoc-2022/day-07"
	day08 "github.com/ruegerj/aoc-2022/day-08"
	day09 "github.com/ruegerj/aoc-2022/day-09"
	"github.com/ruegerj/aoc-2022/util"
	"golang.org/x/exp/slices"
)

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
		3: func() { runDay(3, day03.Part1, day03.Part2) },
		4: func() { runDay(4, day04.Part1, day04.Part2) },
		5: func() { runDay(5, day05.Part1, day05.Part2) },
		6: func() { runDay(6, day06.Part1, day06.Part2) },
		7: func() { runDay(7, day07.Part1, day07.Part2) },
		8: func() { runDay(8, day08.Part1, day08.Part2) },
		9: func() { runDay(9, day09.Part1, day09.Part2) },
	}

	requestedDay := dayRegistry[dayNr]

	if requestedDay == nil {
		fmt.Println("🛠 Not implemented")
		return
	}

	requestedDay()
}

func runDay[TResult int | string](nr int, part1 func(string) TResult, part2 func(string) TResult) {
	input := util.LoadDailyInput(nr)
	normalizedNr := util.PadNumber(nr)

	fmt.Printf("⭐️ Day %s\n", normalizedNr)

	start1 := time.Now()
	fmt.Printf("Part 1: %s (%s)\n", fmt.Sprint(part1(input)), time.Since(start1))

	start2 := time.Now()
	fmt.Printf("Part 2: %s (%s)\n", fmt.Sprint(part2(input)), time.Since(start2))
}
