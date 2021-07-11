package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       uint8
}

func (pers Person) speak() {
	fmt.Printf("Hi, my name is %s %s. I'm %d\n", pers.firstName, pers.lastName, pers.age)
}

func main() {
	structs()
}

func structs() {
	user000001 := Person{
		firstName: "Aleksey",
		lastName:  "Ivanov",
		age:       26,
	}
	fmt.Println(user000001)
	fmt.Println(user000001.lastName)
	myStructPointer := &user000001
	fmt.Println(myStructPointer)
	user000001.speak()

}
