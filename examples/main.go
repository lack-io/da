package main

import "fmt"

func main() {
	a := 0
	a = a<<1 + 0
	a = a<<1 + 1
	a = a<<1 + 1
	fmt.Println(a)
}
