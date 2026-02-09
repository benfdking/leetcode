package leetcode

func maxArea(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[j], height[i]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
