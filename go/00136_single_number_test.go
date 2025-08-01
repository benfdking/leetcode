package _go

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "base",
			nums: []int{1},
			want: 1,
		},
		{
			name: "simplest",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "slightly more complicated",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
