package main

import (
	"sync"
)

func main() {
	p := sync.Pool{}
	p.Get()
}
