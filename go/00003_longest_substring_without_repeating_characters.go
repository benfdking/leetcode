package _go

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left := 0
	right := 0
	maximumLength := 1
	mapValues := map[byte]int{
		s[0]: 0,
	}
	for right < len(s)-1 {
		right++
		if previousPosition, ok := mapValues[s[right]]; ok {
			mapValues[s[right]] = right
			left = previousPosition + 1
			for k, v := range mapValues {
				if v < previousPosition {
					delete(mapValues, k)
				}
			}
		} else {
			mapValues[s[right]] = right
			newLength := right - left + 1
			if newLength > maximumLength {
				maximumLength = newLength
			}
		}
	}
	return maximumLength
}

func lengthOfLongestSubstringSlow(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		m := make(map[uint8]bool)
		for j := i; j < len(s); j++ {
			if m[s[j]] {
				break
			}
			m[s[j]] = true
			length := j - i + 1
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
