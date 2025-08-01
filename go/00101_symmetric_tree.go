package _go

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	return isSymmetric(
		&TreeNode{
			Right: root.Left.Right,
			Left:  root.Right.Left,
		}) && isSymmetric(
		&TreeNode{
			Right: root.Right.Right,
			Left:  root.Left.Left,
		})
}
