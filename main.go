package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello plz work, %q", html.EscapeString(r.URL.Path))
	})
	log.Print("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
