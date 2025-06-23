package main

import (
	"testing"
)

func Test_isRightVal(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				num: 121,
			},
			want: true,
		},
		{
			args: args{
				num: 0,
			},
			want: true,
		},
		{
			args: args{
				num: 1,
			},
			want: true,
		},
		{
			args: args{
				num: 1234,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRightVal(tt.args.num); got != tt.want {
				t.Errorf("isRightVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumCost(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// {
		// 	args: args{
		// 		[]int{101, 115, 116, 120, 122},
		// 	},
		// 	want: 33,
		// },
		{
			args: args{
				[]int{10, 12, 13, 14, 15},
			},
			want: 11,
		},
		// {
		// 	args: args{
		// 		[]int{109, 113, 115, 122, 128},
		// 	},
		// 	want: 11,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.nums); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
