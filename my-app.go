package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var greeting = os.Getenv("GREETING")

func main() {
	log.Printf("Listening port 8080: %v", greeting)
	http.HandleFunc("/", handler)
	panic("TODO")
	_ = http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, greeting)
}
