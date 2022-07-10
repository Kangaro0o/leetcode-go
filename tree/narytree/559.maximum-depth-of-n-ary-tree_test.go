package narytree

import "testing"

func Test_maxDepth(t *testing.T) {
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	type args struct {
		root *Node
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth1(t *testing.T) {
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}

	type args struct {
		root *Node
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth1(tt.args.root); got != tt.want {
				t.Errorf("maxDepth1() = %v, want %v", got, tt.want)
			}
		})
	}
}
