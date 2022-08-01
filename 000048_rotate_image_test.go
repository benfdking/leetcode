package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "base empty",
			matrix: [][]int{},
			want:   [][]int{},
		},
		{
			name:   "base 1",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name: "base 2 by 2",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "base 3 by 3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)

			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("want %+v, got %+v", tt.want, tt.matrix)
			}
		})
	}
}
