package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())       // Инициализация генератора
	targetNumber := rand.Intn(100) + 1 // Генерация случайного числа в диапазоне от 1 до 100
	fmt.Println("Я загадал число в диапазоне от 1 до 100")
	fmt.Println("Какое число я загадал?")
	//fmt.Println(targetNumber)
	success := false
	for guesses := 0; guesses < 10; guesses++ {

		if 10-guesses == 1 {
			fmt.Println("У вас осталась последняя попытка.")
		} else if 10-guesses <= 4 {
			fmt.Println("У вас осталось", 10-guesses, "попытки.")
		} else {
			fmt.Println("У вас осталось", 10-guesses, "попыток.")
		}

		fmt.Print("Ваше предположение: ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input) // ?Преобразование строки к целому числу
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(targetNumber, guess)
		if guess < targetNumber {
			fmt.Println("Ваше число меньше загаданного")
		} else if guess > targetNumber {
			fmt.Println("Ваше число больше загаданного")
		} else {
			fmt.Println("Вы угадали! Это", targetNumber, "!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Не угадал :/")
	}
	// 1. x := 4
	// 2. x <= 6 ?
	// 3. Выполнение команд в теле цикла
	// 4. x++ (x -> 5)
	// 5. x <= 6 ?
	// ...

}
