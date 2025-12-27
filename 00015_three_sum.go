package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// Sort the given array
	sort.Ints(nums)

	// find unique entries
	var outs [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			three := nums[i] + nums[l] + nums[r]
			if three > 0 {
				r--
			} else if three < 0 {
				l++
			} else {
				outs = append(outs, []int{nums[i], nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return outs
}
