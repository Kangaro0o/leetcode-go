package binarytree

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
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
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
