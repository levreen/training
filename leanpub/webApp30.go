// webApp30.go
package main

import (
	. "fmt"
	. "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
	HandleFunc("/hello", Hello)
	ListenAndServe(ADDRESS, nil)
}

func Hello(w ResponseWriter, r *Request) {
	w.Header().Set("Context-Type", "text/plain")
	Fprintf(w, MESSAGE)
}
