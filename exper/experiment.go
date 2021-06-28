package exper

import "fmt"

func PrintOnScreen(value string) {
	fmt.Println("Эта функция вызвана из другого пакета", value)
}

func myFunction() {
	fmt.Println("Privet! Я в файле experiment.go!")
}
