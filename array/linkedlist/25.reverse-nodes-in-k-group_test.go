package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: initLinkedList(1, 2, 3, 4, 5),
				k:    2,
			},
			want: initLinkedList(2, 1, 4, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func initLinkedList(nums ...int) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for _, num := range nums {
		p.Next = &ListNode{
			Val: num,
		}
		p = p.Next
	}
	return dummy.Next
}
