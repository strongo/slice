package slice

// Index returns index of an item  in a slice
func Index[T comparable](items []T, item T) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}
