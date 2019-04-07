package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	fruits2 := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println("Hello World", fruits, fruits2, len(fruits2), fruits2[1:3])
}
