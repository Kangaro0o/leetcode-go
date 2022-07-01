package other

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
		},
		{
			name: "case2",
			args: args{
				nums: []int{0},
			},
		},
		{
			name: "case3",
			args: args{
				nums: []int{1, 2, 3, 0, 10, 9, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
