package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootPathHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received from %q\n", r.URL.Path)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}