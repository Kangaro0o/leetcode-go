package binarytree

import "testing"

func Test_isChild(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isChild(tt.args.root, tt.args.p); got != tt.want {
				t.Errorf("isChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
