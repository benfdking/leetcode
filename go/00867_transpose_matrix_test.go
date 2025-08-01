package _go

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "empty",
			matrix: [][]int{},
			want:   [][]int{},
		},
		{
			name:   "base 1",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name:   "base 1 and 2",
			matrix: [][]int{{1, 2}},
			want:   [][]int{{1}, {2}},
		},
		{
			name:   "base square",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   [][]int{{1, 3}, {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
