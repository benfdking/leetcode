package _go

import "testing"

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "base 0",
			grid: [][]byte{{numIslandsNotLand}},
			want: 0,
		},
		{
			name: "base 1",
			grid: [][]byte{{numIslandsLand}},
			want: 1,
		},
		{
			name: "base 2 vertical",
			grid: [][]byte{
				{numIslandsLand},
				{numIslandsNotLand},
				{numIslandsLand},
			},
			want: 2,
		},
		{
			name: "base 2 horizontal",
			grid: [][]byte{
				{numIslandsLand, numIslandsNotLand, numIslandsLand},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
