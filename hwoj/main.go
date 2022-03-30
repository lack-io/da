package main

import (
	"fmt"
	"log"

	"github.com/lack-io/da/hwoj/simple"
)

func main() {
	fmt.Println("input:")
	outs, err := simple.MemoryAlloc()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\noutput:\n", outs)
}
