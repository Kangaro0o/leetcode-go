package pdd

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		nums  []int
		start int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				nums:  []int{1, 2, 3},
				start: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.nums, tt.args.start)
			fmt.Println("nums:", tt.args.nums)
		})
	}
}

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{nums: []int{4, 2, 0, 2, 3, 2, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			fmt.Println("nums:", tt.args.nums)
		})
	}
}
