package main

import (
	"awesomeProject1/exper"
	"errors"
	"fmt"
	"reflect"
)

var test int8 = 45 // Переменная объявлена на уровне пакета, доступна из всех функций

func sum(a int, b int) int {
	fmt.Println("Результат сложения:", a+b)
	return 0

}

func main() {
	exper.PrintOnScreen("lalabla")
	//fmt.Println(1.0 / 3.0)
	//fmt.Printf("%0.2f\n", 1.0 / 3.0)
	//fmt.Printf("%T", 1.637)
	//sum(3, 1)
	//litersPerWall(4.2, 3.0)
	litersPerWallWithReturn(5.2, 3.5)
	//forFunction := litersPerWallWithReturn(2.8, 1.4)
	//fmt.Println(forFunction)
	//integer := 26
	//fmt.Println(reflect.TypeOf(integer))
	//a, b, c := 10, 12, 64
	//a, b = 23, 46
	//fmt.Println(a, b, c)
	//fmt.Println(test)
	//allWalls(2.3, 5.2, 4.7, 1.5, 9.5, 3.5)
	//fmt.Println()
	//fmt.Print()
	//errCode := sum(4, 5)
	//fmt.Println(errCode)
	varA, varB, varC := 1, true, "string"
	fmt.Println(varA, varB, varC)

}

func litersPerWall(width float64, height float64) {
	area := width * height
	//fmt.Println(reflect.TypeOf(area))
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func litersPerWallWithReturn(width float64, height float64) float64 {
	area := width * height
	err := errors.New("height cant be negative") // Создаём новое значение ошибки
	fmt.Println(reflect.TypeOf(err.Error()))
	fmt.Println(err.Error())
	fmt.Println("Тип значения, возвращаемого errors.New():", reflect.TypeOf(err))
	fmt.Println("Тип значения, возвращаемого методом .Error():", reflect.TypeOf(err.Error()))
	fmt.Println(reflect.TypeOf(reflect.TypeOf(err.Error())))
	fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0
}

func allWalls(width1 float64, height1 float64, width2 float64, height2 float64, width3 float64, height3 float64) {
	wall1 := litersPerWallWithReturn(width1, height1)
	wall2 := litersPerWallWithReturn(width2, height2)
	wall3 := litersPerWallWithReturn(width3, height3)
	fmt.Println("На 3 стены понадобится", wall1+wall2+wall3, "литров краски.")

}

//func status(grade float64) string {
//	if grade < 60.0 {
//		return "failing"
//
//	}
//}
