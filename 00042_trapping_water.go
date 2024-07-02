package leetcode

func trap(height []int) int {
	total := 0
	// TODO This would be nicer with option
	position1 := 0
	position2 := 0
	counting := false

	for {
		// move position 2 forward
		position2++
		if position2 >= len(height) {
			break
		}

		// if not filling
		// TODO optimise
		if !counting {
			// and going up move position
			if height[position1] < height[position2] {
				position1 = position2
			}
			// if goes down, keep position 1 in place and move down
			if height[position1] > height[position2] {
				counting = true
			}
			if height[position1] == height[position2] {
				position1 = position2
			}
		}

		if counting {
			if height[position1] > height[position2] {
				areAllSmallerThanP2 := true
				for i := position2 + 1; i < len(height); i++ {
					if height[i] > height[position2] {
						areAllSmallerThanP2 = false
					}
				}
				if areAllSmallerThanP2 {
					sum := 0
					h := height[position2]
					for i := position1 + 1; i < position2; i++ {
						sum += h - height[i]
					}
					total += sum
					counting = false
					position1 = position2
				}

			}
			if height[position1] < height[position2] {
				sum := 0
				h := height[position1]
				for i := position1 + 1; i < position2; i++ {
					sum += h - height[i]
				}
				total += sum
				counting = false
				position1 = position2
			}
			if height[position1] == height[position2] {
				// Need to fill
				sum := 0
				h := height[position1]
				for i := position1 + 1; i < position2; i++ {
					sum += h - height[i]
				}
				total += sum
				counting = false
				position1 = position2
			}
		}
	}
	return total
}

func trapONMemory(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	maxLeft := make([]int, len(height))
	maxLeft[0] = 0
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	maxRight := make([]int, len(height))
	maxRight[len(height)-1] = 0
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	mins := make([]int, len(height))
	for i := 0; i < len(mins); i++ {
		mins[i] = min(maxRight[i], maxLeft[i])
	}

	sum := 0
	for i := 0; i < len(mins); i++ {
		diff := mins[i] - height[i]
		if diff > 0 {
			sum += diff
		}
	}

	return sum
}

func trapO1Memory(heights []int) int {
	if len(heights) <= 2 {
		return 0
	}

	sum := 0

	leftPointer := 0
	rightPointer := len(heights) - 1
	maxL := heights[leftPointer]
	maxR := heights[rightPointer]

	for leftPointer < rightPointer-1 {
		if maxL < maxR {
			// shift left
			leftPointer++
			height := heights[leftPointer]
			value := maxL - height
			if value > 0 {
				sum += value
			} else {
				maxL = height
			}
		} else {
			// shift right
			rightPointer--
			height := heights[rightPointer]
			value := maxR - height
			if value > 0 {
				sum += value
			} else {
				maxR = height
			}
		}
	}

	return sum
}
