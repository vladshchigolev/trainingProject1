package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.Open("test_file.txt")
	fmt.Println(*file)
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
