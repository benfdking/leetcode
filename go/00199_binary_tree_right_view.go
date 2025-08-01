package _go

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var left, right []int
	if root.Left != nil {
		left = rightSideView(root.Left)
	}
	if root.Right != nil {
		right = rightSideView(root.Right)
	}
	return append([]int{root.Val}, overrideRightArray(left, right)...)
}

func overrideRightArray(left, right []int) []int {
	out := make([]int, max(len(left), len(right)))
	for i := range left {
		out[i] = left[i]
	}
	for i := range right {
		out[i] = right[i]
	}
	return out
}
