package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return &ListNode{
			Val:  list2.Val,
			Next: list2.Next,
		}
	}
	if list2 == nil {
		return &ListNode{
			Val:  list1.Val,
			Next: list1.Next,
		}
	}
	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	}
	return &ListNode{
		Val:  list2.Val,
		Next: mergeTwoLists(list1, list2.Next),
	}
}
