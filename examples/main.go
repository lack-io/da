package main

import "fmt"

func main() {
	bit := 0

	//bit |= 1 << 1
	bit |= 1 << 2

	fmt.Println(bit&(1<<1) == 1<<1)
}
