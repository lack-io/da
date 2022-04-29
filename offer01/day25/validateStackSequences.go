package day25

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/solution/mian-shi-ti-31-zhan-de-ya-ru-dan-chu-xu-lie-mo-n-2/
func validateStackSequences(pushed, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, num := range pushed {
		stack = append(stack, num)                                // num 入栈
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] { // 循环判断与出栈
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
