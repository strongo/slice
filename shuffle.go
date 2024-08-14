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
	n := len(s)
	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// Shuffle really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		s[i], s[j] = s[j], s[i]
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		s[i], s[j] = s[j], s[i]
	}
}
