package _go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[0],
			Right: &TreeNode{
				Val: nums[1],
			},
		}
	}

	middle := len(nums) / 2
	left := nums[0:middle]
	right := nums[middle+1:]

	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
}
