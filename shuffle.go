package slice

import (
	"math/rand"
	"time"
)

// Shuffle pseudo-randomizes the order of elements.
func Shuffle[T any](s []T, r *rand.Rand) {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
