package main

import (
	"fmt"
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Uzai, maior cheater do xadrez")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
