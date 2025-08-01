package _go

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 0",
			s:    "au",
			want: 2,
		},
		{
			name: "Example -1",
			s:    "tmmzuxt",
			want: 5,
		},
		{
			"Example 1",
			"abcabcbb",
			3,
		},
		{
			"Example 2",
			"bbbbb",
			1,
		},
		{
			"Example, 3",
			"pwwkew",
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
