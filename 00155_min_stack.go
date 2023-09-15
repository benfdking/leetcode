package leetcode

type MinStack struct {
	values []int
	min    []int
}

func Constructor() MinStack {
	return MinStack{
		values: []int{},
		min:    []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.values = append(ms.values, val)
	appendValue := val
	if len(ms.min) != 0 {
		v := ms.min[len(ms.min)-1]
		if appendValue > v {
			appendValue = v
		}
	}
	ms.min = append(ms.min, appendValue)
}

func (ms *MinStack) Pop() {
	ms.values = ms.values[0 : len(ms.values)-1]
	ms.min = ms.min[0 : len(ms.min)-1]
}

func (ms *MinStack) Top() int {
	return ms.values[len(ms.values)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
