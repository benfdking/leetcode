package _go

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "base nil",
			root: nil,
			want: []int{},
		},
		{
			name: "base 1",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "base 2 on right",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: []int{1, 2},
		},
		{
			name: "base 2 on left",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			want: []int{1, 2},
		},
		{
			name: "base 3 levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			want: []int{1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
