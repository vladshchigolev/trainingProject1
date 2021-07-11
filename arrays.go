package main

import (
	"fmt"
)

func main() {
	arrayA := [10][5]int{}
	fmt.Println(arrayA)
	arrayA[0] = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < 5; j++ {
			arrayA[i][j] = j + 1
		}
	}
	var arrayB [4]int
	for index, value := range arrayB {
		value = index * 10
		fmt.Println(value)
	}
	fmt.Println(arrayA)
	arrayB[0] = 3
	arrayB[2] = 1
	fmt.Println(arrayB)
	arrayB = [4]int{4}
	fmt.Println(arrayB)
	sliceA := []string{}
	fmt.Println(sliceA)
}
