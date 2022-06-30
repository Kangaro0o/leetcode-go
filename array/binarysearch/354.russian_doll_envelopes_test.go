package binarysearch

import (
	"fmt"
	"sort"
	"testing"
)

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				envelopes: Envelop{
					[]int{5, 4},
					[]int{6, 4},
					[]int{6, 7},
					[]int{2, 3},
				},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				envelopes: Envelop{
					[]int{1, 1},
					[]int{1, 1},
					[]int{1, 1},
					[]int{1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				envelopes: Envelop{
					[]int{1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case4",
			args: args{
				envelopes: Envelop{
					[]int{1, 2},
					[]int{2, 3},
					[]int{3, 4},
					[]int{3, 5},
					[]int{4, 5},
					[]int{5, 5},
					[]int{5, 6},
					[]int{6, 7},
					[]int{7, 8},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EnvelopSort(t *testing.T) {
	envelops := Envelop{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}
	sort.Sort(envelops)
	// 打印
	for _, e := range envelops {
		fmt.Printf("%+v\n", e)
	}

}

func Test_getLengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 3, 5, 4, 5, 6, 5, 7, 8},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("getLengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
