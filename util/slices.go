package util

func Max(slice []int) int {
	max := slice[0]

	for _, item := range slice {
		if item <= max {
			continue
		}

		max = item
	}

	return max
}

func Reverse[TItem any](slice []TItem) []TItem {
	reversed := make([]TItem, len(slice))
	copy(reversed, slice)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
