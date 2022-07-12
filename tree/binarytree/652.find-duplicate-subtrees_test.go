package binarytree

import (
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 4},
		},
	}

	w1 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 4},
	}
	w2 := &TreeNode{Val: 4}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: root1,
			},
			want: []*TreeNode{w1, w2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
