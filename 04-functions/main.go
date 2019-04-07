package main

import "fmt"

func gretting(name string) string {
	return "Hello " + name
}

func sum(value1, value2 int) int {
	return value1 + value2
}

func main() {
	fmt.Println(gretting("Diego"), sum(3, 2))
}
