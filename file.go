package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//num := 5.3
	//fmt.Println(int(num))
	//fmt.Println(reflect.TypeOf(int(num)))
	//var now time.Time = time.Now() // "Time" здесь - это экспортируемый из пакета "time" тип
	//var day int = now.Day()
	//fmt.Println(now)
	//fmt.Println(day)
	//replacer := strings.NewReplacer("s", "o")
	//fmt.Println(replacer)
	//fmt.Println(replacer.Replace("sex drugs rock-n-roll"))
	//var fuel float64 = 10
	//var minus = 3.5
	//fmt.Println(fuel - minus)
	//fmt.Println("Введите ваш процент правильных ответов: ")
	//input := bufio.NewReader(os.Stdin)
	//reader.ReadString('n')
	//fmt.Println(input)
	//i := 100
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Print("Global 'i': ", i)

	fmt.Print("Input a grade: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n') // Читаем введённый с клавиатуры текст
	if err != nil {                                          // Обрабатываем возможную ошибку
		log.Fatal(err)
	}
	//fmt.Println(input, reflect.TypeOf(input))
	input = strings.TrimSpace(input)            // Удаляем все ненужные символы по краям строки, включая перевод строки в конце
	grade, err := strconv.ParseFloat(input, 64) // Преобразуем строку во float64
	if err != nil {                             // Обрабатываем возможную ошибку
		log.Fatal(err)
	}

	var int int = 2
	fmt.Println(int)
	if grade >= 60 {
		i := 2
		fmt.Println(i)
		status := "passing"
		int := "string"
		fmt.Println(int)
		fmt.Println("A grade of", grade, "is", status)
		if true {
			i := 3
			fmt.Println(i)
		}
	} else {
		status := "failing"
		fmt.Println("A grade of", grade, "is", status)
	}

	// предпоследняя команда fmt.Println(c)
	// последняя команда fmt.Println(d)
	// пятая снизу команда fmt.Println(d)

	//fileInfo, _ := os.Stat("test_file.txt")
	//fmt.Println(fileInfo.Name())
	//var var1 int = 1
	//var2 := 3.5 - 1.5
	//var2 = float64(var1)
	//fmt.Println(float64(var1), var2)
	//fmt.Println(var2, reflect.TypeOf(var2))
	//var empty float64
	//fmt.Println("Empty: ", empty)
	//numVar, strVar, boolVar := 23, "str", true
	//fmt.Println(numVar, strVar, boolVar)
}
