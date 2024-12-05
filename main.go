package main

import (
    "fmt"
    "net/http"
)

// Handler function for GET requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        fmt.Fprintf(w, "Hello, World!")
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    // Define a route and associate it with the handler function
    http.HandleFunc("/", helloHandler)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
