package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		assert.EqualValues(t, -1, Index([]string{}, "12"))
		assert.EqualValues(t, 0, Index([]string{"12"}, "12"))
		assert.EqualValues(t, 0, Index([]string{"12", "13"}, "12"))
		assert.EqualValues(t, 1, Index([]string{"13", "12"}, "12"))
	})
	t.Run("int", func(t *testing.T) {
		assert.EqualValues(t, -1, Index([]int{}, 12))
		assert.EqualValues(t, 0, Index([]int{12}, 12))
		assert.EqualValues(t, 0, Index([]int{12, 13}, 12))
		assert.EqualValues(t, 1, Index([]int{13, 12}, 12))
	})
}
