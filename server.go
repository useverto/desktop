package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
)

// Loadview load view
func Loadview() {

	fs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(fs)))
	go func() {
		http.ListenAndServe(":8000", nil)
	}()
}
