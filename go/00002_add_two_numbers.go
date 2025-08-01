package _go

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{}
	l1Sub := l1
	l2Sub := l2
	carry := 0
	start := out
	for {
		total := l1Sub.Val + l2Sub.Val + carry
		carry = total / 10
		start.Val = total % 10

		if l1Sub.Next == nil && l2Sub.Next == nil && carry == 0 {
			return out
		} else {
			if l1Sub.Next == nil {
				l1Sub.Next = &ListNode{}
			}
			if l2Sub.Next == nil {
				l2Sub.Next = &ListNode{}
			}
		}

		l1Sub = l1Sub.Next
		l2Sub = l2Sub.Next
		start.Next = &ListNode{}
		start = start.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
