package leetcode

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "base case 1",
			nums:   []int{1},
			target: 1,
			want:   1,
		},
		{
			name:   "base case 2 numbers",
			nums:   []int{1, 1},
			target: 2,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.nums, tt.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
