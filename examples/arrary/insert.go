package arrary

func Insert(list []int, index, n int) []int {
	length := len(list)

	// 检查边界
	if index < 0 || index > length {
		return list
	}

	// 数组扩展
	list = append(list, n)

	// 元素交换
	for i := index; i <= length; i++ {
		list[i], list[length] = list[length], list[i]
	}

	return list
}

func Remove(list []int, index int) []int {
	length := len(list)

	if index < 0 || index >= length {
		return list
	}

	// 元素交换
	for i := index; i < length-1; i++ {
		list[i] = list[i+1]
	}

	return list[:length-1]
}
