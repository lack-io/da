package day14

import "testing"

func Test_findTarget1(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_findTarget1-1",
			struct {
				root *TreeNode
				k    int
			}{root: &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}, k: 9},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget1(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget1() = %v, want %v", got, tt.want)
			}
		})
	}
}
