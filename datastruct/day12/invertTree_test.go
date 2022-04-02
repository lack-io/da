package day12

import (
	"reflect"
	"testing"
)

func Test_invertTree1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test_invertTree1-1", struct{ root *TreeNode }{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree1() = %v, want %v", got, tt.want)
			}
		})
	}
}
