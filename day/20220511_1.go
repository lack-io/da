package day

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	out := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			if node == nil {
				out = append(out, "null")
			} else {
				out = append(out, strconv.Itoa(node.Val))
				tmp = append(tmp, node.Left)
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	return fmt.Sprintf("%v", strings.Join(out, ","))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.TrimPrefix(data, "[")
	data = strings.TrimSuffix(data, "]")
	parts := strings.Split(data, ",")
	if len(parts) == 0 || strings.TrimSpace(parts[0]) == "" {
		return nil
	}
	nodeVal, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	root := &TreeNode{Val: nodeVal}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 1; i < len(parts); i += 2 {
		node := queue[0]
		queue = queue[1:]

		if val := strings.TrimSpace(parts[i]); val != "null" {
			left := &TreeNode{}
			left.Val, _ = strconv.Atoi(val)
			node.Left = left
			queue = append(queue, left)
		}
		if val := strings.TrimSpace(parts[i+1]); val != "null" {

			right := &TreeNode{}
			right.Val, _ = strconv.Atoi(val)
			node.Right = right
			queue = append(queue, right)
		}
	}

	return root
}
