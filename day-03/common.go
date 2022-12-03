package day03

import "strings"

func getMatches(a string, b string) []string {
	matches := make([]string, 0)

	for _, item := range strings.Split(a, "") {
		if !strings.Contains(b, item) {
			continue
		}

		matches = append(matches, item)
	}

	return matches
}

func calcItemPriority(item string) int {
	itemPriority := 0

	lowerCaseAlphabet := "abcdefghijklmnopqrstuvwxyz"

	if strings.Index(lowerCaseAlphabet, item) == -1 {
		item = strings.ToLower(item)
		itemPriority += 26
	}

	return itemPriority + strings.Index(lowerCaseAlphabet, item) + 1
}
