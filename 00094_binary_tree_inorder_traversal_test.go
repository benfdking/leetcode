package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "base",
			root: nil,
			want: []int{},
		},
		{
			name: "single value",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "left and right value",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
			want: []int{3, 1, 2},
		},
		{
			name: "multiple levels",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{
						Val: 3,
					}},
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
