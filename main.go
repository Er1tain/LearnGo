package main

import (
	"fmt"
)

type Person struct {
	name string
	male string
	age  int
}

func toString(person Person) string {
	res := person.name + " " + person.male + " " + string(person.age)

	fmt.Println(res)
	return res
}

func main() {
	fmt.Println("Hi")

	var a int = 3

	fmt.Println(a)

	Alex := Person{
		name: "Alex",
		male: "man",
		age:  20,
	}

	fmt.Println(Alex)
	toString(Alex)
}
