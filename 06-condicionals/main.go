package main

import "fmt"

func main() {
	x, y := 7, 7

	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 1")
	case 3:
		fmt.Println("x is 1")
	default:
		fmt.Println("x is something else")
	}

	fmt.Println("Hello World")
}
