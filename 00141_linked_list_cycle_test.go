package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "single node",
			head: &ListNode{Val: 1},
			want: false,
		},
		{
			name: "single node",
			head: &ListNode{Val: 1},
			want: false,
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
