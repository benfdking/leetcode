package leetcode

import "testing"

func Test_TrappingWater(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "empty",
			height: []int{},
			want:   0,
		},
		{
			name:   "one",
			height: []int{1},
			want:   0,
		},
		{
			name:   "two",
			height: []int{1, 2},
			want:   0,
		},
		{
			name:   "three with no holding",
			height: []int{1, 2, 3},
			want:   0,
		},
		{
			name:   "three with unequal holding",
			height: []int{4, 2, 3},
			want:   1,
		},
		{
			name:   "three with holding",
			height: []int{2, 1, 2},
			want:   1,
		},
		{
			name:   "multiple with different height edges",
			height: []int{0, 2, 1, 0, 1, 3},
			want:   4,
		},
		{
			name:   "two dips",
			height: []int{2, 1, 2, 2, 1, 2},
			want:   2,
		},
		{
			name:   "sample",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapO1Memory(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
