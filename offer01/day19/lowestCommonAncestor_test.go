package day19

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test_lowestCommonAncestor", struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{root: &TreeNode{3, &TreeNode{Val: 5}, &TreeNode{Val: 1}}, p: &TreeNode{Val: 5}, q: &TreeNode{Val: 1}}, &TreeNode{Val: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
