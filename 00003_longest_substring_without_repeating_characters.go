package leetcode

func lengthOfLongestSubstring(s string) int {
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
