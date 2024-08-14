package slice

func RemoveMulti[T comparable](source, remove []T) (result []T, removedCount int) {
	result = source[:0]
sourceRange:
	for _, s := range source {
		for _, r := range remove {
			if s == r {
				removedCount++
				continue sourceRange
			}
		}
		result = append(result, s)
	}
	return
}
