package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s", r.URL.Path)
	})

	http.ListenAndServe(
		"localhost:80", nil)
}

//This works if you do go build one.go then sudo ./one and browse to the PC ip address /help
// localhost/helpme also works.
