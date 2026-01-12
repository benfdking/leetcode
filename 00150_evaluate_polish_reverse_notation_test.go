package leetcode

import "testing"

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "single value",
			tokens: []string{"3"},
			want:   3,
		},
		{
			name:   "single +",
			tokens: []string{"3", "2", "+"},
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
