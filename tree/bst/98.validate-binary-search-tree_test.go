package bst

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	root2 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}

	root3 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 4},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root: root1,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				root: root2,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				root: root3,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				root: &TreeNode{Val: 0},
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
