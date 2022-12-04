package day04

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeLimits(sectionRange string) (int, int) {
	limits := strings.Split(sectionRange, "-")

	lower, err := strconv.Atoi(limits[0])

	if err != nil {
		fmt.Println("Failed to convert int")
		os.Exit(1)
	}

	upper, err := strconv.Atoi(limits[1])

	if err != nil {
		fmt.Println("Failed to convert int")
		os.Exit(1)
	}

	return lower, upper
}
