package binarytree

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name: "case1",
			args: args{
				root: buildTreeHelper([]int{1, 2, 3, 4, 5, 6, 0}, 0),
			},
			wantCnt: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := countNodes(tt.args.root); gotCnt != tt.wantCnt {
				t.Errorf("countNodes() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
