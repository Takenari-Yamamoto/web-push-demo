package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Takenari san!")
	})

	port := os.Getenv("PORT")
	fmt.Println("ポートだよぉん", port)
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
