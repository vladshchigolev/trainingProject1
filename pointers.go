package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	Попробовать что-то сделать с указателями
	fmt.Println("Hello, Go!")
	myInt := 5
	fmt.Println(&myInt)
	myIntPointer := &myInt
	fmt.Println(reflect.TypeOf(reflect.TypeOf(myIntPointer)))
}
