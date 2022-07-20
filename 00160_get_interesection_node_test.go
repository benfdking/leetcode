package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	singularNode := &ListNode{Val: 1}

	tests := []struct {
		name  string
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		{
			name:  "base with all empty",
			headA: nil,
			headB: nil,
			want:  nil,
		},
		{
			name:  "base no intersection with actual values",
			headA: &ListNode{Val: 1},
			headB: &ListNode{Val: 2},
			want:  nil,
		},
		{
			name:  "base single intersection",
			headA: singularNode,
			headB: singularNode,
			want:  singularNode,
		},
		{
			name:  "base single intersection in chain",
			headA: &ListNode{Val: 1, Next: singularNode},
			headB: &ListNode{Val: 2, Next: singularNode},
			want:  singularNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.headA, tt.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
