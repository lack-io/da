package day26

import "strings"

func isNumber(s string) bool {
	arr := strings.TrimSpace(s)
	if len(arr) == 0 {
		return false
	}
	index := 0
	if arr[index] == '+' || arr[index] == '-' {
		index++
	}
	cur := index
	index = scanUnsignedInteger(arr, index)
	numeric := index > cur
	if index < len(arr) && arr[index] == '.' {
		//略过小数点
		index++
		cur = index
		index = scanUnsignedInteger(arr, index)
		numeric = (index > cur) || numeric
	}
	if index < len(arr) && (arr[index] == 'e' || arr[index] == 'E') {
		//略过e/E
		index++
		//略过符号位
		if index < len(arr) && (arr[index] == '+' || arr[index] == '-') {
			index++
		}
		cur = index
		index = scanUnsignedInteger(arr, index)
		numeric = (index > cur) && numeric
	}
	return numeric && index == len(arr)
}

func scanUnsignedInteger(arr string, index int) int {
	for index < len(arr) && arr[index] >= '0' && arr[index] <= '9' {
		index++
	}
	return index
}
