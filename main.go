package main

import (
	"fmt"
	"log"
	"net/http"

	"GoLangServer/handlers"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api/hello", HelloHandler)
	http.HandleFunc("/api/user", handlers.CreateUser)

	fmt.Println("Server started at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go Api")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Hello World"}`))
}
