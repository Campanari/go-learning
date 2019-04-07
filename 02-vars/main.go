package main

import "fmt"

const name = "Diego"

var age int64 = 1

func main() {
	isCool := true
	//isCool = false
	age = 31

	test1, test2 := 2.8, " "

	fmt.Println("Hello World", name, age, isCool, test1, test2)
	fmt.Printf("%T\n", isCool)

}
