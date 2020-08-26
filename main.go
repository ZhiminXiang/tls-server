package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	err := http.ListenAndServeTLS(":8081", "/var/secret/tls.crt", "/var/secret/tls.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World: HTTPS.\n"))
}
