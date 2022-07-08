package binarytree

import (
	"reflect"
	"testing"
)

func Test_flipMatchVoyage(t *testing.T) {
	root1 := buildTreeHelper([]int{1, 2}, 0)

	root2 := buildTreeHelper([]int{1, 2, 3}, 0)

	root3 := buildTreeHelper([]int{1, 2, 3}, 0)

	root4 := buildTreeHelper([]int{1, 2, 0, 3}, 0)

	root5 := buildTreeHelper(
		[]int{30, 26, 11, 1, 13, 0, 8, 20, 19, 27, 3, 0, 7, 0, 0, 29, 25, 0, 6, 28, 0, 17, 18, 4, 0, 21, 23, 0, 0, 12, 0, 0, 10, 0, 0, 0, 0, 9, 5, 16, 2, 0, 0, 0, 0, 0, 0, 0, 15, 0, 22, 0, 0, 0, 14, 24},
		0,
	)
	type args struct {
		root   *TreeNode
		voyage []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				root:   root1,
				voyage: []int{2, 1},
			},
			want: []int{-1},
		},
		{
			name: "case2",
			args: args{
				root:   root2,
				voyage: []int{1, 3, 2},
			},
			want: []int{1},
		},
		{
			name: "case3",
			args: args{
				root:   root3,
				voyage: []int{1, 2, 3},
			},
			want: nil,
		},
		{
			name: "case4",
			args: args{
				root:   root4,
				voyage: []int{1, 3, 2},
			},
			want: []int{-1},
		},
		{
			name: "case5",
			args: args{
				root:   root5,
				voyage: []int{30, 13, 8, 7, 17, 10, 18, 26, 11, 3, 28, 12, 27, 6, 1, 19, 29, 4, 25, 21, 5, 15, 14, 9, 23, 2, 16, 22, 24, 20},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipMatchVoyage(tt.args.root, tt.args.voyage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipMatchVoyage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTreeHelper(nums []int, i int) *TreeNode {
	if len(nums) == 0 || nums == nil || i >= len(nums) {
		return nil
	}
	if nums[i] == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[i],
	}
	root.Left = buildTreeHelper(nums, 2*i+1)
	root.Right = buildTreeHelper(nums, 2*i+2)
	return root
}
