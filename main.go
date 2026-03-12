package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/", loginPage)
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/create", createUser)

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", r)
}