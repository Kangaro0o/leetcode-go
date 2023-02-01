package binarysearch

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 3,
				a: 2,
				b: 3,
				c: 5,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				n: 4,
				a: 2,
				b: 3,
				c: 4,
			},
			want: 6,
		},
		{
			name: "case3",
			args: args{
				n: 5,
				a: 2,
				b: 11,
				c: 13,
			},
			want: 10,
		},
		{
			name: "case4",
			args: args{
				n: 1000000000,
				a: 2,
				b: 217983653,
				c: 336916467,
			},
			want: 1999999984,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
