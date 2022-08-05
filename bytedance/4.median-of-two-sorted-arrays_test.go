package bytedance

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: float64(2),
		},
		{
			name: "case2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "case3",
			args: args{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	a := 5 / 2
	fmt.Println("a: ", a)
	b1 := 2
	b2 := 3
	b := float64(b1+b2) / 2
	fmt.Println("b: ", b)
}
