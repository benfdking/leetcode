package leetcode

func sortColors(nums []int) {
	count0 := 0
	count1 := 0
	count2 := 0
	for _, v := range nums {
		switch v {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		default:
			panic("bad number")
		}
	}
	for i := 0; i < count0; i++ {
		nums[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		nums[i] = 1
	}
	for i := count0 + count1; i < count0+count1+count2; i++ {
		nums[i] = 2
	}
}
