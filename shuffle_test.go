package slice

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	original := make([]int, 1000)
	for i := range original {
		original[i] = i
	}
	s := make([]int, len(original))
	copy(s, original)
	Shuffle(s, nil)
	changed := false
	for i := range original {
		if original[i] != s[i] {
			changed = true
		}
	}
	if !changed {
		t.Errorf("Shuffle() expected to return a slice with changed order, got same")
	}
	if !SameUniqueValues(original, s) {
		t.Errorf("Shuffle(%v) expected to return a slice with the same unique values, got %v", original, s)
	}
}
