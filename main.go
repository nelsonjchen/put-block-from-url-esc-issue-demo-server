package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/red%2Fblue.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		_, err := io.WriteString(w, "Need %25 in URL!\n")
		if err != nil {
			log.Fatalf("Could not write: %s\n", err.Error())

		}
	})

	http.HandleFunc("/red/blue.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		_, err := io.WriteString(w, "no!\n")
		if err != nil {
			log.Fatalf("Could not write: %s\n", err.Error())

		}
	})

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
