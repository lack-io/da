package listNode

import (
	"reflect"
	"testing"
)

func TestGetElement(t *testing.T) {
	type args struct {
		head  *ListNode
		index int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"TestGetElement-1",
			args{head: GenerateListNode(1, 2, 3), index: 0},
			GenerateListNode(1),
		},
		{
			"TestGetElement-2",
			args{head: GenerateListNode(1, 2, 3), index: 1},
			GenerateListNode(2),
		},
		{
			"TestGetElement-3",
			args{head: GenerateListNode(1, 2, 3), index: 2},
			GenerateListNode(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetElement(tt.args.head, tt.args.index); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("GetElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		head  *ListNode
		index int
		val   int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"TestInsert-1",
			args{
				head:  GenerateListNode(1, 3, 4),
				index: 1,
				val:   2,
			},
			GenerateListNode(1, 2, 3, 4),
		},
		{
			"TestInsert-2",
			args{
				head:  GenerateListNode(1, 3, 4),
				index: 2,
				val:   2,
			},
			GenerateListNode(1, 3, 2, 4),
		},
		{
			"TestInsert-3",
			args{
				head:  GenerateListNode(1, 3, 4),
				index: 0,
				val:   2,
			},
			GenerateListNode(2, 1, 3, 4),
		},
		{
			"TestInsert-4",
			args{
				head:  GenerateListNode(1, 3, 4),
				index: 4,
				val:   2,
			},
			GenerateListNode(1, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.head, tt.args.index, tt.args.val); got.String() != tt.want.String() {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
