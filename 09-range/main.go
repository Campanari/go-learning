package main

import "fmt"

func main() {
	ids := []int{34, 23, 24, 6, 22, 57}

	for i, id := range ids {
		fmt.Printf("%d - %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	emails := map[string]string{"test1": "test@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s %s\n", key, value)
	}
}
