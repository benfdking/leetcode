package leetcode

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "base empty",
			root: nil,
			want: 0,
		},
		{
			name: "base single value",
			root: &TreeNode{Val: 1},
			want: 0,
		},
		{
			name: "base single left",
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			want: 1,
		},
		{
			name: "base single right",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: 1,
		},
		{
			name: "base two left",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 2,
		},
		{
			name: "base double right",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 2,
		},
		{
			name: "triangle",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			want: 2,
		},
		{
			name: "more complex example",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
