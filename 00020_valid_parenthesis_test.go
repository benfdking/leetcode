package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "",
			want: true,
		},
		{
			s:    "()",
			want: true,
		},
		{
			s:    "()[]{}",
			want: true,
		},
		{
			s:    "(]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
