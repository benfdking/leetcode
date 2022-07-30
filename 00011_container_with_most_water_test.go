package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "base 0",
			height: []int{0, 0},
			want:   0,
		},
		{
			name:   "base 0 with elevation",
			height: []int{0, 2},
			want:   0,
		},
		{
			name:   "base 1",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "base 1 with difference",
			height: []int{1, 2},
			want:   1,
		},
		{
			name:   "leetcode example",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
