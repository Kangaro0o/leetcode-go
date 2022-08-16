package bst

import (
	"testing"
)

func Test_maxSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
	}

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 2},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6},
			},
		},
	}

	root3 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:  6,
				Left: &TreeNode{Val: 9},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   -5,
					Right: &TreeNode{Val: -3},
				},
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 10},
				},
			},
		},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: root1,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				root: root2,
			},
			want: 20,
		},
		{
			name: "case3",
			args: args{
				root: root3,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumBST(tt.args.root); got != tt.want {
				t.Errorf("maxSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
