package slice

import "slices"

// SameUniqueValues compares two slices of comparable types.
func SameUniqueValues[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) < len(b) {
		a, b = b, a
	}
	for _, v := range a {
		if !slices.Contains(b, v) {
			return false
		}
	}
	return true
}
