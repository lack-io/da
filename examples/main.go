package main

import (
	"fmt"

	"github.com/lack-io/da"
)

func main() {
	head := da.DPGenerate([]interface{}{1, 2, 3, 4})
	fmt.Println(head)
}
