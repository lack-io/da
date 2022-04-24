package day20

// 根据每棵子树度分为不同的情况：
// 	1. d == 1: return false
//  2. d == 2:
//    	 if i > j: j 为根结点，i 为 右子树
//  3. d == 3:
//       if i > j > k: j 为根结点， k 为左子树，i 为右子树
func verifyPostorder(postorder []int) bool {
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && verify(postorder, i, m-1) && verify(postorder, m, j-1)
}
