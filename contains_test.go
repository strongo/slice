package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		assert.EqualValues(t, false, Contains([]string{}, "12"))
		assert.EqualValues(t, true, Contains([]string{"12"}, "12"))
		assert.EqualValues(t, true, Contains([]string{"12", "13"}, "12"))
		assert.EqualValues(t, true, Contains([]string{"13", "12"}, "12"))
	})
	t.Run("int", func(t *testing.T) {
		assert.EqualValues(t, false, Contains([]int{}, 12))
		assert.EqualValues(t, true, Contains([]int{12}, 12))
		assert.EqualValues(t, true, Contains([]int{12, 13}, 12))
		assert.EqualValues(t, true, Contains([]int{13, 12}, 12))
	})
}
