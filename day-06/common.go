package day06

import "container/list"

func getTopNValues(list *list.List, n int) []any {
	item := list.Front()

	values := make([]any, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			item = item.Next()
		}

		values[i] = item.Value
	}

	return values
}
