package main

import (
	"log"
	"net/http"

	"dev.azure.com/rchi-texas/Golang/server"
)

func main() {
	http.HandleFunc("/", server.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
