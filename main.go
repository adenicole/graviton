package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("hit index")

	fmt.Fprintf(w, "Hello World!")
}

func handle() {
	http.HandleFunc("/", index)

	log.Println("serving on :8080")
	panic(http.ListenAndServe(":8080", nil))
}

func main() {
	handle()
}
