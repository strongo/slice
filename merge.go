package slice

func Merge[T comparable](source []T, slices ...[]T) (result []T, changed bool) {
	result = source[:]
	for _, slice := range slices {
		for _, v := range slice {
			for _, r := range result {
				if r == v {
					goto found
				}
			}
			result = append(result, v)
			changed = true
		found:
		}
	}
	return
}
