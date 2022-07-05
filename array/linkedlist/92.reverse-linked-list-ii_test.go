package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head:  initLinkedList(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: initLinkedList(1, 4, 3, 2, 5),
		},
		{
			name: "case2",
			args: args{
				head:  initLinkedList(5),
				left:  1,
				right: 1,
			},
			want: initLinkedList(5),
		},
		{
			name: "case3",
			args: args{
				head:  initLinkedList(3, 5),
				left:  1,
				right: 2,
			},
			want: initLinkedList(5, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
