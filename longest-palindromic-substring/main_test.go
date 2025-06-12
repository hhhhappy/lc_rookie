package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "babad",
			args: args{
				s: "babad",
			},
			want: "aba",
		},
		{
			name: "cbbd",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "bbbb",
			args: args{
				s: "bbbb",
			},
			want: "bbbb",
		},
		{
			name: "bbb",
			args: args{
				s: "bbb",
			},
			want: "bbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
