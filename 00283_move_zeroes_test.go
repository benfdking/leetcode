package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name             string
		nums             []int
		wantToEqualAfter []int
	}{
		{
			name:             "base empty",
			nums:             []int{},
			wantToEqualAfter: []int{},
		},
		{
			name:             "base single 0 value",
			nums:             []int{0},
			wantToEqualAfter: []int{0},
		},
		{
			name:             "base single non-0 value",
			nums:             []int{1},
			wantToEqualAfter: []int{1},
		},
		{
			name:             "base single already correct",
			nums:             []int{1, 0},
			wantToEqualAfter: []int{1, 0},
		},
		{
			name:             "base single needs adjustment",
			nums:             []int{0, 1},
			wantToEqualAfter: []int{1, 0},
		},
		// TODO: Needs more complicated example
		//{
		//	name:             "base single needs adjustment",
		//	nums:             []int{0, 1},
		//	wantToEqualAfter: []int{1, 0},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			moveZeroes(nums)
			if !reflect.DeepEqual(nums, tt.wantToEqualAfter) {
				t.Error("not equal")
			}
		})
	}
}
