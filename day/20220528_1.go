package day

import "bytes"

func removeOuterParentheses(s string) string {

	ans := bytes.NewBufferString("")
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
			if len(stack) > 1 {
				ans.WriteRune(c)
			}
		}
		if c == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				ans.WriteRune(c)
			}
		}
	}

	return ans.String()
}
