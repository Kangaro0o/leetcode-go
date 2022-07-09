package binarytree

import (
	"fmt"
	"testing"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name: "case1",
			args: args{
				root: buildTreeHelper([]int{1, 2, 3, 4}, 0),
			},
			wantRes: "1(2(4))(3)",
		},
		{
			name: "case2",
			args: args{
				root: buildTreeHelper([]int{1, 2, 3, 0, 4}, 0),
			},
			wantRes: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tree2str(tt.args.root); gotRes != tt.wantRes {
				t.Errorf("tree2str() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func foo() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func TestDefer(t *testing.T) {
	fmt.Println(foo())
}
