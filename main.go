package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		_, err := io.WriteString(w, "Hello world\n")
		if err != nil {
			log.Fatalf("Could not write: %s\n", err.Error())

		}
	})

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
