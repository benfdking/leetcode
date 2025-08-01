package _go

func twoSum2(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum < target:
			i++
		case sum > target:
			j--
		}
	}
}
