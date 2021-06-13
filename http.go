package main

import (
	"io"
	"log"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	const message = "Halo Nusantara\n"
	io.WriteString(w, message)
}

func main() {
	const host = "0.0.0.0:8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greetHandler)

	if err := http.ListenAndServe(host, mux); err  != nil {
		log.Fatalf("server failed: %v", err)
	}
}
