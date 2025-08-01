package _go

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "4",
			arg:  4,
			want: 2,
		},
		{
			name: "6",
			arg:  6,
			want: 2,
		},
		{
			name: "8",
			arg:  8,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.arg); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
