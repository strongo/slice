package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveStringInPlace(t *testing.T) {
	type args struct {
		v     string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{v: "", slice: []string{}},
			want: []string{},
		},
		{
			name: "single",
			args: args{v: "a", slice: []string{"a"}},
			want: []string{},
		},
		{
			name: "all",
			args: args{v: "a", slice: []string{"a", "a", "a"}},
			want: []string{},
		},
		{
			name: "first",
			args: args{v: "a", slice: []string{"a", "b", "c"}},
			want: []string{"b", "c"},
		},
		{
			name: "last",
			args: args{v: "c", slice: []string{"a", "b", "c"}},
			want: []string{"a", "b"},
		},
		{
			name: "middle",
			args: args{v: "b", slice: []string{"a", "b", "c"}},
			want: []string{"a", "c"},
		},
		{
			name: "first_and_last",
			args: args{v: "a", slice: []string{"a", "b", "c", "a"}},
			want: []string{"b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveStringInPlace(tt.args.v, tt.args.slice), "RemoveStringInPlace(%v, %v)", tt.args.v, tt.args.slice)
		})
	}
}
