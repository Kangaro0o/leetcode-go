package bst

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
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
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
