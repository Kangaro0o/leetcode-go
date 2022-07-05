package binarytree

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				root: root,
			},
			want: []int{1, 3, 2},
		},
		{
			name: "case2",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal1(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
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
				root: root,
			},
			wantRes: []int{1, 3, 2},
		},
		{
			name: "case2",
			args: args{
				root: nil,
			},
			wantRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := inorderTraversal1(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("inorderTraversal1() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
