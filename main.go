package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		print(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
