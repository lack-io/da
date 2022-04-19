package day15

import (
	"reflect"
	"testing"
)

func DPGenerate(nums []interface{}) *TreeNode {
	queue := make([]*TreeNode, 0)

	length := len(nums)
	if length == 0 {
		return nil
	}

	assert := func(v interface{}) *TreeNode {
		if v == nil {
			return nil
		}
		if vv, ok := v.(int); ok {
			return &TreeNode{Val: vv}
		}
		return nil
	}

	head := assert(nums[0])
	if head == nil {
		return nil
	}
	queue = append(queue, head)

	i := 1
	for i < length && len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		node := assert(nums[i])
		parent.Left = node
		if node != nil {
			queue = append(queue, node)
		}
		if i < length-1 {
			node = assert(nums[i+1])
			parent.Right = node
			if node != nil {
				queue = append(queue, node)
			}
		}

		i += 2
	}

	return head
}

func Test_pathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test_pathSum-1", struct {
			root   *TreeNode
			target int
		}{root: DPGenerate([]interface{}{1, 2, 2}), target: 3},
			[][]int{{1, 2}, {1, 2}}},
		{"Test_pathSum-2", struct {
			root   *TreeNode
			target int
		}{root: DPGenerate([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), target: 22},
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
