package _go

import "testing"

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "base nil",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "not symmetric - left no right",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
			},
			want: false,
		},
		{
			name: "not symmetric - right no left",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
			want: false,
		},
		{
			name: "symmetric - one level",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 2,
				},
			},
			want: true,
		},
		{
			name: "asymmetric - one level",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			want: false,
		},
		{
			name: "symmetric - two levels",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Right: &TreeNode{Val: 3},
						Left:  &TreeNode{Val: 4},
					},
				},
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
