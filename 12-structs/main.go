package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName string
	gender              string
	age                 int
}

func (person Person) greet() string {
	return "Hello " + person.firstName + strconv.Itoa(person.age)
}

func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getMarried(personB *Person) {
	if person.gender == "f" {
		person.lastName = personB.lastName
	}

	if personB.gender == "f" {
		personB.lastName = person.lastName
	}
}

func main() {
	person1 := Person{firstName: "Diego", lastName: "Campanari", age: 31, gender: "m"}

	fmt.Println("Hello World", person1)

	person2 := Person{"Vitoria", "Leite", "f", 24}

	person1.age++

	fmt.Println("Hello World", person2, person1.firstName, person1.age)

	person1.hasBirthday()

	fmt.Println(person1.greet())

	person1.getMarried(&person2)

	fmt.Println("Hello World", person2)
}
