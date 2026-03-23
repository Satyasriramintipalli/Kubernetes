package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Satya, Go App Running 🚀")
}

func main() {
	http.HandleFunc("/", home)

	// 🔥 THIS IS IMPORTANT
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}