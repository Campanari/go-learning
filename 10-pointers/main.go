package main

import "fmt"

func main() {
	a := 2
	b := &a

	fmt.Println("Hello World", a, b, *b, *&a)
}
