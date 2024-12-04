package leetcode

import "testing"

func Test_longestConsecutive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "single entry",
			nums: []int{1},
			want: 1,
		},
		{
			name: "duplicate entry",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "simple up",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "simple down",
			nums: []int{2, 1},
			want: 2,
		},
		{
			name: "more complicated example",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "failed example",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
