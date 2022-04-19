package da

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
