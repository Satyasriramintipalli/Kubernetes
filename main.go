package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Satya, Go App Running 🚀")
}

func main() {

	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/", loginPage)
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/create", createUser)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}