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

}
