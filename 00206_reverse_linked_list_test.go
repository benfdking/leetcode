package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "base empty",
			head: nil,
			want: nil,
		},
		{
			name: "base 1",
			head: &ListNode{Val: 1},
			want: &ListNode{Val: 1},
		},
		{
			name: "base 2",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			name: "base 3",
			head: &ListNode{Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: &ListNode{Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
