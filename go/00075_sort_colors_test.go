package _go

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "empty",
			nums: []int{},
			want: []int{},
		},
		{
			name: "base with all",
			nums: []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sortColors(tt.nums); !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("climbStairs() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
