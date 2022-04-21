package day17

func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)

	return arr[:k]
}

func quickSort(arr []int, l, r int) {
	// 子数组长度为 1 时终止递归
	if l >= r {
		return
	}
	// 哨兵划分操作（以 arr[l] 作为基准数）
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= arr[l] {
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[l] = arr[l], arr[i]
	// 递归左（右）子数组执行哨兵划分
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
}
