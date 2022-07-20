package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var arrayA []*ListNode
	var arrayB []*ListNode

	holdA := headA
	for holdA != nil {
		arrayA = append(arrayA, headA)
		holdA = holdA.Next
	}
	holdB := headB
	for holdB != nil {
		arrayB = append(arrayB, headB)
		holdB = holdB.Next
	}

	var intersection *ListNode
	for i := 1; i <= min(len(arrayA), len(arrayB)); i++ {
		if arrayA[len(arrayA)-i] == arrayB[len(arrayB)-i] {
			intersection = arrayA[len(arrayA)-i]
		}
	}

	return intersection
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
