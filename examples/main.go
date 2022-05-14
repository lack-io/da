package main

import (
	"fmt"
	"log"
)

func main() {
	a := ""
	n, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a, n)

}
