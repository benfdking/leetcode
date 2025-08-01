package _go

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "empty",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "non empty l1",
			list1: &ListNode{Val: 1},
			list2: nil,
			want:  &ListNode{Val: 1},
		},
		{
			name:  "non empty l2",
			list1: nil,
			list2: &ListNode{Val: 1},
			want:  &ListNode{Val: 1},
		},
		{
			name:  "non empty l1 l2 one length bigger 2",
			list1: &ListNode{Val: 1},
			list2: &ListNode{Val: 2},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name:  "non empty l1 l2 one length bigger 1",
			list1: &ListNode{Val: 2},
			list2: &ListNode{Val: 1},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
