package leetcode

func searchValue(m [][]int, n int) bool {
	var row int
	{
		min, max := 0, len(m)
		for {
			if max-min <= 1 {
				break
			}
			mid := (max - min) / 2
			if n > m[mid][0] {
				min = mid
			} else if n < m[mid][0] {
				max = mid
			} else {
				return true
			}
		}
		row = min
	}

	rowInQuestion := m[row]
	{
		min, max := 0, len(m)
		for {
			if (max - min) < 1 {
				break
			}
			mid := (max - min) / 2
			if n > rowInQuestion[mid] {
				min = mid
			} else if n < rowInQuestion[mid] {
				max = mid
			} else {
				return true
			}
		}
	}

	return false
}
