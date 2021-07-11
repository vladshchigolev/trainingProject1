package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//	Попробовать что-то сделать с указателями
	fmt.Println("Hello, Go!")
	myInt := 5
	fmt.Println(&myInt) // Получаем адрес переменной и выводим его в консоль
	var myIntPointer *int
	fmt.Println(myIntPointer)
	myIntPointer = &myInt
	fmt.Println(myIntPointer, reflect.TypeOf(myIntPointer))
	fmt.Println(stringsReturn())
	myArray := [7]string{"qw", "er", "ty"} // [7] - это количество элементов (длина массива), а индекс последнего элемента будет "6"!
	fmt.Printf("%#v, %#v\n", myArray, myIntPointer)
	fmt.Println(myArray)
	//myArray[0] = "do" // На предыдущей строке мы объявляем переменную с типом [7]string, ВСЁ ЭТО - ТИП! Если будем присваивать переменнной что-то новое, это должен быть НОВЫЙ МАССИВ, состоящий из СЕМИ!! СТРОКОВЫХ!! значений
	//myArray[1] = "re"
	//myArray[2] = "mi"
	//myArray[3] = "fa"
	//myArray[4] = "so"
	//myArray[5] = "la"
	//myArray[6] = "ti"
	//fmt.Println(myArray[3])
	mySlice := []int{1, 2, 3, 4, 5, 6, 23, 45, 74}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 22, 44, 88)
	fmt.Println(mySlice)
	mySlice[11] = 11
	fmt.Println(mySlice)
	fmt.Println(append(mySlice, 33, 55, 99))
	//fmt.Println([6]string{"a", "b", "c", "d"}) // Вернёт безымянный массив [a b c d <пустая строка> <пустая строка>]
	//fmt.Println(myArray)
	//myArray = [7]string{"e", "f", "g"} // Справа от "=" будет создан массив, 3 первых элемента которого инициализируются нашими значениями, а остальные 4 - нулевыми (пустыми строками)
	//for i := 0; i < 7; i++ {
	//	fmt.Println(myArray[i])
	//}

}

func stringsReturn() (string, string, error) {
	err := errors.New("Oops! Something went wrong :(")
	return "Hello,", "Go!", err
}
