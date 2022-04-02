package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)

	for i := 0; i < len(arr); i += 2 {
		if doubled(arr, i, i+1) {
			continue
		}

		exists := false
		for j := i + 2; j < len(arr); j++ {
			if doubled(arr, i, j) {
				exists = true
				arr[i+1], arr[j] = arr[j], arr[i+1]
				break
			} else if doubled(arr, i+1, j) {
				exists = true
				arr[i], arr[j] = arr[j], arr[i]
				break
			}
		}
		if !exists {
			return false
		}
	}

	return true
}

func doubled(arr []int, i, j int) bool {
	if arr[i] == 0 && arr[j] == 0 {
		return true
	}
	if arr[i] == 0 || arr[j] == 0 {
		return false
	}
	return arr[i] == 2*arr[j] || arr[i]*2 == arr[j]
}

func main() {
	fmt.Println(canReorderDoubled([]int{-33, 0}))
}
