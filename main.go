package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("     ___       ______     ______     ___     ___    ___    ___   ")
	fmt.Println("    /   \\     /  __  \\   /      |   |__ \\   / _ \\  |__ \\  |__ \\  ")
	fmt.Println("   /  ^  \\   |  |  |  | |  ,----'      ) | | | | |    ) |    ) | ")
	fmt.Println("  /  /_\\  \\  |  |  |  | |  |          / /  | | | |   / /    / /  ")
	fmt.Println(" /  _____  \\ |  `--'  | |  `----.    / /_  | |_| |  / /_   / /_  ")
	fmt.Println("/__/     \\__\\ \\______/   \\______|   |____|  \\___/  |____| |____| ")
	fmt.Println("-------------------------------")
	fmt.Println("Happy Coding & festive season 🎄")

	daysLeft := daysUntilXMas()
	timeLeftMsg := "XMas has already passed 🎅"

	if daysLeft > 0 {
		timeLeftMsg = fmt.Sprintf("Only ~%d days left until XMas ⏱\n", daysLeft)
	}

	fmt.Println(timeLeftMsg)
	fmt.Println("usage: go run day-<nr>/day<nr>p<1|2>")
}

func daysUntilXMas() int {
	xmas := time.Date(2022, 12, 24, 0, 0, 0, 0, time.Local)
	today := time.Now()
	delta := xmas.Sub(today)

	return int(math.RoundToEven(delta.Hours() / 24))
}
