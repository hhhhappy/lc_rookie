package main

import "testing"

func Test_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				[][]int{
					[]int{10, 20}, []int{30, 200}, []int{400, 50}, []int{30, 20},
				},
			},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCost(tt.args.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
