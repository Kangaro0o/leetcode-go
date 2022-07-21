package swordoffer

import "testing"

func Test_kthLargest(t *testing.T) {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}

	root2 := &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 6},
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
	}
	type args struct {
		root *TreeNode
		k    int
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
				k:    1,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				root: root2,
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
