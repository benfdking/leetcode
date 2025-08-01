package _go

// recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// iterative

// Steps:
// 1. Init holder pointers: init, curr, next
// 2. While loop on current != null
//   1. Save next
//   2. Reverse
//   3. Advance previous and current
// 3. Return prev
//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		nextTemp := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = nextTemp
//	}
//	return prev
//}
