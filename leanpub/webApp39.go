// webApp39
package main

import (
	. "fmt"
	. "net/http"
	"time"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message := "hello pinas!"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	Spawn(
		func() { ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil) },
		func() { ListenAndServe(ADDRESS, nil) },
	)
}

func Spawn(f ...func()) {
	done := make(chan bool)

	for _, s := range f {
		go func() {
			s()
			done <- true
		}()
		time.Sleep(time.Second)
	}

	for l := len(f); l > 0; l-- {
		<-done
	}
}
