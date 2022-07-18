package leetcode

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums)/2-1)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	panic("error invalid")
}
