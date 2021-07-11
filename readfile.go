package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.Open("test_file.txt")
	fmt.Println(file) // В переменной file находится ссылка на структуру, здесь мы выведем ссылку на структуру, но при этом мы получим саму структуру
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner.Scan())
	fmt.Println(scanner.Text())
	fmt.Println(scanner.Scan())
	fmt.Println(scanner.Text())
	fmt.Println(scanner.Scan())
	//////////////////
	now := time.Now()
	fmt.Println(now)
	//////////////////

}
