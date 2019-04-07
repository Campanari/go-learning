package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	i := 1

	for i < 10 {
		fmt.Println(i)
		i = i + 1 //i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}
}
