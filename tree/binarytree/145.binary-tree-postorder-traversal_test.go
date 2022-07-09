package binarytree

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "case1",
			args: args{
				root: root1,
			},
			wantRes: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := postorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("postorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
