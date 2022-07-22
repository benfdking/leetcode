package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + diameterOfBinaryTree(root.Right)
	}
	if root.Right == nil {
		return 1 + diameterOfBinaryTree(root.Left)
	}
	return 2 + diameterOfBinaryTree(root.Left) + diameterOfBinaryTree(root.Right)
}
