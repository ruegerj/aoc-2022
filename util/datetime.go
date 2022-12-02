package util

import (
	"math"
	"time"
)

func DaysUntilXMas() int {
	xmas := time.Date(2022, 12, 24, 0, 0, 0, 0, time.Local)
	today := time.Now()
	delta := xmas.Sub(today)

	return int(math.RoundToEven(delta.Hours() / 24))
}
