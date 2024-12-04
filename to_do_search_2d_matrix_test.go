package leetcode

import "testing"

func Test_searchValue(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		n    int
		want bool
	}{
		{
			name: "base present",
			m:    [][]int{{1}},
			n:    1,
			want: true,
		},
		{
			name: "base not present",
			m:    [][]int{{1}},
			n:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchValue(tt.m, tt.n); got != tt.want {
				t.Errorf("searchValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
