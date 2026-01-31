package main

import (
	"log"
	"net/http"
	"rentboard/internal/route"
)

func main() {
	handler := route.NewRouter()
	err := http.ListenAndServe(":8080", handler)
	log.Fatal(err)
}
