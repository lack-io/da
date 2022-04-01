package day11

import "testing"

func Test_isSymmetric1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			struct{ root *TreeNode }{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
					Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric1(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric1() = %v, want %v", got, tt.want)
			}
		})
	}
}
