package main

import (
	"fmt"
	"net/http"
)

// Fprintf formats by what is specified, and writes to the first parameter
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is amazing!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go by Zanis, it is also amazing")
}

func main() {
	http.HandleFunc("/", index_handler)
	// new handle function for new page.
	http.HandleFunc("/abuout/", about_handler)
	// will listen on 8000 port and for server as it is localhost so will provide nil here
	http.ListenAndServe("8000", nil)
}

// for more information about ResponseWriter https://golang.org/pkg/net/http/
