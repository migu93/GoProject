package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":8080", nil)
}
