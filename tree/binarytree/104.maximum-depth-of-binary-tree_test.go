package binarytree

import "testing"

func Test_maxDepth(t *testing.T) {
	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	type args struct {
		root *TreeNode
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
