package shunfeng

import "testing"

func Test_sumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("sumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
