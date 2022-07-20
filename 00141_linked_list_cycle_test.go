package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "empty",
			head: nil,
			want: false,
		},
		{
			name: "single entry",
			head: &ListNode{Val: 1},
			want: false,
		},
		{
			name: "self referencing",
			head: func() *ListNode {
				value := &ListNode{Val: 1}
				value.Next = value
				return value
			}(),
			want: true,
		},
		{
			name: "single loop",
			head: func() *ListNode {
				value := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
				value.Next.Next = value
				return value
			}(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
