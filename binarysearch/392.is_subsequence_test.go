package binarysearch

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				s: "",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				s: "",
				t: "",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				s: "b",
				t: "abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
