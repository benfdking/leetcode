package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	within, maxValue := diameterOfBinaryTreeSub(root)
	return max(within, maxValue)
}

func diameterOfBinaryTreeSub(root *TreeNode) (maxWithin int, maxCouldBePartOf int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 0, 0
	}
	if root.Left == nil || root.Right == nil {
		return 1, 1
	}
	leftWithin, leftMax := diameterOfBinaryTreeSub(root.Left)
	rightWithin, rightMax := diameterOfBinaryTreeSub(root.Right)
	return rightWithin + leftWithin + 1, max(leftMax, rightMax) + 2
}
