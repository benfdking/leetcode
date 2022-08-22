package leetcode

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {
		var out int
		if target == nums[0] {
			out++
		}
		if target == -nums[0] {
			out++
		}
		return out
	}
	return findTargetSumWays(nums[1:], target-nums[0]) +
		findTargetSumWays(nums[1:], target+nums[0])
}
