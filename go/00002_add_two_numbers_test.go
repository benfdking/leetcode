package _go

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tts := []struct {
		name string
		arg1 *ListNode
		arg2 *ListNode
		want *ListNode
	}{
		{
			name: "base 0s",
			arg1: &ListNode{},
			arg2: &ListNode{},
			want: &ListNode{},
		},
		{
			name: "base 1s",
			arg1: &ListNode{
				Val: 1,
			},
			arg2: &ListNode{
				Val: 1,
			},
			want: &ListNode{
				Val: 2,
			},
		},
		{
			name: "2 layers of 1s",
			arg1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			arg2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "1 longer than 2",
			arg1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			arg2: &ListNode{
				Val: 1,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "2 longer than 1",
			arg1: &ListNode{
				Val: 1,
			},
			arg2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "pure carry addition",
			arg1: &ListNode{
				Val: 6,
			},
			arg2: &ListNode{
				Val: 6,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "carry with value",
			arg1: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 1,
				},
			},
			arg2: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		{
			name: "carry overs with carry",
			arg1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
			arg2: &ListNode{
				Val: 1,
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			out := addTwoNumbers(tt.arg1, tt.arg2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Errorf("want %+v, got %+v", tt.want, out)
			}
		})
	}
}
