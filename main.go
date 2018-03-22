package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[1:]
		out, err := os.Create(file)
		if err != nil {
			log.Printf("error: creating file %q: %v", file, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(out, r.Body)
		if err != nil {
			log.Printf("error: copying request body to %q: %v", file, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
