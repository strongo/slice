package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSameUniqueValues(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "empty_integers", want: true},
		{name: "1_and_2_items", args: args[int]{a: []int{1}, b: []int{1, 2}}, want: false},
		{name: "one_equal", args: args[int]{a: []int{1}, b: []int{1}}, want: true},
		{name: "one_unequal", args: args[int]{a: []int{1}, b: []int{2}}, want: false},
		{name: "3_equal_same_order", args: args[int]{a: []int{1, 3, 5}, b: []int{1, 3, 5}}, want: true},
		{name: "3_equal_reverse_order", args: args[int]{a: []int{1, 3, 5}, b: []int{5, 3, 1}}, want: true},
		{name: "3_equal_different_order", args: args[int]{a: []int{1, 3, 5}, b: []int{3, 5, 1}}, want: true},
		{name: "3_unequal", args: args[int]{a: []int{1, 3, 5}, b: []int{1, 3, 6}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SameUniqueValues(tt.args.a, tt.args.b)
			assert.Equalf(t, tt.want, actual, "Equal(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
