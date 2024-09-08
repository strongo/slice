package slice

// RemoveInPlaceByValue removes an item from a slice in place without new memory allocation
func RemoveInPlaceByValue[T comparable](slice []T, v T) []T {
	items, _ := RemoveInPlace(slice, func(item T) bool {
		return item == v
	})
	return items
}

// RemoveInPlace removes an item from a slice in place without new memory allocation
func RemoveInPlace[T any](slice []T, f func(v T) bool) (items []T, removedCount int) {
	for i, item := range slice {
		if f(item) {
			removedCount++
			continue
		}
		slice[i-removedCount] = item
	}
	return slice[:len(slice)-removedCount], removedCount
}
