package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
}

func main() {
	a := &MyType{}
	b := new(MyType)

	fmt.Println(reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind())
}
