package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/", loginPage)
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/create", createUser)

	// 🔥 Prometheus endpoint
	r.Handle("/metrics", promhttp.Handler())

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}