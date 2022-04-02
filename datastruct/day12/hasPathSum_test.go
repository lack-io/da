package day12

import "testing"

func Test_hasPathSum2(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test_hasPathSum2-1", args: struct {
			root      *TreeNode
			targetSum int
		}{root: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}, targetSum: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum2(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
