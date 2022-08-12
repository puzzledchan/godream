package main

import (
	"log"
	"net/http"

	"main/base"
)

func main() {
	http.HandleFunc("/", base.IndexHandler)
	http.HandleFunc("/post", base.HelloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
