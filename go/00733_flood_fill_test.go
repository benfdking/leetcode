package _go

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "base 0",
			args: args{
				image: [][]int{{0}},
				sr:    0,
				sc:    0,
				color: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "simple horizontal connected 2",
			args: args{
				image: [][]int{{0, 0}},
				sr:    0,
				sc:    0,
				color: 1,
			},
			want: [][]int{{1, 1}},
		},
		{
			name: "simple vertical connected 2",
			args: args{
				image: [][]int{{0}, {0}},
				sr:    0,
				sc:    0,
				color: 1,
			},
			want: [][]int{{1}, {1}},
		},
		{
			name: "simple connected square",
			args: args{
				image: [][]int{{0, 0}, {0, 0}},
				sr:    0,
				sc:    0,
				color: 1,
			},
			want: [][]int{{1, 1}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
