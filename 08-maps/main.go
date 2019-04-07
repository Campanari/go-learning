package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["test1"] = "test@gmail.com"
	emails["test2"] = "test@gmail.com"
	emails["test3"] = "test@gmail.com"
	emails["test4"] = "test@gmail.com"
	emails["test5"] = "test@gmail.com"
	emails["test6"] = "test@gmail.com"
	emails["test7"] = "test@gmail.com"
	emails["test8"] = "test@gmail.com"
	emails["test9"] = "test@gmail.com"

	fmt.Println("Hello World", emails, len(emails), emails["test9"])

	delete(emails, "test8")

	fmt.Println("Hello World", emails, len(emails), emails["test9"])

	emails2 := map[string]string{"test1": "test@gmail.com"}

	fmt.Println("Hello World", emails2, len(emails2))
}
