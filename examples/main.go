package main

import "fmt"

func main() {
	num := 3
	nums := make([]int, 0)

	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	fmt.Println(num)

	fmt.Println(nums)
}
