package main

import (
	"fmt"
	"fun-gram/internal/handlers/get"
	"net/http"
)

func main() {
	http.HandleFunc("/", get.Home)
	http.HandleFunc("/saudi", get.Saudi)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
