package main

import "fmt"

func main() {
	var sliceA []string
	var sliceB []string
	fmt.Println(sliceA)
	sliceB = append(sliceA, "wow")
	fmt.Println("Slice A:", sliceA)
	fmt.Println("Slice B:", sliceB)
}
