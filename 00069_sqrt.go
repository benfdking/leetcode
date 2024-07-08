package leetcode

func mySqrt(x int) int {
	if x < 2 {
		return 1
	}

	left := 1
	right := x / 2
	for left <= right {
		middle := (left + right) / 2
		result := middle * middle

		if result == x {
			return middle
		}
		if result > x {
			right = middle - 1
		} else if result < x {
			left = middle + 1
		}
	}
	return right
}
