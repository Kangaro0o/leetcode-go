package binarytree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
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
				root: root1,
			},
			wantRes: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := preorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("preorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_preorderTraversal1(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
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
				root: root1,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal1() = %v, want %v", got, tt.want)
			}
		})
	}
}
