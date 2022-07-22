package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	return &ListNode{
		Val: head.Next.Val,
		Next: reverseList(
			&ListNode{
				Val:  head.Val,
				Next: head.Next.Next,
			},
		),
	}
}
