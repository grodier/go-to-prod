package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("GET /", helloHandler)
	mux.HandleFunc("GET /{name}/", helloHandler)
	mux.HandleFunc("GET /healthcheck/", healthcheckHandler)

	fmt.Println("Server running on port :8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Issue with server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
