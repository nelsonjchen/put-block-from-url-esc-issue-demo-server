package main

// Please see the README.md file for more information.

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("REQUEST | Escaped Path: %#v | Request Received: %#v", req.URL.EscapedPath(), req)

		if req.URL.EscapedPath() == "/red/blue.txt" {
			w.WriteHeader(404)
			_, err := fmt.Fprintf(w, "This path actually doesn't exist.")
			if err != nil {
				return
			}
			return
		}

		if req.URL.EscapedPath() == "/red%2Fblue.txt" {
			w.WriteHeader(200)
			_, err := fmt.Fprintf(w, "This path exists!")
			if err != nil {
				return
			}
			return
		}

		_, err := fmt.Fprintf(w, "Debug: Escaped Path Recieved\n%#v\n%#v", req.URL, req.URL.EscapedPath())
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
