package math

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
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
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "case4",
			args: args{
				n: 30,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
