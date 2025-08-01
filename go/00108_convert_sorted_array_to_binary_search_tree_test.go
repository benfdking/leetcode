package _go

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			name: "empty",
			nums: []int{},
			want: nil,
		},
		{
			name: "single entry",
			nums: []int{1},
			want: &TreeNode{Val: 1},
		},
		{
			name: "1,3 entry",
			nums: []int{1, 3},
			want: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "multiple structures",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -10,
					Right: &TreeNode{
						Val: -3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
